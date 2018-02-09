FROM golang as builder
WORKDIR /src/hello
COPY hello-world.go .
RUN go build -o /hello

FROM scratch
COPY --from=builder /hello /hello
CMD ["/hello"]
