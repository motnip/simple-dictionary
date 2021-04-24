#https://github.com/docker-library/golang/blob/f12c995e27fef88ccb984605ab4748737ae3a778/1.16/alpine3.13/Dockerfile
FROM golang:1.16.3-alpine3.13
ARG BUILD_VERSION
WORKDIR /go/src/app/
COPY . .
RUN go build -ldflags "-X main.BuildVersion=$BUILD_VERSION" .
EXPOSE 3000
CMD ./sermo
