package subfinder

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"sync"
)

func RunSubfinder(seedDomain string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	cmd := exec.Command("subfinder", "-d", seedDomain)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	for _, domain := range strings.Split(out.String(), "\n") {
		if strings.Contains(domain, seedDomain) {
			results <- domain
		}
	}
}
