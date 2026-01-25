// Package parse handles parsing of cprofile output
package parse

import (
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
	breakIndex := 0
	prLineIndex := 0
	for i, line := range lines {
		if strings.Contains(line, "seconds") {
			prLineIndex = i
		}
		if strings.Contains(line, "ncalls") {
			breakIndex = i + 1
			break
		}
	}
	prLine := lines[prLineIndex]
	lines = lines[breakIndex:]

	functionCalls := []*FunctionCall{}

	for _, line := range lines {
		if line == "" {
			break
		}
		fc := FunctionCall{}
		line = strings.TrimSpace(line)
		vals := strings.Fields(line)
		fc.Ncalls = vals[0]
		fc.Tottime = vals[1]
		fc.TottimePercall = vals[2]
		fc.Cumtime = vals[3]
		fc.CumtimePercall = vals[4]

		remaining := vals[5:]
		lastVal := strings.Join(remaining, " ")

		flf := parseLastColumn(lastVal)
		fc.Filename = flf.Filename
		fc.LineNo = flf.LineNo
		fc.Function = flf.Function

		functionCalls = append(functionCalls, &fc)

	}
	log.Info("Parsed Command")

	return functionCalls, ProfileRun(prLine), nil
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
