package server

import (
	"net/http"

	"github.com/hueypark/mcp/src/tools/greeter"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Server struct {
	mcpServer *mcp.Server
}

func New() *Server {
	mcpServer := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v1.0.0"}, nil)
	mcp.AddTool(mcpServer, &mcp.Tool{Name: "greet", Description: "say hi"}, greeter.SayHi)

	return &Server{mcpServer: mcpServer}
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, mcp.NewSSEHandler(func(*http.Request) *mcp.Server { return s.mcpServer }))
}
