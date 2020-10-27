# letsgo
Golang net/http 摸索，简单的封装等。

> 尝试摸索net/http的使用的项目，任在继续摸索中~。
>
> 小白同行，大佬到了不妨指点一二。

###### 做了如下工作

- 尝试实现了html的读取和简单渲染。
- 尝试实现静态文件的代理。
- 尝试数据库的访问与查询。
- 尝试将单个Handler 封装成结构体，根据r.Method分发至结构体的Get，Post，Put，Delete函数。
- 尝试实现中间件，避免装饰模式编写的不便。
- 尝试编写中间件 - Session。(net/http未提供)
- 尝试实现中间件 - DynamicRouter 动态路由匹配。
