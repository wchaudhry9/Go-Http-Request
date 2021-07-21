FROM golang:1.16-alpine
RUN echo "Starting up server"
WORKDIR /appRoot
COPY . .
CMD ["go", "run", "dockerHttp.go"]
