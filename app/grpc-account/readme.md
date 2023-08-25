账户中心
- 三方扫码登录，保留必要信息
- 颁发双令牌：签发token 刷新token，前端做token刷新逻辑
- 维护casbin接口权限规则
- 账户信息


设计:
- 把 accessToken 与 refreshToken 存在 localStorage 中。既然要 JWT ，那就贯彻到底，给服务端彻底减负。

token里 包含用户id，这些不变信息。其他还是需要查次库

docs里业务流程设计图