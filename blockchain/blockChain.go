package blockchain

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"math/big"
)

const BLOCKCHAIN = "blockchain.db"
const BUCKET_NAME = "blocks"
const LAST_HASH = "lasthash"

var CHAIN *BlockChain

//区块链结构体的定义，代表的是整条区块链
type BlockChain struct {
	LastHash []byte   //区块链中最新区块的hash，用于查询最新的区块内容
	BoltDb   *bolt.DB //区块链中操作区块数据文件的数据库操作对象
}

//创建一个区块链
func NewBlockChain() *BlockChain {
	var bc *BlockChain
	//1、先打开文件
	db, err := bolt.Open(BLOCKCHAIN, 0600, nil)
	//2、查看blockchain.db文件
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME)) //假设有桶
		if bucket == nil {                       //没有桶，要新创建桶
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_HASH))
		if len(lastHash) == 0 { //桶中没有lasthash记录，，需要创建创世区块并保存
			//创世区块
			genesis := CreateGenesisBlock()
			//区块序列化一户的数据
			genesisBytes := genesis.Serialize()
			//创世区块保存到boltdb中
			bucket.Put(genesis.Hash, genesisBytes)
			//更新指向最新区块的lasthash值
			bucket.Put([]byte(LAST_HASH), genesis.Hash)
			bc = &BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
		} else { //桶当中已有lasthash的记录，不再需要创世区块，只需要读取即可
			lasthash1 := bucket.Get([]byte(LAST_HASH))
			bc = &BlockChain{
				LastHash: lasthash1,
				BoltDb:   db,
			}

		}
		return nil
	})

	//if err != nil {
	//	panic(err.Error())
	//}

	////把创世区块保存到数据库文件中去
	//db.Update(func(tx *bolt.Tx) error {
	//	bucket,err := tx.CreateBucket([]byte(BUCKET_NAME))
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	//序列化
	//	genesisBytes := genesis.Serialize()
	//	//把创世区块存储到桶里
	//	bucket.Put(genesis.Hash,genesisBytes)
	//	//更新最新区块的hash值
	//	bucket.Put([]byte(LAST_HASH),genesis.Hash)
	//	return nil
	//})
	//
	CHAIN = bc
	return bc
}

//该方法用于便利区块链chain.db文件，并将所有的区块查出，并返回
func (bc BlockChain) QueryAllBlocks() ([]*Block, error) {
	blocks := make([]*Block, 0) //blocks是一个切片容器，用于盛放查询到的区块

	db := bc.BoltDb
	var err error
	//从chain.db文件查询所有区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("查询区块链数据失败！")
			return err
		}
		//bucket存在
		eachHash := bc.LastHash
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0) //默认值0的大整数
		for {
			//根据区块的hash值获取对应的区块
			eachBlockBytes := bucket.Get(eachHash)
			//反序列化操作
			eachBlock, _ := DeSerialize(eachBlockBytes)
			//将便利到的没哟个区块放入到切片容器中
			blocks = append(blocks, eachBlock)

			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig) == 0 { //找到了创世区块
				break //条促循环
			}
			//不满足条件，没哟找到创世区块
			eachHash = eachBlock.PrevHash
		}
		return nil
	})

	return blocks, err
}

//该方法用于完成根据用户输入的高度查询你对应的区块信息
func (bc BlockChain) QueryBlockByHeight(height int64) (*Block, error) {
	if height < 0 {
		return nil, nil
	}
	db := bc.BoltDb

	var errs error
	var eachBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			errs = errors.New("读取数据失败！")
			return errs
		}
		//each:每一个
		eachHash := bc.LastHash
		for { //获取最后一个区块的hash
			eachBlockBytes := bucket.Get(eachHash)
			//拿到最后一个区块的byte类型
			//eachBlockBytes := bucket.Get(eachBlockHash)
			//反序列化操作
			eachBlock, errs = DeSerialize(eachBlockBytes)
			if errs != nil {
				return errs
			}
			if eachBlock.Height < height {
				break
			}
			if eachBlock.Height == height {
				break
			}
			//如果高度匹配不满足用户条件
			eachHash = eachBlock.PrevHash
		}

		return nil
	})
	return eachBlock, errs
}

//保存数据到区块链中：先生成一个新区块，然后将新区块添加到区块链中
func (bc *BlockChain) SaveData(data []byte) (Block, error) {
	//1、从文件中读取到最新的区块
	db := bc.BoltDb
	var lastBlock *Block
	//error 的自定义
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("读取区块链数据失败")
			return err
			//panic("读取区块链数据失败")
		}
		//lastHash := bucket.Get([]byte(LAST_HASH))
		lastBlockBytes := bucket.Get(bc.LastHash)
		//反序列化

		lastBlock, _ = DeSerialize(lastBlockBytes)
		return nil
	})
	//新建一个区块
	newBlock := NewBlock(lastBlock.Height+1, lastBlock.Hash, data)
	//把新区块放到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//序列化后的区块数据
		blockBytes := newBlock.Serialize()
		//fmt.Println("保存数据到区块，序列化后的区块数据：",blockBytes)
		//把新创建的区块存入到boltdb数据库中
		bucket.Put(newBlock.Hash, blockBytes) //内存中的数据不具备存储型，需要序列化
		//更新LASTHASH对应的值，更新为最新存储的区块的hash值
		bucket.Put([]byte(LAST_HASH), newBlock.Hash)
		bc.LastHash = newBlock.Hash //将区块链实例的LASTHASH值更新为最新区块的hash

		return nil
	})
	//返回值语句，newBlock，err ，其中err可能包含错误信息
	return newBlock, err
}

//该方法用于根据用户输入的认证号查询到对应的区块信息
func (bc BlockChain)QueryBlockByCertId(cert_id string)(*Block,error){
	db := bc.BoltDb
	var err error
	var block *Block
	fmt.Println("查询中",cert_id)
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil{//判断桶是否存在
			err = errors.New("查询链上数据发生错误，请重试！")
			return err
		}
		eachHash := bc.LastHash
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0)
		for{
			eachBlockBytes := bucket.Get(eachHash)
			eachBlock,err := DeSerialize(eachBlockBytes)
			fmt.Println("所有区块信息",eachBlock)

			fmt.Println("所有区块hash信息",string(eachBlock.Data))
			fmt.Println(cert_id)
			if err != nil{
				break
			}

			//将遍历到的区块中的数据跟用户提供的认证号进行比较
			if string(eachBlock.Data) == cert_id{//if成立，找到区块
				block = eachBlock
				break
			}
			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig) == 0{
				break
			}
			eachHash = eachBlock.PrevHash
		}
		return nil
	})
	fmt.Println(block)
	return block,err
}
