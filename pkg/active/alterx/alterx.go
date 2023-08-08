package alterx

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func RunAlterx(domains []string, threads int) []string {
	fmt.Println("Starting permuatation scan!")
	var results []string
	createDomainFile(domains)

	cmd := exec.Command("alterx", "-l", "tempDomains.txt", "-silent", "-o", "alterxDomains.txt")
	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	cmd = exec.Command("dnsx", "-l", "alterxDomains.txt", "-silent", "-a", "-cname", "-aaaa", "-t", strconv.Itoa(threads))
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		panic(err)
	}

	for _, domain := range strings.Split(out.String(), "\n") {
		if len(domain) != 0 {
			results = append(results, domain)
		}
	}
	fmt.Println("ALTERX RESULTS")
	fmt.Println(results)
	os.Remove("tempDomains.txt")
	os.Remove("alterxDomains.txt")

	return results
}

func createDomainFile(domains []string) {
	file, err := os.OpenFile("tempDomains.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range domains {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}
