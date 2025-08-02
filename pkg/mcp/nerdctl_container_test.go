package mcp

import (
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestContainerInspect(t *testing.T) {
	testCase(t, func(c *mcpContext) {
		toolResult, err := c.callTool("container_inspect", map[string]any{
			"name": "example-container",
		})
		t.Run("container_inspect returns OK", func(t *testing.T) {
			if err != nil {
				t.Fatalf("call tool failed %v", err)
			}
			if toolResult.IsError {
				t.Fatalf("call tool failed")
			}
		})
		t.Run("container_inspect inspects provided container", func(t *testing.T) {
			if !strings.HasPrefix(toolResult.Content[0].(mcp.TextContent).Text, "nerdctl inspect example-container") {
				t.Fatalf("unexpected result %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
	})
}

func TestContainerList(t *testing.T) {
	testCase(t, func(c *mcpContext) {
		toolResult, err := c.callTool("container_list", map[string]any{})
		t.Run("container_list returns OK", func(t *testing.T) {
			if err != nil {
				t.Fatalf("call tool failed %v", err)
			}
			if toolResult.IsError {
				t.Fatalf("call tool failed")
			}
			if !strings.HasPrefix(toolResult.Content[0].(mcp.TextContent).Text, "nerdctl container list -a") {
				t.Fatalf("unexpected result %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
	})
}

func TestContainerLogs(t *testing.T) {
	testCase(t, func(c *mcpContext) {
		toolResult, err := c.callTool("container_logs", map[string]any{
			"name": "example-container",
		})
		t.Run("container_logs returns OK", func(t *testing.T) {
			if err != nil {
				t.Fatalf("call tool failed %v", err)
			}
			if toolResult.IsError {
				t.Fatalf("call tool failed")
			}
		})
		t.Run("container_logs retrieves logs from provided container", func(t *testing.T) {
			if !strings.HasPrefix(toolResult.Content[0].(mcp.TextContent).Text, "nerdctl logs example-container") {
				t.Fatalf("unexpected result %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
	})
}

func TestContainerRemove(t *testing.T) {
	testCase(t, func(c *mcpContext) {
		toolResult, err := c.callTool("container_remove", map[string]any{
			"name": "example-container",
		})
		t.Run("container_remove returns OK", func(t *testing.T) {
			if err != nil {
				t.Fatalf("call tool failed %v", err)
			}
			if toolResult.IsError {
				t.Fatalf("call tool failed")
			}
		})
		t.Run("container_remove removes provided container", func(t *testing.T) {
			if !strings.HasPrefix(toolResult.Content[0].(mcp.TextContent).Text, "nerdctl container rm example-container") {
				t.Fatalf("unexpected result %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
	})
}

func TestContainerRun(t *testing.T) {
	testCase(t, func(c *mcpContext) {
		toolResult, err := c.callTool("container_run", map[string]any{
			"imageName": "example.com/org/image:tag",
		})
		t.Run("container_run returns OK", func(t *testing.T) {
			if err != nil {
				t.Fatalf("call tool failed %v", err)
			}
			if toolResult.IsError {
				t.Fatalf("call tool failed")
			}
		})
		t.Run("container_run runs provided image", func(t *testing.T) {
			if !strings.HasSuffix(toolResult.Content[0].(mcp.TextContent).Text, " example.com/org/image:tag\n") {
				t.Fatalf("unexpected result %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
		t.Run("container_run runs in detached mode", func(t *testing.T) {
			if !strings.Contains(toolResult.Content[0].(mcp.TextContent).Text, " -d ") {
				t.Fatalf("unexpected result %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
		t.Run("container_run publishes all exposed ports", func(t *testing.T) {
			if !strings.Contains(toolResult.Content[0].(mcp.TextContent).Text, " --publish-all ") {
				t.Fatalf("unexpected result %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
		toolResult, err = c.callTool("container_run", map[string]any{
			"imageName": "example.com/org/image:tag",
			"ports": []any{
				1337, // Invalid entry to test
				"8080:80",
				"8082:8082",
				"8443:443",
			},
		})
		t.Run("container_run with ports returns OK", func(t *testing.T) {
			if err != nil {
				t.Fatalf("call tool failed %v", err)
			}
			if toolResult.IsError {
				t.Fatalf("call tool failed")
			}
		})
		t.Run("container_run with ports publishes provided ports", func(t *testing.T) {
			if !strings.Contains(toolResult.Content[0].(mcp.TextContent).Text, " --publish=8080:80 ") {
				t.Fatalf("expected port --publish=8080:80, got %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
			if !strings.Contains(toolResult.Content[0].(mcp.TextContent).Text, " --publish=8082:8082 ") {
				t.Fatalf("expected port --publish=8082:8082, got %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
			if !strings.Contains(toolResult.Content[0].(mcp.TextContent).Text, " --publish=8443:443 ") {
				t.Fatalf("expected port --publish=8443:443, got %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
		toolResult, err = c.callTool("container_run", map[string]any{
			"imageName": "example.com/org/image:tag",
			"ports":     []any{"8080:80"},
			"environment": []any{
				"KEY=VALUE",
				"FOO=BAR",
			},
		})
		t.Run("container_run with environment returns OK", func(t *testing.T) {
			if err != nil {
				t.Fatalf("call tool failed %v", err)
			}
			if toolResult.IsError {
				t.Fatalf("call tool failed")
			}
		})
		t.Run("container_run with environment sets provided environment variables", func(t *testing.T) {
			if !strings.Contains(toolResult.Content[0].(mcp.TextContent).Text, " --env KEY=VALUE ") {
				t.Fatalf("expected env --env KEY=VALUE, got %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
			if !strings.Contains(toolResult.Content[0].(mcp.TextContent).Text, " --env FOO=BAR ") {
				t.Fatalf("expected env --env FOO=BAR, got %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
	})
}

func TestContainerStop(t *testing.T) {
	testCase(t, func(c *mcpContext) {
		toolResult, err := c.callTool("container_stop", map[string]any{
			"name": "example-container",
		})
		t.Run("container_stop returns OK", func(t *testing.T) {
			if err != nil {
				t.Fatalf("call tool failed %v", err)
			}
			if toolResult.IsError {
				t.Fatalf("call tool failed")
			}
		})
		t.Run("container_stop stops provided container", func(t *testing.T) {
			if !strings.HasPrefix(toolResult.Content[0].(mcp.TextContent).Text, "nerdctl container stop example-container") {
				t.Fatalf("unexpected result %v", toolResult.Content[0].(mcp.TextContent).Text)
			}
		})
	})
}
