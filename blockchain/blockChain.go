package blockchain

import "github.com/boltdb/bolt"

const BLOCKCHAIN  = "blockchain.db"
const BUCKET_NAME  = "blocks"
const LAST_HASH  = "lasthash"

//区块链结构体的定义，代表的是整条区块链
type BlockChain struct {
	LastHash []byte   //区块链中最新区块的hash，用于查询最新的区块内容
	BoltDb   *bolt.DB //区块链中操作区块数据文件的数据库操作对象
}

//创建一个区块链
func NewBlockChain() BlockChain {
	//创世区块
	genesis := CreateGenesisBlock()
	db, err := bolt.Open(BLOCKCHAIN, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	bc := BlockChain{
		LastHash: genesis.Hash,
		BoltDb:   db,
	}
	//把创世区块保存到数据库文件中去
	db.Update(func(tx *bolt.Tx) error {
		bucket,err := tx.CreateBucket([]byte(BUCKET_NAME))
		if err != nil {
			panic(err.Error())
		}
		//序列化
		genesisBytes := genesis.Serialize()
		//把创世区块存储到桶里
		bucket.Put(genesis.Hash,genesisBytes)
		//更新最新区块的hash值
		bucket.Put([]byte(LAST_HASH),genesis.Hash)
		return nil
	})
	return bc
}

//保存数据到区块链中：先生成一个新区块，然后将新区块添加到区块链中
func (bc BlockChain) SaveData(data []byte) {
	//1、从文件中读取到最新的区块
	db := bc.BoltDb
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil{
			panic("读取区块链数据失败")
		}
		lastHash := bucket.Get([]byte(LAST_HASH))
		lastBlockBytes := bucket.Get(lastHash)
		//反序列化

		lastBlock,_ = DeSerialize(lastBlockBytes)
		return nil
	})
	//新建一个区块
	newBlock := NewBlock(lastBlock.Height+1, lastBlock.Hash, data)
	//把新区块放到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//把新创建的区块存入到boltdb数据库中
		bucket.Put(newBlock.Hash,newBlock.Serialize())//内存中的数据不具备存储型，需要序列化
		//更新LASTHASH对应的值，更新为最新存储的区块的hash值

		return nil
	})
}
