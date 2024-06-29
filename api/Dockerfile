FROM golang:1.20-bullseye

ENV ROOT=/api

WORKDIR ${ROOT}

COPY api/go.mod api/go.sum ./

RUN apt-get update

RUN apt-get install git && \
    apt-get install curl && \
    apt-get install -y tree

COPY . .

RUN go mod download
RUN go install github.com/cosmtrek/air@v1.29.0



CMD ["air", "-c", ".air.toml"]
