package amass

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/g0ldencybersec/EasyEASM/pkg/utils"
)

func RunAmass(seedDomain string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Running Amass on %s\n", seedDomain)
	binPath := utils.GetGoEnvPathBin()
	cmd := exec.Command(binPath+"/amass", "enum", "--passive", "-nocolor", "-d", seedDomain)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
	cmd = exec.Command(binPath+"/oam_subs", "-names", "-d", seedDomain)
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
