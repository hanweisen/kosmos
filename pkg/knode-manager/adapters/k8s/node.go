package k8sadapter

import (
	"context"
	"errors"

	corev1 "k8s.io/api/core/v1"
)

var (
	// ErrNotImplemented is returned when a requested feature is not implemented.
	ErrNotImplemented = errors.New("not implemented")
)

type NodeAdapter struct {
}

func NewNodeAdapter() (*NodeAdapter, error) {
	return nil, ErrNotImplemented
}

func (n *NodeAdapter) Configure(ctx context.Context, node *corev1.Node) {
}

func (n *NodeAdapter) Trace(_ context.Context) error {
	return ErrNotImplemented
}

func (n *NodeAdapter) NotifyStatus(_ context.Context, _ func(*corev1.Node)) {
}
