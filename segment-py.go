package main

import (
	"os/exec"
	"strings"
)

func segmentPy(p *powerline) {
	out, err := exec.Command("python", "--version").CombinedOutput()
	if err != nil {
		return
	}

	p.appendSegment("py-version", segment{
		content:    strings.TrimSuffix(strings.Replace(string(out), "Python ", "", 1), "\n"),
		foreground: p.theme.PathFg,
		background: p.theme.PathBg,
	})
}
