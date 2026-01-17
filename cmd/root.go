package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "envman",
	Short:   "a tool for managing .env files",
	Long:    `envman is a CLI to hotswap in different .env files when working with mutiple environments, particularly useful in a monorepo`,
	Version: Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
