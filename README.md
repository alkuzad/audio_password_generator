# Random MP3 password generator

Small program to generate sound-based passwords of given size

## Using

You can choose if you want to install liblame in your system (libmp3lame-dev on Ubuntu for example) or work on wave files and optionally convert manually.

### without MP3 support

1. `go build`
2. `./audio_password_generator -size 50`

###  with MP3 support

1. `go build -tags lame`
2. `./audio_password_generator -mp3 -size 50`

### Compilation using Docker

1. `docker build -t audio_password_generator .`
2. `docker run --rm -v ./output:/tmp/output audio_password_generator 50`
