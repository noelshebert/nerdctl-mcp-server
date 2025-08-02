package mcp

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) initNerdctlImage() []server.ServerTool {
	return []server.ServerTool{
		{mcp.NewTool("image_build",
			mcp.WithDescription("Build a Nerdctl image from a Dockerfile"),
			mcp.WithString("dockerFile", mcp.Description("The absolute path to the Dockerfile to build the image from"), mcp.Required()),
			mcp.WithString("imageName", mcp.Description("Specifies the name which is assigned to the resulting image if the build process completes successfully (--tag, -t)")),
		), s.imageBuild},
		{mcp.NewTool("image_list",
			mcp.WithDescription("List the Nerdctl images on the local machine"),
		), s.imageList},
		{mcp.NewTool("image_pull",
			mcp.WithDescription("Copies (pulls) a Nerdctl container image from a registry onto the local machine storage"),
			mcp.WithString("imageName", mcp.Description("Nerdctl container image name to pull"), mcp.Required()),
		), s.imagePull},
		{mcp.NewTool("image_push",
			mcp.WithDescription("Pushes a Nerdctl container image, manifest list or image index from local machine storage to a registry"),
			mcp.WithString("imageName", mcp.Description("Nerdctl container image name to push"), mcp.Required()),
		), s.imagePush},
		{mcp.NewTool("image_remove",
			mcp.WithDescription("Removes a Nerdctl image from the local machine storage"),
			mcp.WithString("imageName", mcp.Description("Nerdctl container image name to remove"), mcp.Required()),
		), s.imageRemove},
	}
}

func (s *Server) imageBuild(_ context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	imageName := ctr.GetArguments()["imageName"]
	if _, ok := imageName.(string); !ok {
		imageName = ""
	}
	return NewTextResult(s.nerdctl.ImageBuild(ctr.GetArguments()["dockerFile"].(string), imageName.(string))), nil
}

func (s *Server) imageList(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return NewTextResult(s.nerdctl.ImageList()), nil
}

func (s *Server) imagePull(_ context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return NewTextResult(s.nerdctl.ImagePull(ctr.GetArguments()["imageName"].(string))), nil
}

func (s *Server) imagePush(_ context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return NewTextResult(s.nerdctl.ImagePush(ctr.GetArguments()["imageName"].(string))), nil
}

func (s *Server) imageRemove(_ context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return NewTextResult(s.nerdctl.ImageRemove(ctr.GetArguments()["imageName"].(string))), nil
}
