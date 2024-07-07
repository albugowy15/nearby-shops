FROM golang:1.22.4 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /nearby-shops-api cmd/web/main.go

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM gcr.io/distroless/static-debian12 AS build-release-stage
WORKDIR /
COPY --from=build-stage /nearby-shops-api /nearby-shops-api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/nearby-shops-api"]
