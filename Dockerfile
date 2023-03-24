FROM golang:1.20 as build

WORKDIR /go/src/frogbot
COPY . .

RUN go mod download
#RUN go vet -v
#RUN go test -v

RUN ls /go/src/frogbot
RUN CGO_ENABLED=0 go build -o /go/bin/frogbot ./bot

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/src/frogbot/frogs /frogs/
COPY --from=build /go/bin/frogbot /
CMD ["/frogbot"]
