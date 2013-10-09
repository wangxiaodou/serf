package cli

import (
	"flag"
	"fmt"
	"strings"
)

// JoinCommand is a Command implementation that tells a running Serf
// agent to join another.
type JoinCommand struct{}

func (c *JoinCommand) Help() string {
	helpText := `
Usage: serf join [options] address ...

  Tells a running Serf agent (with "serf agent") to join the cluster
  by specifying at least one existing member.

Options:

  -rpc-addr=127.0.0.1:7373  RPC address of the Serf agent.
`
	return strings.TrimSpace(helpText)
}

func (c *JoinCommand) Run(args []string, ui Ui) int {
	cmdFlags := flag.NewFlagSet("join", flag.ContinueOnError)
	cmdFlags.Usage = func() { ui.Output(c.Help()) }
	rpcAddr := RPCAddrFlag(cmdFlags)
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	client, err := RPCClient(*rpcAddr)
	if err != nil {
		ui.Error(fmt.Sprintf("Error connecting to Serf agent: %s", err))
		return 1
	}
	defer client.Close()

	return 0
}

func (c *JoinCommand) Synopsis() string {
	return "Tell Serf agent to join cluster"
}
