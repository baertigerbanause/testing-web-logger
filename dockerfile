FROM golang:1.15.12-apline AS build
WORKDIR /testing-web-logger
COPY . .
RUN go get -d ./...
ADD *.go //testing-web-logger/
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest
COPY --from=build /testing-web-logger .
ENTRYPOINT [ "./testing-web-logger" ]
CMD [ "./testing-web-logger" ]