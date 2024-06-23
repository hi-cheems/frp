package frpc

import (
	"embed"

	"github.com/hi-cheems/frp/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
