package server

import (
	"context"
	"fmt"
	"gatekeeper-core/internal/network"
	"gatekeeper-core/pkg/iso8583"
	"gatekeeper-core/pkg/logging"
	"net"
)

func (s *Server) handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	logging.Log.Debug(fmt.Sprintf("new connection established from %s", conn.RemoteAddr()))

	for {
		rawBytes, err := network.ReadISO8583Frame(conn)
		if err != nil {
			return // при закрытии соединения клиентом - ничего не делаем
		}

		msg := &iso8583.Message{}
		err = msg.Unmarshal(rawBytes)
		if err != nil {
			// todo базовый обработчик, не принимающий сообщение, которые отвечает format error
			logging.Log.Error("parsing message error: %v", err)
			continue
		}

		s.outboundChan <- s.router.Handle(ctx, msg)
	}
}
