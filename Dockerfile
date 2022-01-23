FROM golang:1.17

WORKDIR /vm

COPY .. .

RUN make build

CMD ["make", "run"]
