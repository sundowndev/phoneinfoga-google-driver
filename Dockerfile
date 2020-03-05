FROM golang:1.14-alpine

LABEL maintainer="Sundowndev" \
  org.label-schema.build-date="2020-03-04T21:20:49Z" \
  org.label-schema.name="phoneinfoga" \
  org.label-schema.description="Advanced information gathering & OSINT tool for phone numbers." \
  #org.label-schema.version=$VERSION \
  org.label-schema.url="https://github.com/sundowndev/PhoneInfoga" \
  #org.label-schema.vcs-ref=$VCS_REF \
  org.label-schema.vcs-url="https://github.com/sundowndev/PhoneInfoga" \
  org.label-schema.vendor="Sundowndev" \
  org.label-schema.schema-version="1.0"

ADD . /opt/phoneinfoga

RUN go get -d -v ./...

RUN go build -o phoneinfoga

WORKDIR /opt/phoneinfoga

CMD ["phoneinfoga"]