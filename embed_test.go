package embed

import (
	"embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

//go:embed testdata
var assets embed.FS

func TestGetStaticFS(t *testing.T) {
	// Load embedded filesystem
	sfs, err := GetStaticFS(assets, "testdata")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Open a file
	f, err := sfs.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Read the file
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.Equal(t, "Hello, World!\r\n", string(data))
}
