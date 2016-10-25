package enable

import (
	"github.com/emersion/go-imap/server"
)

type Conn interface {}

type conn struct {
	server.Conn

	extensions []server.Extension
}

type Handler struct {
	Command

	ext *extension
}

func (h *Handler) Handle(conn server.Conn) error {
	var enabled []string
	for _, cap := range h.Capabilities {
		// TODO
	}

	return conn.WriteResp(&Response{})
}

type extension struct {
	extensions []server.Extension
}

func (ext *extension) Capabilities(c server.Conn) []string {
	return []string{Capability}
}

func (ext *extension) Command(name string) server.HandlerFactory {
	if name != commandName {
		return nil
	}

	return func() server.Handler {
		return &Handler{ext: ext}
	}
}

func (ext *extension) NewConn(c server.Conn) server.Conn {
	return &conn{Conn: c}
}

func NewExtension(serverID ID) server.Extension {
	return &extension{serverID}
}
