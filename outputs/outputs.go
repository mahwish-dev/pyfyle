// Package outputs handles creating markdown and csv files
package outputs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"pyfyle/config"
	"pyfyle/parse"

	"github.com/gocarina/gocsv"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
)

func CreateOutputs(functionCalls []*parse.FunctionCall, config config.Config) error {
	now := time.Now()

	timestamp := now.Format("2006-01-02_15-04-05")

	filenameCSV := fmt.Sprintf("profile_%s.csv", timestamp)
	filenameMD := fmt.Sprintf("profile_%s.md", timestamp)
	fileCSV, err := os.Create(filenameCSV)
	if err != nil {
		return err
	}
	fileMD, err := os.Create(filenameMD)
	if err != nil {
		return err
	}
	err = gocsv.MarshalFile(&functionCalls, fileCSV)
	if err != nil {
		return err
	}

	for i, fc := range functionCalls {
		function := fc.Function
		function = strings.ReplaceAll(function, "<", "\\<")
		function = strings.ReplaceAll(function, ">", "\\>")
		functionCalls[i].Function = function

	}
	table := tablewriter.NewTable(fileMD, tablewriter.WithRenderer(renderer.NewMarkdown()))
	table.Header([]string{"Filename", "LineNo", "Function", "Ncalls", "Tottime", "TottimePercall", "Cumtime", "CumtimePercall"})
	err = table.Bulk(functionCalls)
	if err != nil {
		return err
	}
	err = table.Render()
	if err != nil {
		return err
	}
	return nil
}
