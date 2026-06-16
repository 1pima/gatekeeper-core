package clientpool

import (
	"context"
	"fmt"
	"gatekeeper-core/internal/config"
	"gatekeeper-core/internal/network" // Твой пакет с ReadISO8583Frame и WriteISO8583Frame
	"gatekeeper-core/pkg/iso8583"
	"gatekeeper-core/pkg/logging"
	"net"
	"time"
)

type Worker struct {
	id           int
	address      string
	outboundChan <-chan *iso8583.Message
	conn         net.Conn
}

func NewWorker(id int, address string, outChan <-chan *iso8583.Message) *Worker {
	return &Worker{
		id:           id,
		address:      address,
		outboundChan: outChan,
	}
}

func (w *Worker) Start(ctx context.Context) {
	// 1. Сходу пытаемся подключиться к банку
	w.connect(ctx)

	// 2. Запускаем встроенный пингер (0800) для этого воркера
	go w.startPinger(ctx)

	// 3. Запускаем цикл чтения из сокета (Банк ведь тоже может нам что-то слать в этот канал)
	go w.readLoop(ctx)

	// 4. Основной цикл: слушаем канал ответов от нашего сервера и шлем в банк
	for {
		select {
		case <-ctx.Done():
			w.closeConn()
			return
		case msg, ok := <-w.outboundChan:
			if !ok {
				return
			}
			w.send(ctx, msg)
		}
	}
}

func (w *Worker) connect(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Стучимся в банк с таймаутом коннекта
			conn, err := net.DialTimeout("tcp", w.address, 5*time.Second)
			if err != nil {
				logging.Log.Error(fmt.Sprintf("worker %d connection failed, retrying in 5s", w.id), err)

				time.Sleep(5 * time.Second)
				continue
			}
			w.conn = conn
			logging.Log.Info(fmt.Sprintf("worker %d connected to bank successfully 🔌", w.id))
			return
		}
	}
}

func (w *Worker) send(ctx context.Context, msg *iso8583.Message) {
	if w.conn == nil {
		logging.Log.Error(fmt.Sprintf("worker %d cannot send, connection dead", w.id), nil)
		return
	}

	bytes, err := msg.Marshal()
	if err != nil {
		logging.Log.Error("failed to pack message to bytes", err)
		return
	}

	err = network.WriteISO8583Frame(w.conn, bytes)
	if err != nil {
		logging.Log.Error(fmt.Sprintf("worker %d write error, triggering reconnect", w.id), err)
		w.handleDisconnect(ctx)
	}
}

func (w *Worker) readLoop(ctx context.Context) {
	for {
		if w.conn == nil {
			time.Sleep(1 * time.Second)
			continue
		}

		// Читаем входящие пакеты от банка (например, ответы 0810 на наши пинги)
		rawBytes, err := network.ReadISO8583Frame(w.conn)
		if err != nil {
			select {
			case <-ctx.Done():
				return // Завершаемся, если приложение тушат
			default:
				logging.Log.Error(fmt.Sprintf("worker %d read error, connection dropped by bank", w.id), err)

				w.handleDisconnect(ctx)
				continue
			}
		}

		msg := iso8583.Message{}
		err = msg.Unmarshal(rawBytes)
		if err != nil {
			logging.Log.Error("failed to parse incoming bank message", err)
			continue
		}

		if msg.MTI == "0810" {
			logging.Log.Debug(fmt.Sprintf("worker %d pong (0810) received, connection is healthy.", w.id))
		}
	}
}

func (w *Worker) handleDisconnect(ctx context.Context) {
	w.closeConn()
	w.connect(ctx)
}

func (w *Worker) startPinger(ctx context.Context) {
	var pingInterval = time.Duration(config.Settings.PingInterval) * time.Second
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.sendEchoRequest(ctx)
		}
	}
}

func (w *Worker) sendEchoRequest(ctx context.Context) {
	if w.conn == nil {
		return
	}

	// собираем 0800
	STAN := fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
	msg := iso8583.Message{MTI: iso8583.MTINetworkManagementRequest, NetworkManagementInfo: "301", STAN: STAN}

	w.send(ctx, &msg)
}

func (w *Worker) closeConn() {
	if w.conn != nil {
		_ = w.conn.Close()
		w.conn = nil
	}
}
