package ipc

import "net"

type IPCClient interface {
	Send(msg string) error
	Read() (string, error)
	Close()
}
type ipcClient struct {
	C net.Conn
}

func NewClient(file string) (IPCClient, error) {
	conn, err := net.Dial("unix", file)
	if err != nil {
		return nil, err
	}
	return &ipcClient{C: conn}, nil
}

func (client *ipcClient) Send(msg string) error {
	_, err := client.C.(*net.UnixConn).Write([]byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func (client *ipcClient) Read() (string, error) {
	return "", nil
}

func (client *ipcClient) Close() {
	client.Close()
}
