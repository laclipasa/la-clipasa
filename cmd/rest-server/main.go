package main

import (
	"flag"
	"log"
	"strings"

	"github.com/laclipasa/la-clipasa/internal/http"
)

func main() {
	var env string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.Parse()

	var errs []string

	if env == "" {
		errs = append(errs, "    - env is required but unset")
	}

	if len(errs) > 0 {
		log.Fatalf("%s", "error: \n"+strings.Join(errs, "\n"))
	}

	http.RunServer(env)
}
