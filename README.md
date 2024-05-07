## Genitive Backend 


后端端口：8006

`Mint`: 铸造代币

`Burn`: 销毁代币


接口：

/api/mint 
- request pamram
  - address string [地址]
  - amount bigint [金额 btc 8位对齐]
- response result
  - status bool [true 成功 false 失败 ]

成功返回状态码200， 成功

/api/burn

方法：
POST
