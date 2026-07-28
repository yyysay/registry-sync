package model

type Target struct {
	Name      string
	Registry  string
	Namespace string
	Flatten   bool
	Auth      *Auth
}
