package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
)

type Info struct {
	Version   string  `json:"version"`
	GoVersion string  `json:"go_version"`
	Compiler  string  `json:"compiler"`
	Platform  string  `json:"platform"`
	VCS       VCSInfo `json:"vcs"`
}

type VCSInfo struct {
	Revision string `json:"revision"`
	Time     string `json:"time"`
	Modified bool   `json:"modified"`
}

func Get() Info {
	info, _ := debug.ReadBuildInfo()
	m := make(map[string]string, len(info.Settings))
	for _, s := range info.Settings {
		m[s.Key] = s.Value
	}
	return Info{
		Version:   info.Main.Version,
		GoVersion: info.GoVersion,
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", m["GOOS"], m["GOARCH"]),
		VCS: VCSInfo{
			Revision: m["vcs.revision"],
			Time:     m["vcs.time"],
			Modified: func() bool { r, _ := strconv.ParseBool(m["vcs.modified"]); return r }(),
		},
	}
}
