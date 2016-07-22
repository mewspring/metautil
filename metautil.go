// Package metautil provides utility functions for handling FLAC metadata
// blocks.
package metautil

import (
	"log"

	"github.com/mewkiz/flac"
	"github.com/mewkiz/flac/meta"
)

// AppendBlock appends the given metadata block to the FLAC stream.
func AppendBlock(stream *flac.Stream, block *meta.Block) {
	stream.Blocks = append(stream.Blocks, block)
}

// AppendBlockBody appends the given metadata block body to the FLAC stream.
func AppendBlockBody(stream *flac.Stream, body interface{}) {
	// NOTE: The block header omit values for IsLast and Length. These are added
	// by flac.Encode when encoding the FLAC stream. Whether this is a good
	// approach or not is yet to be decided.
	block := new(meta.Block)
	switch body := body.(type) {
	case *meta.StreamInfo:
		block.Type = meta.TypeStreamInfo
		block.Body = body
	case *meta.Application:
		block.Type = meta.TypeApplication
		block.Body = body
	case *meta.SeekTable:
		block.Type = meta.TypeSeekTable
		block.Body = body
	case *meta.VorbisComment:
		block.Type = meta.TypeVorbisComment
		block.Body = body
	case *meta.CueSheet:
		block.Type = meta.TypeCueSheet
		block.Body = body
	case *meta.Picture:
		block.Type = meta.TypePicture
		block.Body = body
	default:
		log.Printf("support for appending metadata block body of type %T not yet implemented", body)
		return
	}
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
	Body *meta.VorbisComment
}

// NewCommentBlock returns a new Vorbis comment metadata block.
func NewCommentBlock(vendor string) *CommentBlock {
	return &CommentBlock{
		Body: &meta.VorbisComment{
			Vendor: vendor,
		},
	}
}

// Add adds the given key-value pair to the Vorbis comment metadata block.
func (c *CommentBlock) Add(key, val string) {
	tag := [2]string{key, val}
	c.Body.Tags = append(c.Body.Tags, tag)
}
