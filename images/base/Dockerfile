FROM kalilinux/kali-rolling

LABEL maintainer="hadenlabs"

ENV BASE_DEPS \
    bash \
    kali-linux-large \
    pciutils \
    bash-completion

ENV BUILD_DEPS \
    fakeroot \
    curl \
    openssl

RUN apt-get -y update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
    $BASE_DEPS \
    $BUILD_DEPS \
    && apt-get clean \
    && apt-get purge -y \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* \
    printf "alias ll='ls $LS_OPTIONS -l'\nalias l='ls $LS_OPTIONS -lA'\n\n# enable bash completion in interactive shells\nif [ -f /etc/bash_completion ] && ! shopt -oq posix; then\n    . /etc/bash_completion\nfi\n" > /root/.bashrc

WORKDIR /data

CMD ["/bin/bash"]
