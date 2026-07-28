package planner_test

import (
	"fmt"
	"os"
	"testing"

	"registry-sync/config"
	"registry-sync/planner"

	"gopkg.in/yaml.v3"
)

func TestBuild(t *testing.T) {

	data, _ := os.ReadFile("../ToAliyun.yaml")
	// data, _ := os.ReadFile("../test.yaml")

	var cfg config.Config

	yaml.Unmarshal(data, &cfg)

	plans := planner.Build(&cfg)

	fmt.Println(planner.Dump(plans))
}
