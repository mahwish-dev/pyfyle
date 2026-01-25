// Package parse handles parsing of cprofile output
package parse

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
)

type FunctionCall struct {
	Filename       string `csv:"filename"`
	LineNo         string `csv:"line"`
	Function       string `csv:"function"`
	Ncalls         string `csv:"ncalls"`
	Tottime        string `csv:"tottime"`
	TottimePercall string `csv:"tottime_percall"`
	Cumtime        string `csv:"cumtime"`
	CumtimePercall string `csv:"cumtime_percall"`
}

type filenameLineNoFunc struct {
	Filename string
	LineNo   string
	Function string
}

type ProfileRun string

func Parse(rawOutput string) ([]*FunctionCall, ProfileRun, error) {
	lines := strings.Split(rawOutput, "\n")
	callsRe := regexp.MustCompile(`(\d+) function calls in (\d+\.\d+) seconds`)
	breakIndex := 0
	var pr ProfileRun
	for i, line := range lines {
		matches := callsRe.FindStringSubmatch(line)
		if len(matches) > 0 {
			breakIndex = i
			pr = ProfileRun(matches[0])

			break
		}

	}
	startIndex := 0
	for i, line := range lines {
		if strings.Contains(line, "ncalls") {
			startIndex = i
			break
		}
	}

	lines = lines[startIndex+1:]
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			breakIndex = i
			break
		}
	}
	lines = lines[:breakIndex]
	functionCalls := []*FunctionCall{}

	for _, line := range lines {
		fc := FunctionCall{}
		line = strings.TrimSpace(line)
		vals := strings.SplitN(line, "    ", 6)
		lastTwo := strings.SplitN(strings.TrimSpace(vals[len(vals)-1]), " ", 2)

		vals = append(vals[:len(vals)-1], lastTwo...)
		lastVal := vals[5]
		flf := parseLastColumn(lastVal)
		fc.Ncalls = vals[0]
		fc.Tottime = vals[1]
		fc.TottimePercall = vals[2]
		fc.Cumtime = vals[3]
		fc.CumtimePercall = vals[4]
		fc.Filename = flf.Filename
		fc.LineNo = flf.LineNo
		fc.Function = flf.Function

		functionCalls = append(functionCalls, &fc)

	}
	log.Info("Parsed Command")

	return functionCalls, pr, nil
}

func parseLastColumn(val string) filenameLineNoFunc {
	out := filenameLineNoFunc{Filename: "~", LineNo: "~", Function: "~"}
	if val[0] == '{' {
		out.Function = val
	} else {
		filnameIndex := strings.IndexRune(val, ':')
		filename := val[:filnameIndex]
		out.Filename = filename
		remaining := val[filnameIndex+1:]
		lineNoIndex := strings.IndexRune(remaining, '(')
		out.LineNo = remaining[:lineNoIndex]
		remaining = remaining[lineNoIndex+1:]
		remaining = remaining[:len(remaining)-1]
		out.Function = remaining
	}

	return out
}
