FROM golang:1.15.7

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

COPY . /app

ENV APP_HOME /app

WORKDIR ${APP_HOME}

RUN go mod vendor
RUN go build -o natsApp

EXPOSE 8080

ENTRYPOINT [ "./${APP_HOME}/natsApp" ]