package cli

import beer "github.com/patriciabonaldy/punkapi/internal"

// BeerRepo definiton of methods to access a data beer
type BeerRepo interface {
	GetBeers() ([]beer.Beer, error)
}
