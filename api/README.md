
``
// path: 接口地址的定义
// in：参数来源(query,path,不写是body)
// - path推荐写法：(path:"/article/{id}")
// - query是请?连接&号拼接，如（path:"/article?name={name}&age={age}"）
// method:请求方法
// json：属性名,字段名
// tags: 接口分类标签（swagger文档生成时的分组）
// summary: 接口描述，最好理解为接口标题
// dc：字段描述，描述这个字段是什么意思
// v: 数据校验()
// - v:"required|length:6,30|same:password1#请确认密码|密码长度不够|两次密码不一致"

``