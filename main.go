package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
	"github.com/ztrue/tracerr"
)

const version = "1.1.0"

var opts struct {
	Verbose  bool   `short:"v" long:"verbose" description:"Print verbose output"`
	Version  bool   `long:"version" description:"Print version and exit"`
	Color    bool   `short:"c" long:"color" description:"Color output"`
	Username string `short:"u" long:"username" description:"Username of duinocoin user"`
}

func main() {
	if _, err := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash).Parse(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if opts.Version {
		if opts.Color {
			fmt.Fprintln(color.Output, y("ducolookup"), w("version"), version)
			fmt.Fprintln(color.Output, c("Go"), w("version"), runtime.Version())
		} else {
			fmt.Println("ducolookup version", version)
			fmt.Println("Go version", runtime.Version())
		}

		os.Exit(0)
	}

	if err := printUserInfo(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}
}
