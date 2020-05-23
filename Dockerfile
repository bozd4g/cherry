FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /server .

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=0 /server /server

COPY ./templates/* /templates/
COPY ./templates/layout/* /templates/layout/
COPY ./static/js/* /static/js/
COPY ./static/style/* /static/style/

CMD ["/server"]


