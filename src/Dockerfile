FROM golang:latest as builder
RUN cd ..
RUN mkdir src
WORKDIR employee
COPY src ./
ARG version=dev
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -ldflags "-X main.version=$version" -o src ./cmd/src/main.go

FROM alpine
COPY --from=builder /go/employee/test .
CMD ["./test"]
