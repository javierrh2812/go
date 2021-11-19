package project

import "sync"

type URLStore struct {
	urls map[string]string // map from short url to long url
	mu sync.RWMutex
}