package main

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/CopperCarbonateCollective/ollama-DreamingMetal/cmd"
)

func main() {
	cobra.CheckErr(cmd.NewCLI().ExecuteContext(context.Background()))
}
