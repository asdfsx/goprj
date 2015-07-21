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

	t := map[string]string{"test":"test"}

	if val, ok := t["1"]; ok{
		fmt.Println(val, ok)
	} else{
		fmt.Println(val, ok)
	}

	a, b := getTest()
	fmt.Println(a, b)
}

func getTest() (string, int){
return "test",0
}
