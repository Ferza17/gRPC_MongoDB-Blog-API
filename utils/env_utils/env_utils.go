package env_utils

import (
	"fmt"
	"os"
)

func SetEnvironmentVariable() error {
	if err := os.Setenv("PORT", fmt.Sprint(50051)); err != nil {
		return err
	}

	return nil
}
