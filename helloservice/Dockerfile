FROM golang:1.12 AS build
# Use Go Modules, work outside GOPATH
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM scratch
COPY --from=build /app/main .
ENTRYPOINT [ "./main" ]