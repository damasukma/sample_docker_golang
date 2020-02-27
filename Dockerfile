# Builder
FROM golang:1.12.8-alpine3.10 as builder

RUN apk update && apk upgrade && \
    apk --update add git make


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

COPY . .   

EXPOSE 9000

ENTRYPOINT CompileDaemon --build="go build -o engine /app/main.go" --command=./engine

# RUN make engine
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o engine main.go



# Distribution
# FROM alpine:latest

# RUN apk update && apk upgrade && \
#     apk --update --no-cache add tzdata && \
#     mkdir /app 

# WORKDIR /root/ 

# COPY --from=builder /app/engine .

# EXPOSE 9000

# CMD ["./engine"]
