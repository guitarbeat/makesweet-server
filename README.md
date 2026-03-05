# MakeSweet Server 🎨

A Go server wrapping [paulfitz/makesweet](https://github.com/paulfitz/makesweet) for generating animated GIF memes. Originally by [Maheshivara](https://github.com/Maheshivara/makesweet-server), forked and configured for one-click Render deployment.

## Deploy

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/guitarbeat/makesweet-server)

> ⏱️ **First build takes 5–10 minutes** due to C++ compilation of the makesweet core.

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

## Running Locally

### Docker (recommended)

```bash
git clone https://github.com/guitarbeat/makesweet-server.git
cd makesweet-server
cp .env.example .env
docker compose up
```

The server runs at `http://localhost:8080/api`.

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
