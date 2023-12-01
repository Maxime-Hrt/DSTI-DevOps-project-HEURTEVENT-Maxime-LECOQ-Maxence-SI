FROM golang:1.21

WORKDIR /user
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o /usr/local/bin/devops-project .

EXPOSE 8080
CMD ["/usr/local/bin/devops-project"]
