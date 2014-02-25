package doctor

import (
	"fmt"
	"io"
	"io/ioutil"
)

// Cache is a key-value store intended to be used as an in memory file store.
type Cache map[string][]byte

// Load will load data from the given reader and store it at given key
func (c Cache) Load(key string, data io.Reader) error {
	if key == "" {
		return fmt.Errorf("no")
	}
	b, err := ioutil.ReadAll(data)
	if err != nil {
		return nil
	}
	c[key] = b
	return nil
}

// LoadFile will load the entire contents of the file given by fname
// and store it in the Cache keyed with its filename
func (c Cache) LoadFile(fname string) error {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	c[fname] = b
	return nil
}

// Get value from cache at given key.
func (c Cache) Get(key string) ([]byte, error) {
	b, ok := c[key]
	if !ok {
		return nil, fmt.Errorf("no value found for %s", key)
	}
	return b, nil
}