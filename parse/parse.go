// Package parse handles parsing of cprofile output
package parse

import (
	"regexp"
	"strings"
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

func Parse(rawOutput string) ([]*FunctionCall, error) {
	lines := strings.Split(rawOutput, "\n")
	callsRe := regexp.MustCompile(`(\d+) function calls in (\d+\.\d+) seconds`)
	lineNoRe := regexp.MustCompile(`:(\d+)\(`)
	fileNameRe := regexp.MustCompile(`^([^:]+):`)
	functionNameRe := regexp.MustCompile(`\(([^)]+)\)$`)
	breakIndex := 0
	for i, line := range lines {
		matches := callsRe.FindStringSubmatch(line)
		if len(matches) > 0 {
			breakIndex = i
			break
		}

	}
	lines = lines[breakIndex+5:]
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			breakIndex = i
			break
		}
	}
	lines = lines[:breakIndex]
	functionCalls := []*FunctionCall{}

	for _, line := range lines {
		lineNum := "~"
		fileName := "~"
		var functionName string
		fc := FunctionCall{}
		line = strings.TrimSpace(line)
		vals := strings.SplitN(line, "    ", 6)
		lastTwo := strings.SplitN(strings.TrimSpace(vals[len(vals)-1]), " ", 2)

		vals = append(vals[:len(vals)-1], lastTwo...)
		lastVal := vals[5]
		matches := lineNoRe.FindStringSubmatch(lastVal)
		if len(matches) >= 2 {
			lineNum = matches[1]
		}
		matches = fileNameRe.FindStringSubmatch(lastVal)
		if len(matches) > 1 {
			fileName = matches[1]
			fileNameMatches := functionNameRe.FindStringSubmatch(lastVal)
			functionName = fileNameMatches[1]

		} else {
			functionName = lastVal
		}
		fc.Ncalls = vals[0]
		fc.Tottime = vals[1]
		fc.TottimePercall = vals[2]
		fc.Cumtime = vals[3]
		fc.CumtimePercall = vals[4]
		fc.Filename = fileName
		fc.LineNo = lineNum
		fc.Function = functionName

		functionCalls = append(functionCalls, &fc)

	}

	return functionCalls, nil
}
