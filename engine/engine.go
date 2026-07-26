package engine

type Engine struct {
	copier Copier

	cache Cache

	resolver Resolver
}

func New(
	copier Copier,
	cache Cache,
	resolver Resolver,
) *Engine {

	return &Engine{

		copier: copier,

		cache: cache,

		resolver: resolver,
	}
}
