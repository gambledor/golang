// Package main provides networking tools
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gambledor/golang/net/lookup"
	"github.com/gambledor/golang/net/serverinfo"
)

var (
	command  *string
	hostname *string
	netType  *string
	service  *string
)

func init() {
	command = flag.String("cmd", "lu", "The command to execute.\n\tAd exapmle: lu (lookup), lucn (lookup CNAME), luport (lookup port)")
	hostname = flag.String("h", "", "The hostname")
	netType = flag.String("nt", "", "The network type (tcp, udp)")
	service = flag.String("s", "", "The netword service (telnet, imap, ect.)")
}

func main() {
	flag.Parse()

	var err error

	switch *command {
	case "lu":
		if hostname != nil {
			err = lookup.LookupHost(*hostname)
		} else {
			fmt.Fprintf(os.Stderr, "Usage: %s lookup <hostname>\n", os.Args[0])
		}
	case "lucn":
		if hostname != nil {
			err = lookup.LookupCname(*hostname)
		} else {
			fmt.Fprintf(os.Stderr, "Usage: %s lookupCname <hostname>\n", os.Args[0])
		}
	case "luport":
		if netType != nil && service != nil {
			err = lookup.LookupPort(*netType, *service)
		} else {
			fmt.Fprintf(os.Stderr, "Usage: %s luport <networkType> <service>\n", os.Args[0])
		}
	case "headinfo":
		if service != nil {
			err = serverinfo.GetHeadInfo(*service)
		} else {
			fmt.Fprintf(os.Stderr, "Usage: %s headinfo <service>\n", os.Args[0])
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s", *command)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
