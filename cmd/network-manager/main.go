package main

import (
	"os"

	apiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/cli"

	"github.com/kosmos.io/kosmos/cmd/network-manager/app"
)

func main() {
	ctx := apiserver.SetupSignalContext()
	cmd := app.NewAgentCommand(ctx)
	code := cli.Run(cmd)
	os.Exit(code)
}
