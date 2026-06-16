package main

import (
	"context"
	"gatekeeper-core/internal/config"
	"gatekeeper-core/internal/transport"
	"gatekeeper-core/pkg/setup"
	"os"
	"os/signal"
	"syscall"

	"gatekeeper-core/internal/network/clientpool"
	"gatekeeper-core/internal/network/server"
	"gatekeeper-core/pkg/iso8583"
	"gatekeeper-core/pkg/logging"

	"golang.org/x/sync/errgroup"
)

func main() {
	// 1. Собираем конфиг приложения
	err := setup.Configure(&config.Settings)
	if err != nil {
		panic(err)
	}

	// 2. Собираем логи и спеку iso8583
	logging.Init()
	iso8583.Init()

	err = transport.Init()
	if err != nil {
		panic(err)
	}

	// 3. Собираем канал для передачи ответов от сервера в пул и инициализируем сеть
	outboundChan := make(chan *iso8583.Message, 1000)
	tcpPool := clientpool.New(5, outboundChan)

	// routes.go
	tcpServer := server.New(isoRouter(), outboundChan)

	// 4. Запуск через errgroup для координации горутин
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return tcpPool.Start(gCtx) // Остановится, когда gCtx будет отменен
	})
	g.Go(func() error {
		return tcpServer.Start(gCtx) // Остановится, когда gCtx будет отменен
	})

	// 5. Ожидание завершения
	if err := g.Wait(); err != nil {
		logging.Log.Error("app shutdown with error: " + err.Error())
	} else {
		logging.Log.Error("app gracefully stopped")
	}
}
