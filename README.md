# learn_hertz

本项目由[hertz](https://github.com/cloudwego/hertz)框架驱动

- 支持pg、mysql、sqlite3数据库切换
- 支持yaml配置文件

## 初始化

初始化账号为: admin/admin123456(什么配置文件都不改的情况下)

如需要修改在配置文件中修改

### 初始化管理员账号

```yaml
admin:
  username: admin
  password: admin123456
```

### 初始化数据库

默认不配置的情况下位sqlite3数据库

需要pg或者mysql的话修改以下配置

```yaml
db:
  type: postgres
  host: xxx
  port: 5432
  user: postgres
  password: xxx
  database: test
```

### 修改服务端口

默认端口为8888

```yaml
server:
  port: 8888
```