package subfinder

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func RunSubfinder(seedDomain string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("    => Running subfinder on %s\n", seedDomain)
	cmd := exec.Command("subfinder", "-d", seedDomain, "-silent")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	for _, domain := range strings.Split(out.String(), "\n") {
		if strings.Contains(domain, seedDomain) && len(domain) != 0 {
			results <- domain
		}
	}
	// fmt.Printf("  => subfinder run completed for %s\n", seedDomain)
}
