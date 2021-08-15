package assets

import "embed"

//go:embed "controls/*"
var FileSystem embed.FS

const (
	ImageClose = "controls/image-close.png"
)
