FROM golang:1.16.5-buster 
WORKDIR $GOPATH/src/gitlab.com/APWHY/app-rokt-api
COPY . . 
RUN GOOS=linux GOARCH=amd64 go build -o /main
# could copy this over to a scratch image if we wanted to optimise size
EXPOSE 8080
ENTRYPOINT ["/main"]
