package lookup

import (
	"fmt"
	"net"
)

// LookupHost to lookup DNS host
func LookupHost(hostname string) error {
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		return err
	}

	// print addresses
	for _, s := range addrs {
		fmt.Println(s)
	}

	return nil
}

// LookupCname to lookup DNS common name
func LookupCname(hostname string) error {
	cname, err := net.LookupCNAME(hostname)
	if err != nil {
		return err
	}
	fmt.Println(cname)

	return nil
}

func LookupPort(netType, service string) error {
	port, err := net.LookupPort(netType, service)
	if err != nil {
		return err
	}
	fmt.Println("Service port:", port)

	return nil
}
