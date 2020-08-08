package main

import (
	"os"
	"runtime/pprof"

	fetch "github.com/patriciabonaldy/punkapi/application/fetching"
	service "github.com/patriciabonaldy/punkapi/application/usescases"
	repo "github.com/patriciabonaldy/punkapi/infrastructure/adapter/repository"
	"github.com/patriciabonaldy/punkapi/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	//CPU profiling code start here
	f, _ := os.Create("beers.cpu.prof")
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	repo := repo.NewRepository()
	fetch := fetch.NewFetchy(repo)
	service := service.NewService(fetch)
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(service))
	rootCmd.Execute()

	//Memory profiling code start here
	f2, _ := os.Create("beers.mem.prof")
	defer f2.Close()
	pprof.WriteHeapProfile(f2)

}
