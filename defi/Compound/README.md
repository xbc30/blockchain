### [Compound](https://compound.finance/)

**组成部分:**

```javascript
Comptroller + RateModel + PriceOracle + UniToken = Compound
Comptroller(强制用户保持足够的抵押物余额)
RateModel(根据给定市场的当前利用率在算法上确定利率)
PriceOracle(代币价格预言机-平均来自三个报告者的价格数据)
cToken(defi上流通的价值token)

rDai -- 代币本息所有权分离 
```

**术语列表**

##### cToken / cEther

##### CDP(抵押债仓)

##### 放贷
> 将基础资产转进/转出市场

* Mint(抵押eth铸dai币以获得cToken/cEther)

* Redeem(赎回抵押)

* Redeem Underlying(赎回部分抵押)

##### 借贷

* Borrow(借款)

* Repay Borrow(偿还借款)

* Repay Borrow Behalf(偿还部分借款)

##### 清算

* Liquidate Borrow(清算借款)

* Liquidate 清算激励 / 清算罚金

##### 资金

* Total Borrow(总借贷)

* Total Supply(总放贷)

* Total Reserves(总准备金)

* Get Cash(合约所剩准备金)

##### 基本模型

* Exchange Rate(市场汇率)

> exchangeRate = (getCash() + totalBorrows() - totalReserves()) / totalSupply()

> 市场汇率 = ( 合约所剩储备金 + 总借贷 - 总准备金 ) / 总放贷

* Borrow Rate / Supply Rate (借贷利率 / 放贷利率)

> BorrowingInterestRate = 2.5% + U * 20%

> U(利用率) = Borrow / (Borrow + Cash)

* Reserve Factor(准备金率)
> Reserve Factor = Reserve / totalSupply

* Collateral Factor(抵押率)
> 抵押物和乘以抵押因子即等于用户的可借贷数额

##### 市场动态模型

1.从初始利率开始，每次发生交易时，资产的利率指数都会更新，以复利利息，某期间的利率（以r * t表示）使用逐块利率计算得出：

Index(n) = Index(n-1) * (r * t)

2.市场的未偿还借款总额会更新为包括自上一个指数至今的应计利息：

totalBorrowBalance(n) = totalBorrowBalance * (1 + r * t) 

3.一部分应计利息被保留（留作储备），由reserveFactor确定，范围为0到1：

reserve(n) = reserve(n-1) + totalBorrowBalance(n-1) * (r * t * reserveFactor)