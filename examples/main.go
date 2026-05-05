package main

import (
	"fmt"
	"log"
	"os"

	"github.com/prakash-in21/pixelapi-go"
)

func main() {
	apiKey := os.Getenv("PIXELAPI_KEY")
	if apiKey == "" {
		log.Fatal("set PIXELAPI_KEY env var (get one free at https://pixelapi.dev/app)")
	}

	client := pixelapi.NewClient(apiKey)

	// 1. Generate an AI image
	r, err := client.Generate("product photo of red sneakers, white background, studio lighting")
	if err != nil {
		log.Fatalf("generate failed: %v", err)
	}
	fmt.Printf("generated: %s (used %.4f credits)\n", r.OutputURL, r.CreditsUsed)
	_ = client.Save(r, "out/sneakers.png")

	// 2. Remove background from a photo
	r, err = client.RemoveBackground("https://pixelapi.dev/demo/photo.jpg")
	if err != nil {
		log.Fatalf("remove-background failed: %v", err)
	}
	fmt.Printf("bg-removed: %s (used %.4f credits)\n", r.OutputURL, r.CreditsUsed)
}
