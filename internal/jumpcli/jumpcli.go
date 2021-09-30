package jumpcli

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/liucxer/jumpserver/internal/jumpclient"
	"github.com/liucxer/jumpserver/internal/msg"
	"github.com/liucxer/jumpserver/pkg/tcp_client"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

type CliConf struct {
	JumpServer string `json:"jumpServer"`
	ToHost     string `json:"toHost"`
}

var cliConf = CliConf{}

func init() {
	conf := ".jumpcli.conf"
	bts, err := ioutil.ReadFile(conf)
	if err != nil {
		panic(fmt.Sprintf("read conf error err:%v", err))
	}

	err = json.Unmarshal(bts, &cliConf)
	if err != nil {
		panic(fmt.Sprintf("json.Unmarshal error err:%v", err))
	}
}

func Run(cmd string) error {
	logrus.Infof("cmd:%s", cmd)
	client := tcp_client.New(cliConf.JumpServer)

	client.OnNewMessage(func(message string) {
		// new message received
		logrus.Infof("message")
		var m msg.Msg
		err := json.Unmarshal([]byte(message), &m)
		if err != nil {
			return
		}
		fmt.Println(m.CmdResult)
		os.Exit(1)
	})

	err := client.Connect()
	if err != nil {
		return err
	}

	clientName := uuid.New().String()
	err = jumpclient.Register(client, clientName)
	if err != nil {
		return err
	}

	m := msg.Msg{
		MsgID:          uuid.New().String(),
		FromClientName: clientName,
		ToClientName:   cliConf.ToHost,
		Cmd:            cmd,
	}

	bts, err := json.Marshal(m)
	if err != nil {
		logrus.Errorf("err :%v", err)
		return err
	}
	err = client.SendBytes(bts)
	if err != nil {
		return err
	}
	logrus.Info("send success")

	time.Sleep(3 * time.Second)
	logrus.Info("timeout")
	return nil
}
