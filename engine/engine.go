package engine

type Engine struct {
	copier Copier

	cache Cache
}

func New(
	copier Copier,
	cache Cache,
) *Engine {

	return &Engine{
		copier: copier,
		cache:  cache,
	}
}
