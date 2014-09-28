package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type TablePrinter interface {
	Print(value string)
	BreakField()
	BreakRecord()
}

type FmtPrinter struct {
	writer     io.Writer
	separatoer string
}

func (pr FmtPrinter) Print(value string) {
	fmt.Fprint(pr.writer, value)
}

func (pr FmtPrinter) BreakField() {
	fmt.Fprint(pr.writer, pr.separatoer)
}

func (pr FmtPrinter) BreakRecord() {
	fmt.Fprintln(pr.writer)
}

func main() {
	var data []map[string]interface{}
	dec := json.NewDecoder(os.Stdin)
	dec.Decode(&data)

	printer := FmtPrinter{os.Stdout, "\t"}
	printTable(data, printer, true)
}

func printTable(data []map[string]interface{}, tp TablePrinter, header bool) {
	if data == nil || len(data) == 0 {
		return
	}
	// Use the keys of the first element as the header.
	head := data[0]
	keys := make([]string, 0)
	for k, _ := range head {
		keys = append(keys, k)
	}
	if header {
		for i, k := range keys {
			tp.Print(k)
			if i < len(keys)-1 {
				tp.BreakField()
			}
		}
		tp.BreakRecord()
	}
	for i, row := range data {
		for j, k := range keys {
			val, exists := row[k]
			if exists {
				tp.Print(fmt.Sprintf("%v", val))
			}
			if j < len(keys)-1 {
				tp.BreakField()
			}
		}

		if i < len(data)-1 {
			tp.BreakRecord()
		}
	}
}
