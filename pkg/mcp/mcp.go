package mcp

import (
	"slices"

	"github.com/noelshebert/nerdctl-mcp-server/pkg/nerdctl"
	"github.com/noelshebert/nerdctl-mcp-server/pkg/version"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server  *server.MCPServer
	nerdctl nerdctl.Nerdctl
}

func NewSever() (*Server, error) {
	s := &Server{
		server: server.NewMCPServer(
			version.BinaryName,
			version.Version,
			server.WithResourceCapabilities(true, true),
			server.WithPromptCapabilities(true),
			server.WithToolCapabilities(true),
			server.WithLogging(),
		),
	}
	var err error
	if s.nerdctl, err = nerdctl.NewNerdctl(); err != nil {
		return nil, err
	}
	s.server.AddTools(slices.Concat(
		s.initNerdctlContainer(),
		s.initNerdctlImage(),
		s.initNerdctlNetwork(),
		s.initNerdctlVolume(),
	)...)
	return s, nil
}

func (s *Server) ServeStdio() error {
	return server.ServeStdio(s.server)
}

func (s *Server) ServeSse(baseURL string) *server.SSEServer {
	options := make([]server.SSEOption, 0)
	if baseURL != "" {
		options = append(options, server.WithBaseURL(baseURL))
	}
	return server.NewSSEServer(s.server, options...)
}

func NewTextResult(content string, err error) *mcp.CallToolResult {
	if err != nil {
		return &mcp.CallToolResult{
			IsError: true,
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: err.Error(),
				},
			},
		}
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: content,
			},
		},
	}
}
