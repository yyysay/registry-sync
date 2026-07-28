package model

type Image struct {
	Registry   string
	Repository string
	Tag        string
	Platform   []string
}
