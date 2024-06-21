# Package embed provides access to files embedded in the running Go program.

This package is a wrapper around the [embed](https://golang.org/pkg/embed/) package, with cache functionality.

## Installation

```bash
go get github.com/inovacc/embed
```

## Usage

```go
package main

import (
	"embed"
	"fmt"
	efs "github.com/inovacc/embed"
	"io"
)

//go:embed testdata
var assets embed.FS

func main() {
	// Load embedded filesystem
	sfs, err := efs.GetStaticFS(assets, "testdata")
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

	fmt.Println(string(data))
}
```
