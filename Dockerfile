FROM golang:1.21-alpine as build

RUN apk --no-cache add curl && \
    curl -L -o /tmp/pkl https://github.com/apple/pkl/releases/download/0.25.2/pkl-alpine-linux-amd64

WORKDIR /worker

COPY go.* ./
RUN go mod download && go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -v -a -trimpath -ldflags="-s -w" -o dist/ \
    cmd/basilisk/basilisk.go


FROM alpine:3.19

COPY --from=build /tmp/pkl /usr/bin/pkl
RUN chmod +x /usr/bin/pkl

WORKDIR /app

COPY --from=build /worker/config/AppConfig.pkl /app/config/AppConfig.pkl
COPY --from=build /worker/dist/basilisk /usr/bin/basilisk

EXPOSE 8080

CMD [ "basilisk" ]
