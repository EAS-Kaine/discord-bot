FROM golang:1.18.1-alpine3.15 as builder

WORKDIR /tmp/go
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -buildvcs=false -o app

FROM scratch
COPY --from=builder /tmp/go/app /bin/app
CMD [ "/bin/app" ]