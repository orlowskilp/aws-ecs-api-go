ARG VERSION

FROM centos:${VERSION}

# Needs to be defined after FROM
ARG USERNAME

RUN useradd -ms /bin/bash ${USERNAME} && \
# Install all the language specific packages
    yum groupinstall -y "Development tools" && \
    yum install -y golang && \
# Install extensions for VSCode
    go get -v github.com/sqs/goreturns && \
    go get -v github.com/rogpeppe/godef && \
    go get -v github.com/ramya-rao-a/go-outline && \
    go get -v github.com/mdempsky/gocode && \
    go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs && \
    go get -v github.com/stamblerre/gocode && \
    go get -v golang.org/x/lint/golint && \
    go get -v golang.org/x/tools/cmd/gorename && \
    go get -v github.com/go-delve/delve/cmd/dlv && \
# Install project-specific dependencies
    go get -v github.com/gin-gonic/gin && \
    go get -v github.com/stretchr/testify/assert && \
# Move files to user directory
    cp -r /root/go /home/${USERNAME} && \
    chown -R ${USERNAME}. /home/${USERNAME}/go
USER ${USERNAME}
WORKDIR /home/${USERNAME}
