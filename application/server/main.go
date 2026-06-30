package main

import (
	"application/api"
	"application/config"
	"application/pkg/fabric"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// ============================================================
// 云岭天眼 - 后端服务入口
// 基于 Gin Web 框架，提供 RESTful API
// 三个组织各有一个 API 分组，对应不同的 Fabric 权限
// ============================================================

func main() {
	// 初始化配置（从 config.yaml 加载端口、Fabric 连接参数等）
	if err := config.InitConfig(); err != nil {
		log.Fatalf("初始化配置失败：%v", err)
	}

	// 初始化 Fabric 客户端（建立与 Peer 的 gRPC 连接）
	// 详见 pkg/fabric/fabric.go
	if err := fabric.InitFabric(); err != nil {
		log.Fatalf("初始化Fabric客户端失败：%v", err)
	}

	// 创建 Gin 路由
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	apiGroup := r.Group("/api")

	// 创建各组织的 API Handler
	agricultureHandler := api.NewAgricultureHandler()   // 🌿 农业局（Org1）
	inspectionHandler := api.NewInspectionHandler()      // 🔬 检测认证中心（Org2）
	supplyChainHandler := api.NewSupplyChainHandler()    // 📦 供应链平台（Org3）

	// ========== 🌿 农业局（Org1MSP）接口 ==========
	// 权限：农场注册/管理、产品注册/采收、溯源查询
	agriculture := apiGroup.Group("/agriculture")
	{
		agriculture.POST("/farm/register", agricultureHandler.RegisterFarm)
		agriculture.POST("/farm/update", agricultureHandler.UpdateFarm)
		agriculture.POST("/farm/suspend", agricultureHandler.SuspendFarm)
		agriculture.GET("/farm/:id", agricultureHandler.QueryFarm)
		agriculture.GET("/farm/list", agricultureHandler.QueryFarmList)

		agriculture.POST("/product/register", agricultureHandler.RegisterProduct)
		agriculture.POST("/product/harvest", agricultureHandler.RecordHarvest)
		agriculture.GET("/product/:id", agricultureHandler.QueryProduct)
		agriculture.GET("/product/list", agricultureHandler.QueryProductList)
		agriculture.GET("/product/by-farm/:farmId", agricultureHandler.QueryProductsByFarm)

		agriculture.GET("/trace/:productId", agricultureHandler.GetFullTraceability)
		agriculture.GET("/trace/batch/:batchNo", agricultureHandler.VerifyProductByBatch)
		agriculture.GET("/block/list", agricultureHandler.QueryBlockList)
		agriculture.GET("/contracts/info", agricultureHandler.GetContractsInfo)
	}

	// ========== 🔬 检测认证中心（Org2MSP）接口 ==========
	// 权限：创建检测报告、查询报告/证书
	inspection := apiGroup.Group("/inspection")
	{
		inspection.POST("/report/create", inspectionHandler.CreateInspection)
		inspection.POST("/report/fail", inspectionHandler.FailInspection)
		inspection.GET("/report/:id", inspectionHandler.QueryInspection)
		inspection.GET("/report/list", inspectionHandler.QueryInspectionList)
		inspection.GET("/certificate/list", inspectionHandler.QueryCertificates)
		inspection.GET("/block/list", inspectionHandler.QueryBlockList)
	}

	// ========== 📦 供应链平台（Org3MSP）接口 ==========
	// 权限：物流管理、采购订单、产品查询、溯源
	supplyChain := apiGroup.Group("/supply-chain")
	{
		supplyChain.POST("/logistics/create", supplyChainHandler.CreateLogisticsRecord)
		supplyChain.POST("/logistics/update-status", supplyChainHandler.UpdateLogisticsStatus)
		supplyChain.GET("/logistics/:id", supplyChainHandler.QueryLogisticsRecord)
		supplyChain.GET("/logistics/list", supplyChainHandler.QueryLogisticsList)
		supplyChain.GET("/logistics/by-product/:productId", supplyChainHandler.QueryLogisticsByProduct)

		supplyChain.GET("/product/:id", supplyChainHandler.QueryProduct)
		supplyChain.GET("/product/list", supplyChainHandler.QueryProductList)

		supplyChain.POST("/order/create", supplyChainHandler.CreatePurchaseOrder)
		supplyChain.POST("/order/update-status", supplyChainHandler.UpdateOrderStatus)
		supplyChain.GET("/order/:id", supplyChainHandler.QueryPurchaseOrder)
		supplyChain.GET("/order/list", supplyChainHandler.QueryPurchaseOrderList)

		supplyChain.GET("/trace/:productId", supplyChainHandler.GetFullTraceability)
		supplyChain.GET("/block/list", supplyChainHandler.QueryBlockList)
	}

	// 启动 HTTP 服务器
	addr := fmt.Sprintf(":%d", config.GlobalConfig.Server.Port)
	log.Printf("云岭天眼高原特色农产品溯源系统启动中，监听端口 %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("启动服务器失败：%v", err)
	}
}
