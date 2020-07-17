package cli

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	punkapiEndpoint = "https://api.punkapi.com/v2/beers"
)

type repository struct {
	url string
}

// NewRepository initialize csv repository
func NewRepository() BeerRepo {
	return &repository{url: punkapiEndpoint}
}

// SaveBeers  data from csv
func SaveBeers(nameFile string, beers []Beer) error {
	var arraysBeers [][]string

	for _, beer := range beers {
		arraysBeers = append(arraysBeers, beer.BeerRow())
		//arraysBeers = append(arraysBeers, beerscli.ToArray(beer))
		//beer.CSVrow(f)
	}
	saveToCsv(nameFile, arraysBeers)

	return nil
}

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]Beer, error) {
	var beers []Beer
	response, err := http.Get(fmt.Sprintf("%v", r.url))

	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, err
	}

	return beers, nil
}

func saveToCsv(nameFile string, records [][]string) error {
	f, err := os.Create(nameFile)
	defer f.Close()
	CSVheader(f)
	csvWriter := csv.NewWriter(f)

	csvWriter.WriteAll(records)
	csvWriter.Flush()

	return err
}
