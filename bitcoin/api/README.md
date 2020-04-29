
### Single Block
> 请求指定区块hash对应的数据
```
curl https://blockchain.info/rawblock/$block_hash
```

```
// curl https://blockchain.info/rawblock/0000000000000bae09a7a393a8acded75aa67e46cb81f7acaa5ad94f9eacd103

{
    "hash":"0000000000000bae09a7a393a8acded75aa67e46cb81f7acaa5ad94f9eacd103",
    "ver":1,
    "prev_block":"00000000000007d0f98d9edca880a6c124e25095712df8952e0439ac7409738a",
    "mrkl_root":"935aa0ed2e29a4b81e0c995c39e06995ecce7ddbebb26ed32d550a72e8200bf5",
    "time":1322131230,
    "bits":437129626,
    "nonce":2964215930,
    "n_tx":22,
    "size":9195,
    "block_index":818044,
    "main_chain":true,
    "height":154595,
    "received_time":1322131301,
    "relayed_by":"108.60.208.156",
    "tx":[--Array of Transactions--]
}
```

### Single Transaction

> 请求指定交易hash对应的数据

```
curl https://blockchain.info/rawtx/$tx_hash
```

```
// curl https://blockchain.info/rawtx/b6f6991d03df0e2e04dafffcd6bc418aac66049e2cd74b80f14ac86db1e3f0da

{
   "ver":1,
   "inputs":[
      {
         "sequence":4294967295,
         "witness":"",
         "prev_out":{
            "spent":true,
            "spending_outpoints":[
               {
                  "tx_index":0,
                  "n":0
               }
            ],
            "tx_index":0,
            "type":0,
            "addr":"1FwYmGEjXhMtxpWDpUXwLx7ndLNfFQncKq",
            "value":100000000,
            "n":2,
            "script":"76a914a3e2bcc9a5f776112497a32b05f4b9e5b2405ed988ac"
         },
         "script":"48304502210098a2851420e4daba656fd79cb60cb565bd7218b6b117fda9a512ffbf17f8f178022005c61f31fef3ce3f906eb672e05b65f506045a65a80431b5eaf28e0999266993014104f0f86fa57c424deb160d0fc7693f13fce5ed6542c29483c51953e4fa87ebf247487ed79b1ddcf3de66b182217fcaf3fcef3fcb44737eb93b1fcb8927ebecea26"
      }
   ],
   "weight":1032,
   "block_height":154598,
   "relayed_by":"0.0.0.0",
   "out":[
      {
         "spent":true,
         "spending_outpoints":[
            {
               "tx_index":0,
               "n":3
            }
         ],
         "tx_index":0,
         "type":0,
         "addr":"14pDqB95GWLWCjFxM4t96H2kXH7QMKSsgG",
         "value":98000000,
         "n":0,
         "script":"76a91429d6a3540acfa0a950bef2bfdc75cd51c24390fd88ac"
      },
      {
         "spent":true,
         "spending_outpoints":[
            {
               "tx_index":0,
               "n":5
            }
         ],
         "tx_index":0,
         "type":0,
         "addr":"13AMPUTTwryLGX3nrMvumaerSqNXkL3gEV",
         "value":2000000,
         "n":1,
         "script":"76a91417b5038a413f5c5ee288caa64cfab35a0c01914e88ac"
      }
   ],
   "lock_time":0,
   "size":258,
   "block_index":0,
   "time":1322135154,
   "tx_index":0,
   "vin_sz":1,
   "hash":"b6f6991d03df0e2e04dafffcd6bc418aac66049e2cd74b80f14ac86db1e3f0da",
   "vout_sz":2
}
```

### Block Height

> 请求指定区块高度的数据(数据量比较大)

```
curl https://blockchain.info/block-height/$block_height?format=json
```

```
// curl https://blockchain.info/block-height/600000?format=json
...
```

### Single Address

> 请求单个地址对应区块数据

```
curl https://blockchain.info/rawaddr/$bitcoin_address
```

```
// curl https://blockchain.info/rawaddr/1AJbsFZ64EpEfS5UAjAfcUG8pH8Jn3rn1F
{
    "hash160":"660d4ef3a743e3e696ad990364e555c271ad504b",
    "address":"1AJbsFZ64EpEfS5UAjAfcUG8pH8Jn3rn1F",
    "n_tx":17,
    "n_unredeemed":2,
    "total_received":1031350000,
    "total_sent":931250000,
    "final_balance":100100000,
    "txs":[--Array of Transactions--]
}
```

### Multi Address

> 请求多个地址对应区块数据

```
curl https://blockchain.info/multiaddr?active=$address|$address
```

```
{
    "addresses":[

    {
        "hash160":"641ad5051edd97029a003fe9efb29359fcee409d",
        "address":"1A8JiWcwvpY7tAopUkSnGuEYHmzGYfZPiq",
        "n_tx":4,
        "total_received":1401000000,
        "total_sent":1000000,
        "final_balance":1400000000
    },

    {
        "hash160":"ddbeb8b1a5d54975ee5779cf64573081a89710e5",
        "address":"1MDUoxL1bGvMxhuoDYx6i11ePytECAk9QK",
        "n_tx":0,
        "total_received":0,
        "total_sent":0,
        "final_balance":0
    },

    "txs":[--Latest 50 Transactions--]
}
```

### Unspent outputs

> 请求查询地址未花费的交易输出

```
curl https://blockchain.info/unspent?active=$address
```

```
{
    "unspent_outputs":[
        {
            "tx_age":"1322659106",
            "tx_hash":"e6452a2cb71aa864aaa959e647e7a4726a22e640560f199f79b56b5502114c37",
            "tx_index":"12790219",
            "tx_output_n":"0",
            "script":"76a914641ad5051edd97029a003fe9efb29359fcee409d88ac", (Hex encoded)
            "value":"5000661330"
        }
    ]
}
```

### Balance

> 请求查询余额数据

```
curl https://blockchain.info/balance?active=$address
```

```
{
    "1MDUoxL1bGvMxhuoDYx6i11ePytECAk9QK": {
        "final_balance": 0,
        "n_tx": 0,
        "total_received": 0
    },
    "15EW3AMRm2yP6LEF5YKKLYwvphy3DmMqN6": {
        "final_balance": 0,
        "n_tx": 2,
        "total_received": 310630609
    }
}
```