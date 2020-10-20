package blockchain

import (
	"DataCertPlatform/utils"
	"bytes"
	"time"
)

//定义区块结构体，用于表示区块
type Block struct {
	Height    int64  //区块的高度，第几个区块
	TimeStamp int64  //时间戳
	PrevHash  []byte //前一个区块的hash值
	Data      []byte //数据字段
	Hash      []byte //当前区块的hash值
	Version   string //版本号
}
//创建一个新区块
func NewBlock(height int64,prevHash []byte,data []byte)Block{
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
		//Hash:      nil,
		Version:   "0x01",
	}
	heightBytes,_ := utils.Int64ToByte(block.Height)
	timeStampBytes,_ := utils.Int64ToByte(block.TimeStamp)
	versionBytes := utils.StringToBytes(block.Version)

	var blockBytes []byte
	//bytes.Join拼接
	bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		block.PrevHash,
		block.Data,
		versionBytes,
	},[]byte{})
	block.Hash = utils.SHA256HashBlock(blockBytes)
	return block
}

//创建创世区块
func CreateGenesisBlock()Block{
	genesisBlock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
	return genesisBlock
}