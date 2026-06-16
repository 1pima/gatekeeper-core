package logging

import (
	"gatekeeper-core/internal/config"
	"io"
	"log/slog"
	"net"
	"os"
	"time"
)

var Log *slog.Logger

// Init инициализирует логгер
func Init() {
	var writers []io.Writer

	// базово пишем в stdout, если есть syslog - также транслируем туда
	writers = append(writers, os.Stdout)

	if config.Settings.SYSLogAddr != "" {
		conn, err := net.DialTimeout("tcp", config.Settings.SYSLogAddr, 2*time.Second)
		if err != nil {
			slog.Warn("failed to connect to log socket, falling back to stdout only",
				slog.String("addr", config.Settings.SYSLogAddr),
				slog.String("error", err.Error()),
			)
		} else {
			writers = append(writers, newAsyncSocketWriter(conn))
		}
	}

	multiWriter := io.MultiWriter(writers...)
	baseHandler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})
	handler := &ContextHandler{Handler: baseHandler}

	Log = slog.New(handler)
	slog.SetDefault(Log)
}

// asyncSocketWriter обеспечивает неблокирующую отправку логов по сети
type asyncSocketWriter struct {
	ch   chan []byte
	conn net.Conn
}

func newAsyncSocketWriter(conn net.Conn) *asyncSocketWriter {
	w := &asyncSocketWriter{
		// Буфер на случай, если сеть встанет - 1024 по дефолту
		ch:   make(chan []byte, 1024),
		conn: conn,
	}
	go w.start()
	return w
}

// Write реализует интерфейс io.Writer
func (w *asyncSocketWriter) Write(p []byte) (n int, err error) {
	// slog переиспользует буфер p, поэтому для передачи в другой горутине копируем данные
	msg := make([]byte, len(p))
	copy(msg, p)

	select {
	case w.ch <- msg:
		return len(p), nil
	default:
		return len(p), nil
	}
}

func (w *asyncSocketWriter) start() {
	defer w.conn.Close()
	for msg := range w.ch {
		_, err := w.conn.Write(msg)
		if err != nil {
			continue
		}
	}
}
