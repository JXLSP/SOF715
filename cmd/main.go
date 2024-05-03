package main

import (
	"os"
	"sof/internal/sofapp"
)

func main() {
    if err := sofapp.NewAppServe(); err != nil {
        os.Exit(1)
    }
}

