all: build
build: goimports
	go install gologrotate
test: goimports
	go test -v gologrotate

goimports: godeps imports
	./vendor/bin/goimports -w src/gologrotate/

godeps:
	go get golang.org/x/tools/cmd/goimports

imports:
	go install golang.org/x/tools/cmd/goimports

