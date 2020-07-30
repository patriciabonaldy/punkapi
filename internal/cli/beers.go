package cli

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	beerscli "github.com/patriciabonaldy/punkapi/internal"
	service "github.com/patriciabonaldy/punkapi/internal/cli/fetching"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "idFlag"
const nameFileFlag = "nameFile"

// InitBeersCmd initialize beers command
func InitBeersCmd(fetching service.Service) *cobra.Command {

	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(fetching),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id de la beer")
	beersCmd.Flags().StringP(nameFileFlag, "n", "", "name of file")

	return beersCmd
}

func runBeersFn(fetching service.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		nameFile := ""
		id, _ := cmd.Flags().GetString(idFlag)
		var beers []beerscli.Beer
		var err error

		if id == "" {
			beers, err = fetching.FetchBeers()

		} else {
			var beer beerscli.Beer
			i, _ := strconv.Atoi(id)
			beer, err = fetching.FetchByID(i)
			beers = append(beers, beer)
		}

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(beers)
			nameFile, _ = cmd.Flags().GetString(nameFileFlag)
			if nameFile == "" {
				nameFile = "beers.csv"
			}
			SaveBeers(nameFile, beers)
		}

	}
}

// SaveBeers  data from csv
func SaveBeers(nameFile string, beers []beerscli.Beer) error {
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
	saveToCsv(nameFile, arraysBeers)

	return errString
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
