FROM golang

ADD . /src/

WORKDIR /src/
RUN go get .
RUN go build -o xsscatcher .
RUN mkdir -p /app/
RUN mkdir -p /var/log
RUN cp xsscatcher /app/.
CMD ["/app/xsscatcher"]