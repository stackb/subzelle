package subzelle

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/proto"

	pb "github.com/stackb/subzelle/proto"
)

var ConfigurationBasename = ".subzelle.config"

var active pb.Configuration

func GetActiveConfiguration() *pb.Configuration {
	return &active
}

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not get working directory: %v", err)
		os.Exit(1)
	}
	config, err := readConfiguration(cwd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read configuration file: %v", err)
		os.Exit(1)
	}
	active = *config
}

func readConfiguration(dirname string) (*pb.Configuration, error) {
	// looking for `${workspaceRoot}/.subzelle.prototext`
	filename := filepath.Join(dirname, ConfigurationBasename)

	// read the file as a string
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		d := filepath.Dir(dirname)
		if d == dirname {
			return nil, err
		}
		return readConfiguration(d)
	}

	// unmarshal proto and finish
	var config pb.Configuration
	if err := proto.UnmarshalText(string(in), &config); err != nil {
		return nil, err
	}

	return &config, nil
}
