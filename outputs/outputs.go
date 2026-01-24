// Package outputs handles creating markdown and csv files
package outputs

import (
	"fmt"
	"os"
	"path/filepath"
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

	timestamp := now.Format("2006-01-02T15:04:05-07:00")

	filenameCSV := fmt.Sprintf("profile_%s.csv", timestamp)

	err := createCSV(filenameCSV, &functionCalls)
	if err != nil {
		return err
	}
	if config.OutputMarkdown {

		err = createMD(timestamp, functionCalls)
		if err != nil {
			return err
		}
	}
	return nil
}

func createMD(timestamp string, data []*parse.FunctionCall) error {
	filename := fmt.Sprintf("profile_%s.md", timestamp)
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(cwd, "site", "content", "posts", filename)
	fmt.Println(path)

	fileMD, err := os.Create(path)
	if err != nil {
		return err
	}
	format := `+++
	title = 'My First Post'
	date = %s
	draft = true
+++
	`

	format = fmt.Sprintf(format, timestamp)
	_, err = fileMD.WriteString(format)
	if err != nil {
		return err
	}

	for i, fc := range data {
		function := fc.Function
		function = strings.ReplaceAll(function, "<", "\\<")
		function = strings.ReplaceAll(function, ">", "\\>")
		data[i].Function = function

	}
	table := tablewriter.NewTable(fileMD, tablewriter.WithRenderer(renderer.NewMarkdown()))
	table.Header([]string{"Filename", "LineNo", "Function", "Ncalls", "Tottime", "TottimePercall", "Cumtime", "CumtimePercall"})
	err = table.Bulk(data)
	if err != nil {
		return err
	}
	err = table.Render()
	if err != nil {
		return err
	}
	return nil
}

func createCSV(filename string, data *[]*parse.FunctionCall) error {
	fileCSV, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = gocsv.MarshalFile(data, fileCSV)
	if err != nil {
		return err
	}
	return nil
}
