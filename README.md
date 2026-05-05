# PixelAPI — Official Go SDK

Official Go SDK for [PixelAPI](https://pixelapi.dev) — *AI image, video, audio, and 3D API for builders.*

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![pixelapi.dev](https://img.shields.io/badge/docs-pixelapi.dev-blue)](https://pixelapi.dev/docs)

## Why PixelAPI

- **From $0.001 per image** — image generation, background removal at $0.010, upscaling at $0.060
- **15 AI tools, one API key** — image gen, BG removal, upscale, face restore, object removal, image edit, audio, voice, 3D, content moderation, smart-generate, ad creative, image search, portrait studio, interior design, photo relighting
- **No cold starts** — always-warm infrastructure, sub-3-second response on most endpoints
- **100 free credits on signup** — no credit card required
- **Production-ready** — same endpoints serve our paying customers at thousands of calls per day
- **GST invoice (India)** — registered Indian business, GST-compliant

## Installation

```bash
go get github.com/prakash-in21/pixelapi-go
```

## Quickstart

```go
package main

import (
    "fmt"
    "github.com/prakash-in21/pixelapi-go"
)

func main() {
    client := pixelapi.NewClient("YOUR_KEY")
    r, _ := client.Generate("product photo of red sneakers, white background")
    fmt.Println(r.OutputURL)
    client.Save(r, "out/sneakers.png")
}
```

## Run

```bash
go run examples/main.go
```

## Endpoints covered

This SDK wraps PixelAPI's REST API. Full endpoint documentation: [pixelapi.dev/docs](https://pixelapi.dev/docs).

| Endpoint | Description | Price |
|---|---|---|
| `POST /v1/image/generate` | AI image generation | $0.001/image |
| `POST /v1/image/remove-background` | Background removal | $0.010/image |
| `POST /v1/image/replace-background` | Background replacement | $0.075/image |
| `POST /v1/image/upscale` | 4× upscaling | $0.060/image |
| `POST /v1/image/face-restore` | Face restoration | $0.005/face |
| `POST /v1/image/remove-object` | Object removal | $0.025/image |
| `POST /v1/image/edit` | Prompt-driven edit | $0.020/edit |
| `POST /v1/image/portrait` | AI portrait studio | $2.39/portrait |
| `POST /v1/image/relight` | Photo relighting | $0.018/image |
| `POST /v1/image/interior-design` | Interior design / staging | $0.075/render |
| `POST /v1/image/moderate` | NSFW / content moderation | $0.0005/image |
| `POST /v1/audio/generate` | AI music generation | $0.007/clip |
| `POST /v1/tts/generate` | Text-to-speech (23 languages) | $0.015/30s |
| `POST /v1/3d/generate` | 3D model generation (GLB) | $0.50/model |

## Pricing

| Plan | Monthly | Credits | Best for |
|---|---|---|---|
| Free | $0 | 100 credits | Try the API |
| Starter | $10 | 10,000 credits | Side projects, low volume |
| Pro | $50 | 60,000 credits | Production apps |
| Scale | $200 | 300,000 credits | High-volume / agencies |

INR billing with GST invoice via Razorpay; international USD via Stripe / PayPal.

## Authentication

```
export PIXELAPI_KEY='your-key'
```

Get your API key at [pixelapi.dev/app](https://pixelapi.dev/app) (free, no credit card).

Pass it via the `X-API-Key` header (or use the SDK constructor — handled automatically).

## Rate limits

- Free: 10 requests / minute
- Starter: 60 requests / minute
- Pro: 300 requests / minute
- Scale: unlimited (fair-use)

## Errors

The SDK raises typed errors for 401 (auth), 402 (insufficient credits), 429 (rate-limited), and 5xx (server-side).

## Support

- Docs: [pixelapi.dev/docs](https://pixelapi.dev/docs)
- Email: support@pixelapi.dev

## License

MIT — see `LICENSE`.

## Contributing

Pull requests welcome. For major changes, open an issue first to discuss.

---

Built by the [PixelAPI](https://pixelapi.dev) team. Indian business, registered with GST.
