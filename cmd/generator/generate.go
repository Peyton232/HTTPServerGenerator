package main

import (
	"fmt"
	"os"

	"github.com/Peyton232/HTTPServerGenerator/pkg/codegen"
)

func errExit(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

var (
	flagOutputFile     string
	flagConfigFile     string
	flagOldConfigStyle bool
	flagOutputConfig   bool
	flagPrintVersion   bool
	flagPackageName    string

	// The options below are deprecated, and they will be removed in a future
	// release. Please use the new config file format.
	flagGenerate           string
	flagIncludeTags        string
	flagExcludeTags        string
	flagTemplatesDir       string
	flagImportMapping      string
	flagExcludeSchemas     string
	flagResponseTypeSuffix string
	flagAliasTypes         bool
)

type configuration struct {
	codegen.Configuration `yaml:",inline"`

	// OutputFile is the filename to output.
	OutputFile string `yaml:"output,omitempty"`
}

//CMD
// go generate github.com.peyton232/HTTPserverGenerator/cmd/generate -i=<pathToOpenAPISpec> o=<pathToRootOfProject> -[db Flag]
func main() {
	// get file path from input cmd
	// get output path for repo root
	// read in openAPI yaml file

	// generate api and model config files

	// generate the api and model (this one first) in respective folders
	// minor todo fix model alias issue

	//	generate handler section from stitching of different templates
	// figure out how to handle changes of generation, initially just skip this step if file exiss

	//	generate repo main from template

	// read in types and seperate by id/name types
	// call DB generator with supplied types
	// same issue with handler generation
}
