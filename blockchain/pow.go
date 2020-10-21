package blockchain

import "math/big"

//工作量证明算法
type ProofOfWork struct {
	Target *big.Int //系统的目标值
	Block  Block   //要找的nonce值对应的区块
}

//实例化一个Pow算法的实例
func NewPow(block Block) ProofOfWork {
	t := big.NewInt(1)
	t = t.Lsh(t, 255)
	pow := ProofOfWork{
		Target: t,
		Block:  block,
	}
	return pow
}
