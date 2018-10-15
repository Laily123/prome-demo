## prometheus 监控 demo

### 数据上报形式
监控数据上报到 prometheus 分两种形式。
1. 服务开启一个 http 服务，暴露出一个 html 页面，由 prometheus 定时抓取
2. 服务将需要监控的数据以 http 的形式发送给 pushgateway。pushgateway 在提供给 prometheus。

### 暴露 html 页面方式
1. 启动服务时向 prometheus 注册一个监控目标(这一步也可以优化成根据 etcd 自动更新)  
注册 api
```javascript
    // POST /add-target
    // request body
    {
	"addr":"", // 暴露的页面 host，页面路径默认是 /metrics
    "job_name": "service", // prometheus 抓取的 job 名称，根据 prometheus 配置获得,(可以理解为顶级分类)
	"target_id": "120.79.193.99#8081#2", // 抓取目标id，用于唯一区分抓取目标，可以用 hostname+端口的形式
	"group": "" // 抓取目标分类，可以为空
    "instance": "xxx", // 服务实例名称，也可以用 hostname
    }
```
2. 构建 metrics。  
prometheus 提供了四种 metrics。介绍 https://prometheus.io/docs/concepts/metric_types/。  
使用示例：  
https://github.com/prometheus/client_golang/blob/master/examples/random/main.go  
或者该项目的 export.go  

### 数据发送给 pushgateway
1. 构造监控数据  
文档 https://github.com/prometheus/pushgateway  
或者见本项目 push.go  
url 里提供 job_name, group, instance 等  
body 里提供数据   
2. 发送监控数据到 pushgateway 地址 
