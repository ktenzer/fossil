FROM fedora:latest


LABEL ios.k8s.display-name="fossul-app" \
    maintainer="Keith Tenzer <ktenzer@redhat.com>"

RUN dnf install -y golang

RUN mkdir -p /go/src/fossul
RUN mkdir /app

WORKDIR /go/src/fossul

ENV FOSSUL_BUILD_PLUGIN_DIR=/app/plugins
ENV GOPATH=/go
ENV GOBIN=/app

RUN curl -L https://github.com/openshift/origin/releases/download/v3.11.0/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz |tar xz;cp openshift*/oc /app;cp openshift*/kubectl /app

RUN curl https://raw.githubusercontent.com/golang/dep/v0.5.1/install.sh | sh

#RUN go get -u github.com/gorilla/mux

COPY . /go/src/fossul

#RUN cd /app && go build engine/util && go install engine/app
RUN /go/src/fossul/fossul-app-build.sh

RUN chown -R 1001:0 /app && \
    chmod -R 775 /app
 #   chmod -R ug+rwX /app

RUN chmod -R 777 /tmp

RUN echo "1.0" > /etc/imageversion

RUN ls -l /app

USER 1001

WORKDIR /app

CMD /app/fossul-app-startup.sh
