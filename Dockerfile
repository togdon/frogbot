FROM golang:1.27.2 as build

WORKDIR /go/src/frogbot
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/frogbot ./bot

FROM gcr.io/distroless/static-debian13

COPY --from=build /go/src/frogbot/frogs /frogs/
COPY --from=build /go/bin/frogbot /
CMD ["/frogbot"]
