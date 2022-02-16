package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

/* Blockchain bloğumuzun üç ayrı alana sahip olması gerekecek.
Birincisi, blok zincirinin ana satış noktası güvenliği olduğundan, biraz karmaya
ihtiyacımız olacak. İkincisi, içinde bir şekilde korumaya veya muhafaza etmeye değer
bazı veriler olmadıkça, hash'i kullanamazdık, bu nedenle bazı verilere de ihtiyacımız olacak.
Son faktör, bloklarımızın birbirleriyle nasıl etkileşime girdiği olacaktır.*/
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

/*Bir blok zincirinin zinciri, basitçe bir blok gruplamasıdır.
Yani bu bana bir dizi için güzel bir uygulama gibi geldi.*/
type BlockChain struct {
	blocks []*Block
}

/*Bloğuna verdiğimiz verileri hash etmenin başlamak için iyi bir yer olacağına inanıyorum.
Bir bağımlılık zinciri oluşturmak için, bloktaki mevcut verileri ve üstümüzdeki bloktaki
verileri hash edeceğiz. Blok zincirinin güvenilirliği, önceki bloğun bu karmasına bağlıdır.
Bu bağlantı dizisi olmadan sistemin bütünlüğünü programlı olarak kontrol edemeyiz.*/
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// This will join our previous block hash with current block data

	hash := sha256.Sum256(info)
	//The actual hashing algorithm

	b.Hash = hash[:]
}
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	//It is simple subtituing value to block
	block.DeriveHash()
	return block
}

/*Şimdi, her şeyi birbirine bağlamak için bir blok zincirine ihtiyacımız var.
 */
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
func main() {

	chain := InitBlockChain()

	chain.AddBlock("first block")
	chain.AddBlock("second block")
	chain.AddBlock("third block")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}

}
