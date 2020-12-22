[Ondrej Sika (sika.io)](https://sika.io) | <ondrej@sika.io>

# [example] Docker Monitoring Using Prometheus & cAdvisor

    2020 Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/example-docker-monitoring

## Usage

### Start

```
make up
```

### Generate Load

```
make loadgen
```

### See

- Prometheus: <http://localhost:9090>
- cAdvisor: <http://localhost:8081>
- Alert Manager: <http://localhost:9093>
- Maildev: <http://localhost:8082>
- Traefik Dashboard: <http://localhost:8080>
- Web 1: <http://web1-127-0-0-1.nip.io>
- Web 2: <http://web2-127-0-0-1.nip.io>

### Stop

```
make down
```
