FROM golang:latest

RUN mkdir -p /usr/local/go/src/NASDAQ_Slot_Machine
WORKDIR /usr/local/go/src/NASDAQ_Slot_Machine
ADD . /usr/local/go/src/NASDAQ_Slot_Machine

RUN go mod download
RUN go build ./main.go

EXPOSE 8080
CMD ["./main"]