FROM golang:latest AS builder
RUN mkdir /artifactory-exporter
COPY . /artifactory-exporter/
WORKDIR /artifactory-exporter
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o artifactory-exporter .

FROM scratch
COPY --from=builder /artifactory-exporter/artifactory-exporter .
ENTRYPOINT ["./artifactory-exporter"]