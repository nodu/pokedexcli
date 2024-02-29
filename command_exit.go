package main

import (
	"os"
)

func commandExit(cfg *config) error { // TODO help and exit shouldn't have config?
	os.Exit(0)

	return nil
}
