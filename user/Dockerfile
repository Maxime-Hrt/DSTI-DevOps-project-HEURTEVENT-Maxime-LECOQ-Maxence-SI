FROM golang:1.21

WORKDIR /user
COPY ./user/go.mod ./user/go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/devops-project .

EXPOSE 8080
CMD ["/usr/local/bin/devops-project"]
