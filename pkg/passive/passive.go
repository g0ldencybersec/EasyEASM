package passive

import (
	"sync"

	"github.com/g0ldencybersec/EasyEASM/pkg/passive/amass"
	"github.com/g0ldencybersec/EasyEASM/pkg/passive/subfinder"
)

type Runner struct {
	Seed_domain string
	Results     int
	Data        []string
}

func (r *Runner) Run() <-chan string {
	var wg sync.WaitGroup
	results := make(chan string)

	go func() {
		defer func() {
			close(results)
		}()
		wg.Add(2)

		go subfinder.RunSubfinder(r.Seed_domain, results, &wg)
		go amass.RunAmass(r.Seed_domain, results, &wg)

		wg.Wait()

	}()

	return results
}
