// Package serverinfo provides functions to get remote server info
package serverinfo

import (
	"fmt"
	"io/ioutil"
	"net"
)

// GetHeadInfo to get head info from server. service has to be of the form "host:port"
func GetHeadInfo(service string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		return err
	}
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return err
	}
	fmt.Println(string(result))

	return nil
}
