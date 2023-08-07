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

func (r *ActiveRunner) RunActiveEnum() []string {
	return dnsx.RunDnsx(r.SeedDomains)
}

func (r *ActiveRunner) RunPermutationScan() []string {
	return alterx.RunAlterx(r.Subdomains)
}

func (r *ActiveRunner) RunHttpx() {
	httpx.RunHttpx(r.Subdomains)
}
