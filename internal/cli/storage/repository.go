package cli

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/patriciabonaldy/punkapi/internal/errors"

	beer "github.com/patriciabonaldy/punkapi/internal"
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

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]beer.Beer, error) {
	var beers []beer.Beer
	response, err := http.Get(fmt.Sprintf("%v", r.url))

	if err != nil {
		return nil, errors.WrapUnreacheableBeerErr(err, "error obteniendo endpoint %v", r.url)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapUnreacheableBeerErr(err, "error leyendo el response %v", r.url)
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, errors.WrapUnreacheableBeerErr(err, "error parsing to beers")
	}

	return beers, nil
}

func saveToCsv(nameFile string, records [][]string) error {
	f, err := os.Create(nameFile)
	defer f.Close()
	//CSVheader(f)
	csvWriter := csv.NewWriter(f)

	csvWriter.WriteAll(records)
	csvWriter.Flush()

	return err
}
