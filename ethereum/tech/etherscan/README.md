### Etherscan

#### Contract Address
> 通过合约创建者和这笔交易的nonce，生成address
```go
// CreateAddress creates an ethereum address given the bytes and the nonce
func CreateAddress(b common.Address, nonce uint64) common.Address {
    data, _ := rlp.EncodeToBytes([]interface{}{b, nonce})
    return common.BytesToAddress(Keccak256(data)[12:])
}
```

#### Contract Source Code

* SOL2UML

* OutLine

#### EVM bytecode decompiler
> EVM 字节码反编译器

#### Swarm Source
> 部署合同后，源代码随后存储在Swarm中(分散式数据存储和分发),"bzzr地址"是合约的文件ID，编译器将元数据文件的Swarm哈希附加到每个合约的字节码末尾，以便您可以以经过身份验证的方式检索文件，而无需求助于集中式数据提供程序
