FROM golang:1.18 as BUILDER

ENV GO111MODULE=on

ENV CGO_ENABLED=0
ENV GOOS=$GOOS
ENV GOARCH=$GOARCH

WORKDIR /pokomand-go
COPY . .
RUN go mod download \
    && go mod verify \
    && go build -o /build/builldedApp main/main.go

FROM scratch as FINAL

WORKDIR /main
COPY --from=BUILDER /build/builldedApp .

ENTRYPOINT ["./builldedApp"]