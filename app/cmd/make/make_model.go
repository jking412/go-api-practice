package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Create a new model",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1),
}

func runMakeModel(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	dir := fmt.Sprintf("app/models/%s/", model.PackageName)
	os.Mkdir(dir, os.ModePerm)

	createFileFromStub(dir+model.PackageName+"_model.go", "model/model", model)
	createFileFromStub(dir+model.PackageName+"_util.go", "model/model_util", model)
	createFileFromStub(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)
}
