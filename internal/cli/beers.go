package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const nameFileFlag = "nameFile"

// InitBeersCmd initialize beers command
func InitBeersCmd(repository BeerRepo) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(repository),
	}

	beersCmd.Flags().StringP(nameFileFlag, "n", "", "name of file")

	return beersCmd
}

func runBeersFn(repository BeerRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		nameFile := ""
		beers, err := repository.GetBeers()
		if err != nil {
			fmt.Println("Fallo la api")
		}
		fmt.Println(beers)
		nameFile, _ = cmd.Flags().GetString(nameFileFlag)
		if nameFile == "" {
			nameFile = "beers.csv"
		}
		SaveBeers(nameFile, beers)
	}
}
