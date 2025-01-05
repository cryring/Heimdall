package store

import "github.com/cryring/Heimdall/internal/node"

type Store struct {
}

func New() *Store {
	return &Store{}
}

func (s *Store) Add(n *node.Node) error {
	return nil
}

func (s *Store) Get(id string) (*node.Node, error) {
	return nil, nil
}
