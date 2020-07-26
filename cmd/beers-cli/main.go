package main

import (
	"github.com/patriciabonaldy/punkapi/internal/cli"
	service "github.com/patriciabonaldy/punkapi/internal/cli/fetching"
	storage "github.com/patriciabonaldy/punkapi/internal/cli/storage"
	"github.com/spf13/cobra"
)

func main() {

	repo := storage.NewRepository()
	fetch := service.NewFetchy(repo)
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(fetch))
	rootCmd.Execute()

}
