package gomavlib

import (
	"math/rand"
	"net"
	"time"
)

// netTimedConn forces a net.Conn to use timeouts
type netTimedConn struct {
	conn net.Conn
}

func (c *netTimedConn) Close() error {
	return c.conn.Close()
}

func (c *netTimedConn) Read(buf []byte) (int, error) {
	err := c.conn.SetReadDeadline(time.Now().Add(netReadTimeout))
	if err != nil {
		return 0, err
	}
	return c.conn.Read(buf[:])
}

func (c *netTimedConn) Write(buf []byte) (int, error) {
	err := c.conn.SetWriteDeadline(time.Now().Add(netWriteTimeout))
	if err != nil {
		return 0, err
	}
	return c.conn.Write(buf)
}

func randomByte() byte {
	buf := make([]byte, 1)
	rand.Read(buf)
	return buf[0]
}

func uint24Decode(in []byte) uint32 {
	return uint32(in[2])<<16 | uint32(in[1])<<8 | uint32(in[0])
}

func uint24Encode(in uint32) []byte {
	ret := make([]byte, 3)
	ret[0] = byte(in)
	ret[1] = byte(in >> 8)
	ret[2] = byte(in >> 16)
	return ret
}

func uint48Decode(in []byte) uint64 {
	return uint64(in[5])<<40 | uint64(in[4])<<32 | uint64(in[3])<<24 |
		uint64(in[2])<<16 | uint64(in[1])<<8 | uint64(in[0])
}

func uint48Encode(in uint64) []byte {
	ret := make([]byte, 6)
	ret[0] = byte(in)
	ret[1] = byte(in >> 8)
	ret[2] = byte(in >> 16)
	ret[3] = byte(in >> 24)
	ret[4] = byte(in >> 32)
	ret[5] = byte(in >> 40)
	return ret
}