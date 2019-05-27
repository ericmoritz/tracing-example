FROM golang:1.12.5-stretch
ADD . /src
WORKDIR /src
RUN go build .

FROM golang:1.12.5-stretch
COPY --from=0 /src/tracing-example ./tracing-example
ENTRYPOINT [ "./tracing-example" ]
