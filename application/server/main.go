package main

import (
	"application/api"
	"application/config"
	"application/pkg/fabric"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	if err := config.InitConfig(); err != nil {
		log.Fatalf("初始化配置失败：%v", err)
	}

	// 初始化 Fabric 客户端
	if err := fabric.InitFabric(); err != nil {
		log.Fatalf("初始化Fabric客户端失败：%v", err)
	}

	// 创建 Gin 路由
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	apiGroup := r.Group("/api")

	// 注册路由
	agricultureHandler := api.NewAgricultureHandler()
	inspectionHandler := api.NewInspectionHandler()
	supplyChainHandler := api.NewSupplyChainHandler()

	// ========== 农业局（Org1）接口 ==========
	agriculture := apiGroup.Group("/agriculture")
	{
		// 农场管理
		agriculture.POST("/farm/register", agricultureHandler.RegisterFarm)
		agriculture.POST("/farm/update", agricultureHandler.UpdateFarm)
		agriculture.POST("/farm/suspend", agricultureHandler.SuspendFarm)
		agriculture.GET("/farm/:id", agricultureHandler.QueryFarm)
		agriculture.GET("/farm/list", agricultureHandler.QueryFarmList)

		// 产品管理
		agriculture.POST("/product/register", agricultureHandler.RegisterProduct)
		agriculture.POST("/product/harvest", agricultureHandler.RecordHarvest)
		agriculture.GET("/product/:id", agricultureHandler.QueryProduct)
		agriculture.GET("/product/list", agricultureHandler.QueryProductList)
		agriculture.GET("/product/by-farm/:farmId", agricultureHandler.QueryProductsByFarm)

		// 溯源查询
		agriculture.GET("/trace/:productId", agricultureHandler.GetFullTraceability)
		agriculture.GET("/trace/batch/:batchNo", agricultureHandler.VerifyProductByBatch)

		// 区块查询
		agriculture.GET("/block/list", agricultureHandler.QueryBlockList)

		// 合约信息
		agriculture.GET("/contracts/info", agricultureHandler.GetContractsInfo)
	}

	// ========== 检测认证中心（Org2）接口 ==========
	inspection := apiGroup.Group("/inspection")
	{
		// 检测管理
		inspection.POST("/report/create", inspectionHandler.CreateInspection)
		inspection.POST("/report/fail", inspectionHandler.FailInspection)
		inspection.GET("/report/:id", inspectionHandler.QueryInspection)
		inspection.GET("/report/list", inspectionHandler.QueryInspectionList)

		// 证书管理
		inspection.GET("/certificate/list", inspectionHandler.QueryCertificates)

		// 区块查询
		inspection.GET("/block/list", inspectionHandler.QueryBlockList)
	}

	// ========== 供应链平台（Org3）接口 ==========
	supplyChain := apiGroup.Group("/supply-chain")
	{
		// 物流管理
		supplyChain.POST("/logistics/create", supplyChainHandler.CreateLogisticsRecord)
		supplyChain.POST("/logistics/update-status", supplyChainHandler.UpdateLogisticsStatus)
		supplyChain.GET("/logistics/:id", supplyChainHandler.QueryLogisticsRecord)
		supplyChain.GET("/logistics/list", supplyChainHandler.QueryLogisticsList)
		supplyChain.GET("/logistics/by-product/:productId", supplyChainHandler.QueryLogisticsByProduct)

		// 产品查询
		supplyChain.GET("/product/:id", supplyChainHandler.QueryProduct)
		supplyChain.GET("/product/list", supplyChainHandler.QueryProductList)

		// 采购订单管理
		supplyChain.POST("/order/create", supplyChainHandler.CreatePurchaseOrder)
		supplyChain.POST("/order/update-status", supplyChainHandler.UpdateOrderStatus)
		supplyChain.GET("/order/:id", supplyChainHandler.QueryPurchaseOrder)
		supplyChain.GET("/order/list", supplyChainHandler.QueryPurchaseOrderList)

		// 溯源查询
		supplyChain.GET("/trace/:productId", supplyChainHandler.GetFullTraceability)

		// 区块查询
		supplyChain.GET("/block/list", supplyChainHandler.QueryBlockList)
	}

	// 启动服务器
	addr := fmt.Sprintf(":%d", config.GlobalConfig.Server.Port)
	log.Printf("云岭天眼高原特色农产品溯源系统启动中，监听端口 %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("启动服务器失败：%v", err)
	}
}
