package fetching

import (
	"fmt"
	"math"

	beerEntity "github.com/patriciabonaldy/punkapi/domain/entity"
	port "github.com/patriciabonaldy/punkapi/domain/ports"
	"github.com/patriciabonaldy/punkapi/infrastructure/errors"
)

//Service struct services beer
type Service struct {
	bR port.BeerRepo
}

// NewFetchy initialize csv repository
func NewFetchy(r port.BeerRepo) port.Fetching {
	return &Service{r}
}

//FetchBeers service get all beers
func (r *Service) FetchBeers() ([]beerEntity.Beer, error) {
	beers, err := r.bR.GetBeers()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return beers, nil
}

//FetchByID service get beer by id
func (r *Service) FetchByID(id int) (beerEntity.Beer, error) {
	beers, err := r.FetchBeers()

	if err != nil {
		return beerEntity.Beer{}, err
	}

	beersPerRoutine := 10
	numRoutines := numOfRoutines(len(beers), beersPerRoutine)

	b := make(chan beerEntity.Beer)
	done := make(chan bool, numRoutines)

	for i := 0; i < numRoutines; i++ {
		toSearch := make([]beerEntity.Beer, beersPerRoutine)
		copy(toSearch[:], beers[i:i+beersPerRoutine])

		go func(beers []beerEntity.Beer, b chan beerEntity.Beer, done chan bool) {
			for _, beer := range beers {
				if beer.ProductID == id {
					b <- beer
				}
			}
			done <- true
		}(toSearch, b, done)
	}

	var beer beerEntity.Beer
	i := 0
	for i < numRoutines {
		select {
		case beer = <-b:
			return beer, nil
		case <-done:
			i++
		}
	}

	return beerEntity.NewBeerEmpty(), errors.NewUnreacheableBeerErr("No existe la beer con id %v", id)
}

func numOfRoutines(numOfBeers, beersPerRoutine int) int {
	return int(math.Ceil(float64(numOfBeers) / float64(beersPerRoutine)))
}
