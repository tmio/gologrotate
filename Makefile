all: build
build: goimports
	go install gologrotate
test: goimports
	go test -v gologrotate

goimports: imports
	./vendor/bin/goimports -w src/gologrotate/

imports:
	go install golang.org/x/tools/cmd/goimports

