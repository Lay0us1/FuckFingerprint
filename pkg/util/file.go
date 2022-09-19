package util

import (
	"bufio"
	"github.com/yhy0/FuckFingerprint/pkg/logging"
	"os"
)

// LoadFile content to slice
func LoadFile(filename string) (lines []string) {
	f, err := os.Open(filename)
	if err != nil {
		logging.Logger.Fatal("LoadFile err, ", err)
		return
	}
	defer f.Close() //nolint
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return
}

func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}
