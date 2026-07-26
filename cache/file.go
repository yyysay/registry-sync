package cache

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type FileCache struct {
	path string

	data map[string]bool

	mu sync.Mutex
}

func NewFileCache(
	path string,
) *FileCache {

	return &FileCache{

		path: path,

		data: make(
			map[string]bool,
		),
	}
}

func (c *FileCache) Check(
	ctx context.Context,
	key string,
) bool {

	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.data) == 0 {

		c.load()
	}

	return c.data[key]
}

func (c *FileCache) Save(
	ctx context.Context,
	key string,
) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = true

	return c.flush()
}

func (c *FileCache) load() {

	content, err := os.ReadFile(
		c.path,
	)

	if err != nil {

		return
	}

	_ = json.Unmarshal(
		content,
		&c.data,
	)
}

func (c *FileCache) flush() error {

	dir := filepath.Dir(
		c.path,
	)

	err := os.MkdirAll(
		dir,
		0755,
	)

	if err != nil {

		return err
	}

	content, err := json.MarshalIndent(
		c.data,
		"",
		"  ",
	)

	if err != nil {

		return err
	}

	return os.WriteFile(
		c.path,
		content,
		0644,
	)
}
