package tcp_server

import (
	"bufio"
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"net"
)

// Client holds info about connection
type Client struct {
	conn   net.Conn
	Server *server
}

// TCP server
type server struct {
	address                  string // Address to open connection: localhost:9999
	config                   *tls.Config
	onNewClientCallback      func(c *Client) error
	onClientConnectionClosed func(c *Client, err error) error
	onNewMessage             func(c *Client, message string) error
}

// Read client data from channel
func (c *Client) listen() {
	err := c.Server.onNewClientCallback(c)
	if err != nil {
		logrus.Errorf("c.Server.onNewClientCallback err. err:%v", err)
	}

	reader := bufio.NewReader(c.conn)
	for {
		message, err := reader.ReadString(0x0)
		if err != nil {
			c.conn.Close()
			err = c.Server.onClientConnectionClosed(c, err)
			if err != nil {
				logrus.Errorf("c.Server.onClientConnectionClosed err. err:%v", err)
			}
			return
		}

		err = c.Server.onNewMessage(c, message[:len(message)-1])
		if err != nil {
			logrus.Errorf("c.Server.onNewMessage err. err:%v", err)
			break
		}
	}
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
		err = c.Server.onClientConnectionClosed(c, err)
		if err != nil {
			logrus.Errorf("c.Server.onClientConnectionClosed err. err:%v", err)
		}
	}
	return err
}

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// Called right after server starts listening new client
func (s *server) OnNewClient(callback func(c *Client) error) {
	s.onNewClientCallback = callback
}

// Called right after connection closed
func (s *server) OnClientConnectionClosed(callback func(c *Client, err error) error) {
	s.onClientConnectionClosed = callback
}

// Called when Client receives new message
func (s *server) OnNewMessage(callback func(c *Client, message string) error) {
	s.onNewMessage = callback
}

// Listen starts network server
func (s *server) Listen() {
	var listener net.Listener
	var err error
	if s.config == nil {
		listener, err = net.Listen("tcp", s.address)
	} else {
		listener, err = tls.Listen("tcp", s.address, s.config)
	}
	if err != nil {
		logrus.Errorf("Error starting TCP server. err:%v", err)
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		client := &Client{
			conn:   conn,
			Server: s,
		}
		go client.listen()
	}
}

// Creates new tcp server instance
func New(address string) *server {
	logrus.Infof("Creating server with address:%s", address)
	server := &server{
		address: address,
	}

	return server
}

func NewWithTLS(address, certFile, keyFile string) *server {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		logrus.Errorf("Error loading certificate files. Unable to create TCP server with TLS functionality.err:%v", err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	server := New(address)
	server.config = config
	return server
}
