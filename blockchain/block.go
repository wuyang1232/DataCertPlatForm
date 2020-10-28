package blockchain

import (
	"bytes"
	"encoding/gob"
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
	Nonce     int64  //区块对应的nonce值
}

//创建一个新区块
func NewBlock(height int64, prevHash []byte, data []byte) Block {
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
		//Hash:      nil,
		Version: "0x01",
	}

	//找nonce值，通过工作量证明算法寻找
	pow := NewPow(block) //pow相当于工人
	hash,nonce := pow.Run()
	block.Nonce = nonce
	block.Hash = hash

	////将结构体数据转换为[]byte类型
	//heightBytes, _ := utils.Int64ToByte(block.Height)
	//timeStampBytes, _ := utils.Int64ToByte(block.TimeStamp)
	//versionBytes := utils.StringToBytes(block.Version)
	//nonceBytes, _ := utils.Int64ToByte(block.Nonce)
	//
	//var blockBytes []byte
	////bytes.Join拼接
	//bytes.Join([][]byte{
	//	heightBytes,
	//	timeStampBytes,
	//	block.PrevHash,
	//	block.Data,
	//	versionBytes,
	//	nonceBytes,
	//}, []byte{})
	////调用hash值进行计算
	//block.Hash = utils.SHA256HashBlock(blockBytes)

	//挖矿竞争，和获得记账权
	return block
}

//创建创世区块
func CreateGenesisBlock() Block {
	genesisBlock := NewBlock(0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil)
	return genesisBlock
}

//对区块进行序列化
func (b Block) Serialize()[]byte{
	buff :=  new(bytes.Buffer)//缓冲区
	encoder := gob.NewEncoder(buff)
	encoder.Encode(b)//将区块b放入到序列化编码器中
	return buff.Bytes()
}
//区块的反序列操作
func DeSerialize(data []byte)(*Block,error){
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil{
		return nil,err
	}
	return &block,nil
}