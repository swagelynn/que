<p align="center">
  <h1>que</h1>
  <br>
  a lightweight solution to q&a on a personal site
  <br>
</p>


## Docker Compose Quick Start
```yaml
services:
  que:
    image: ghcr.io/swagelynn/que:latest
    volumes:
      - ./data:/quedata
    ports:
      - 5438:5438
    environment:
      - DISCORD_WEBHOOK=https://discord.com/api/webhooks/../.. \# (optional)
```