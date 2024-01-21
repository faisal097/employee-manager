# Specifies a parent image
FROM golang:1.20-bullseye
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .
 
# Installs Go dependencies
RUN go mod download
 
# Builds your app with optional configuration
RUN go build -o server cmd/main.go
 
# Tells Docker which network port your container listens on
#EXPOSE 8080
 
# Specifies the executable command that runs when the container starts
#CMD [ “/server” ]

#docker build --rm -t faisal0097/experiments:server-dev-v1 .
#docker run -it --rm --name devserver -p 3333:3333 faisal0097/experiments:server-dev-v1 /app/server -conf ./config.json