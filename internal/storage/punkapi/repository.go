package punkapi

import intf "github.com/patriciabonaldy/punkapi/internal/cli"

const (
	punkapiEndpoint = "https://api.punkapi.com/v2/beers"
)

type repository struct {
	url string
}

// NewRepository initialize csv repository
func NewRepository() intf.BeerRepo {
	return &repository{url: punkapiEndpoint}
}

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]beerscli.Beer, error) {

}
