package lightqueryapi

import (
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
}

func NewClient(
	addr net.UDPAddr,
) (*Client, error) {
	conn, err := net.DialUDP("udp", nil, &addr)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) String() string {
	return fmt.Sprintf("Client{conn: %v}", c.conn)
}
