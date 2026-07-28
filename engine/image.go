package engine

import (
	"strings"
)

func BuildImageName(
	registry string,
	repository string,
	tag string,
) string {

	if registry == "" {
		return repositoryWithTag(repository, tag)
	}

	return registry + "/" + repositoryWithTag(repository, tag)
}

func repositoryWithTag(
	repository string,
	tag string,
) string {

	if tag == "" {
		return repository
	}

	return repository + ":" + tag
}

func flattenRepository(
	repository string,
) string {

	parts := strings.Split(repository, "/")

	return parts[len(parts)-1]
}
