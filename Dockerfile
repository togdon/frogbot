FROM golang:1.22 as build

WORKDIR /go/src/frogbot
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/frogbot ./bot

FROM gcr.io/distroless/static-debian12

COPY --from=build /go/src/frogbot/frogs /frogs/
COPY --from=build /go/bin/frogbot /
CMD ["/frogbot"]
