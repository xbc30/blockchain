## Address

## 校验和地址

```javascript
const createKeccakHash = require('keccak')

function toChecksumAddress (address) {
  address = address.toLowerCase().replace('0x', '')
  var hash = createKeccakHash('keccak256').update(address).digest('hex')
  console.log(hash)
  // 9210d83a63ada9beb95a9a2df4e8dd7164c2ed5b71ebaff87a0691cb4ec3f345
  var ret = '0x'

  for (var i = 0; i < address.length; i++) {
    if (parseInt(hash[i], 16) >= 8) {
      ret += address[i].toUpperCase()
    } else {
      ret += address[i]
    }
  }
  console.log(ret)
  // 0x941aFAeBe68977C31Fe0A4cF3e65F7c7c8Cc3B11
  return ret
}

toChecksumAddress('0x941afaebe68977c31fe0a4cf3e65f7c7c8cc3b11')
```