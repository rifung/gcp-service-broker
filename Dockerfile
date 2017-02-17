FROM golang:latest

RUN curl https://glide.sh/get | sh

WORKDIR /gcp-sb/src/gcp-service-broker
COPY auth ./auth/
COPY brokerapi ./brokerapi/
COPY cmd ./cmd/
COPY creds ./creds/
COPY db_service ./db_service/
COPY glide.* ./
COPY db_service ./db_service
COPY *.go ./
COPY utils ./utils

ENV GOPATH /gcp-sb
RUN glide install
RUN go build

# Note that this also sets it for the service broker
ENV PORT 8080
EXPOSE ${PORT}

ENTRYPOINT ["./gcp-service-broker"]
