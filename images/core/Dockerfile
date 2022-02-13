FROM kalilinux/kali-rolling:arm64

ARG VERSION=core

LABEL maintainer="Team Hadenlabs <support@hadenlabs.com>" \
      org.label-schema.vcs-url="https://github.com/hadenlabs/docker-kali-linux" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

# set environment variables
ENV DEBIAN_FRONTEND noninteractive

ENV BASE_DEPS \
    bash \
    kali-linux-core \
    pciutils \
    bash-completion

ENV BUILD_DEPS \
    fakeroot \
    curl \
    openssl

RUN apt-get -y update \
    && apt-get install -y --no-install-recommends \
    ${BASE_DEPS} \
    ${BUILD_DEPS} \
    && apt-get clean \
    && apt-get purge -y \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /data

CMD ["/bin/bash"]
