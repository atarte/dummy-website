# Build
FROM golang:alpine AS build

ARG APP_VERSION

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
ADD assets ./assets
# ADD handlers/ ./handlers
# ADD models/ ./models
# ADD templates/ ./templates
# ADD views/ ./views

RUN go build -o /tmp/dummy-website -ldflags "-X 'main.Version=$APP_VERSION' -X 'main.Address=""' -X 'main.Port="80"' " ./main.go

# Dev
FROM alpine:latest

WORKDIR /app

COPY --from=build /tmp/dummy-website ./dummy-website
# ADD templates/ ./templates
ADD assets ./assets

RUN chmod 555 ./dummy-website
USER 1000

CMD [ "./dummy-website" ]
