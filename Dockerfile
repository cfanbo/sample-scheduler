FROM --platform=$TARGETPLATFORM ubuntu:20.04
WORKDIR .
COPY bin/sample-scheduler /usr/local/bin
CMD ["sample-scheduler"]