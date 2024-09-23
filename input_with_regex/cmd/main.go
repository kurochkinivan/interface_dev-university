package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kurochkinivan/input_with_regex/internal/constants"
	outputfuncs "github.com/kurochkinivan/input_with_regex/internal/output_funcs"
	"github.com/kurochkinivan/input_with_regex/internal/parser"
	flag "github.com/spf13/pflag"
)

var (
	path, oper  string
	showVersion bool
	f           func(ps []parser.Point, ls []parser.Line, cs []parser.Circle)
)

func main() {
	err := parser.ParseObjectsFromFile(path, f)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	flag.StringVarP(&path, "file", "f", "", "name of the destination file")
	flag.StringVarP(&oper, "operation", "o", "", "operation with file: print/count")
	flag.BoolVarP(&showVersion, "version", "v", false, "print project version")
	flag.Parse()

	if showVersion {
		fmt.Printf("current version is: %s\n", constants.Version)
		return
	}

	if path == "" {
		fmt.Fprintln(os.Stderr, "no path to file provided")
		flag.Usage()
		return
	}

	switch oper {
	case "count":
		f = outputfuncs.Count
	case "print":
		f = outputfuncs.Print
	default:
		fmt.Fprintln(os.Stderr, "wrong operation flag")
		flag.Usage()
		return
	}
}