FROM golang:1.21-alpine3.17 AS build
WORKDIR /build

COPY go.* .
RUN go mod download

ENV CGO_ENABLED=0
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /app ./cmd/app/main.go

FROM scratch

COPY --from=build /app /app
COPY ./swagger ./swagger

COPY env/app.env /

EXPOSE ${APP_PORT}

ENTRYPOINT ["/app", "--config=/app.env"]