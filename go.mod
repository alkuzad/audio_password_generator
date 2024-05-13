module github.com/alkuzad/audio_password_generator

go 1.22.2

require (
	github.com/go-audio/audio v1.0.0
	github.com/go-audio/wav v1.1.0
	github.com/viert/go-lame v0.0.0-20201108052322-bb552596b11d // go:build lame
)

require github.com/go-audio/riff v1.0.0 // indirect
