package jumpclient

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/liucxer/jumpserver/internal/msg"
	"github.com/liucxer/jumpserver/pkg/tcp_client"
	"github.com/sirupsen/logrus"
)

func Register(client *tcp_client.Client, clientName string) error {
	m := msg.Msg{
		MsgID:          uuid.New().String(),
		Register:       true,
		FromClientName: clientName,
	}

	bts, err := json.Marshal(m)
	if err != nil {
		logrus.Errorf("err :%v", err)
		return err
	}
	_ = client.SendBytes(bts)

	return err
}

func ExecCmd(client *tcp_client.Client, ToClient string, cmStr string) error {
	return nil
}
