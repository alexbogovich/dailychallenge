package main

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestIsUnivalFromSample(t *testing.T) {

	node := NewNode(
		"0",
		NewNode("1", nil, nil),
		NewNode("0",
			NewNode("1", NewNode("1", nil, nil), NewNode("1", nil, nil)),
			NewNode("0", nil, nil),
		),
	)

	_, count := IsUnival(node)

	assert.Equal(t, 5, count)
}

func TestIsUnivalSimpleLine(t *testing.T) {

	node := NewNode(
		"0",
		nil,
		NewNode("0",
			NewNode("1", NewNode("1", nil, nil), NewNode("1", nil, nil)),
			NewNode("0", nil, nil),
		),
	)

	_, count := IsUnival(node)

	assert.Equal(t, 4, count)
}

func TestIsUnivalNone(t *testing.T) {

	node := NewNode(
		"1",
		nil,
		NewNode("2",
			NewNode("3", NewNode("5", nil, nil), NewNode("6", nil, nil)),
			NewNode("4", nil, nil),
		),
	)

	_, count := IsUnival(node)

	assert.Equal(t, 3, count)
}
