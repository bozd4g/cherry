FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /server .

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=0 /server /server

ADD ./templates/ /templates/
ADD ./static/ /static/

CMD ["/server"]


