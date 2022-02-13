FROM hadenlabs/kali-linux-full:core as base

# set environment variables
ENV DEBIAN_FRONTEND noninteractive

ENV BASE_DEPS \
    bash \
    pciutils \
    bash-completion

FROM base as kali-meta-packages

ENV KALI_META_PACKAGES \
    kali-tools-top10 \
    kali-tools-web

RUN apt-get -y update \
    && apt-get install -y --no-install-recommends \
    ${BASE_DEPS} \
    ${KALI_META_PACKAGES} \
    && apt-get clean \
    && apt-get purge -y \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


FROM kali-meta-packages as crossover

ARG VERSION=0.0.0

LABEL maintainer="Team Hadenlabs <support@hadenlabs.com>" \
      org.label-schema.vcs-url="https://github.com/hadenlabs/docker-kali-linux" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

# set environment variables
ENV DEBIAN_FRONTEND noninteractive

ENV KALI_PACKAGES \
    dnsmap \
    dnswalk

RUN apt-get -y update \
    && apt-get install -y --no-install-recommends \
    ${KALI_PACKAGES} \
    && apt-get clean \
    && apt-get purge -y \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /data

CMD ["/bin/bash"]
