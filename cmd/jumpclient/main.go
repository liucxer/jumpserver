package main

import (
	"encoding/json"
	"github.com/liucxer/confmiddleware/conflogger"
	"github.com/liucxer/jumpserver/internal/jumpclient"
	"github.com/liucxer/jumpserver/internal/msg"
	"github.com/liucxer/jumpserver/pkg/tcp_client"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"time"
)

func init() {
	var logger = conflogger.Log{
		Name:  "jumpClient",
		Level: "Debug",
	}
	logger.SetDefaults()
	logger.Init()

}

func main() {
	if len(os.Args) < 3 {
		panic("os.Args")
		return
	}
		host := os.Args[1]
		clientName := os.Args[2]
	client := tcp_client.New(host)

	client.OnNewMessage(func(message string) {
		var m msg.Msg
		err := json.Unmarshal([]byte(message), &m)
		if err != nil {
			return
		}

		logrus.Infof("message :%s", m.Cmd)
		cmd := exec.Command("sh", "-c", m.Cmd)
		out, err := cmd.Output()
		if err != nil {
			return
		}

		m.CmdResult = string(out)
		m.FromClientName, m.ToClientName = m.ToClientName, m.FromClientName
		bts, err := json.Marshal(m)
		if err != nil {
			return
		}

		err = client.SendBytes(bts)
		if err != nil {
			return
		}

	})

	err := client.Connect()
	if err != nil {
		return
	}

	time.Sleep(time.Second)
	jumpclient.Register(client, clientName)
	for {
		time.Sleep(time.Second)
	}
}
