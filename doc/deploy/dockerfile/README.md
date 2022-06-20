基于 `docker` 的 `go` 本地开发环境。

## 使用

### 1. 按需修改 .env 配置

~~~env
# 设置时区
TZ=Asia/Shanghai

# 宿主机上Mysql/Reids/MQ数据存放的目录
DATA_PATH_HOST=E:\dockerdata


~~~

### 2.启动服务

- 启动全部服务

```bash
docker-compose up -d
```

- 按需启动部分服务

```bash
docker-compose up -d etcd golang mysql redis
```

