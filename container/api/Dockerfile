FROM golang:1.15.7 as build

WORKDIR /go/src/github.com/akaimo/nginx-lua-sample

COPY . .

RUN set -x \
  && go build -o api main.go


FROM gcr.io/distroless/base-debian10

COPY --from=build /go/src/github.com/akaimo/nginx-lua-sample/api /api

CMD ["/api"]
