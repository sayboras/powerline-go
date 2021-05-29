package main

import (
	"os"
	"os/exec"
	"strings"
	"time"

	pwl "github.com/justjanne/powerline-go/powerline"
)

const timeFormat = "2006-01-02T15:04:05Z07:00"

func segmentAWS(p *powerline) []pwl.Segment {
	profile := os.Getenv("AWS_PROFILE")
	region := os.Getenv("AWS_DEFAULT_REGION")
	if profile == "" || isExpired(profile) {
		return []pwl.Segment{}
	}
	var r string
	if region != "" {
		r = " (" + region + ")"
	} else {
		if region, err := getAWSKey(profile, "region"); err == nil {
			r = " (" + region + ")"
		}
	}

	return []pwl.Segment{{
		Name:       "aws",
		Content:    profile + r,
		Foreground: p.theme.AWSFg,
		Background: p.theme.AWSBg,
	}}
}

func isExpired(profile string) bool {
	command := exec.Command("aws", "configure", "get", "x_security_token_expires", "--profile", profile)
	output, err := command.Output()
	if err != nil {
		return false
	}
	if len(output) == 0 {
		return false
	}
	expiredTimeStr, err := getAWSKey(profile, "x_security_token_expires")
	if err != nil {
		return false
	}
	expiredTime, err := time.Parse(timeFormat, expiredTimeStr)
	if err != nil {
		return false
	}
	return expiredTime.Before(time.Now())
}

func getAWSKey(profile string, key string) (string, error) {
	command := exec.Command("aws", "configure", "get", key, "--profile", profile)
	output, err := command.Output()
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(string(output), "\n", ""), nil
}