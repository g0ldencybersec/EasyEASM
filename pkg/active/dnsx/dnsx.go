package dnsx

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func RunDnsx(seedDomains []string, wordlist string, threads int) []string {
	fmt.Printf("Runing Bruteforce!")
	var results []string
	for _, domain := range seedDomains {
		if domain != "" {
			cmd := exec.Command("dnsx", "-d", domain, "-silent", "-w", wordlist, "-a", "-cname", "-aaaa", "-t", strconv.Itoa(threads))

			var out bytes.Buffer
			cmd.Stdout = &out

			err := cmd.Run()

			if err != nil {
				panic(err)
			}

			for _, domain := range strings.Split(out.String(), "\n") {
				if len(domain) != 0 {
					results = append(results, domain)
				}
			}
		}

	}

	fmt.Println("ACTIVE RESULTS")
	fmt.Println(results)
	os.Remove("tempDomains.txt")
	return results
}
