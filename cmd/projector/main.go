package main

import (
    "github.com/gkarthikreddi/projector/pkg/config"
    "log"
    "fmt"
    "encoding/json"
)

func main() {
    opts, err := config.GetOpts()
    if err != nil {
        log.Fatalf("Unable to get options %v", err)
    }
}