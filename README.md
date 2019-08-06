# Artifactory Exporter for Prometheus

An alternative Artifactory Exporter written in Go.

### Usage:

```
Flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
      --web.listen-address=":9607"  
                                 Address:Port to listen on for web interface and telemetry.
      --web.telemetry-path="/metrics"  
                                 Path under which to expose metrics.
      --web.artifactory-api-url="http://localhost:8081/artifactory/api/storageinfo"  
                                 Artifactory REST API URL.
  -u, --web.username="admin"     Artifactory API user.
  -p, --web.password="password"  Artifactory API user password.
```