FROM golang:1.16
WORKDIR /etc/source
COPY . /etc/source
RUN make build
RUN cp /etc/source/bin/thanos /usr/bin/thanos