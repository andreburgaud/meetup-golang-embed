package main

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed config.json
var f embed.FS

type Config struct {
	Env string `json:"env"`
	Url string `json:"url"`
}

func main() {
	var configs []Config
	data, _ := f.ReadFile("config.json")
	json.Unmarshal(data, &configs)

	for _, c := range configs {
		fmt.Printf("Environment %4s uses URL %s\n", c.Env, c.Url)
	}
}
