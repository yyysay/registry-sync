package model

type ResolvedImage struct {
	Registry string

	Repository string

	Tag string

	Digest string

	Platform []string
}
