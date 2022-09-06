package make

import (
	"embed"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"go-api-practice/pkg/console"
	"go-api-practice/pkg/file"
	"go-api-practice/pkg/str"
	"strings"
)

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

//go:embed stubs
var stubsFS embed.FS

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	CmdMake.AddCommand(
		CmdMakeCMD,
		CmdMakeModel,
		CmdMakeAPIController,
		CmdMakeRequest,
	)
}

// makeModelFromString 格式化用户输入的内容
func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = str.Singular(strcase.ToCamel(name))
	model.StructNamePlural = str.Plural(model.StructName)
	model.TableName = str.Snake(model.StructNamePlural)
	model.VariableName = str.LowerCamel(model.StructName)
	model.PackageName = str.Snake(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)
	return model
}

// createFileFromStub 读取 stub 文件并进行变量替换
// 最后一个选项可选，如若传参，应传 map[string]string 类型，作为附加的变量搜索替换
func createFileFromStub(filePath string, stubName string, model Model, variables ...interface{}) {
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}

	if file.Exists(filePath) {
		console.Exit("file already exists: " + filePath)
	}

	modelData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}
	modelStub := string(modelData)
	replaces["{{TableName}}"] = model.TableName
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{PackageName}}"] = model.PackageName

	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}

	err = file.Put([]byte(modelStub), filePath)
	if err != nil {
		console.Exit(err.Error())
	}

	console.Success("created: " + filePath)
}
