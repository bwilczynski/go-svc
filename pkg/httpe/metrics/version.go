package metrics

import (
	"strconv"

	"github.com/bwilczynski/go-svc/pkg/version"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	buildInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "version_info",
			Help: "A metric with a constant '1' value labeled by version information.",
		},
		[]string{"version", "go_version", "compiler", "platform", "vcs_revision", "vcs_time", "vcs_modified"},
	)
)

func init() {
	prometheus.MustRegister(buildInfo)
	info := version.Get()
	buildInfo.WithLabelValues(info.Version,
		info.GoVersion,
		info.Compiler,
		info.Platform,
		info.VCS.Revision,
		info.VCS.Time,
		strconv.FormatBool(info.VCS.Modified)).Set(1)
}
