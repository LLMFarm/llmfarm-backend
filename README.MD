# 平台简介
  
大模型农场LLMFarm平台是全球领先的大模型中间层开发架构，帮助企业快速对接私有数据和各类海内外大模型，并可视化无代码化构建各类 AI First 应用，低成本落地各行业各类AI场景。本项目由宜创科技开发。
  
## 核心价值
  
### 可视化构建Prompt Chain
  
LLMFarm平台提供了可视化构建Prompt Chain的功能，使用户及企业能够更加直观地设计和搭建AI应用场景。通过简单的拖拽和连接，企业可以自定义不同的对话流程、大模型调用逻辑，对接企业各类私有数据，提供不同的 API 接口，实现个性化的AI场景应用。这种可视化构建的方式，无需程序员，10 倍降低开发成本，使更多的用户及企业能够参与到AI应用场景的构建中。
  
### 可扩展能力
  
LLM Farm平台支持扩充对接各类国内外大模型，包括开源大模型，如 海外LLaMa 系列，国内ChatGLM ，Baichuan系列等，也支持对接文生图大模型如 Stable Diffusion 等，也支持图生文模型 VisualGLM 等。同时也对接了 ChatGPT，百度文心一言，ChatGLM等基础模型 API。支持全类型 SQL，NOSQL，VectorDB等数据库的支持。支持企业内外部各类 API 的集成调用，各类云服务市场的 API 集成调用。
  
### 行业多种类型的AI应用场景
  
无论是客服服务、知识库、SQL查询、技术支持还是智能助手等，企业都可以根据自身需求，灵活地构建和定制不同类型的AI应用场景。这种灵活性使企业能够根据自身行业和业务特点，个性化地应用AI技术。
  
## 官方网站
  
[https:www.llmfarm.com](https:www.llmfarm.com)

## 文档地址
  
[LLMFarm新手指南](https://lxjcwfwipaa.feishu.cn/wiki/JiIBwo42wi53VEk4av4c8qRmnDc?from=from_copylink)
  
## 演示图
![链接](https://hexyun.oss-cn-beijing.aliyuncs.com/hwy/20230811164552.jpg)
![链接](https://hexyun.oss-cn-beijing.aliyuncs.com/hwy/20230811164500.jpg)
  
## 版本说明：
  
- 开源版

    本项目提供的开源版本。目前仅包含提示词节点、GPT节点、输入节点。Chain管理、Bot管理、对话功能。后续版本会持续迭代更新功能。
- SaaS版

    SaaS版作为本产品功能最全，更新最及时的版本。提供包含多种模型节点、向量数据库相关节点、逻辑处理相关节点。提供开箱即用的机器人模板市场，LLMFarm平台已经提供500+的应用场景模板，用户只需下载即可一键使用，并可以进行二次开发和调整。
    目前SaaS版对个人和企业开放注册。
地址：[LLMFarm](http://chat.llmfarm.com)
- 企业版
    企业版包含SaaS版的全部功能。并根据SaaS版的更新进度进行更新。企业版可以进行企业私有化部署。保证企业的数据安全。

## 安装流程

### 依赖
Go 1.19+
[GoFrame](https://https://goframe.org/pages/viewpage.action?pageId=1114399) 2.0+ 
  
### 配置文件
manifest/config/config.yaml.bak 复制改为manifest/config/config.yaml
```c  
server:
address: ":8000"  #运行端口
openapiPath: "/api.json"
swaggerPath: "/swagger"
  
logger:
level : "all"
stdout: true
  
  
database:
logger:
level: "all"
stdout: true
default:
link: "mysql:root:123456@tcp(127.0.0.1:3306)/llmfarm?parseTime=true&loc=Local"  #数据库连接配置
debug: true
charset: "utf8mb4"
api:
chainApiDomain: "http://127.0.0.1:8080"  #llmfarm-chain的端口号
  
versionType: "individual"
fileStorageType: "local"  #存储方式 local为本地存储 oss为阿里云oss存储 obs为华为云存储
  
ossconfig:
accessKeyId: ""
accessKeySecret: ""
bucket: ""
type: "aliyun"
isHttps: false
endpoint: ""
domain: ""
obsconfig:
accessKeyId: ""
accessKeySecret: ""
bucket: ""
type: "huaweiyun"
endpoint: ""
domain: ""
localconfig:
path: "/Users/Downloads"  # 本地存储绝对路径
domain: "localhost:8000"  # 本地存储访问路径 
```

### 导入mysql初始化表

项目数据库文件`resource/data/llmfarm.sql`创建数据库

### 运行项目

在项目目录下执行
```c
go mod tidy
go mod vendor
go run main.go
```

### 运行其他端

项目为前后端分离，llmfarm-front地址：

github地址：[https://github.com/LLMFarm/llmfarm-front](https://github.com/LLMFarm/llmfarm-front)

gitee地址：[https://gitee.com/llmfarm/llmfarm-front](https://gitee.com/llmfarm/llmfarm-front)

llmfarm-chain地址：

github地址：[https://github.com/LLMFarm/llmfarm-chain](https://github.com/LLMFarm/llmfarm-chain)

gitee地址：[https://gitee.com/llmfarm/llmfarm-chain](https://gitee.com/llmfarm/llmfarm-chain)