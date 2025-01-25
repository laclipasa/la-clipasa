package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	var env, specPath string

	flag.StringVar(&env, "env", "", "Environment Variables filename")
	flag.StringVar(&specPath, "spec-path", "openapi.yaml", "OpenAPI specification filepath")
	flag.Parse()

	var errs []string

	if env == "" {
		errs = append(errs, "    - env is required but unset")
	}

	if len(errs) > 0 {
		log.Fatalf("error: \n" + strings.Join(errs, "\n"))
	}

	// errC, err := server.Run(env, specPath)
	// if err != nil {
	// 	log.Fatalf("Couldn't run: %s", err)
	// }

	// fmt.Println("\n" + colors.G + colors.Bold +
	// 	"Visit the docs: \n\t" +
	// 	colors.G + internal.BuildAPIURL("docs") + "\n\t" +
	// 	colors.G + internal.BuildAPIURL("docs-redoc") + " " +
	// 	colors.Off + "\n")

	// if err := <-errC; err != nil {
	// 	log.Fatalf("Error while running: %s", err)
	// }
}
