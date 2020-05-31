FROM golang:latest

RUN go mod vendor

RUN go build .

EXPOSE  9000

CMD [ "./changecoder.ws" ]