go run cmd/beers-cli/main.go beers

#go run cmd/beers-cli/main.go beers -i 1 -n file.csv

go tool pprof cmd/beers-cli/main.go beers.cpu.prof
#go tool pprof cmd/beers-cli/main.go beers.mem.prof
#top 5