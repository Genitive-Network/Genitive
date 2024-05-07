## Genitive Backend 


后端端口：8006

`Mint`: 铸造代币

`Burn`: 销毁代币


接口：

/api/mint 
- request pamram
  - address string [地址]
  - amount decimal [金额]
- response result
  - status bool [true 成功 false 失败 ]

/api/burn

方法：
POST
