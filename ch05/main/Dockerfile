# docker build -t goPost .
# docker run -it goPost

FROM 1.25.1-alpine3.22 AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /pro
ADD ./usePost05.go /pro/
WORKDIR /pro
RUN go mod tidy
RUN go build -o server usePost05.go

FROM alpine:latest

RUN mkdir /pro
COPY --from=builder /pro/server /pro/server
WORKDIR /pro
CMD ["/pro/server"]