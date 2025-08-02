package mcp

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) initNerdctlNetwork() []server.ServerTool {
	return []server.ServerTool{
		{mcp.NewTool("network_list",
			mcp.WithDescription("List all the available Nerdctl networks"),
		), s.networkList},
	}
}

func (s *Server) networkList(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return NewTextResult(s.nerdctl.NetworkList()), nil
}
