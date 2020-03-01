// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"

	"github.com/voidfiles/magda/graph/generated"
	"github.com/voidfiles/magda/graph/model"
)

func (r *queryResolver) GetEntry(ctx context.Context, id string) (model.Entry, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
