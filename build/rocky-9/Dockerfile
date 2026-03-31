FROM rockylinux:9
LABEL org.opencontainers.image.authors="support@halon.io"

ARG HALON_REPO_USER
ARG HALON_REPO_PASS

RUN yum update -y && yum install -y epel-release

RUN echo "[halon]" >> /etc/yum.repos.d/Halon.repo
RUN echo "name=Halon repository" >> /etc/yum.repos.d/Halon.repo
RUN echo "baseurl=https://${HALON_REPO_USER}:${HALON_REPO_PASS}@repo.halon.io/dists/rhel9/amd64/stable" >> /etc/yum.repos.d/Halon.repo
RUN echo "gpgcheck=1" >> /etc/yum.repos.d/Halon.repo
RUN echo "gpgkey=https://raw.githubusercontent.com/halon/changelog/master/pgp-keys/7F0A73B5.asc" >> /etc/yum.repos.d/Halon.repo
RUN echo "enabled=1" >> /etc/yum.repos.d/Halon.repo
RUN yum install -y halon-6.10.0

RUN yum install -y cmake3 make gcc-c++ rpm-build

RUN echo -n "ROCKY9" > /OSRELEASE.txt

RUN mkdir /tmp/halon
RUN cd /tmp/halon \
    && curl -O -L https://go.dev/dl/go1.25.5.linux-amd64.tar.gz \
    && echo "9e9b755d63b36acf30c12a9a3fc379243714c1c6d3dd72861da637f336ebb35b go1.25.5.linux-amd64.tar.gz" | sha256sum -c \
    && tar -C /usr/local -xzf go1.25.5.linux-amd64.tar.gz
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
