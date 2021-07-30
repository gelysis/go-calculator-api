FROM golang:1-alpine AS GO_BUILD
COPY . /catathon2021
WORKDIR /catathon2021/src
RUN go build -o /catathon2021/catathon2021-server

FROM alpine:3
WORKDIR app
COPY --from=GO_BUILD /catathon2021/catathon2021-server ./
EXPOSE 8181
CMD ./catathon2021-server
