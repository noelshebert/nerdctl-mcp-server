package mcp

import (
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestVolumeList(t *testing.T) {
	testCase(t, func(c *mcpContext) {
		toolResult, err := c.callTool("volume_list", map[string]any{})
		t.Run("volume_list returns OK", func(t *testing.T) {
			if err != nil {
				t.Fatalf("call tool failed %v", err)
			}
			if toolResult.IsError {
				t.Fatalf("call tool failed")
			}
		})
		t.Run("volume_list lists all available volumes", func(t *testing.T) {
			if !strings.HasPrefix(toolResult.Content[0].(mcp.TextContent).Text, "nerdctl volume ls") {
				t.Fatalf("unexpected result %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
	})
}
