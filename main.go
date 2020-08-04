package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vimicasa/go-api-demo-skeleton/app"
	"github.com/vimicasa/go-api-demo-skeleton/app/router"
	"github.com/vimicasa/go-api-demo-skeleton/config"
)

func main() {

	opts := config.ConfYaml{}

	var (
		showVersion bool
		configFile  string
	)

	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.StringVar(&configFile, "c", "", "Configuration file path.")
	flag.StringVar(&configFile, "config", "", "Configuration file path.")
	flag.StringVar(&opts.Core.Address, "A", "", "address to bind")
	flag.StringVar(&opts.Core.Address, "address", "", "address to bind")
	flag.StringVar(&opts.Core.Port, "p", "", "port number")
	flag.StringVar(&opts.Core.Port, "port", "", "port number")

	flag.Usage = usage
	flag.Parse()

	// Show version and exit
	app.SetVersion(Version)
	if showVersion {
		app.PrintVersion()
		os.Exit(0)
	}

	var err error
	// set default parameters.
	app.AppConf, err = config.LoadConf(configFile)
	if err != nil {
		log.Printf("Load yaml config file error: '%v'", err)

		return
	}

	// overwrite server port and address
	if opts.Core.Address != "" {
		app.AppConf.Core.Address = opts.Core.Address
	}
	if opts.Core.Port != "" {
		app.AppConf.Core.Port = opts.Core.Port
	}

	if err = app.InitLog(); err != nil {
		log.Fatalf("Can't load log module, error: %v", err)
	}

	// Routes.
	router.RunHTTPServer()

}

// Version control.
var Version = "0.0.1"

var usageStr = `

Usage: app [options]

Server Options:
    -A, --address <address>          Address to bind (default: any)
    -p, --port <port>                Use port for clients (default: 8088)
    -c, --config <file>              Configuration file path
Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
`

// usage will print out the flag options for the server.
func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}
