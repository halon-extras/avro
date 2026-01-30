FROM --platform=linux/amd64 mcr.microsoft.com/azurelinux/base/core:3.0
LABEL org.opencontainers.image.authors="support@halon.io"

ARG HALON_REPO_USER
ARG HALON_REPO_PASS

RUN tdnf install -y ca-certificates azurelinux-repos-extended

RUN echo "[halon]" >> /etc/yum.repos.d/Halon.repo
RUN echo "name=Halon repository" >> /etc/yum.repos.d/Halon.repo
RUN echo "baseurl=https://${HALON_REPO_USER}:${HALON_REPO_PASS}@repo.halon.io/dists/azure3/amd64/stable" >> /etc/yum.repos.d/Halon.repo
RUN echo "gpgcheck=1" >> /etc/yum.repos.d/Halon.repo
RUN echo "gpgkey=https://raw.githubusercontent.com/halon/changelog/master/pgp-keys/7F0A73B5.asc" >> /etc/yum.repos.d/Halon.repo
RUN echo "enabled=1" >> /etc/yum.repos.d/Halon.repo
RUN tdnf install -y halon-6.7.0

RUN tdnf install -y cmake3 make gcc-c++ rpm-build clang build-essential

RUN echo -n "AZURE3" > /OSRELEASE.txt

RUN mkdir /tmp/halon
RUN cd /tmp/halon \
    && curl -O -L https://go.dev/dl/go1.23.4.linux-amd64.tar.gz \
    && echo "6924efde5de86fe277676e929dc9917d466efa02fb934197bc2eba35d5680971 go1.23.4.linux-amd64.tar.gz" | sha256sum -c \
    && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
RUN /usr/local/go/bin/go install github.com/CycloneDX/cyclonedx-gomod/cmd/cyclonedx-gomod@latest
ENV PATH="$PATH:/usr/local/go/bin:/root/go/bin"

RUN mkdir /licenses
RUN cp /usr/local/go/LICENSE /licenses/LICENSE.go
RUN cd /tmp/halon \
    && curl -O -L https://github.com/davecgh/go-spew/blob/master/LICENSE \
    && cp LICENSE /licenses/LICENSE.davecgh-go-spew
RUN cd /tmp/halon \
    && curl -O -L https://raw.githubusercontent.com/golang/snappy/refs/heads/master/LICENSE \
    && cp LICENSE /licenses/LICENSE.golang-snappy
RUN cd /tmp/halon \
    && curl -O -L https://raw.githubusercontent.com/linkedin/goavro/refs/heads/master/LICENSE \
    && cp LICENSE /licenses/LICENSE.linkedin-goavro
RUN cd /tmp/halon \
    && curl -O -L https://raw.githubusercontent.com/pmezard/go-difflib/refs/heads/master/LICENSE \
    && cp LICENSE /licenses/LICENSE.pmezard-go-difflib
RUN cd /tmp/halon \
    && curl -O -L https://raw.githubusercontent.com/stretchr/objx/refs/heads/master/LICENSE \
    && cp LICENSE /licenses/LICENSE.stretchr-objx
RUN cd /tmp/halon \
    && curl -O -L https://raw.githubusercontent.com/stretchr/testify/refs/heads/master/LICENSE \
    && cp LICENSE /licenses/LICENSE.stretchr-testify
RUN cd /tmp/halon \
    && curl -O -L https://raw.githubusercontent.com/go-check/check/refs/heads/v1/LICENSE \
    && cp LICENSE /licenses/LICENSE.go-check-check
RUN cd /tmp/halon \
    && curl -O -L https://raw.githubusercontent.com/go-yaml/yaml/refs/tags/v3.0.1/LICENSE \
    && cp LICENSE /licenses/LICENSE.go-yaml-yaml

RUN rm -rf /tmp/halon

COPY build.sh /build.sh
CMD ["/build.sh"]
