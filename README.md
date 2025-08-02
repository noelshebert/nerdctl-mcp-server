# nerdctl-mcp-server

MCP Server for Nerdctl coded in Go

nerdctl-mcp-server is a fork of podman-mcp-server (<https://github.com/manusa/podman-mcp-server>)
with all the non-Go stuff stripped away. To be fair, most of the work put into
this fork has been to find/replace any mention of Podman/Docker with Nerdctl.
As Nerdctl like Podman has a command-line API copied from the Docker CLI, this
was not difficult. There was some cleanup of syntax of inline descriptions to
make sense, and removal of comments referring to Podman documentation. Lastly,
code hygene offered by the Go LSP (interface{} replaced by any for example) 
were applied.

WHY?

Well, the answer is historical. I started using Nerdctl as a rootless alternative
to Docker (which was not ready at the time) and Podman at that time was flakey.
In addition, the most current version of Podman was not easily available on Ubuntu,
my distribution of choice. 

Now, Docker has a working rootless mode. I still don't use it because I just don't.
I do use Podman most of the time, however I still use Nerdctl on my main server and it's 
fun just to be able to say "I use nerdctl, by-the-way!" ;-)

I am actively learning all things AI, especially integrating it into my workflow. Most
AI assistive tools make use of MCP (Model Context Protocol) servers and a Podman server
has already done by Marc Nuri (manusa). Why not Nerdctl?

TODO:  Expand this README with installation instructions, configuraton instructions
for Claude Code and Gemini (to name two) and provide formal release targets for Mac 
and Windows. Comment the crap out of the code (for my own benefit) and add some new
features...
