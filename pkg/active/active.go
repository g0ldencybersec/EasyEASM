package active

import (
	"github.com/g0ldencybersec/EasyEASM/pkg/active/alterx"
	"github.com/g0ldencybersec/EasyEASM/pkg/active/dnsx"
	"github.com/g0ldencybersec/EasyEASM/pkg/active/httpx"
)

type ActiveRunner struct {
	SeedDomains []string
	Results     int
	Subdomains  []string
}

func (r *ActiveRunner) RunActiveEnum(wordlist string, threads int) []string {
	return dnsx.RunDnsx(r.SeedDomains, wordlist, threads)
}

func (r *ActiveRunner) RunPermutationScan(threads int) []string {
	return alterx.RunAlterx(r.Subdomains, threads)
}

func (r *ActiveRunner) RunHttpx() {
	httpx.RunHttpx(r.Subdomains)
}
