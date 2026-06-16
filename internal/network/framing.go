package network

import (
	"encoding/binary"
	"gatekeeper-core/pkg/exc"
	"io"
	"net"
)

func ReadISO8583Frame(conn net.Conn) ([]byte, error) {
	header := make([]byte, 2)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, exc.NewInternalError(err.Error())
	}

	msgLength := binary.BigEndian.Uint16(header)
	if msgLength == 0 {
		return nil, exc.NewConflict("received empty message")
	}

	payload := make([]byte, msgLength)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, exc.NewInternalError("failed to read payload: " + err.Error())
	}
	return payload, nil
}

func WriteISO8583Frame(conn net.Conn, payload []byte) error {
	length := len(payload)
	frame := make([]byte, 2+length)
	binary.BigEndian.PutUint16(frame[:2], uint16(length))
	copy(frame[2:], payload)

	_, err := conn.Write(frame)
	return err
}
