## [Infura](https://infura.io/docs)
> ETH网络中最经常被使用的API服务（DApp/归集/监听链上数据/数据上链）

### [eth_call](https://infura.io/docs/ethereum/json-rpc/eth_call)
> 立即执行新的消息调用，而无需在区块链上创建事务
```
POST https://<network>.infura.io/v3/YOUR-PROJECT-ID
```
**Headers:**
Content-Type: application/json  

**Parameters:**  

* TRANSACTION CALL OBJECT(required)  

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
from | 20 Bytes | NO | The address the transaction is sent from
to | 20 Bytes | YES | The address the transaction is directed to
gas | Integer  | NO | Integer of the gas provided for the transaction execution. eth_call consumes zero gas, but this parameter may be needed by some executions.
gasPrice | Integer | NO | Integer of the gasPrice used for each paid gas
value | Integer | NO | Integer of the value sent with this transaction
data | hash | NO | Hash of the method signature and encoded parameters

* BLOCK PARAMETER(required)  
an integer block number, or the string "latest", "earliest" or "pending"  

**EXAMPLE:**  

```javascript
## JSON-RPC over HTTPS POST
## Replace YOUR-PROJECT-ID with a Project ID from your Infura Dashboard
## You can also replace mainnet with a different supported network
curl https://mainnet.infura.io/v3/YOUR-PROJECT-ID \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_call","params": [{"from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155","to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567","gas": "0x76c0","gasPrice": "0x9184e72a000","value": "0x9184e72a","data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"}, "latest"],"id":1}'
    
## JSON-RPC over websockets
## Replace YOUR-PROJECT-ID with a Project ID from your Infura Dashboard
## You can also replace mainnet with a different supported network
wscat -c wss://mainnet.infura.io/ws/v3/YOUR-PROJECT-ID
>{"jsonrpc":"2.0","method":"eth_call","params": [{"from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155","to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567","gas": "0x76c0","gasPrice": "0x9184e72a000","value": "0x9184e72a","data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"}, "latest"],"id":1}
```

**RESPONSE:**  
```javascript
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0x"
}
```

### [eth_getLogs](https://infura.io/docs/ethereum/json-rpc/eth_getLogs)
> 返回与给定过滤器对象匹配的所有日志的数组，常用于查询合约日志记录
```
POST https://<network>.infura.io/v3/YOUR-PROJECT-ID
```
**Headers:**
Content-Type: application/json  

**Parameters:**  

* FILTER OBJECT(required)  

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
address | 20 Bytes | NO | a string representing the address (20 bytes) to check for balance
fromBlock | Integer | NO | an integer block number, or the string "latest", "earliest" or "pending"
toBlock  | Integer  | NO | an integer block number, or the string "latest", "earliest" or "pending"
topics | Array | NO | Array of 32 Bytes DATA topics. Topics are order-dependent
blockhash |  | NO | restricts the logs returned to the single block with the 32-byte hash blockHash

**EXAMPLE:**  

```javascript
## JSON-RPC over HTTPS POST
## Replace YOUR-PROJECT-ID with a Project ID from your Infura Dashboard
## You can also replace mainnet with a different supported network
curl https://mainnet.infura.io/v3/YOUR-PROJECT-ID \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"blockHash": "0x7c5a35e9cb3e8ae0e221ab470abae9d446c3a5626ce6689fc777dcffcab52c70", "topics":["0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80"]}],"id":1}'

## JSON-RPC over websockets
## Replace YOUR-PROJECT-ID with a Project ID from your Infura Dashboard
## You can also replace mainnet with a different supported network
wscat -c wss://mainnet.infura.io/ws/v3/YOUR-PROJECT-ID
>{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"blockHash": "0x7c5a35e9cb3e8ae0e221ab470abae9d446c3a5626ce6689fc777dcffcab52c70", "topics":["0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80"]}],"id":1}
```

**RESPONSE:**  
```javascript
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": [
    {
      "address": "0x1a94fce7ef36bc90959e206ba569a12afbc91ca1",
      "blockHash": "0x7c5a35e9cb3e8ae0e221ab470abae9d446c3a5626ce6689fc777dcffcab52c70",
      "blockNumber": "0x5c29fb",
      "data": "0x0000000000000000000000003e3310720058c51f0de456e273c626cdd35065700000000000000000000000000000000000000000000000000000000000003185000000000000000000000000000000000000000000000000000000000000318200000000000000000000000000000000000000000000000000000000005c2a23",
      "logIndex": "0x1d",
      "removed": false,
      "topics": [
        "0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80"
      ],
      "transactionHash": "0x3dc91b98249fa9f2c5c37486a2427a3a7825be240c1c84961dfb3063d9c04d50",
      "transactionIndex": "0x1d"
    },
    {
      "address": "0x06012c8cf97bead5deae237070f9587f8e7a266d",
      "blockHash": "0x7c5a35e9cb3e8ae0e221ab470abae9d446c3a5626ce6689fc777dcffcab52c70",
      "blockNumber": "0x5c29fb",
      "data": "0x00000000000000000000000077ea137625739598666ded665953d26b3d8e374400000000000000000000000000000000000000000000000000000000000749ff00000000000000000000000000000000000000000000000000000000000a749d00000000000000000000000000000000000000000000000000000000005c2a0f",
      "logIndex": "0x57",
      "removed": false,
      "topics": [
        "0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80"
      ],
      "transactionHash": "0x788b1442414cb9c9a36dba2abe250763161a6f6395788a2e808f1b34e92beec1",
      "transactionIndex": "0x54"
    }
  ]
}
```

### [eth_getStorageAt](https://infura.io/docs/ethereum/json-rpc/eth_getStorageAt)
> 从给定合约地址的存储位置返回值，常用于查询当前合约的状态数据
```
POST https://<network>.infura.io/v3/YOUR-PROJECT-ID
```
**Headers:**
Content-Type: application/json  

**Parameters:**  

* ADDRESS(required)  
a string representing the address (20 bytes) of the storage

* STORAGE POSITION(required)  
a hex code of the position in the storage

* BLOCK PARAMETER(required)  
an integer block number, or the string "latest", "earliest" or "pending"  

**EXAMPLE:**  

```javascript
## JSON-RPC over HTTPS POST
## Replace YOUR-PROJECT-ID with a Project ID from your Infura Dashboard
## You can also replace mainnet with a different supported network
curl https://mainnet.infura.io/v3/YOUR-PROJECT-ID \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_getStorageAt","params": ["0x295a70b2de5e3953354a6a8344e616ed314d7251", "0x6661e9d6d8b923d5bbaab1b96e1dd51ff6ea2a93520fdc9eb75d059238b8c5e9", "0x65a8db"],"id":1}'

## JSON-RPC over websockets
## Replace YOUR-PROJECT-ID with a Project ID from your Infura Dashboard
## You can also replace mainnet with a different supported network
wscat -c wss://mainnet.infura.io/ws/v3/YOUR-PROJECT-ID
>{"jsonrpc":"2.0","method":"eth_getStorageAt","params": ["0x295a70b2de5e3953354a6a8344e616ed314d7251", "0x6661e9d6d8b923d5bbaab1b96e1dd51ff6ea2a93520fdc9eb75d059238b8c5e9", "0x65a8db"],"id":1}
```

**RESPONSE:**  
```javascript
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0x0000000000000000000000000000000000000000000000000000000000000000"
}
```
### [eth_sendRawTransaction](https://infura.io/docs/ethereum/json-rpc/eth_sendRawTransaction)
> 提交预签名的交易，以广播到以太坊网络
```
POST https://<network>.infura.io/v3/YOUR-PROJECT-ID
```
**Headers:**
Content-Type: application/json  

**Parameters:**  

* TRANSACTION DATA(required)  
The signed transaction data. 

**EXAMPLE:**  

```javascript
## JSON-RPC over HTTPS POST
## Replace YOUR-PROJECT-ID with a Project ID from your Infura Dashboard
## You can also replace mainnet with a different supported network
curl https://mainnet.infura.io/v3/YOUR-PROJECT-ID \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"],"id":1}'

## JSON-RPC over websockets
## Replace YOUR-PROJECT-ID with a Project ID from your Infura Dashboard
## You can also replace mainnet with a different supported network
wscat -c wss://mainnet.infura.io/ws/v3/YOUR-PROJECT-ID
>{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"],"id":1}
```

**RESPONSE:**  
```javascript
{
  "id":1,
  "jsonrpc": "2.0",
  "result": "0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331"
}
```