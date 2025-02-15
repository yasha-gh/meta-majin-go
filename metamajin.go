package majin

import (
	"context"
	"fmt"
	"net/http"
	// "connectrpc.com/connect"
	// 	userv1 "github.com/yasha-gh/meta-majin/gen/user/v1"
	"github.com/yasha-gh/meta-majin/gen/user/v1/userv1connect"
)

func NewMetaMajinClient(hostname *string, port *uint, protocol *string) *MetaMajin {
	newHost := "127.0.0.1"
	var newPort uint = 9111
	newProtocol := "http"

	if hostname != nil {
		newHost = *hostname
	}
	if port != nil {
		newPort = *port
	}
	if protocol != nil {
		newProtocol = *protocol
	}
	clients := initMetaMajinClients(newHost, newPort, newProtocol)
	return &MetaMajin{
		Hostname: newHost,
		Port: newPort,
		Protocol: newProtocol,
		Clients: clients,
	}
}

type MetaMajin struct {
	Hostname string `json:"hostname"`
	Port uint `json:"port"`
	Protocol string `json:"protocol"`
	ctx context.Context `json:"-"`
	Clients MajinClients `json:"-"`
}

// MetaMajin service clients
type MajinClients struct {
	User userv1connect.UserServiceClient
}

func (m *MetaMajin) SetContext(ctx context.Context) {
	m.ctx = ctx
}

func (m *MetaMajin) GetConnectString() string {
	// return fmt.Sprintf("%v://%v:%v", m.Protocol, m.Hostname, m.Port)
	return getMajinConnectString(m.Hostname, m.Port, m.Protocol)
}

func getMajinConnectString(hostname string, port uint, protocol string) string {
	return fmt.Sprintf("%v://%v:%v", protocol, hostname, port)
}

func initMetaMajinClients(hostname string, port uint, protocol string) MajinClients {
	connectString := getMajinConnectString(hostname, port, protocol)
	userClient := userv1connect.NewUserServiceClient(
		http.DefaultClient,
		connectString,
	)
	return MajinClients{
		User: userClient,
	}
}
