#build 
FROM alpine:latest
MAINTAINER Derek <peter890701@gmail.com>

WORKDIR /root

# fix library dependencies
# otherwise golang binary may encounter 'not found' error
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# fix certificate
RUN apk --update upgrade && \
  apk add curl ca-certificates && \
  update-ca-certificates && \
  rm -rf /var/cache/apk/*

# Copy binary files
COPY bin/* /root/

# Command 
CMD [ "./main.exe" ]
