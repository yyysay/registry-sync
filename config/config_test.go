package config_test

import (
	"fmt"
	"os"
	"testing"

	"registry-sync/config"

	"gopkg.in/yaml.v3"
)

func TestLoadConfig(t *testing.T) {

	data, err := os.ReadFile("../test.yaml")
	if err != nil {
		t.Fatal(err)
	}

	var cfg config.Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)
}
