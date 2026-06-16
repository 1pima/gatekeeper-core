package server

import (
	"context"
	"fmt"
	"gatekeeper-core/pkg/logging"
	"net"
)

func (s *Server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		logging.Log.Error(fmt.Sprintf("TCP Server error %s", err.Error()))
		return err
	}

	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)
	logging.Log.Info(fmt.Sprintf("TCP Server listening on %s", s.address))

	go func() {
		<-ctx.Done()
		logging.Log.Info("TCP Server shutting down")

		defer func(listener net.Listener) {
			_ = listener.Close()
		}(listener)
	}()

	// Бесконечный цикл приема новых подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return nil // Завершаемся штатно
			default:
				logging.Log.Error("failed to accept incoming connection", err.Error())
				continue
			}
		}

		// Каждое подключение обслуживаем в отдельной горутине
		go s.handleConnection(ctx, conn)
	}
}
