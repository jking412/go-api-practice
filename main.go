package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-api-practice/app/cmd"
	"go-api-practice/bootstrap"
	btsconfig "go-api-practice/config"
	"go-api-practice/pkg/config"
	"go-api-practice/pkg/console"
	"os"
)

func init() {
	btsconfig.Initialize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "go-api-practice",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		PersistentPreRun: func(command *cobra.Command, args []string) {
			config.InitConfig(cmd.Env)

			bootstrap.SetupLogger()

			bootstrap.SetupDB()

			bootstrap.SetupRedis()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
	)

	cmd.RegisterDefaultCommand(rootCmd, cmd.CmdServe)

	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
