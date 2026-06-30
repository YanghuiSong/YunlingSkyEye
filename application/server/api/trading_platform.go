package api

import (
	"application/service"
	"application/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InspectionHandler 检测认证中心 API 处理器
// 文件名 trading_platform 因历史原因命名，实际内容为检测认证中心（Org2MSP）接口
type InspectionHandler struct {
	svc *service.InspectionService
}

func NewInspectionHandler() *InspectionHandler {
	return &InspectionHandler{
		svc: &service.InspectionService{},
	}
}

// CreateInspection 创建检测报告
func (h *InspectionHandler) CreateInspection(c *gin.Context) {
	var req struct {
		ID               string `json:"id"`
		ProductID        string `json:"productId"`
		ProductName      string `json:"productName"`
		Inspector        string `json:"inspector"`
		InspectorOrg     string `json:"inspectorOrg"`
		PesticideResidue string `json:"pesticideResidue"`
		HeavyMetal       string `json:"heavyMetal"`
		Microorganism    string `json:"microorganism"`
		Conclusion       string `json:"conclusion"`
		CertNumber       string `json:"certNumber"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "检测报告格式错误")
		return
	}
	if err := h.svc.CreateInspection(req.ID, req.ProductID, req.ProductName, req.Inspector, req.InspectorOrg,
		req.PesticideResidue, req.HeavyMetal, req.Microorganism, req.Conclusion, req.CertNumber); err != nil {
		utils.ServerError(c, "创建检测报告失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "检测报告创建成功", nil)
}

// FailInspection 检测不通过
func (h *InspectionHandler) FailInspection(c *gin.Context) {
	var req struct {
		ID               string `json:"id"`
		ProductID        string `json:"productId"`
		ProductName      string `json:"productName"`
		Inspector        string `json:"inspector"`
		InspectorOrg     string `json:"inspectorOrg"`
		PesticideResidue string `json:"pesticideResidue"`
		HeavyMetal       string `json:"heavyMetal"`
		Microorganism    string `json:"microorganism"`
		Conclusion       string `json:"conclusion"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "检测报告格式错误")
		return
	}
	if err := h.svc.FailInspection(req.ID, req.ProductID, req.ProductName, req.Inspector, req.InspectorOrg,
		req.PesticideResidue, req.HeavyMetal, req.Microorganism, req.Conclusion); err != nil {
		utils.ServerError(c, "提交检测报告失败: "+err.Error())
		return
	}
	utils.SuccessWithMessage(c, "检测报告已提交", nil)
}

// QueryInspection 查询检测报告
func (h *InspectionHandler) QueryInspection(c *gin.Context) {
	id := c.Param("id")
	report, err := h.svc.QueryInspection(id)
	if err != nil {
		utils.ServerError(c, "查询检测报告失败: "+err.Error())
		return
	}
	utils.Success(c, report)
}

// QueryInspectionList 分页查询检测报告列表
func (h *InspectionHandler) QueryInspectionList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	bookmark := c.DefaultQuery("bookmark", "")
	resultFilter := c.DefaultQuery("result", "")
	result, err := h.svc.QueryInspectionList(int32(pageSize), bookmark, resultFilter)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}

// QueryCertificates 查询认证证书列表
func (h *InspectionHandler) QueryCertificates(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	bookmark := c.DefaultQuery("bookmark", "")
	result, err := h.svc.QueryCertificates(int32(pageSize), bookmark)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}

// QueryBlockList 分页查询区块列表
func (h *InspectionHandler) QueryBlockList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	result, err := h.svc.QueryBlockList(pageSize, pageNum)
	if err != nil {
		utils.ServerError(c, err.Error())
		return
	}
	utils.Success(c, result)
}
