FROM golang:1.15.7

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

# set env variables
ENV APP_HOME=/app

# Copy the source from the current directory to the working Directory inside the container
ADD . /${APP_HOME}

# set work directory
WORKDIR /${APP_HOME}/.

# install dependencies
RUN go mod vendor

# build binary file
RUN go build -o natsApp

# expose app to world
EXPOSE 8080

# Command to run the executable
CMD [ "./natsApp" ]