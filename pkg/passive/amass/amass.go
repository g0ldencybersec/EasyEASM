package amass

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func RunAmass(seedDomain string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Running Amass on %s\n", seedDomain)
	cmd := exec.Command("amass", "enum", "--passive", "-nocolor", "-d", seedDomain)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
	cmd = exec.Command("amass", "db", "-names", "-d", seedDomain)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	for _, domain := range strings.Split(out.String(), "\n") {
		if strings.Contains(domain, seedDomain) && len(domain) != 0 {
			results <- domain
		}
	}
	fmt.Printf("Amass Run completed for %s\n", seedDomain)
}
