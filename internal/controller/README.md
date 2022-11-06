# controller层
## 1.将用户传入的json(req)传给service层 
## 2.在logic进入service层的依赖注入, 进行业务逻辑的操作
## 3.其中在logic层调用dao层与数据库进行通信获取数据,并返回获取到的数据