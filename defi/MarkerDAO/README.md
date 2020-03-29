## [MakerDAO](https://docs.makerdao.com/)

### 模块

#### Core Module
> 核心模块对系统至关重要，它包含了Maker协议的整个状态，并在系统处于预期的正常运行状态时控制系统的中心机制

#### Collateral Module
> 抵押品模块是为每一个新的ilk(抵押品类型)添加到Vat。它包含一个特定抵押品类型的所有适配器和拍卖合约

#### Dai Module
> DAI的起源是为了表示核心系统认为与内部债务单位价值相等的任何标记。因此，DAI模块包含DAI令牌契约和所有适配DaiJoin的适配器

#### System Stabilizer Module
> 系统稳定器模块的作用是当系统的稳定性受到威胁，抵押品的价值低于清算水平(由治理决定)时，对系统进行修正。系统稳定器模块通过参与债务和盈余拍卖，激励拍卖管理员(外部参与者)介入并推动系统回到安全状态(系统平衡)，进而通过这样做赚取利润

#### Oracle Module
> 每个抵押类型都部署了一个oracle模块，将对应抵押类型的价格数据提供给增值税。Oracle模块引入了地址的白名单，这使得它们可以广播价格更新的消息，然后将这些消息传入一个median，然后再将其拉入OSM

#### MKR Module
> MKR模块包含MKR令牌，这是一个已部署的Ds-Token契约。它是一个提供标准ERC20令牌接口的ERC20令牌。它还包含用于销毁和授权生成MKR的逻辑

#### Governance Module
> 治理模块包含促进MKR投票、提案执行和Maker协议的投票安全性的契约

#### Rates Module
> MCD系统的一个基本特征是积累保险库债务余额的稳定费用，以及Dai存款利率(DSR)存款的利息。用于执行这些积累功能的机制受到一个重要的约束:积累必须是关于保险库数量和DSR存款数量的恒定时间操作。

#### Proxy Module
> 创建代理模块是为了方便用户/开发人员与Maker协议进行交互。它包含保险库管理和制造商治理所需的契约接口、代理和别名

#### Emergency Shutdown Module
> 紧急关闭模块(ESM)是一个能够调用End.cage()来触发关闭Maker协议的合约

### 思维导图

![image](../../pic/MCD-System-2.0.png)

### 相关科普文章

* [深入浅出理解 MakerDAO: 不止于稳定币](https://www.jianshu.com/p/8dd963f39795)