### [Compound](https://compound.finance/)

**组成部分:**

```javascript
Comptroller + RateModel + PriceOracle + UniToken = Compound
Comptroller(强制用户保持足够的抵押物余额)
RateModel(根据给定市场的当前利用率在算法上确定利率)
PriceOracle(代币价格预言机-平均来自三个报告者的价格数据)
UniToken(defi上流通的价值token)

rDai -- 代币本息所有权分离 
```