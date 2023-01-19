package version

import (
	"bytes"
	"runtime"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
)

var (
	Version    = "1.0"
	Revision   string
	Branch     string
	BuildTiime string
	BuildUser  string
	GoVersion  = runtime.Version()
)

func String() string {
	wr := new(bytes.Buffer)
	twr := tabwriter.NewWriter(wr, 0, 0, 2, ' ', 0)
	tb := tabby.NewCustom(twr)
	tb.AddLine("Version:", Version)
	tb.AddLine("Revision:", Revision)
	tb.AddLine("Branch:", Branch)
	tb.AddLine("BuildTiime:", BuildTiime)
	tb.AddLine("BuildUser:", BuildUser)
	tb.AddLine("GoVersion:", GoVersion)
	twr.Flush()
	return wr.String()
}
