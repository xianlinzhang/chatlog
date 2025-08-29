package silk

import (
	"fmt"
)

// Mock implementation that just returns an error
func Silk2MP3(data []byte) ([]byte, error) {
	return nil, fmt.Errorf("silk to mp3 conversion not available: go-silk library not installed")
}
