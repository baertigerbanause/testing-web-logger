FROM golang:1.15.12-alpine
WORKDIR /testing-web-logger
COPY . .
RUN go get -d ./...
RUN go build
CMD [ "./testing-web-logger" ]