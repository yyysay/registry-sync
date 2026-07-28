package engine

type Engine struct {
	copier Copier

	resolver Resolver

	metadataResolver MetadataResolver
}

func New(
	copier Copier,
	resolver Resolver,
	metadataResolver MetadataResolver,
) *Engine {

	return &Engine{

		copier: copier,

		resolver: resolver,

		metadataResolver: metadataResolver,
	}
}
