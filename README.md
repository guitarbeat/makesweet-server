# MakeSweet Server + Slack Bot 🎨

A Go API wrapping [paulfitz/makesweet](https://github.com/paulfitz/makesweet) for generating animated GIF memes, plus a Slack bot that turns emoji reactions on images into MakeSweet GIFs.

## Repository Layout

```text
.
├── makesweet/       # C++ GIF rendering engine and templates
├── server/          # Go HTTP API
├── slack-bot/       # Python Slack Socket Mode integration
├── Dockerfile       # API image
├── compose.yml      # API + bot local orchestration
└── render.yaml      # API + bot Render blueprint
```

The Slack bot was imported from [`guitarbeat/makesweet-slack-bot`](https://github.com/guitarbeat/makesweet-slack-bot) with its Git history preserved. New development for both components can happen here.

## Deploy on Render

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/guitarbeat/makesweet-server)

> ⏱️ **First build takes 5–10 minutes** due to C++ compilation of the makesweet core.

The blueprint creates separate `makesweet-server` and `makesweet-slack-bot` web services. Set the bot's secret `SLACK_BOT_TOKEN` and `SLACK_APP_TOKEN` values, then set `MAKESWEET_URL` to the public URL Render assigns to `makesweet-server` (for example, `https://makesweet-server.onrender.com`). The bot exposes `/health` on Render's assigned `PORT` while its Slack connection runs in Socket Mode.

## API Endpoints

All endpoints accept `POST` with `multipart/form-data`.

| Template | Endpoint | Form Fields |
|----------|----------|-------------|
| Heart Locket | `/api/gif/heart-locket` | `image-left`, `image-right` |
| Billboard | `/api/gif/billboard` | `image` |
| Flag | `/api/gif/flag` | `image` |
| Flying Bear | `/api/gif/flying-bear` | `image` |
| Nesting Doll | `/api/gif/nesting-doll` | `image-left`, `image-mid`, `image-right` |
| Circuit Board | `/api/gif/circuit` | `image` ⚠️ *currently broken* |

### Example

```bash
curl -X POST https://your-server.onrender.com/api/gif/flag \
  -F "image=@photo.png" \
  -o flag.gif
```

## Running Both Components Locally

Docker Compose is the recommended path:

```bash
git clone https://github.com/guitarbeat/makesweet-server.git
cd makesweet-server
cp .env.example .env
```

Add your Slack credentials to `.env`:

```dotenv
SLACK_BOT_TOKEN=xoxb-your-token
SLACK_APP_TOKEN=xapp-your-token
```

Then start both services:

```bash
docker compose up
```

The API runs at `http://localhost:8080/api`, and the bot health endpoint is available at `http://localhost:3000/health`. Inside Compose, the bot calls the API at `http://makesweet-server:8080`.

To run only the API, no Slack credentials are needed:

```bash
docker compose up makesweet-server
```

For Slack app creation, scopes, events, and non-Docker development, see [`slack-bot/README.md`](slack-bot/README.md).

### Swagger Docs

Once running, visit `/api/docs/index.html` for interactive API documentation.

## Tools

<div>
  <table>
    <tr>
      <th style="text-align:center">Golang</th>
      <th style="text-align:center">Docker</th>
      <th style="text-align:center">Gin</th>
      <th style="text-align:center">Swaggo</th>
    </tr>
    <tr>
      <td style="text-align:center"><a href="https://go.dev"><img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" height="90" alt="Golang" /></a></td>
      <td style="text-align:center"><a href="https://www.docker.com"><img src="https://uxwing.com/wp-content/themes/uxwing/download/brands-and-social-media/docker-icon.svg" height="90" alt="Docker" /></a></td>
      <td style="text-align:center"><a href="https://gin-gonic.com"><img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" height="90" alt="Gin" /></a></td>
      <td style="text-align:center"><a href="https://github.com/swaggo/swag"><img src="https://raw.githubusercontent.com/swaggo/swag/master/assets/swaggo.png" height="90" alt="Swaggo" /></a></td>
    </tr>
  </table>
</div>

## Credits

- [paulfitz/makesweet](https://github.com/paulfitz/makesweet) — The core C++ GIF generation engine
- [Maheshivara/makesweet-server](https://github.com/Maheshivara/makesweet-server) — Original Go server implementation
- [`guitarbeat/makesweet-slack-bot`](https://github.com/guitarbeat/makesweet-slack-bot) — Original Slack bot repository
