package tcp_client

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"net"
)

// TCP server
type Client struct {
	conn         net.Conn
	address      string // Address to open connection: localhost:9999
	onNewMessage func(message string)
}

// Send text message to client
func (c *Client) Send(message string) error {
	return c.SendBytes([]byte(message))
}

// Send bytes to client
func (c *Client) SendBytes(b []byte) error {
	_, err := c.conn.Write(append(b, 0x0))
	if err != nil {
		logrus.Errorf("c.conn.Write err:%v", err)
		c.conn.Close()
	}
	return err
}

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// Called when Client receives new message
func (s *Client) OnNewMessage(callback func(message string)) {
	s.onNewMessage = callback
}

func (s *Client) Connect() error {
	conn, err := net.Dial("tcp", s.address)
	if err != nil {
		logrus.Errorf("Connect net.Dial err. err:%v", err)
		return err
	}
	s.conn = conn

	reader := bufio.NewReader(s.conn)
	go func() {
		for {
			message, err := reader.ReadString(0x0)
			if err != nil {
				s.conn.Close()
				return
			}

			s.onNewMessage(message[:len(message)-1])
		}
	}()

	return err
}

// Creates new tcp client instance
func New(address string) *Client {
	client := &Client{
		address: address,
	}

	return client
}
