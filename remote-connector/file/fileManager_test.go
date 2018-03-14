// Package file
package file

import (
	"fmt"
	"os"
	"testing"
)

func TestReadConfigFileWithSuccess(t *testing.T) {
	got, err := ReadConfigFile("/home/blasting", ".remote_connections")
	if err != nil {
		panic(err)
	}
	if got == nil {
		t.Error("Got an nil configuration")
	}
	if len(got) == 0 {
		t.Error("Got an empty configuration")
	}
	// t.Logf("%v", got)
}

func TestReadConfigFileWithNoFileFound(t *testing.T) {
	var path string = os.Getenv("HOME")
	var fileName string = "doesnotexeites"

	expected := fmt.Sprintf("open %s/%s: no such file or directory", path, fileName)
	_, err := ReadConfigFile("/home/blasting", "doesnotexeites")
	if err.Error() != expected {
		t.Error(err.Error())
		t.Error("The config file should have been not found")
	}
}
