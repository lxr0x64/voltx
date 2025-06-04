package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/lxr0x64/voltx/pkg/probe"
	// Ajusta conforme seu go.mod
)

func banner() {
	fmt.Println(`
__      __       _   __  __
\ \    / /__ _ _| |_|  \/  | __ _ _ __   ___
 \ \/\/ / _ \ '_|  _| |\/| |/ _` + "`" + ` | '  \ (_-<
  \_/\_/\___/_|  \__|_|  |_|\__,_|_|_|_/__/
               V O L T X

projectdiscovery.io

Use with caution. You are responsible for your actions.
Developers assume no liability and are not responsible for any misuse or damage.
	`)
}

func main() {
	banner()

	if len(os.Args) < 2 {
		fmt.Println("Uso: voltx <url1> <url2> ...")
		os.Exit(1)
	}

	urls := os.Args[1:]
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			status, err := probe.HTTPProbe(u)
			if err != nil {
				fmt.Printf("[ERRO] %s: %v\n", u, err)
				return
			}

			fmt.Printf("[OK] %s: %d\n", u, status)
		}(url)
	}

	wg.Wait()
}
