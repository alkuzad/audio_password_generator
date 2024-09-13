package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/alkuzad/audio_password_generator/mp3"
	"github.com/alkuzad/audio_password_generator/nato"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

var WAVBuffers = make(map[string]*audio.IntBuffer)

func main() {
	sizeFlag := flag.Int("size", 0, "Size of the password")
	mp3EnabledFlag := flag.Bool("mp3", false, "Enable wave to mp3 compression")
	dirFlag := flag.String("dir", "", "Target directory to write output to")

	flag.Parse()

	if *sizeFlag <= 0 {
		fmt.Printf("Invalid size %d. Please provide a positive integer for size.\n", *sizeFlag)
		return
	}

	// Pre-load all relevant WAV files
	if err := preloadWAVFiles(); err != nil {
		fmt.Printf("Error preloading WAV files: %v\n", err)
		return
	}

	input := generateRandomStringFromNatoMap(*sizeFlag)
	natoString := nato.ToNato(input)

	// Collect all buffers
	var buffers []*audio.IntBuffer

	for _, word := range natoString {
		if buffer, ok := WAVBuffers[word]; ok {
			buffers = append(buffers, buffer)
		} else {
			log.Fatalf("Invalid word %s", word)
		}
	}

	addSuccess(&buffers)

	mergedBuffer := mergeBuffers(buffers)

	wavFilePath := filepath.Join(*dirFlag, "password.wav")

	writeWAVFile(wavFilePath, mergedBuffer)
	if *mp3EnabledFlag == true {
		fh, _ := os.Open(wavFilePath)
		defer fh.Close()
		defer os.Remove(fh.Name())
		if err := mp3.EncodeToMp3AndSave(fh, filepath.Join(*dirFlag, "password.mp3")); err != nil {
			log.Fatalf("Error happened on wave->mp3 conversion\nerr: %s", err)
		}
	}

	os.WriteFile(filepath.Join(*dirFlag, "password.txt"), []byte(input), 0644)
}

func addSuccess(buffers *[]*audio.IntBuffer ) error {
	buffer, err := readWAVFile("sound/success.wav")
	if err != nil {
		return err
	}
	*buffers = append(*buffers, buffer)
	return nil
}

func preloadWAVFiles() error {
	for _, word := range nato.NatoMap {
		filename := fmt.Sprintf("sound/%s.wav", word)
		buffer, err := readWAVFile(filename)
		if err != nil {
			return err
		}
		WAVBuffers[word] = buffer
	}
	return nil
}

// Generate a random string using only characters from NatoMap keys
func generateRandomStringFromNatoMap(size int) string {
	var chars []rune
	for char := range nato.NatoMap {
		chars = append(chars, char)
	}

	var result strings.Builder
	result.Grow(size)
	for i := 0; i < size; i++ {
		result.WriteRune(chars[rand.Intn(len(chars))])
	}
	return result.String()
}

// readWAVFile reads a WAV file and returns its buffer
func readWAVFile(filename string) (*audio.IntBuffer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := wav.NewDecoder(file)
	buffer, err := decoder.FullPCMBuffer()
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// mergeBuffers merges multiple IntBuffers into one
func mergeBuffers(buffers []*audio.IntBuffer) *audio.IntBuffer {
	if len(buffers) == 0 {
		return nil
	}

	// Assume all buffers have the same format
	totalLength := 0
	for _, buffer := range buffers {
		totalLength += len(buffer.Data)
	}

	mergedData := make([]int, totalLength)
	offset := 0
	for _, buffer := range buffers {
		copy(mergedData[offset:], buffer.Data)
		offset += len(buffer.Data)
	}

	return &audio.IntBuffer{Data: mergedData, Format: buffers[0].Format, SourceBitDepth: buffers[0].SourceBitDepth}
}

func writeWAVFile(filename string, buffer *audio.IntBuffer) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := wav.NewEncoder(file, buffer.Format.SampleRate, buffer.SourceBitDepth, buffer.Format.NumChannels, 1)
	if err := encoder.Write(buffer); err != nil {
		return err
	}
	return encoder.Close()
}
