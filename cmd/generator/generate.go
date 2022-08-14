package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/Peyton232/HTTPServerGenerator/pkg/codegen"
)

func errExit(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

var (
	flagOutputDirec  string
	flagConfigFile   string
	flagPrintVersion bool
	flagGenerateDB   bool
)

type configuration struct {
	codegen.Configuration `yaml:",inline"`

	// OutputFile is the filename to output.
	OutputFile string `yaml:"output,omitempty"`
}

//CMD
// go generate github.com.peyton232/HTTPserverGenerator/cmd/generate -i=<pathToOpenAPISpec> o=<pathToRootOfProject> -[db Flag]
func main() {
	// cmd flags
	flag.StringVar(&flagOutputDirec, "o", "", "Where to output generated code")
	flag.StringVar(&flagConfigFile, "config", "", "a YAML config file that controls oapi-codegen behavior")
	flag.BoolVar(&flagPrintVersion, "version", false, "when specified, print version and exit")
	flag.BoolVar(&flagGenerateDB, "database", true, "when specified generate db code")

	flag.Parse()

	if flagPrintVersion {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Fprintln(os.Stderr, "error reading build info")
			os.Exit(1)
		}
		fmt.Println(bi.Main.Path + "/cmd/generator")
		fmt.Println(bi.Main.Version)
		return
	}

	if flag.NArg() < 1 {
		fmt.Println("Please specify a path to a OpenAPI 3.0 spec file")
		os.Exit(1)
	}

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
