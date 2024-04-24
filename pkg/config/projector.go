package config

import (
    "encoding/json"
    "os"
    "path"
)

type Data struct {
    Projector map[string]map[string]string `json:"projector"`
}

type Projector struct {
    data   Data
    conf *Config
}

func (p *Projector) GetValue(key string) (string, bool) {
    found := false
    out := ""

    curr := p.conf.Pwd
    prev := ""

    for ; curr != prev; {
        if dir, ok := p.data.Projector[curr]; ok {
            if value, ok := dir[key]; ok {
                found = true
                out = value
                break
            }
        }
        prev = curr
        curr = path.Dir(curr)
    }

    return out, found
}

func (p *Projector) GetValueAll() map[string]string {
    out := map[string]string{}
    
    paths := []string{}
    curr := p.conf.Pwd
    prev := ""
    for ; curr != prev; {
        paths = append(paths, curr)
        prev = curr
        curr = path.Dir(curr)
    }

    for i := len(paths) - 1; i>=0; i-- {
        if dir, ok := p.data.Projector[paths[i]]; ok {
            for k, v := range dir {
                out[k] = v
            }
        }
    }

    return out
}

func (p *Projector) SetValue(key string, value string) {
    pwd := p.conf.Pwd
    if _, ok := p.data.Projector[pwd]; !ok {
        p.data.Projector[pwd] = map[string]string{}
    }
    p.data.Projector[pwd][key] = value
}

func (p *Projector) RemoveValue(key string) {
    pwd := p.conf.Pwd
    if dir, ok := p.data.Projector[pwd]; ok {
        delete(dir, key)
    }
}

func (p *Projector) Save() error {
    dir := path.Dir(p.conf.Config)
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        err := os.MkdirAll(dir, 0755)
        if err != nil {
            return err
        }
    }
    jsonString, err := json.Marshal(p.data)
        if err != nil {
            return err
        }
    os.WriteFile(p.conf.Config, jsonString, 0755)
    return nil
}
func defaultProjector(conf *Config) *Projector {
    return &Projector{
        conf: conf,
        data:   Data{
            Projector: map[string]map[string]string{}, 
        },
    }
}

func NewProjector(conf *Config) *Projector{
    if _, err := os.Stat(conf.Config); err == nil{
        content, err := os.ReadFile(conf.Config)
        if err != nil {
            return defaultProjector(conf)
        }

        var data Data
        err = json.Unmarshal(content, &data)
        if err != nil {
            return defaultProjector(conf)
        }
        return &Projector {
            data: data,
            conf: conf,
        }            
    }
    return defaultProjector(conf)
}
