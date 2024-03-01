package main

import (
	"os"
)

func commandExit(cfg *config, args ...string) error { // NOTE help and exit shouldn't have config? // Seems instructor uses this...
	os.Exit(0)

	return nil
}
