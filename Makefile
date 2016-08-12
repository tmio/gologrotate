SHELL := /bin/bash
all: build

clean:
	find src/gologrotate -name "*.gz" | xargs rm

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

docker: 
	docker build -t gologrotate .

run_docker_it:
	docker run -it -v "`pwd`/src/gologrotate/tests/find_all:/logs" gologrotate

check_gz_files_present:
	./check_gz_files_present.sh
	
it: clean build test docker run_docker_it check_gz_files_present
