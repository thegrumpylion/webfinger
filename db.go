package webfinger

import (
	"encoding/json"
	"errors"
)

// ErrResNotFound is thrown when resource is not in DB
var ErrResNotFound = errors.New("resource not found")

// DB storage interface
type DB interface {
	// Get the resource from DB
	Get(q *Query) (*Resource, error)
}

// MemDB memory implementation of DB interface
type MemDB struct {
	m map[string][]byte
}

// NewMemDB returns a new memory DB
func NewMemDB(m map[string][]byte) *MemDB {
	return &MemDB{
		m: m,
	}
}

// Get the resource from DB
func (db *MemDB) Get(q *Query) (*Resource, error) {
	data, ok := db.m[q.Resource]
	if !ok {
		return nil, ErrResNotFound
	}

	// decode resource
	res := &Resource{}
	err := json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	// filter links
	links := res.Links
	if q.Rel != nil {
		links = []*Link{}
		idx := map[string]*Link{}

		for _, l := range res.Links {
			idx[l.Rel] = l
		}

		for _, rel := range q.Rel {
			if l, ok := idx[rel]; ok {
				links = append(links, l)
			}
		}
	}
	res.Links = links

	return res, nil
}
