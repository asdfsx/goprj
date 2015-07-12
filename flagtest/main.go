package main

import (
	"flagtest/flagtest"
	"fmt"
	"os"
)

func main() {
	fmt.Println("flag test")
	cfg := flagtest.NewConfig()
	err := cfg.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", cfg)
}
