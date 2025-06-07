package main

import (
	"fmt"

	"github.com/dev-vamsi/taskmate/internal/config"
)

func main() {
	cfg := &config.Config{}
	err := cfg.Load("./config/default.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.Database.URL)
}
