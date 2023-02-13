FROM golang:1.19.4

RUN go install github.com/xo/xo@latest

RUN mkdir -p /xo

ADD ./deployment/wait-for-it.sh /wait-for-it.sh
ADD ./deployment/xo-generate.sh /xo-generate.sh

RUN chmod +x /wait-for-it.sh
RUN chmod +x /xo-generate.sh
