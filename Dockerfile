FROM golang:alpine as builder

WORKDIR /builder
COPY ./ /builder

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o maintenance cmd/main.go

FROM scratch
COPY --from=builder /builder/maintenance /maintenance

EXPOSE 8080

ENTRYPOINT [ "/maintenance" ]
