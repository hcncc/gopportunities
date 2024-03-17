package main

import (
	"fmt"

	"github.com/hcncc/gopportunities/config"
	"github.com/hcncc/gopportunities/router"
)

func main() {
	// Initialize Router in main file
	err := config.Init()

	if err != nil {
		fmt.Println("Error in initialize configurations", err)
		return
	}

	// Initialize Router in main file
	router.Initialize()
}
