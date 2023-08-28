FROM golang:1.20 AS build

WORKDIR /usr/src/game-server

COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go build -o bin .

# --------- #
# --------- #

FROM golang:1.20 as state

WORKDIR /usr/src/game-server

COPY --from=build /usr/src/game-server/bin .

# CMD [ "/usr/src/game-server/bin" ]

CMD /usr/src/game-server/bin
