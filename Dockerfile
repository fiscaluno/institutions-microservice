FROM golang:1.10

<<<<<<< HEAD
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

WORKDIR /go/src/bin

CMD ["app"]
=======
WORKDIR /go/src/github.com/fiscaluno/hyoga
COPY . .

RUN go get -u github.com/kardianos/govendor
RUN govendor sync

>>>>>>> c42ff3d26dc59b82a981fd4173284a377c11c78d
