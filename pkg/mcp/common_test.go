package mcp

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"
	"os/exec"
	"path"
	"runtime"
	"testing"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type mcpContext struct {
	nerdctlBinaryDir string
	ctx              context.Context
	cancel           context.CancelFunc
	mcpServer        *Server
	mcpHTTPServer    *httptest.Server
	mcpClient        *client.Client
}

func (c *mcpContext) beforeEach(t *testing.T) {
	var err error
	c.ctx, c.cancel = context.WithCancel(context.Background())
	if c.mcpServer, err = NewSever(); err != nil {
		t.Fatal(err)
		return
	}
	c.mcpHTTPServer = server.NewTestServer(c.mcpServer.server)
	if c.mcpClient, err = client.NewSSEMCPClient(c.mcpHTTPServer.URL + "/sse"); err != nil {
		t.Fatal(err)
		return
	}
	if err = c.mcpClient.Start(c.ctx); err != nil {
		t.Fatal(err)
		return
	}
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{Name: "test", Version: "1.33.7"}
	_, err = c.mcpClient.Initialize(c.ctx, initRequest)
	if err != nil {
		t.Fatal(err)
		return
	}
}

func (c *mcpContext) afterEach() {
	c.cancel()
	_ = c.mcpClient.Close()
	c.mcpHTTPServer.Close()
}

func testCase(t *testing.T, test func(c *mcpContext)) {
	mcpCtx := &mcpContext{
		nerdctlBinaryDir: withNerdctlBinary(t),
	}
	mcpCtx.beforeEach(t)
	defer mcpCtx.afterEach()
	test(mcpCtx)
}

// callTool helper function to call a tool by name with arguments
func (c *mcpContext) callTool(name string, args map[string]any) (*mcp.CallToolResult, error) {
	callToolRequest := mcp.CallToolRequest{}
	callToolRequest.Params.Name = name
	callToolRequest.Params.Arguments = args
	return c.mcpClient.CallTool(c.ctx, callToolRequest)
}

func (c *mcpContext) withNerdctlOutput(outputLines ...string) {
	if len(outputLines) > 0 {
		f, _ := os.Create(path.Join(c.nerdctlBinaryDir, "output.txt"))
		defer f.Close()
		for _, line := range outputLines {
			_, _ = f.WriteString(line + "\n")
		}
	}
}

func withNerdctlBinary(t *testing.T) string {
	binDir := t.TempDir()
	binary := "nerdctl"
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	output, err := exec.
		Command("go", "build", "-o", path.Join(binDir, binary),
			path.Join("..", "..", "testdata", "nerdctl", "main.go")).
		CombinedOutput()
	if err != nil {
		panic(fmt.Errorf("failed to generate nerdctl binary: %w, output: %s", err, string(output)))
	}
	if os.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH")) != nil {
		panic("failed to set PATH")
	}
	return binDir
}
