FROM golang:alpine

WORKDIR /go

RUN apk add --no-cache git curl python3 py3-pip
RUN curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-389.0.0-linux-x86_64.tar.gz
RUN tar -xf google-cloud-cli-389.0.0-linux-x86_64.tar.gz
RUN ./google-cloud-sdk/install.sh
RUN rm google-cloud-cli-389.0.0-linux-x86_64.tar.gz

ENV PATH $PATH:/go/google-cloud-sdk/bin
RUN gcloud components install kubectl

COPY . .
RUN chmod +x pintu/startup-script.sh

CMD ["sh", "pintu/startup-script.sh"]

