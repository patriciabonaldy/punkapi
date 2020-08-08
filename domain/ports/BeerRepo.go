package ports

import beer "github.com/patriciabonaldy/punkapi/domain/entity"

// BeerRepo definiton of methods to access a data beer
type BeerRepo interface {
	GetBeers() ([]beer.Beer, error)
}
