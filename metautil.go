// Package metautil provides utility functions for handling FLAC metadata
// blocks.
package metautil

import (
	"github.com/mewkiz/flac"
	"github.com/mewkiz/flac/meta"
)

// AppendBlock appends the given metadata block to the FLAC stream.
func AppendBlock(stream *flac.Stream, block *meta.Block) {
	stream.Blocks = append(stream.Blocks, block)
}

// RemoveBlock removes the given metadata block from the FLAC stream.
func RemoveBlock(stream *flac.Stream, block *meta.Block) {
	var blocks []*meta.Block
	for _, b := range stream.Blocks {
		if b == block {
			continue
		}
		blocks = append(blocks, b)
	}
	stream.Blocks = blocks
}

// RemoveBlockType removes the given metadata block type from the FLAC stream.
func RemoveBlockType(stream *flac.Stream, typ meta.Type) {
	var blocks []*meta.Block
	for _, block := range stream.Blocks {
		if block.Type == typ {
			continue
		}
		blocks = append(blocks, block)
	}
	stream.Blocks = blocks
}

// A CommentBlock represents a Vorbis comment metadata block.
type CommentBlock struct {
	Block *meta.VorbisComment
}

// NewCommentBlock returns a new Vorbis comment metadata block.
func NewCommentBlock(vendor string) *CommentBlock {
	return &CommentBlock{
		Block: &meta.VorbisComment{
			Vendor: vendor,
		},
	}
}

// Add adds the given key-value pair to the Vorbis comment metadata block.
func (c *CommentBlock) Add(key, val string) {
	tag := [2]string{key, val}
	c.Block.Tags = append(c.Block.Tags, tag)
}
