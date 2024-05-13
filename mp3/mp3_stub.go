//go:build !lame
// +build !lame

package mp3

import (
	"errors"
	"os"
)

func EncodeToMp3AndSave(inp *os.File, out string) error {
	return errors.New("Cannot use -mp3 when compiled without -tags lame")
}
