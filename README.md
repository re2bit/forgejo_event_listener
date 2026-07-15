# Dynamic Runner Subsystem

Ein Service zum Verarbeiten von Forgejo Runners.

### Informationen
- Der Service ist standardmäßig auf Port **8080** freigegeben.
- Ein Health-Check ist unter `http://localhost:8080/health` erreichbar.

### Starten
Lokal:
```bash
go run main.go serve
```

Mit Docker Compose:
```bash
docker compose up
```

### Multi Arch Build
Create Multi Arch Builder:
```bash
docker buildx create --name container-builder --driver docker-container --bootstrap --use
```

Do Multi Arch Build:
```bash
docker buildx bake --push
```

### Docker Hub
https://hub.docker.com/r/re2bit/dynamic-runner-subsystem