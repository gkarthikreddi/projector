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

    conf, err := config.NewConfig(opts)
    if err != nil {
        log.Fatalf("Unable to get config %v", err)
    }
    
    proj := config.NewProjector(conf)
    if conf.Operation == config.Print {
        if len(conf.Args) == 0 {
            data := proj.GetValueAll()
            jsonString, err := json.Marshal(data)
            if err != nil {
                log.Fatal("This line should never occur")
            }
            fmt.Printf("%v", string(jsonString))
        } else if value, ok := proj.GetValue(conf.Args[0]); ok {
            fmt.Printf("%v", value)
        }
    }

    if conf.Operation == config.Add {
        proj.SetValue(conf.Args[0], conf.Args[1])
        proj.Save()
    }

    if conf.Operation == config.Remove {
        proj.RemoveValue(conf.Args[0])
        proj.Save()
    }

}
