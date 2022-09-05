package cmd

import (
	"github.com/spf13/cobra"
	"go-api-practice/helpers"
	"go-api-practice/pkg/console"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate the application key",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs,
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("APP key:")
	console.Success(helpers.RandomString(32))
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}
