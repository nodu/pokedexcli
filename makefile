run:
	go build && ./pokedexcli

test:
	go test ./...

auto:
	# ls *.go 2>/dev/null | entr -r go build -v main.go
	ls *.go | entr -c sh -c 'go build && ./pokedexcli'
