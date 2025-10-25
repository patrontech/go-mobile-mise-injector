package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Tools map[string]string `json:"tools" toml:"tools"`
}

func main() {
	var required = map[string]string{
		"jq":         "1",
		"xmlstarlet": "1",
		"turbopump":  "1",
	}
	if raw, ok := os.LookupEnv("MISE_INJECTOR_REQUIRED"); ok {
		err := json.Unmarshal([]byte(raw), &required)
		if err != nil {
			die(err)
		}
	}
	var path string
	if len(os.Args) < 2 {
		path = ".mise.toml"
	} else {
		path = os.Args[1]
	}
	raw, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		die(err)
	}
	var cfg Config
	if err == nil {
		// Exists
		err = toml.Unmarshal(raw, &cfg)
		if err != nil {
			die(err)
		}
		for key, value := range required {
			if _, ok := cfg.Tools[key]; !ok {
				cfg.Tools[key] = value
			}
		}
	} else if os.IsNotExist(err) {
		// Not Exists
		cfg = Config{Tools: required}
	} else {
		die(err)
	}
	raw, err = toml.Marshal(cfg)
	if err != nil {
		die(err)
	}
	err = os.WriteFile(path, raw, 0644)
	if err != nil {
		die(err)
	}
	fmt.Println(cfg)
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
