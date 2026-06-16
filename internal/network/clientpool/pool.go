package clientpool

import (
	"context"
	"gatekeeper-core/internal/config"
	"gatekeeper-core/pkg/iso8583"
	"gatekeeper-core/pkg/logging"
)

type Pool struct {
	address      string
	workerCount  int
	outboundChan <-chan *iso8583.Message
	workers      []*Worker
}

func New(workerCount int, outboundChan <-chan *iso8583.Message) *Pool {
	return &Pool{
		address:      config.Settings.BankAddr,
		workerCount:  workerCount,
		outboundChan: outboundChan,
		workers:      make([]*Worker, 0, workerCount),
	}
}

func (p *Pool) Start(ctx context.Context) error {
	logging.Log.Debug("initializing TCP client pool")

	// Запускаем воркеров
	for i := 1; i <= p.workerCount; i++ {
		w := NewWorker(i, p.address, p.outboundChan)
		p.workers = append(p.workers, w)

		// Каждый работает в своей независимой горутине
		go w.Start(ctx)
	}

	<-ctx.Done() // Ждем сигнала шатдауна от errgroup в main.go
	logging.Log.Debug("stopping TCP client pool")
	return nil
}
