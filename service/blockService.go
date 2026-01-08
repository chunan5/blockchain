package service

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64  `gorm:"type:bigint;not null"`
	Data      []byte `gorm:"type:varchar(255);not null"`
	PrevHash  []byte `gorm:"type:varchar(255);not null"`
	Hash      []byte `gorm:"type:varchar(255);not null"`
}

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func createBlock(data string, prevBlockHash []byte, hash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.setHash()
	return block
}

func (b *Block) setPrevHash() {}
