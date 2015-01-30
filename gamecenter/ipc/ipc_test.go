package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"200", "ECHO:" + method + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	resp1, err1 := client1.Call("method1", "params1")
	resp2, err2 := client2.Call("method2", "params2")
	if resp1.Body != "ECHO:method1params1" || resp2.Body != "ECHO:method2params2" || err1 != nil || err2 != nil {
		t.Error("IpcClient.Call failed. resp1:", resp1, err1, "resp2:", resp2, err2)
	}
	client1.Close()
	client2.Close()
}
