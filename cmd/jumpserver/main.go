package main

import (
	"github.com/liucxer/confmiddleware/conflogger"
	"github.com/liucxer/jumpserver/internal/jumpserver"
	"github.com/liucxer/jumpserver/pkg/tcp_server"
	"os"
)

func init() {
	var logger = conflogger.Log{
		Name:  "jumpServer",
		Level: "Debug",
	}
	logger.SetDefaults()
	logger.Init()
}

func main() {
	port := os.Args[1]
	server := tcp_server.New("localhost:"+ port)

	server.OnNewClient(jumpserver.NewClientHandle)
	server.OnNewMessage(jumpserver.NewMessageHandle)
	server.OnClientConnectionClosed(jumpserver.ClientConnectionClosedHandle)

	server.Listen()
}
