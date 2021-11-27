FROM golang:1.17.3-alpine3.14 AS build

RUN mkdir -p /home/app

COPY src /home/app/src
COPY go.mod go.sum /home/app/

WORKDIR /home/app

RUN go build src/iot-data-viewer-backend.go

FROM alpine:3.14 AS prod

RUN mkdir -p /home/app

COPY --from=build /home/app/iot-data-viewer-backend /home/app/

WORKDIR /home/app

CMD [ "/iot-data-viewer-backend" ]