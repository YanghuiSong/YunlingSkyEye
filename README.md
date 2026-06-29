# 👁️ 云岭天眼 - 高原特色农产品溯源系统

<div align="center">

**基于 Hyperledger Fabric 联盟链的农产品全程溯源平台**

![Version](https://img.shields.io/badge/version-1.0.0-blue)
![Fabric](https://img.shields.io/badge/Hyperledger_Fabric-2.5.10-green)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8)
![Vue](https://img.shields.io/badge/Vue-3.3-4FC08D)
![License](https://img.shields.io/badge/License-Apache_2.0-yellow)

**[快速开始](#-快速部署)** · **[系统架构](#-系统架构)** · **[API文档](#-api接口)** · **[常见问题](#-常见问题与解决方案)**

</div>

---

## 📋 项目简介

**云岭天眼** 聚焦云南高原特色农业，深度融合区块链技术，实现从**产地种植 → 检测认证 → 冷链物流 → 销售交易 → 扫码溯源**的全链路数据上链与可信追溯。

系统采用 Hyperledger Fabric 联盟链架构，由 **农业局**、**检测认证中心**、**供应链平台** 三个组织共同维护，部署 **6 大智能合约**，覆盖农产品溯源全生命周期。

### 🌄 覆盖云南高原特色农产品（49种）

| 类别 | 代表产品 |
|------|----------|
| 🍵 **茶叶** | 普洱茶、古树滇红、景迈山茶 |
| ☕ **咖啡** | 云南小粒咖啡、铁毕卡咖啡 |
| 🍄 **野生菌** | 香格里拉松茸、楚雄牛肝菌、干巴菌、松露 |
| 🥩 **肉类** | 宣威火腿、诺邓火腿、丽江腊排骨 |
| 🌿 **中药材** | 文山三七、昭通天麻、铁皮石斛 |
| 🍎 **水果** | 昭通苹果、蒙自石榴、丽江雪桃、冰糖橙 |
| 🌾 **粮食** | 八宝贡米、梯田红米、紫米 |
| 🌹 **花卉/其他** | 鲜切花、乳扇、鲜花饼、核桃、咖啡 |

---

## ✨ 核心特性

| 特性 | 说明 |
|------|------|
| **📦 多合约架构** | 6大智能合约分工协作（Farm/Product/Inspection/Logistics/Trace/Trade） |
| **🔗 区块链存证** | 所有数据上链存储，区块哈希链确保不可篡改 |
| **🏛️ 联盟治理** | 三组织分级权限，MSP 证书身份认证，操作安全可控 |
| **🔍 全程溯源** | 支持完整溯源链查询（产品→产地→检测→物流→交易） |
| **🧪 质量检测** | 农残、重金属、微生物多维度检测认证 |
| **🚚 冷链监控** | 物流环节温湿度数据上链，保障高原农产品品质 |
| **👁️ 链上可视化** | 独立的深色主题可视化页面，实时展示区块链数据结构 |

---

## 🏗️ 系统架构

### 联盟组织架构

```
┌─────────────────────────────────────────────────────────────┐
│                    🔧 Orderer 集群 (Raft)                    │
│          orderer1 · orderer2 · orderer3 (3节点CFT)          │
└─────────────────────────────────────────────────────────────┘
                              │
        ┌─────────────────────┼─────────────────────┐
        │                     │                     │
   ┌────▼────┐          ┌────▼────┐          ┌────▼────┐
   │  🌿     │          │  🔬     │          │  📦     │
   │ 农业局  │          │检测中心 │          │供应链平台│
   │ Org1MSP │          │ Org2MSP │          │ Org3MSP │
   │         │          │         │          │         │
   │ peer0   │          │ peer0   │          │ peer0   │
   │ peer1   │          │ peer1   │          │ peer1   │
   └────┬────┘          └────┬────┘          └────┬────┘
        │                     │                     │
        └─────────────────────┼─────────────────────┘
                              │
                   ┌──────────▼──────────┐
                   │   📡 mychannel 通道  │
                   │  所有组织数据同步    │
                   └─────────────────────┘
```

### 六大智能合约

| 合约 | 函数 | 说明 | 调用权限 |
|------|------|------|----------|
| **FarmContract** | `RegisterFarm` / `UpdateFarm` / `SuspendFarm` / `QueryFarm` / `QueryFarmList` | 高原农场注册与管理 | Org1MSP |
| **ProductContract** | `RegisterProduct` / `UpdateProductStatus` / `RecordHarvest` / `QueryProduct` / `QueryProductList` / `QueryProductsByFarm` | 农产品注册与生命周期管理 | Org1/Org3 |
| **InspectionContract** | `CreateInspection` / `FailInspection` / `QueryInspection` / `QueryInspectionList` / `QueryCertificates` | 质量检测、认证评级、证书颁发 | Org2MSP |
| **LogisticsContract** | `CreateLogisticsRecord` / `UpdateLogisticsStatus` / `QueryLogisticsRecord` / `QueryLogisticsList` / `QueryLogisticsByProduct` | 高原冷链物流追踪 | Org3MSP |
| **TraceContract** | `GetFullTracebility` / `VerifyProductByBatch` | 全链路溯源查询、批次验证 | 所有组织 |
| **TradeContract** | `CreatePurchaseOrder` / `UpdateOrderStatus` / `QueryPurchaseOrder` / `QueryPurchaseOrderList` | 采购订单管理 | Org3MSP |

### 溯源全流程

```
🌱 产地种植 ──→ 🧪 检测认证 ──→ 🚚 冷链物流 ──→ 🛒 销售交易 ──→ 🔍 扫码溯源
     │                │               │               │               │
     ▼                ▼               ▼               ▼               ▼
  FarmContract   InspectionContract  LogisticsContract  TradeContract  TraceContract
```

---

## 🚀 快速部署

### 环境要求

| 工具 | 最低版本 | 验证命令 |
|------|----------|----------|
| Docker | 24.0.7 ⚠️ | `docker --version` |
| Docker Compose | 2.x+ | `docker compose version` |
| Go | 1.23+ | `go version` |
| Node.js | 18+ | `node --version` |
| npm | 9+ | `npm --version` |

> ⚠️ **重要：Docker 版本兼容性**
> Docker 29.x **不兼容** Fabric 2.5.10，链码安装时会报 `write unix @->/run/docker.sock: write: broken pipe`。
> 请使用 Docker 24.0.7：
> ```bash
> sudo apt-get install -y docker.io=24.0.7-0ubuntu4
> ```

### 部署步骤

#### 第1步：设置脚本权限

```bash
cd yunling-sky-eye
find . -name "*.sh" -exec chmod +x {} \;
```

#### 第2步：启动区块链网络

```bash
cd network
./install.sh
```

脚本会自动执行以下操作（共16步）：

| 步骤 | 操作 | 说明 |
|------|------|------|
| 1-3 | 环境检查 + 工具容器启动 | 检查 Docker，启动 CLI 容器 |
| 4 | 生成证书和密钥 | `cryptogen` 生成 MSP 材料 |
| 5-7 | 创世区块 + 通道配置 + 锚节点 | `configtxgen` 生成配置 |
| 8 | 启动所有节点 | 3 Orderer + 6 Peer |
| 9-11 | 创建通道 + 节点加入 + 锚节点更新 | 所有节点加入 `mychannel` |
| 12-13 | 打包并安装链码 | 安装云岭天眼链码 |
| 14-15 | 各组织批准 + 提交链码 | 3组织签名提交 |
| 16 | 初始化并验证 | 验证 `Hello` 返回 |

> ⏳ 整个过程约 5-10 分钟（取决于网络状况和 Docker 镜像拉取速度）。

#### 第3步：启动后端服务

```bash
cd application/server
go run main.go
```

启动成功输出：
```
云岭天眼高原特色农产品溯源系统启动中，监听端口 :8888
已保存组织[org1]的区块[0]
已保存组织[org2]的区块[0]
已保存组织[org3]的区块[0]
...
```

#### 第4步：启动前端服务

```bash
cd application/web
npm install
npm run dev
```

启动成功输出：
```
VITE v4.5.5  ready in 312 ms
➜  Local:   http://localhost:5173/
```

#### 第5步：访问系统

打开浏览器访问 **http://localhost:5173**

---

## 🎯 使用指南

### 各组织入口

| 组织 | 页面路径 | 主要功能 |
|------|----------|----------|
| 🌿 **农业局** | `/agriculture` | 注册高原农场 → 注册特色产品 → 记录采收 → 溯源查询 → 合约浏览 |
| 🔬 **检测中心** | `/inspection` | 创建检测报告 → 录入农残/重金属数据 → 颁发认证证书 |
| 📦 **供应链平台** | `/supply-chain` | 创建物流单 → 冷链温湿度追踪 → 创建采购订单 → 溯源查询 |

### 典型业务操作流程

```
示例：普洱茶从茶园到茶杯的全链路上链

1. 🌿 农业局 → 注册农场 "普洱茶山生态庄园"
2. 🌿 农业局 → 注册产品 "古树普洱茶"（选择关联农场）
3. 🌿 农业局 → 记录采收（更新产品状态为"已采收"）
4. 🔬 检测中心 → 搜索产品ID → 创建检测报告（农残未检出、有机认证）
5. 📦 供应链平台 → 创建物流单（普洱→昆明，冷链2-8°C）
6. 📦 供应链平台 → 更新物流状态：待发货 → 运输中 → 已送达
7. 📦 供应链平台 → 创建采购订单（茶企向农户采购）
8. 🌿/🔬/📦 任意组织 → 输入产品ID → 查看完整溯源信息
9. 👁️ 左下角 → "链上架构可视化" → 查看实时区块数据
```

### 界面功能

| 功能 | 说明 |
|------|------|
| **ID复制** | 鼠标悬停表格ID列，点击复制按钮一键复制 |
| **区块查看** | 每个页面右侧都有区块图标，点击查看链上区块 |
| **状态筛选** | 支持按状态筛选列表（正常/暂停/已采收/运输中等） |
| **精确搜索** | 支持输入 ID 精确搜索单条记录 |
| **链上可视化** | 首页左下角入口，展示实时区块哈希、组织架构、交易流程 |

---

## 🌐 API 接口

所有接口统一前缀 `/api`，响应格式：

```json
{
  "code": 200,
  "message": "成功",
  "data": { ... }
}
```

### 农业局接口 `/api/agriculture`

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/farm/register` | 注册农场 |
| POST | `/farm/update` | 更新农场 |
| POST | `/farm/suspend` | 暂停农场 |
| GET | `/farm/:id` | 查询农场 |
| GET | `/farm/list` | 分页查询农场列表 |
| POST | `/product/register` | 注册产品 |
| POST | `/product/harvest` | 记录采收 |
| GET | `/product/:id` | 查询产品 |
| GET | `/product/list` | 分页查询产品列表 |
| GET | `/product/by-farm/:farmId` | 查询某农场的所有产品 |
| GET | `/trace/:productId` | 完整溯源查询 |
| GET | `/trace/batch/:batchNo` | 批次号验证 |
| GET | `/block/list` | 区块列表（分页） |
| GET | `/contracts/info` | 合约信息 |

### 检测中心接口 `/api/inspection`

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/report/create` | 创建检测报告 |
| POST | `/report/fail` | 检测不通过 |
| GET | `/report/:id` | 查询检测报告 |
| GET | `/report/list` | 分页查询检测报告 |
| GET | `/certificate/list` | 认证证书列表 |
| GET | `/block/list` | 区块列表 |

### 供应链平台接口 `/api/supply-chain`

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/logistics/create` | 创建物流记录 |
| POST | `/logistics/update-status` | 更新物流状态 |
| GET | `/logistics/:id` | 查询物流记录 |
| GET | `/logistics/list` | 分页查询物流列表 |
| GET | `/logistics/by-product/:productId` | 查询产品物流 |
| POST | `/order/create` | 创建采购订单 |
| POST | `/order/update-status` | 更新订单状态 |
| GET | `/order/:id` | 查询订单 |
| GET | `/order/list` | 分页查询订单列表 |
| GET | `/trace/:productId` | 溯源查询 |
| GET | `/product/:id` | 查询产品 |
| GET | `/product/list` | 分页查询产品列表 |
| GET | `/block/list` | 区块列表 |

---

## ❗ 常见问题与解决方案

### 1️⃣ 权限问题：Permission denied

**现象**：
```
Error: open data/blocks/blocks.db: permission denied
Error: open ../../network/crypto-config/.../priv_sk: permission denied
```

**原因**：Docker 或 `sudo` 创建的文件属于 `root` 用户，当前用户无法读写。

**解决方案**：
```bash
# 修复应用数据目录权限
sudo chown -R $USER:$USER application/server/data/

# 修复网络证书目录权限
sudo chown -R $USER:$USER network/data/
sudo chown -R $USER:$USER network/config/
sudo chown -R $USER:$USER network/crypto-config/
```

### 2️⃣ 端口占用：address already in use

**现象**：
```
Error: listen tcp :8888: bind: address already in use
```

**解决方案**：
```bash
# 查找占用端口的进程并杀死
lsof -ti:8888 | xargs -r kill -9

# 或使用 fuser
fuser -k 8888/tcp
```

### 3️⃣ 链码安装失败：No such exec instance

**现象**：
```
Error response from daemon: No such exec instance: cebec55cb4f...
```

**原因**：CLI 容器 (`cli.togettoyou.com`) 在脚本执行过程中退出。

**解决方案**：
```bash
# 启动所有容器
cd network
docker compose up -d

# 或完全清理后重来
docker compose down --volumes --remove-orphans
sudo rm -rf data/ config/ crypto-config/
./install.sh
```

### 4️⃣ InitLedger 找不到

**现象**：
```
Error: endorsement failure during invoke. response: status:500 
message:"Function InitLedger not found in contract FarmContract"
```

**解决方案**：已修复（`FarmContract` 中添加了 `InitLedger` 方法）。如果使用旧版链码，请重新安装：
```bash
cd chaincode
go build

# 重新运行 install.sh
cd network
./install.sh
```

### 5️⃣ Hello 验证失败

**现象**：
```
[ERROR] 【警告】network 未部署成功，请检查日志定位具体问题。
```
但实际 `status:200` 初始化成功。

**原因**：新版链码中 `Hello` 返回 `"云岭天眼高原特色农产品溯源系统 v1.0"`，但脚本 grep 匹配的是旧版 `"hello"`。

**解决方案**：已修复（修改了 `install.sh` 中的 grep 条件）。如有需要可手动验证：
```bash
docker exec cli.togettoyou.com bash -c "CORE_PEER_ADDRESS=peer0.org1.togettoyou.com:7051 ... peer chaincode query -C mychannel -n mychaincode -c '{\"Args\":[\"BlockQueryContract:Hello\"]}'"
# 返回: 云岭天眼高原特色农产品溯源系统 v1.0
```

### 6️⃣ 前端 ProductOutlined 图标错误

**现象**：点击"农业局"卡片页面空白，控制台报错：
```
SyntaxError: The requested module does not provide an export named 'ProductOutlined'
```

**解决方案**：已修复（替换为 `AppstoreOutlined`）。如果使用旧代码，请拉取最新版本。

### 7️⃣ Docker 镜像拉取慢

**解决方案**：
```bash
cd network
# 重新运行 install.sh，当询问"是否使用镜像加速"时输入 y
./install.sh
```

### 8️⃣ 网络清理不干净

**现象**：重新安装时出现残留容器或网络冲突。

**解决方案**：
```bash
cd network

# 彻底清理
docker compose down --volumes --remove-orphans 2>/dev/null
docker container prune -f
docker network prune -f

# 删除 root 权限文件
sudo rm -rf data/ config/ crypto-config/

# 重新安装
./install.sh
```

### 9️⃣ 后端数据库 timeout

**现象**：
```
初始化区块监听器失败: 打开数据库失败：timeout
```

**原因**：BoltDB 文件被上一个进程锁定。

**解决方案**：
```bash
# 删除旧的 bolt 数据库文件
rm -f application/server/data/blocks/blocks.db

# 或检查是否有残留进程
lsof application/server/data/blocks/blocks.db
```

### 🔟 链码安装失败：`write unix @->/run/docker.sock: write: broken pipe`

**现象**：
```
Error: chaincode install failed ... docker image build failed: write unix @->/run/docker.sock: write: broken pipe
```

**根因**：Fabric 2.5.10 使用第三方 Docker 客户端库 `github.com/fsouza/go-dockerclient`，该库与 Docker 29.x 的 API v1.52 不兼容。

**解决方案**：
```bash
# 降级 Docker 到 24.0.7（Ubuntu 仓库提供）
sudo apt-get remove -y docker.io
sudo apt-get install -y docker.io=24.0.7-0ubuntu4

# 重启 Docker
sudo systemctl restart docker
```

### 1️⃣1️⃣ 链码崩溃：`reflect: Call using string as type main.OrderStatus`

**现象**：调用 `UpdateOrderStatus`、`UpdateProductStatus` 等接口时链码容器退出（`Exited (2)`），API 返回 500。

**根因**：`contractapi` 框架的反射机制无法将 `string` 自动转换为自定义类型（`OrderStatus`、`ProductStatus`、`LogisticsStatus`）。

**解决方案**：已修复。如果使用旧版链码，请重新打包安装：
```bash
cd chaincode && go build -mod=vendor -o /dev/null .
# 然后重新打包、安装、批准、提交
```

### 1️⃣2️⃣ 溯源查询 500：`inspection is required / logistics is required`

**现象**：
```
value did not match schema: inspection is required, logistics is required
```

**根因**：`TraceabilityInfo` 结构体的 `Inspection` 和 `Logistics` 字段使用了 `omitempty`，空值时被省略。`contractapi` 框架的 JSON Schema 验证将其视为必需字段，拒绝响应。

**解决方案**：已修复（移除 `omitempty`）。如使用旧版链码请重新部署。

### 1️⃣3️⃣ 富查询失败：`ExecuteQuery not supported for leveldb`

**现象**：
```
Failed to handle GET_QUERY_RESULT. error: ExecuteQuery not supported for leveldb
```

**根因**：链码中使用 `GetQueryResult()`（CouchDB 富查询），但 Peer 配置为 LevelDB。

**解决方案**：部署时自动使用 CouchDB。如需手动配置：
```yaml
# docker-compose-base.yaml
CORE_LEDGER_STATE_STATEDATABASE=CouchDB
CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb0.org1.togettoyou.com:5984
CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=admin
CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD=adminpw
```

---

## ⚙️ 配置说明

### CouchDB 配置

系统使用 CouchDB 作为状态数据库以支持富查询。每个 Peer 对应一个 CouchDB 实例：

| Peer | CouchDB 地址 | 端口 |
|------|-------------|------|
| peer0.org1 | `couchdb0.org1.togettoyou.com` | 5984 |
| peer1.org1 | `couchdb1.org1.togettoyou.com` | 15984 |
| peer0.org2 | `couchdb0.org2.togettoyou.com` | 25984 |
| peer1.org2 | `couchdb1.org2.togettoyou.com` | 35984 |
| peer0.org3 | `couchdb0.org3.togettoyou.com` | 45984 |
| peer1.org3 | `couchdb1.org3.togettoyou.com` | 55984 |

### 演示数据

系统内置 `InitLedger` 初始化 2 个农场、2 个产品、1 份检测报告、1 条物流记录和 1 个采购订单。
部署完成后调用：
```bash
docker exec cli.togettoyou.com bash -c "... peer chaincode invoke ... -c '{\"Args\":[\"FarmContract:InitLedger\"]}'"
```

---

## 🛠️ 技术栈

| 层 | 技术 | 版本 |
|-----|------|------|
| **区块链** | Hyperledger Fabric | 2.5.10 |
| **链码语言** | Go (contract-api) | 1.23+ |
| **后端** | Gin (Go web framework) | 1.10.0 |
| **后端 SDK** | fabric-gateway | 1.7.0 |
| **前端** | Vue 3 + TypeScript | 3.3.8 |
| **UI 组件** | Ant Design Vue | 3.2.20 |
| **构建工具** | Vite | 4.5.0 |
| **数据存储** | BoltDB (区块监听) + LevelDB (Fabric) | - |

---

## 📁 项目结构

```
yunling-sky-eye/
├── chaincode/                    # 智能合约（链码）
│   ├── chaincode.go             # 6大合约实现
│   └── go.mod
├── network/                      # Fabric 区块链网络
│   ├── docker-compose.yaml      # 网络容器编排
│   ├── configtx.yaml            # 通道配置
│   ├── crypto-config.yaml       # 证书配置
│   ├── install.sh               # 网络部署脚本（16步）
│   └── uninstall.sh             # 网络卸载脚本
├── application/
│   ├── server/                   # 后端 Go 服务
│   │   ├── main.go              # 入口（Gin 路由）
│   │   ├── api/                 # API 处理器
│   │   │   ├── agriculture.go   # 农业局接口
│   │   │   ├── inspection.go    # 检测中心接口
│   │   │   └── supply_chain.go  # 供应链平台接口
│   │   ├── service/             # 业务逻辑层
│   │   ├── config/              # 配置
│   │   ├── pkg/fabric/          # Fabric 客户端
│   │   └── utils/               # 工具
│   └── web/                     # 前端 Vue 应用
│       ├── src/
│       │   ├── views/
│       │   │   ├── Home.vue           # 首页
│       │   │   ├── Agriculture.vue    # 农业局面板
│       │   │   ├── Inspection.vue     # 检测中心面板
│       │   │   ├── SupplyChain.vue    # 供应链面板
│       │   │   └── BlockchainViz.vue  # 链上可视化
│       │   ├── api/index.ts          # API 接口
│       │   ├── types/index.ts        # 类型定义
│       │   └── utils/                # 工具函数
│       └── index.html
├── dev.md                          # 开发者指南
├── install.sh                      # 一键安装脚本
├── uninstall.sh                    # 一键卸载脚本
└── README.md                       # 本文件
```

---

## 📄 许可证

本项目基于 Apache 2.0 许可证开源。

---

## 🙏 致谢

- 本项目基于 [fabric-realty](https://github.com/togettoyou/fabric-realty) 改造，原项目为房地产交易系统
- 感谢 Hyperledger Fabric 社区提供的优秀区块链框架
