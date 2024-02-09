FROM golang:1.21-alpine as build

WORKDIR /worker

COPY go.* ./
RUN go mod download && go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -v -a -trimpath -ldflags="-s -w" -o dist/ \
    cmd/basilisk/basilisk.go


FROM golang:1.21-alpine

WORKDIR /app

COPY --from=build /worker/dist/basilisk /usr/bin/basilisk

# EXPOSE 8080

CMD [ "basilisk" ]
