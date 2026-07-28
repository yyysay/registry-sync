package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type ImageConfig struct {
	Tags     []string `yaml:"tags,omitempty"`
	Platform []string `yaml:"platform,omitempty"`
}

func (i *ImageConfig) UnmarshalYAML(value *yaml.Node) error {

	switch value.Kind {

	case yaml.ScalarNode:

		// 简写:
		// image: latest
		// image: main
		if value.Value == "" {
			return nil
		}

		i.Tags = []string{value.Value}

		return nil

	case yaml.SequenceNode:

		// 简写:
		// image:
		//   - latest
		//   - latest-testing

		for _, node := range value.Content {
			i.Tags = append(i.Tags, node.Value)
		}

		return nil

	case yaml.MappingNode:

		// 完整:
		// image:
		//   tags:
		//     - latest
		//     - latest-testing
		//   platform:
		//     - linux/amd64

		var data struct {
			Tags     []string `yaml:"tags"`
			Platform []string `yaml:"platform"`
		}

		if err := value.Decode(&data); err != nil {
			return err
		}

		i.Tags = data.Tags
		i.Platform = data.Platform

		return nil

	case yaml.DocumentNode:

		if len(value.Content) == 0 {
			return nil
		}

		return i.UnmarshalYAML(value.Content[0])

	default:

		// 空:
		// image:
		//
		// 保留 nil，由 planner 决定默认 tag=latest

		if value.Tag == "!!null" {
			return nil
		}

		return fmt.Errorf(
			"unsupported image format: kind=%v tag=%s",
			value.Kind,
			value.Tag,
		)
	}
}
