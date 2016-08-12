FROM golang:1.6

COPY . /
WORKDIR /
RUN . ./gopath.sh && go install gologrotate 


ENTRYPOINT ["/bin/gologrotate"]

CMD ["-now", "/logs"]
