package ports

import beer "github.com/patriciabonaldy/punkapi/domain/entity"

// Fetching provides beer fetching operations
type Fetching interface {
	// FetchBeers fetch all beers from repository
	FetchBeers() ([]beer.Beer, error)
	// FetchByID filter all beers and get only the beer that match with given id
	FetchByID(id int) (beer.Beer, error)
}
