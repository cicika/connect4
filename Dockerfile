FROM golang:1.12-stretch
## ENV GO111MODULE=on
## ENV GOPROXY=direct
## RUN git config --global url."git@github.com/".insteadOf "https://github.com/"
WORKDIR /app/
## COPY go.mod .
## COPY go.sum .
COPY connect4.sh .
COPY connect4 .
RUN chmod u+x connect4.sh
## RUN go mod download -json
## RUN go install github.com/cicika/connect4

ENTRYPOINT ./connect4.sh
