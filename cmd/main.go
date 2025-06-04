package main

import (
	"fmt"
	"os"

	"github.com/lxr0x64/voltx/pkg/probe"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: voltx <url>")
		os.Exit(1)
	}
	url := os.Args[1]

	status, err := probe.HTTPProbe(url)
	if err != nil {
		fmt.Printf("Erro ao fazer probe: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Status da URL %s: %d\n", url, status)
}
