package maps

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("not found")

func (d Dictionary) Search(searchTerm string) (string, error) {
	result, ok := d[searchTerm]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}
