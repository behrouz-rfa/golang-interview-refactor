package static

import (
	"embed"
)

//go:embed images/*
//go:embed *.html
var WebUI embed.FS
