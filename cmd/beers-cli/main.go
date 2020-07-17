package main

import (
	cli "github.com/patriciabonaldy/punkapi/internal/cli"
	"github.com/spf13/cobra"
)

func main() {

	repo := cli.NewRepository()
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(repo))
	rootCmd.Execute()

}
