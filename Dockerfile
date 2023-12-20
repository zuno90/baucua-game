FROM golang:1.20 AS build

WORKDIR /usr/src/app

COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ .
RUN go build -o bin .

# --------- #
# --------- #

FROM golang:1.20 as state

WORKDIR /usr/src/app
COPY --from=build /usr/src/app/bin .

# CMD /usr/src/app/bin
