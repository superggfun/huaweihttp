# huaweihttp
华为云函数多IP爬虫

## 部署在华为云
1. [进入华为云函数](https://console.huaweicloud.com/functiongraph)
2. 创建函数
3. V2版本，事件函数，Go 1.x
4. 上传ZIP文件
5. 设置，执行超时时间改成90s
6. 设置，并发。
7. 设置，触发器	API 网关服务(APIG)，安全认证：None
8. 绑定API网关独立域名，SSL
9. 域名后面加/实例名称/url=?https://www.baidu.com/
