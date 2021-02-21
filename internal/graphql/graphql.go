package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/layer5io/meshery/internal/graphql/generated"
	"github.com/layer5io/meshery/internal/graphql/resolver"
	"github.com/layer5io/meshkit/database"
	mesherykube "github.com/layer5io/meshkit/utils/kubernetes"
)

type Options struct {
	DBHandler  *database.Handler
	KubeClient *mesherykube.Client
	URL        string
}

// New returns a graphql handler instance
func New(opts Options) http.Handler {
	resolver := &resolver.Resolver{
		DBHandler:  opts.DBHandler,
		KubeClient: opts.KubeClient,
	}

	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))
}

// NewPlayground returns a graphql playground instance
func NewPlayground(opts Options) http.Handler {
	return playground.Handler("GraphQL playground", opts.URL)
}
