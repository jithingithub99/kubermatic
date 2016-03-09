package provider

import (
	"bufio"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// DigitaloceanSpec describes a digital ocean datacenter
type DigitaloceanSpec struct {
	Region string `yaml:"region"`

	Seed struct {
		SSHKeys []string `yaml:"sshKeys"`
	} `yaml:"seed"`
}

// BringYourOwnSpec describes a datacenter our of bring your own nodes
type BringYourOwnSpec struct {
	Seed struct {
		PrivateIntf string `yaml:"privateInterface"`
	} `yaml:"seed"`
}

// DatacenterSpec describes mutually points to provider datacenter spec
type DatacenterSpec struct {
	Digitalocean *DigitaloceanSpec `yaml:"digitalocean"`
	BringYourOwn *BringYourOwnSpec `yaml:"bringyourown"`
}

// DatacenterMeta describes a Kubermatic datacenter.
type DatacenterMeta struct {
	Location string         `yaml:"location"`
	Country  string         `yaml:"country"`
	Spec     DatacenterSpec `yaml:"spec"`
}

// datacentersMeta describes a number of Kubermatic datacenters.
type datacentersMeta struct {
	Datacenters map[string]DatacenterMeta `yaml:"datacenters"`
}

// DatacentersMeta loads datacenter metadata from the given path.
func DatacentersMeta(path string) (map[string]DatacenterMeta, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(bufio.NewReader(f))
	if err != nil {
		return nil, err
	}

	dcs := datacentersMeta{}
	err = yaml.Unmarshal(bytes, &dcs)
	if err != nil {
		return nil, err
	}

	return dcs.Datacenters, nil
}
