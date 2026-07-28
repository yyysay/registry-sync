package model

type ResolvedImage struct {

	// 实际解析成功的源地址
	//
	// 例如:
	// ghcr.io/sagernet/sing-box:latest
	//
	Source string

	Registry string

	Repository string

	Tag string

	Digest string

	Platform []string
}
