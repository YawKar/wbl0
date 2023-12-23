# Build the executable
FROM golang:1.21.5-alpine3.19 AS build-stage
WORKDIR /app

COPY . /app/

WORKDIR /app/cmd/server

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o serverx

# Run the executable
FROM build-stage AS run-server-stage
WORKDIR /
COPY --from=build-stage /app/cmd/server/serverx /serverx

ENTRYPOINT [ "/serverx" ]
