FROM golang:1.20-bullseye

ENV ROOT=/workspace

WORKDIR ${ROOT}

COPY go.mod go.sum ./

RUN apt-get update

RUN apt-get install git && \
    apt-get install curl && \
    apt-get install -y tree

RUN go mod download
RUN go install github.com/cosmtrek/air@v1.29.0

# COPY . .

CMD ["air", "-c", ".air.toml"]
