# 🌾 智行云岭 - 本地开发指南

## 项目简介

智行云岭是一个基于 Hyperledger Fabric 的农产品区块链溯源系统，包含6大智能合约，覆盖从产地到餐桌的全链路追踪。

## 本地开发环境要求

- Go 1.23+
- Node.js 18+
- npm 9+
- Docker & Docker Compose

## 快速开始

### 1. 设置脚本权限

```bash
cd fabric-realty
find . -name "*.sh" -exec chmod +x {} \;
```

### 2. 启动区块链网络

```bash
cd network
./install.sh
```

### 3. 安装并部署链码

进入 chaincode 目录，确保依赖正确：

```bash
cd chaincode
go build
```

然后将链码安装部署到 Fabric 网络中。

### 4. 启动后端服务

```bash
cd application/server
go run main.go
```

后端服务运行在 8888 端口。

### 5. 启动前端服务

```bash
cd application/web
npm install
npm run dev
```

前端开发服务器运行在 5173 端口。

### 6. 访问系统

http://localhost:5173

## 系统组织结构

| 组织 | 路由路径 | 主要功能 |
|------|----------|----------|
| 🌿 农业局 | `/agriculture` | 农场注册、产品管理、溯源查询、合约浏览 |
| 🔬 检测认证中心 | `/inspection` | 创建检测报告、认证证书管理 |
| 📦 供应链平台 | `/supply-chain` | 物流管理、采购订单、产品查询、溯源 |

## 智能合约接口

使用 `ContractName:MethodName` 格式调用：

- `FarmContract:RegisterFarm` - 注册农场
- `FarmContract:QueryFarm` - 查询农场
- `ProductContract:RegisterProduct` - 注册产品
- `ProductContract:RecordHarvest` - 记录采收
- `InspectionContract:CreateInspection` - 创建检测报告
- `LogisticsContract:CreateLogisticsRecord` - 创建物流记录
- `TraceContract:GetFullTraceability` - 获取完整溯源信息
- `TradeContract:CreatePurchaseOrder` - 创建采购订单

## 注意事项

1. 修改后端代码后需重启 `go run main.go`
2. 前端代码支持 Vite 热更新
3. 链码更新需重新部署到 Fabric 网络
4. 首次部署需确保 CouchDB 索引已创建（用于富查询）
