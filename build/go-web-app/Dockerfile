FROM golang:1.18-alpine AS build

LABEL maintainer="Nuno Guimarães"

WORKDIR /go/src/app

COPY . .

# It will download all the dependencies needed and run
RUN go get -d -v ...

# It will build our app into an exec file
RUN go build -v ./cmd/go-web-app/main.go

FROM golang:1.18-alpine

COPY --from=build /go/src/app/main go-web-app

ENTRYPOINT [ "./go-web-app" ]
