
### TLS SSL SSH
> 证书验证（非对称加密） 数据传输（对称加密）SSH（应用层协议）

### CRT PEM KEY
> 公钥格式三者都行 私钥不能CRT

### base16(hex) base32 base64
> 大小分别为 2/1 8/5 4/3 base编码常用于不可见字符的编码

### csv excel
> 逗号分隔值有很多优势

### 大小端对UTF编码的影响
> UTF-8编码是以1个字节为单位进行处理的，不会受CPU大小端的影响；需要考虑下一位时就地址 + 1。
  
> UTF-16、UTF-32是以2个字节和4个字节为单位进行处理的，即1次读取2个字节或4个字节，这样一来，在存储和网络传输时就要考虑1个单位内2个字节或4个字节之间顺序的问题。

### BOM
> BOM是为UTF-16和UTF-32准备的，用户标记字节序（byte order）

### zlib gzip deflate
> deflate(RFC1951):一种压缩算法，使用LZ77和哈弗曼进行编码；

> zlib(RFC1950):一种格式，是对deflate进行了简单的封装，他也是一个实现库(delphi中有zlib,zlibex)

> gzip(RFC1952):一种格式，也是对deflate进行的封装。
  
> gzip = gzip头 + deflate编码的实际内容 + gzip尾

> zlib = zlib头 + deflate编码的实际内容 + zlib尾

### routine runtime
> 协程 调度器

### channel+select context

