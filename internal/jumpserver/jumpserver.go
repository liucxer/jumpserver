package jumpserver

import (
	"encoding/json"
	"github.com/liucxer/jumpserver/internal/msg"
	"github.com/liucxer/jumpserver/pkg/tcp_server"
	"github.com/sirupsen/logrus"
)

var ClientMap = map[string]*tcp_server.Client{}

func NewClientHandle(c *tcp_server.Client) error {
	logrus.Infof("new client connect from:%+v", c.Conn().LocalAddr())
	return nil
}

func NewMessageHandle(c *tcp_server.Client, message string) error {
	m := msg.Msg{}
	err := json.Unmarshal([]byte(message), &m)
	if err != nil {
		logrus.Errorf("NewMessageHandle json.Unmarshal err:%v, message:%s", err, message)
		return err
	}
	logrus.Infof("new message m:%+v", m)

	if m.Register {
		logrus.Infof("register client")
		ClientMap[m.FromClientName] = c
	}

	if m.ToClientName != "" {
		if toClient, ok := ClientMap[m.ToClientName]; ok {
			err = toClient.Send(message)
			logrus.Infof("send to ToClientName:%s", m.ToClientName)
		} else {
			logrus.Errorf("ToClientNotExist :%s", m.ToClientName)
			m.Error = "ToClientNotExist " + m.ToClientName
			if m.CmdResult != "" {
				return nil
			}

			bts, err := json.Marshal(m)
			if err != nil {
				logrus.Errorf("NewMessageHandle json.Marshal err:%v", err)
				return err
			}

			err = c.SendBytes(bts)
		}
	}

	return err
}

func ClientConnectionClosedHandle(c *tcp_server.Client, err error) error {
	logrus.Infof("client connection closed")
	return nil
}
