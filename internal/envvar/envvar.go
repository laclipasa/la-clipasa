package envvar

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/laclipasa/la-clipasa/internal"
)

// Load reads the env filename and loads it into ENV for the current process.
// It also initializes/replaces app configuration.
func Load(filename string) error {
	if err := godotenv.Load(filename); err != nil {
		return internal.NewErrorf(internal.ErrorCodeUnknown, fmt.Sprintf("loading %s env var file: %s", filename, err))
	}

	if err := internal.NewAppConfig(); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "internal.NewAppConfig")
	}

	return nil
}

// GetEnv returns an environment variable's value or a default if empty.
func GetEnv(key, dft string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return dft
	}

	return v
}
