package main

import (
	"os/exec"
	"strings"
)

func segmentAWSAccount(p *powerline) {
	account, err := exec.Command("aws", "sts", "get-caller-identity", "--output", "json", "--query", "Account").CombinedOutput()
	if err != nil {
		return
	}
	if len(account) > 0 {
		p.appendSegment("aws-account", segment{
			content:    strings.TrimSuffix(strings.Replace(string(account), "\"", "", -1), "\n"),
			foreground: p.theme.AWSFg,
			background: p.theme.AWSBg,
		})
	}
}
