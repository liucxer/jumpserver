package msg

type ClientType string

const (
	Client ClientType = "Client"
	Cli    ClientType = "Cli"
)

type Msg struct {
	MsgID          string `json:"msgID"`
	ClientType     string `json:"clientType"`
	Register       bool   `json:"register"`
	FromClientName string `json:"fromClientName"`
	Cmd            string `json:"cmd"`
	CmdResult      string `json:"cmdResult"`
	ToClientName   string `json:"toClientName"`
	Error          string `json:"error"`
}
