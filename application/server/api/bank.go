package api

import (
	"application/service"
	"application/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupplyChainHandler struct {
	svc *service.SupplyChainService
}

func NewSupplyChainHandler() *SupplyChainHandler {
	return &SupplyChainHandler{
		svc: &service.SupplyChainService{},
	}
}

// ==================== 物流管理 ====================

// CreateLogisticsRecord 创建物流记录
func (h *SupplyChainHandler) CreateLogisticsRecord(c *gin.Context) {
	var req struct {
		ID            string  `json:"id"`
		ProductID     string  `json:"productId"`
		ProductName   string  `json:"productName"`
		Quantity      float64 `json:"quantity"`
		FromProvince  string  `json:"fromProvince"`
		FromCity      string  `json:"fromCity"`
		ToProvince    string  `json:"toProvince"`
		ToCity        string  `json:"toCity"`
		FromAddress   string  `json:"fromAddress"`
		ToAddress     string  `json:"toAddress"`
		Transporter   string  `json:"transporter"`
		TransportMode string  `json:"transportMode"`
		Temperature   string  `json:"temperature"`
		Humidity      string  `json:"humidity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "物流信息格式错误")
		return
	}
	if err := h.svc.CreateLogisticsRecord(req.ID, req.ProductID, req.ProductName, req.Quantity,
		req.FromProvince, req.FromCity, req.ToProvince, req.ToCity, req.FromAddress, req.ToAddress,
		req.Transporter, req.TransportMode, req.Temperature, req.Humidity); err != nil {
		utils.ServerError(c, "创建物流记录失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "物流记录创建成功", nil)
}

// UpdateLogisticsStatus 更新物流状态
func (h *SupplyChainHandler) UpdateLogisticsStatus(c *gin.Context) {
	var req struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数格式错误")
		return
	}
	if err := h.svc.UpdateLogisticsStatus(req.ID, req.Status); err != nil {
		utils.ServerError(c, "更新物流状态失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "物流状态更新成功", nil)
}

// QueryLogisticsRecord 查询物流记录
func (h *SupplyChainHandler) QueryLogisticsRecord(c *gin.Context) {
	id := c.Param("id")
	record, err := h.svc.QueryLogisticsRecord(id)
	if err != nil {
		utils.ServerError(c, "查询物流记录失败: "+err.Error())
		return
	}
	utils.Success(c, record)
}

// QueryLogisticsList 分页查询物流列表
func (h *SupplyChainHandler) QueryLogisticsList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	bookmark := c.DefaultQuery("bookmark", "")
	status := c.DefaultQuery("status", "")
	result, err := h.svc.QueryLogisticsList(int32(pageSize), bookmark, status)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}

// QueryLogisticsByProduct 查询某产品的物流记录
func (h *SupplyChainHandler) QueryLogisticsByProduct(c *gin.Context) {
	productID := c.Param("productId")
	result, err := h.svc.QueryLogisticsByProduct(productID)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}

// ==================== 产品查询 ====================

// QueryProduct 查询产品
func (h *SupplyChainHandler) QueryProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.svc.QueryProduct(id)
	if err != nil {
		utils.ServerError(c, "查询产品失败: "+err.Error())
		return
	}
	utils.Success(c, product)
}

// QueryProductList 分页查询产品列表
func (h *SupplyChainHandler) QueryProductList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	bookmark := c.DefaultQuery("bookmark", "")
	status := c.DefaultQuery("status", "")
	result, err := h.svc.QueryProductList(int32(pageSize), bookmark, status)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}

// ==================== 采购订单管理 ====================

// CreatePurchaseOrder 创建采购订单
func (h *SupplyChainHandler) CreatePurchaseOrder(c *gin.Context) {
	var req struct {
		ID          string  `json:"id"`
		ProductID   string  `json:"productId"`
		ProductName string  `json:"productName"`
		Quantity    float64 `json:"quantity"`
		BuyerName   string  `json:"buyerName"`
		BuyerOrg    string  `json:"buyerOrg"`
		SellerName  string  `json:"sellerName"`
		SellerOrg   string  `json:"sellerOrg"`
		TotalPrice  float64 `json:"totalPrice"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "订单信息格式错误")
		return
	}
	if err := h.svc.CreatePurchaseOrder(req.ID, req.ProductID, req.ProductName, req.Quantity,
		req.BuyerName, req.BuyerOrg, req.SellerName, req.SellerOrg, req.TotalPrice); err != nil {
		utils.ServerError(c, "创建采购订单失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "采购订单创建成功", nil)
}

// UpdateOrderStatus 更新订单状态
func (h *SupplyChainHandler) UpdateOrderStatus(c *gin.Context) {
	var req struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数格式错误")
		return
	}
	if err := h.svc.UpdateOrderStatus(req.ID, req.Status); err != nil {
		utils.ServerError(c, "更新订单状态失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "订单状态更新成功", nil)
}

// QueryPurchaseOrder 查询采购订单
func (h *SupplyChainHandler) QueryPurchaseOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.svc.QueryPurchaseOrder(id)
	if err != nil {
		utils.ServerError(c, "查询采购订单失败: "+err.Error())
		return
	}
	utils.Success(c, order)
}

// QueryPurchaseOrderList 分页查询采购订单列表
func (h *SupplyChainHandler) QueryPurchaseOrderList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	bookmark := c.DefaultQuery("bookmark", "")
	status := c.DefaultQuery("status", "")
	result, err := h.svc.QueryPurchaseOrderList(int32(pageSize), bookmark, status)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}

// ==================== 溯源查询 ====================

// GetFullTraceability 获取完整溯源信息
func (h *SupplyChainHandler) GetFullTraceability(c *gin.Context) {
	productID := c.Param("productId")
	result, err := h.svc.GetFullTraceability(productID)
	if err != nil {
		utils.ServerError(c, "溯源查询失败: "+err.Error())
		return
	}
	utils.Success(c, result)
}

// ==================== 区块查询 ====================

// QueryBlockList 分页查询区块列表
func (h *SupplyChainHandler) QueryBlockList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	result, err := h.svc.QueryBlockList(pageSize, pageNum)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}
