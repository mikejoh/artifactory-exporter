# Artifactory Exporter for Prometheus

An alternative Artifactory Exporter written in Go.

## Notes
* This exporter is currently in active development and only exports storage info data fetched via the REST API. At the moment there's only a couple of metrics exported.
* Tested on Artifactory OSS 6.11.1

## How-to

1. Run `make build` to create a single binary (in `./bin`) or to create a Docker image run `make build-docker-linux`
2. To run the exporter as-is you can do `./bin/artifactory-exporter`. See the flags below for the available options:
```
Flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
      --web.listen-address=":9627"  
                                 Address:Port to listen on for web interface and telemetry.
      --web.telemetry-path="/metrics"  
                                 Path under which to expose metrics.
      --web.artifactory-api-url="http://localhost:8081/artifactory/api/storageinfo"  
                                 Artifactory REST API URL.
  -u, --web.username="admin"     Artifactory API user.
  -p, --web.password="password"  Artifactory API user password.
```

To run the Docker image:
```
docker run -d -p 9627:9627 artifactory-exporter:latest <FLAGS>
```