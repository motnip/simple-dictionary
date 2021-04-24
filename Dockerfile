#Details of golang image
#https://github.com/docker-library/golang/blob/f12c995e27fef88ccb984605ab4748737ae3a778/1.16/alpine3.13/Dockerfile
FROM golang:1.16.3-alpine3.13

#At the moment we skip go env settings

#The WORKDIR instruction sets the working directory for any RUN, CMD, ENTRYPOINT, COPY and ADD instructions.
WORKDIR /go/src/app/

COPY . .

#This dowload all the project's dependencies
RUN go mod download

#The run instruction executes when we build the image.
#That means the command passed to run executes on top
#of the current image in a new layer.
#Then the result is committed to the image
RUN go build .

EXPOSE 3000

#A default command that executes when the container is starting.
#This command will be overrifhted by the argument passed to
#the docker run command
CMD ./sermo
