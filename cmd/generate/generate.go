package main

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
