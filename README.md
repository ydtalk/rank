# 英雄联盟排位数据分析工具

这是一个基于 Riot Games 的 LCU API 的英雄联盟排位数据分析工具的后端服务。该服务提供了获取游戏状态、排位数据和最近比赛记录的功能。

## 功能特性

- 获取当前游戏状态
- 获取排位数据统计
- 获取最近比赛记录

## 系统要求

- Go 1.21 或更高版本
- 英雄联盟客户端必须正在运行
- Windows 操作系统

## 安装

1. 克隆仓库：
```bash
git clone https://github.com/yourusername/lol-rank.git
cd lol-rank
```

2. 安装依赖：
```bash
go mod download
```

## 运行

1. 确保英雄联盟客户端已经启动
2. 运行服务：
```bash
go run main.go
```

服务将在 http://localhost:8080 上启动

## API 端点

### 获取游戏状态
```
GET /api/game-status
```

### 获取排位数据
```
GET /api/ranked-stats
```

### 获取最近比赛记录
```
GET /api/recent-matches
```

## 注意事项

- 该服务需要英雄联盟客户端正在运行才能正常工作
- 服务会自动从英雄联盟客户端的 lockfile 中获取认证信息
- 默认使用端口 2999 连接 LCU API

## 许可证

MIT 