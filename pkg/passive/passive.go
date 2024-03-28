package passive

import (
	"fmt"
	"sync"

	"github.com/sethlaw/EasyEASM/pkg/passive/amass"
	"github.com/sethlaw/EasyEASM/pkg/passive/httpx"
	"github.com/sethlaw/EasyEASM/pkg/passive/subfinder"
)

type PassiveRunner struct {
	SeedDomains []string
	Results     int
	Subdomains  []string
}

func (r *PassiveRunner) RunPassiveEnum() []string {
	fmt.Println("  => Running Passive Sources")
	var wg sync.WaitGroup
	sf_results := make(chan string)
	amass_results := make(chan string)
	for _, domain := range r.SeedDomains {
		wg.Add(2)
		go subfinder.RunSubfinder(domain, sf_results, &wg)
		go amass.RunAmass(domain, amass_results, &wg)
	}

	var results []string
	done := make(chan bool)

	go func() {
		for msg := range sf_results {
			results = append(results, msg)
		}
		done <- true
	}()

	go func() {
		for msg := range amass_results {
			results = append(results, msg)
		}
		done <- true
	}()

	go func() {
		wg.Wait()
		close(sf_results)
		close(amass_results)
	}()

	<-done
	<-done
	return results
}

func (r *PassiveRunner) RunHttpx() {
	httpx.RunHttpx(r.Subdomains)
}
