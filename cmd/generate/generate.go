package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/Peyton232/HTTPServerGenerator/pkg/codegen"
	"github.com/Peyton232/HTTPServerGenerator/pkg/util"
	"github.com/bitfield/script"
	"gopkg.in/yaml.v3"
)

var MODELS_CONFIG codegen.Configuration = codegen.Configuration{
	PackageName: "models",
	Generate: codegen.GenerateOptions{
		Models: true,
	},
}

var API_CONFIG codegen.Configuration = codegen.Configuration{}

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

	yamlPath string
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
		printVersion()
		return
	}

	if flag.NArg() < 1 {
		fmt.Println("Please specify a path to a OpenAPI 3.0 spec file")
		os.Exit(1)
	}

	// Read in config file from disk
	var opts configuration
	err := readConfig(&opts)
	if err != nil {
		errExit("config error: %v", err)
	}

	// Ensure default values are set if user hasn't specified some needed fields.
	opts.Configuration = opts.UpdateDefaults()

	// TODO make an output direc interpreter
	opts.OutputDirec = "../"

	// Now, ensure that the config options are valid.
	if err := opts.Validate(); err != nil {
		errExit("configuration error: %v", err)
	}

	// Generate code
	modelCode, serverCode, err := opts.GenerateCode()
	if err != nil {
		errExit("generation error: %v", err)
	}

	//setup go project
	err = opts.InitializeProject()
	if err != nil {
		errExit("setup error: %v", err)
	}

	// write contents fo files
	err = os.WriteFile("./models/models.gen.go", []byte(modelCode), 0644)
	if err != nil {
		errExit("error writing generated code to file: %s", err)
	}
	err = os.WriteFile("./api/api.gen.go", []byte(serverCode), 0644)
	if err != nil {
		errExit("error writing generated code to file: %s", err)
	}

	// Generate spec folder and generation commands
	opts.specGeneration()

	//TODO generate handlers

	//TODO generate main

	// cleanup go errors
	script.Exec("go get -u ./...").Wait()
	script.Exec("go mod tidy").Wait()

	// read in types and seperate by id/name types
	// call DB generator with supplied types
	// same issue with handler generation
}

func printVersion() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Fprintln(os.Stderr, "error reading build info")
		os.Exit(1)
	}
	fmt.Println(bi.Main.Path + "/cmd/generator")
	fmt.Println(bi.Main.Version)
}

func readConfig(opts *configuration) error {
	if flagConfigFile == "" {
		return fmt.Errorf("no config file proovided")
	}

	buf, err := os.ReadFile(flagConfigFile)
	if err != nil {
		return fmt.Errorf("error reading config file '%s': %v", flagConfigFile, err)
	}
	err = yaml.Unmarshal(buf, &opts)
	if err != nil {
		return fmt.Errorf("error parsing'%s' as YAML: %v", flagConfigFile, err)
	}

	// save yaml path for later
	opts.yamlPath, _ = filepath.Abs(flag.Arg(0))

	return nil
}

func (opts *configuration) GenerateCode() (modelCode string, serverCode string, err error) {
	swagger, err := util.LoadSwagger(flag.Arg(0))
	if err != nil {
		return "", "", fmt.Errorf("error loading swagger spec in %s\n: %s", flag.Arg(0), err)
	}

	// Get model code
	modelCode, err = codegen.Generate(swagger, MODELS_CONFIG)
	if err != nil {
		return "", "", fmt.Errorf("error generating model code: %s", err)
	}

	// Get server code
	opts.Configuration.PackageName = "api"
	opts.Configuration.AdditionalImports = append(opts.Configuration.AdditionalImports, codegen.AdditionalImport{Alias: ".", Package: opts.ProjectName + "/models"}) //TODO this is confusing
	API_CONFIG = opts.Configuration
	serverCode, err = codegen.Generate(swagger, opts.Configuration)
	if err != nil {
		return "", "", fmt.Errorf("error generating server code: %s", err)

	}

	return modelCode, serverCode, nil
}

func (opts *configuration) InitializeProject() error {
	newPath := filepath.Join(opts.OutputDirec, opts.RootName)
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating root directory of project: %s", err)
	}

	//TODO add error checking
	os.Chdir(opts.OutputDirec + opts.RootName)
	script.Exec("go mod init " + opts.ProjectName).Wait()
	script.Exec("go mod tidy").Wait()

	// setup file direcs
	newPath = filepath.Join("./models")
	err = os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating models directory of project: %s", err)
	}

	newPath = filepath.Join("./api")
	err = os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating api directory of project: %s", err)
	}

	return nil
}

func (opts *configuration) specGeneration() error {
	err := os.MkdirAll("spec", os.ModePerm)
	if err != nil {
		return fmt.Errorf("error making spec directory: %s", err)
	}

	// cp yaml spec file into this direc
	script.Exec("cp " + opts.yamlPath + " ./spec/spec.yaml").Stdout()

	// generate model yaml
	modelsConfigYaml, _ := yaml.Marshal(&MODELS_CONFIG)
	err = os.WriteFile("./spec/modelsConfig.yaml", modelsConfigYaml, 0644)
	if err != nil {
		return fmt.Errorf("error writing modelsConfig.yaml to file: %s", err)
	}

	// generate api yaml
	apiConfigYaml, _ := yaml.Marshal(&API_CONFIG)
	err = os.WriteFile("./spec/apiConfig.yaml", apiConfigYaml, 0644)
	if err != nil {
		return fmt.Errorf("error writing apiConfig.yaml to file: %s", err)
	}

	// generate file with commands to call generation commands
	genCode, _ := codegen.GenerateCommands()
	err = os.WriteFile("./spec/gen.go", []byte(genCode), 0644)
	if err != nil {
		return fmt.Errorf("error writing gen.go to spec directory: %s", err)
	}

	// TODO these commands don't work
	return nil
}
