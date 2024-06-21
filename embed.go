package embed

import (
	"io/fs"
	"sync"
)

// Cache struct to hold the cached files
type Cache struct {
	mu    sync.Mutex
	files map[string]fs.File
}

// afs struct to implement custom file system with cache
type afs struct {
	cache Cache
	subFS fs.FS
}

// Open method to open a file and cache it if not already cached
func (a *afs) Open(name string) (fs.File, error) {
	a.cache.mu.Lock()
	defer a.cache.mu.Unlock()

	// Check if the file is in the cache
	if file, ok := a.cache.files[name]; ok {
		return file, nil
	}

	// If not in the cache, load it from the embedded file system
	f, err := a.subFS.Open(name)
	if err != nil {
		return nil, err
	}

	// Cache the file
	a.cache.files[name] = f
	return f, nil
}

// GetStaticFS returns the static files with caching
func GetStaticFS(fsys fs.FS, dir string) (fs.FS, error) {
	sfs, err := fs.Sub(fsys, dir)
	if err != nil {
		return nil, err
	}

	return &afs{
		cache: Cache{
			files: make(map[string]fs.File),
		},
		subFS: sfs,
	}, nil
}
