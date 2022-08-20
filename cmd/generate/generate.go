package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/Peyton232/HTTPServerGenerator/pkg/codegen"
	"github.com/bitfield/script"
	"github.com/deepmap/oapi-codegen/pkg/util"
	"gopkg.in/yaml.v3"
)

func errExit(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

var (
	flagOutputDirec  string
	flagPrintVersion bool
	flagGenerateDB   bool
	flagPackageName  string
	flagConfigFile   string
)

type configuration struct {
	codegen.Configuration `yaml:",inline"`

	// OutputDirec is where the root of the project will be generated
	OutputDirec string `yaml:"output,omitempty"`
}

// CMD
// go generate github.com/peyton232/HTTPserverGenerator/cmd/generate -i=<pathToOpenAPISpec> o=<pathToRootOfProject> -[db Flag]
func main() {
	// cmd flags (--name= type)
	flag.StringVar(&flagOutputDirec, "o", "", "Where to output generated code")
	flag.StringVar(&flagConfigFile, "config", "", "a YAML config file that controls oapi-codegen behavior")
	flag.BoolVar(&flagPrintVersion, "version", false, "when specified, print version and exit")
	flag.BoolVar(&flagGenerateDB, "database", true, "when specified generate db code") //TODO make this modular and a call for it in generations of new project
	flag.StringVar(&flagPackageName, "package", "", "The package name for generated code")

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

	// Read in config file from disk
	var opts configuration
	if flagConfigFile != "" {
		buf, err := ioutil.ReadFile(flagConfigFile)
		if err != nil {
			errExit("error reading config file '%s': %v", flagConfigFile, err)
		}
		err = yaml.Unmarshal(buf, &opts)
		if err != nil {
			errExit("error parsing'%s' as YAML: %v", flagConfigFile, err)
		}
	}

	// Ensure default values are set if user hasn't specified some needed
	// fields.
	opts.Configuration = opts.UpdateDefaults()

	// TODO make an output direc interpreter
	opts.OutputDirec = "../"

	// Now, ensure that the config options are valid.
	if err := opts.Validate(); err != nil {
		errExit("configuration error: %v", err)
	}

	swagger, err := util.LoadSwagger(flag.Arg(0))
	if err != nil {
		errExit("error loading swagger spec in %s\n: %s", flag.Arg(0), err)
	}

	// Get model code
	modelCode, err := codegen.Generate(swagger, codegen.Configuration{
		PackageName: "models",
		Generate: codegen.GenerateOptions{
			Models: true,
		},
	})
	if err != nil {
		errExit("error generating model code: %s\n", err)
	}

	// Get model code
	opts.Configuration.PackageName = "api"
	opts.Configuration.AdditionalImports = append(opts.Configuration.AdditionalImports, codegen.AdditionalImport{Alias: ".", Package: opts.ProjectName + "/models"})
	serverCode, err := codegen.Generate(swagger, opts.Configuration)
	// output: ../api/petstore-server.gen.go
	if err != nil {
		errExit("error generating server code: %s\n", err)
	}

	//setup root direc
	newpath := filepath.Join(opts.OutputDirec, opts.RootName)
	err = os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		errExit("error generating root directory: %s\n", err)
	}

	//TODO generate go mod and sum
	os.Chdir(opts.OutputDirec + opts.RootName)
	pipe := script.Exec("ls -lah")
	pipe.Stdout()

	pipe = script.Exec("go mod init " + opts.ProjectName)
	pipe.Stdout()

	pipe = script.Exec("go mod tidy")
	pipe.Stdout()

	//setup file direcs
	newpath = filepath.Join("./models")
	err = os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		errExit("error generating models directory: %s\n", err)
	}

	newpath = filepath.Join("./api")
	err = os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		errExit("error generating api directory: %s\n", err)
	}

	// write contents fo files
	err = ioutil.WriteFile("./models/models.gen.go", []byte(modelCode), 0644)
	if err != nil {
		errExit("error writing generated code to file: %s", err)
	}

	err = ioutil.WriteFile("./api/api.gen.go", []byte(serverCode), 0644)
	if err != nil {
		errExit("error writing generated code to file: %s", err)
	}

	// TODO generate spec folder and generation

	//TODO generate main

	//TODO generate handlers

	// cleanup
	script.Exec("go get -u ./...").Wait()
	script.Exec("go mod tidy").Wait()

	// read in types and seperate by id/name types
	// call DB generator with supplied types
	// same issue with handler generation
}