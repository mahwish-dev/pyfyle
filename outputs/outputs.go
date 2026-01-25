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

	"github.com/charmbracelet/log"
	"github.com/gocarina/gocsv"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
)

func CreateOutputs(functionCalls []*parse.FunctionCall, pr parse.ProfileRun, config config.Config) (string, error) {
	now := time.Now()

	timestamp := now.Format("2006-01-02T15:04:05-07:00")
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	cwd = filepath.Join(cwd, "pyfyle")

	filenameCSV := fmt.Sprintf("profile_%s.csv", timestamp)

	err = createCSV(filenameCSV, cwd, &functionCalls)
	if err != nil {
		return "", err
	}
	log.Info("Created CSV")
	if config.OutputMarkdown && config.DashboardEnabled {

		err = createMD(timestamp, cwd, functionCalls, pr)
		if err != nil {
			return "", err
		}
		log.Info("Created MD")
	}
	log.Info("Created outputs")
	return filenameCSV, nil
}

func createMD(timestamp string, cwd string, data []*parse.FunctionCall, pr parse.ProfileRun) error {
	filename := fmt.Sprintf("profile_%s.md", timestamp)
	mdDir := filepath.Join(cwd, "site", "content", "posts")
	err := os.MkdirAll(mdDir, 0o755)
	if err != nil {
		return err
	}
	log.Info("Created MD Directory")
	path := filepath.Join(mdDir, filename)

	fileMD, err := os.Create(path)
	if err != nil {
		return err
	}
	log.Info("Created MD File")
	format := `+++
	title = 'Profile Run On : %s'
	date = %s
	draft = true
+++

%s

`

	format = fmt.Sprintf(format, timestamp, timestamp, pr)
	_, err = fileMD.WriteString(format)
	if err != nil {
		return err
	}
	log.Info("Wrote title and timestamp")

	for i, fc := range data {
		function := fc.Function
		function = strings.ReplaceAll(function, "<", "\\<")
		function = strings.ReplaceAll(function, "_", "\\_")
		function = strings.ReplaceAll(function, ">", "\\>")
		filename := fc.Filename
		filename = strings.ReplaceAll(filename, "<", "\\<")
		filename = strings.ReplaceAll(filename, ">", "\\>")
		data[i].Function = function
		data[i].Filename = filename

	}
	table := tablewriter.NewTable(fileMD, tablewriter.WithRenderer(renderer.NewMarkdown()))
	table.Header([]string{"Filename", "LineNo", "Function", "Ncalls", "Tottime", "TottimePercall", "Cumtime", "CumtimePercall"})
	err = table.Bulk(data)
	if err != nil {
		return err
	}
	log.Info("Marshalled Data")
	err = table.Render()
	if err != nil {
		return err
	}
	log.Info("Wrote table")
	return nil
}

func createCSV(filename string, cwd string, data *[]*parse.FunctionCall) error {
	csvDir := filepath.Join(cwd, "csv")
	err := os.MkdirAll(csvDir, 0o755)
	if err != nil {
		return err
	}
	log.Info("Created CSV Directory")
	path := filepath.Join(csvDir, filename)
	fileCSV, err := os.Create(path)
	if err != nil {
		return err
	}
	log.Info("Created CSV File")
	err = gocsv.MarshalFile(data, fileCSV)
	if err != nil {
		return err
	}
	log.Info("Wrote CSV File")
	return nil
}
