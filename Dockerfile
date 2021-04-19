#Details of golang image
#https://github.com/docker-library/golang/blob/f12c995e27fef88ccb984605ab4748737ae3a778/1.16/alpine3.13/Dockerfile
FROM golang:1.16.3-alpine3.13

#At the moment we skip go env settings

#The WORKDIR instruction sets the working directory for any RUN, CMD, ENTRYPOINT, COPY and ADD instructions.
WORKDIR /go/src/app

COPY . .

COPY go.mod .
COPY go.sum .

#This dowload all the project's dependencies
RUN go mod download

RUN go build .

EXPOSE 3000

CMD ./sermo