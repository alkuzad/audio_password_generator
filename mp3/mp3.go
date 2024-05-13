//go:build lame
// +build lame

package mp3

import (
	"bufio"
	"os"

	"github.com/go-audio/wav"
	lame "github.com/viert/go-lame"
)

func EncodeToMp3AndSave(inp *os.File, out string) error {
	of, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer of.Close()
	enc := lame.NewEncoder(of)
	defer enc.Close()

	d := wav.NewDecoder(inp)
	d.ReadInfo()
	enc.SetInSamplerate(int(d.SampleRate))
	enc.SetNumChannels(int(d.NumChans))

	r := bufio.NewReader(inp)
	r.WriteTo(enc)

	return nil
}
