package main

import (
	"github.com/spf13/cobra"

	"github.com/openshift/installer/pkg/web"
)

func newWebCmd() *cobra.Command {
	return web.NewCmd()
}
