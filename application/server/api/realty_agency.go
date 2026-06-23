package api

import (
	"application/service"
	"application/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AgricultureHandler struct {
	svc *service.AgricultureService
}

func NewAgricultureHandler() *AgricultureHandler {
	return &AgricultureHandler{
		svc: &service.AgricultureService{},
	}
}

// ==================== 农场管理 ====================

// RegisterFarm 注册农场
func (h *AgricultureHandler) RegisterFarm(c *gin.Context) {
	var req struct {
		ID         string  `json:"id"`
		Name       string  `json:"name"`
		Province   string  `json:"province"`
		City       string  `json:"city"`
		District   string  `json:"district"`
		Address    string  `json:"address"`
		Area       float64 `json:"area"`
		OwnerName  string  `json:"ownerName"`
		OwnerPhone string  `json:"ownerPhone"`
		CertLevel  string  `json:"certLevel"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "农场信息格式错误")
		return
	}
	if err := h.svc.RegisterFarm(req.ID, req.Name, req.Province, req.City, req.District, req.Address, req.Area, req.OwnerName, req.OwnerPhone, req.CertLevel); err != nil {
		utils.ServerError(c, "注册农场失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "农场注册成功", nil)
}

// UpdateFarm 更新农场
func (h *AgricultureHandler) UpdateFarm(c *gin.Context) {
	var req struct {
		ID         string  `json:"id"`
		Name       string  `json:"name"`
		Province   string  `json:"province"`
		City       string  `json:"city"`
		District   string  `json:"district"`
		Address    string  `json:"address"`
		Area       float64 `json:"area"`
		OwnerName  string  `json:"ownerName"`
		OwnerPhone string  `json:"ownerPhone"`
		CertLevel  string  `json:"certLevel"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "农场信息格式错误")
		return
	}
	if err := h.svc.UpdateFarm(req.ID, req.Name, req.Province, req.City, req.District, req.Address, req.Area, req.OwnerName, req.OwnerPhone, req.CertLevel); err != nil {
		utils.ServerError(c, "更新农场失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "农场更新成功", nil)
}

// SuspendFarm 暂停农场
func (h *AgricultureHandler) SuspendFarm(c *gin.Context) {
	var req struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数格式错误")
		return
	}
	if err := h.svc.SuspendFarm(req.ID); err != nil {
		utils.ServerError(c, "暂停农场失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "农场已暂停", nil)
}

// QueryFarm 查询农场
func (h *AgricultureHandler) QueryFarm(c *gin.Context) {
	id := c.Param("id")
	farm, err := h.svc.QueryFarm(id)
	if err != nil {
		utils.ServerError(c, "查询农场失败: "+err.Error())
		return
	}
	utils.Success(c, farm)
}

// QueryFarmList 分页查询农场列表
func (h *AgricultureHandler) QueryFarmList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	bookmark := c.DefaultQuery("bookmark", "")
	status := c.DefaultQuery("status", "")
	result, err := h.svc.QueryFarmList(int32(pageSize), bookmark, status)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}

// ==================== 产品管理 ====================

// RegisterProduct 注册产品
func (h *AgricultureHandler) RegisterProduct(c *gin.Context) {
	var req struct {
		ID           string  `json:"id"`
		Name         string  `json:"name"`
		Category     string  `json:"category"`
		FarmID       string  `json:"farmId"`
		FarmName     string  `json:"farmName"`
		BatchNo      string  `json:"batchNo"`
		Quantity     float64 `json:"quantity"`
		Unit         string  `json:"unit"`
		Description  string  `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "产品信息格式错误")
		return
	}
	if err := h.svc.RegisterProduct(req.ID, req.Name, req.Category, req.FarmID, req.FarmName, req.BatchNo, req.Quantity, req.Unit, req.Description); err != nil {
		utils.ServerError(c, "注册产品失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "产品注册成功", nil)
}

// RecordHarvest 记录采收
func (h *AgricultureHandler) RecordHarvest(c *gin.Context) {
	var req struct {
		ProductID string `json:"productId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数格式错误")
		return
	}
	if err := h.svc.RecordHarvest(req.ProductID); err != nil {
		utils.ServerError(c, "记录采收失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "采收记录成功", nil)
}

// QueryProduct 查询产品
func (h *AgricultureHandler) QueryProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.svc.QueryProduct(id)
	if err != nil {
		utils.ServerError(c, "查询产品失败: "+err.Error())
		return
	}
	utils.Success(c, product)
}

// QueryProductList 分页查询产品列表
func (h *AgricultureHandler) QueryProductList(c *gin.Context) {
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

// QueryProductsByFarm 查询某农场的所有产品
func (h *AgricultureHandler) QueryProductsByFarm(c *gin.Context) {
	farmID := c.Param("farmId")
	result, err := h.svc.QueryProductsByFarm(farmID)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}

// ==================== 溯源查询 ====================

// GetFullTraceability 获取完整溯源信息
func (h *AgricultureHandler) GetFullTraceability(c *gin.Context) {
	productID := c.Param("productId")
	result, err := h.svc.GetFullTraceability(productID)
	if err != nil {
		utils.ServerError(c, "溯源查询失败: "+err.Error())
		return
	}
	utils.Success(c, result)
}

// VerifyProductByBatch 通过批次号验证
func (h *AgricultureHandler) VerifyProductByBatch(c *gin.Context) {
	batchNo := c.Param("batchNo")
	result, err := h.svc.VerifyProductByBatch(batchNo)
	if err != nil {
		utils.ServerError(c, "批次查询失败: "+err.Error())
		return
	}
	utils.Success(c, result)
}

// ==================== 区块和合约 ====================

// QueryBlockList 分页查询区块列表
func (h *AgricultureHandler) QueryBlockList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	result, err := h.svc.QueryBlockList(pageSize, pageNum)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}

// GetContractsInfo 获取合约信息
func (h *AgricultureHandler) GetContractsInfo(c *gin.Context) {
	result, err := h.svc.GetContractsInfo()
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}
