package mcp

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) initNerdctlVolume() []server.ServerTool {
	return []server.ServerTool{
		{mcp.NewTool("volume_list",
			mcp.WithDescription("List all the available Nerdctl volumes"),
		), s.volumeList},
	}
}

func (s *Server) volumeList(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return NewTextResult(s.nerdctl.VolumeList()), nil
}
