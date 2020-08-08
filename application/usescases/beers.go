package application

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	beer "github.com/patriciabonaldy/punkapi/domain/entity"
	svc "github.com/patriciabonaldy/punkapi/domain/ports"
)

// BeerInterface definiton of methods to access a data beer
type BeerInterface interface {
	FindBeers(id string) ([]beer.Beer, error)
	SaveBeers(nameFile string, beers []beer.Beer) error
	saveToCsv(nameFile string, records [][]string) error
}

//BeerService struct services beer
type BeerService struct {
	fetching svc.Fetching
}

// NewService initialize
func NewService(f svc.Fetching) BeerInterface {
	return &BeerService{f}
}

//FindBeers get beers all or by id
func (b *BeerService) FindBeers(id string) ([]beer.Beer, error) {
	var beers []beer.Beer
	var err error

	if id == "" {
		beers, err = b.fetching.FetchBeers()

	} else {
		var beer beer.Beer
		i, _ := strconv.Atoi(id)
		beer, err = b.fetching.FetchByID(i)
		beers = append(beers, beer)
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return beers, nil
}

// SaveBeers  data from csv
func (b *BeerService) SaveBeers(nameFile string, beers []beer.Beer) error {
	var arraysBeers [][]string
	var errString error = nil

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	for _, beer := range beers {
		bArray := beer.BeerRow()
		arraysBeers = append(arraysBeers, bArray)
	}
	b.saveToCsv(nameFile, arraysBeers)

	return errString
}

func (b *BeerService) saveToCsv(nameFile string, records [][]string) error {
	f, err := os.Create(nameFile)
	defer f.Close()
	//CSVheader(f)
	csvWriter := csv.NewWriter(f)

	csvWriter.WriteAll(records)
	csvWriter.Flush()

	return err
}
