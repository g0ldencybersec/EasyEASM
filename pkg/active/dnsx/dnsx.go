package dnsx

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunDnsx(seedDomains []string) []string {
	fmt.Printf("Runing Bruteforce!")
	var results []string
	for _, domain := range seedDomains {
		if domain != "" {
			cmd := exec.Command("dnsx", "-d", domain, "-silent", "-w", "subdomains.txt", "-a", "-cname", "aaaa")

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
