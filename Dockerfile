FROM golang:1.10-stretch AS builder

COPY . /go/src/github.com/konsorten/junos-run-command/

WORKDIR /go/src/github.com/konsorten/junos-run-command/

RUN go get && go build

FROM golang:1.10-stretch

ENV JUNIPER_HOST=
ENV JUNIPER_USER=root
ENV JUNIPER_PASSWORD=
ENV JUNIPER_COMMAND=

COPY --from=builder /go/src/github.com/konsorten/junos-run-command/junos-run-command /go/bin/junos-run-command

ENTRYPOINT [ "/go/bin/junos-run-command" ]
