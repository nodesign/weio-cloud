/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/mainflux/mainflux-core/api"
	"github.com/mainflux/mainflux-core/config"
	"github.com/mainflux/mainflux-core/db"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var usageStr = `
Usage: mainflux [options]
Server Options:
    -a, --addr <host>                Bind to host address (default: 0.0.0.0)
    -p, --port <port>                Use port for clients (default: 4222)
    -P, --pid <file>                 File to store PID
    -c, --config <file>              Configuration file
Logging Options:
    -l, --log <file>                 File to redirect log output
    -T, --logtime                    Timestamp log entries (default: true)
    -D, --debug                      Enable debugging output
    -V, --trace                      Trace the raw protocol
    -DV                              Debug and trace
Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
`

// usage will print out the flag options for the server.
func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

// PrintServerAndExit will print our version and exit.
func PrintServerAndExit() {
	fmt.Printf("Mainflux version %s\n", Version)
	os.Exit(0)
}

func main() {
	// Server Options
	opts := Options{}

	var showVersion bool
	var debugAndTrace bool
	var configFile string

	// Parse flags
	flag.IntVar(&opts.Port, "port", 0, "Port to listen on.")
	flag.IntVar(&opts.Port, "p", 0, "Port to listen on.")
	flag.StringVar(&opts.Host, "host", "", "Network host to listen on.")
	flag.StringVar(&opts.Host, "h", "", "Network host to listen on.")
	flag.StringVar(&opts.Host, "net", "", "Network host to listen on.")
	flag.BoolVar(&opts.Debug, "D", false, "Enable Debug logging.")
	flag.BoolVar(&opts.Debug, "debug", false, "Enable Debug logging.")
	flag.BoolVar(&opts.Trace, "V", false, "Enable Trace logging.")
	flag.BoolVar(&opts.Trace, "trace", false, "Enable Trace logging.")
	flag.BoolVar(&debugAndTrace, "DV", false, "Enable Debug and Trace logging.")
	flag.BoolVar(&opts.Logtime, "T", true, "Timestamp log entries.")
	flag.BoolVar(&opts.Logtime, "logtime", true, "Timestamp log entries.")
	flag.StringVar(&opts.Username, "user", "", "Username required for connection.")
	flag.StringVar(&opts.Password, "pass", "", "Password required for connection.")
	flag.StringVar(&opts.Authorization, "auth", "", "Authorization token required for connection.")
	flag.StringVar(&configFile, "c", "", "Configuration file.")
	flag.StringVar(&configFile, "config", "", "Configuration file.")
	flag.StringVar(&opts.PidFile, "P", "", "File to store process pid.")
	flag.StringVar(&opts.PidFile, "pid", "", "File to store process pid.")
	flag.StringVar(&opts.LogFile, "l", "", "File to store logging output.")
	flag.StringVar(&opts.LogFile, "log", "", "File to store logging output.")
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")

	flag.Usage = usage

	flag.Parse()

	// Show version and exit
	if showVersion {
		PrintServerAndExit()
	}

	// One flag can set multiple options.
	if debugAndTrace {
		opts.Trace, opts.Debug = true, true
	}

	// Process args looking for non-flag options,
	// 'version' and 'help' only for now
	for _, arg := range flag.Args() {
		switch strings.ToLower(arg) {
		case "version":
			PrintServerAndExit()
		case "help":
			usage()
		}
	}

	// Parse config
	var cfg config.Config
	cfg.Parse()

	// MongoDb
	db.InitMongo(cfg.MongoHost, cfg.MongoPort, cfg.MongoDatabase)

	// Serve HTTP
	go api.HTTPServer(cfg)

	// Print banner
	color.Cyan(banner)
	color.Cyan("Magic happens on port " + strconv.Itoa(cfg.HTTPPort))

	/** Keep main() runnig */
	runtime.Goexit()
}

var banner = `
8   8  8      8  8"""88    8""""8                         
8   8  8 eeee 8  8    8    8    " e     eeeee e   e eeeee 
8e  8  8 8    8e 8    8    8e     8     8  88 8   8 8   8 
88  8  8 8eee 88 8    8    88     8e    8   8 8e  8 8e  8 
88  8  8 88   88 8    8    88   e 88    8   8 88  8 88  8 
88ee8ee8 88ee 88 8eeee8    88eee8 88eee 8eee8 88ee8 88ee8 

			   == WeIO goes high ==
       
    Made with <3 by Drasko DRASKOVIC and Uros PETREVSKI
[w] http://weio.net
[t] @weionoet

`
