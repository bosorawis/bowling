# syntax=docker/dockerfile:1
FROM golang:1.16 AS Builder
ADD . /app
WORKDIR /app
RUN make test && make build


FROM scratch
COPY --from=Builder /app/bin/bowling .
CMD ["./bowling"]
