package cli

import (
	"fmt"

	service "github.com/patriciabonaldy/punkapi/application/usescases"
	beerscli "github.com/patriciabonaldy/punkapi/domain/entity"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "idFlag"
const nameFileFlag = "nameFile"

// InitBeersCmd initialize beers command
func InitBeersCmd(service service.BeerInterface) *cobra.Command {

	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(service),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id de la beer")
	beersCmd.Flags().StringP(nameFileFlag, "n", "", "name of file")

	return beersCmd
}

func runBeersFn(service service.BeerInterface) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		nameFile, _ := cmd.Flags().GetString(nameFileFlag)
		id, _ := cmd.Flags().GetString(idFlag)
		var beers []beerscli.Beer
		var err error

		beers, err = service.FindBeers(id)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(beers)
			if nameFile == "" {
				nameFile = "beers.csv"
			}
			service.SaveBeers(nameFile, beers)
		}

	}
}
