package amass

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"sync"
)

func RunAmass(seedDomain string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	cmd := exec.Command("amass", "enum", "-passive", "-nocolor", "-d", seedDomain)
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
