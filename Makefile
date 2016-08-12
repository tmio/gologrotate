all: build
docker: 
	docker build -t gologrotate .

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

run_docker:
	docker run -it -v "`pwd`/src/gologrotate/tests/find_all:/logs" gologrotate

it: build test docker run_docker
