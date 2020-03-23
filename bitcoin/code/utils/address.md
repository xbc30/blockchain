## Address

### P2PKH
> 比特币的原始地址格式，P2PKH代表Pay-to-Pubkey Hash，即支付接收者公钥的哈希值，例如1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2

### P2SH
> P2SH代表对脚本哈希的支付，常用于隔离见证和多重签名，例如3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy

### bech32
> Bech32是本地segwit地址格式，并支持受到了广大的软件和硬件的钱包，每个都以“bc1”开头，并且由于此前缀而比传统或P2SH地址长，例如bc1qf3uwcxaz779nxedw0wry89v9cjh9w2xylnmqc3

地址以”bc1“开头。Bech32编码的地址，是专为SegWit开发的地址格式。Bech32在2017年底在BIP173被定义，该格式的主要特点之一是它不区分大小写（地址中只包含0-9，az），因此在输入时可有效避免混淆且更加易读。由于地址中需要的字符更少，地址使用Base32编码而不是传统的Base58，计算更方便、高效。数据可以更紧密地存储在二维码中。Bech32提供更高的安全性，更好地优化校验和错误检测代码，将出现无效地址的机会降到最低。

Bech32地址本身与SegWit兼容。不需要额外的空间来将SegWit地址放入P2SH地址，因此使用Bech32格式地址，手续费会更低。

Bech32地址比旧的Base58（Base58Check编码用于将比特币中的字节数组编码为人类可编码的字符串）地址有几个优点：

* QR码更小

* 更好地防错

* 更加安全

* 不区分大小写；

* 只由小写字母组成，所以在阅读、输入和理解时更容易