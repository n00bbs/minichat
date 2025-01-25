package minichat

import (
	"embed"
)

//go:embed migrations/*
var EmbeddedMigrationFs embed.FS
