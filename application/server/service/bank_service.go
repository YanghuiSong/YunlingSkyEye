package service

import (
	"application/pkg/fabric"
	"encoding/json"
	"fmt"
	"time"
)

// InspectionService 检测认证中心业务层
// 负责：创建检测报告、不合格检测、查询报告/证书
// 对应链码：InspectionContract（Org2MSP 专用）
// 调用组织：检测认证中心（org2）
type InspectionService struct{}

const INSPECTION_ORG = "org2" // 检测认证中心组织名称，用于获取 Fabric 合约实例

// CreateInspection 创建检测报告（提交交易，需要记账背书）
// 参数说明：
//   - id: 检测报告唯一ID
//   - productId/productName: 被检测产品信息
//   - inspector/inspectorOrg: 检测机构和检测人
//   - pesticideResidue/heavyMetal/microorganism: 农残/重金属/微生物检测结果
//   - conclusion: 检测结论
//   - certNumber: 认证证书编号（通过时颁发）
// 调用链码：InspectionContract:CreateInspection（写入操作，需要 Org2 背书）
func (s *InspectionService) CreateInspection(id, productId, productName, inspector, inspectorOrg,
	pesticideResidue, heavyMetal, microorganism, conclusion, certNumber string) error {
	contract := fabric.GetContract(INSPECTION_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("InspectionContract:CreateInspection",
		id, productId, productName, inspector, inspectorOrg, now,
		pesticideResidue, heavyMetal, microorganism, conclusion, certNumber, now)
	if err != nil {
		return fmt.Errorf("创建检测报告失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

// FailInspection 检测不通过（与CreateInspection的区别是不生成证书编号）
// 当产品检测不合格时，记录不合格原因，不颁发认证证书
func (s *InspectionService) FailInspection(id, productId, productName, inspector, inspectorOrg,
	pesticideResidue, heavyMetal, microorganism, conclusion string) error {
	contract := fabric.GetContract(INSPECTION_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("InspectionContract:FailInspection",
		id, productId, productName, inspector, inspectorOrg, now,
		pesticideResidue, heavyMetal, microorganism, conclusion, now)
	if err != nil {
		return fmt.Errorf("提交检测报告失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

// QueryInspection 查询单个检测报告详情
// 使用 EvaluateTransaction（只读查询，不经过排序共识，响应更快）
func (s *InspectionService) QueryInspection(id string) (map[string]interface{}, error) {
	contract := fabric.GetContract(INSPECTION_ORG)
	result, err := contract.EvaluateTransaction("InspectionContract:QueryInspection", id)
	if err != nil {
		return nil, fmt.Errorf("查询检测报告失败: %s", fabric.ExtractErrorMessage(err))
	}
	var report map[string]interface{}
	if err := json.Unmarshal(result, &report); err != nil {
		return nil, fmt.Errorf("解析检测报告数据失败: %v", err)
	}
	return report, nil
}

func (s *InspectionService) QueryInspectionList(pageSize int32, bookmark string, resultFilter string) (map[string]interface{}, error) {
	contract := fabric.GetContract(INSPECTION_ORG)
	result, err := contract.EvaluateTransaction("InspectionContract:QueryInspectionList",
		fmt.Sprintf("%d", pageSize), bookmark, resultFilter)
	if err != nil {
		return nil, fmt.Errorf("查询检测报告列表失败: %s", fabric.ExtractErrorMessage(err))
	}
	var queryResult map[string]interface{}
	if err := json.Unmarshal(result, &queryResult); err != nil {
		return nil, fmt.Errorf("解析查询结果失败: %v", err)
	}
	return queryResult, nil
}

// QueryCertificates 查询已颁发的认证证书列表（分页）
// 分页方式：pageSize 指定每页条数，bookmark 为 CouchDB 游标
func (s *InspectionService) QueryCertificates(pageSize int32, bookmark string) (map[string]interface{}, error) {
	contract := fabric.GetContract(INSPECTION_ORG)
	result, err := contract.EvaluateTransaction("InspectionContract:QueryCertificates",
		fmt.Sprintf("%d", pageSize), bookmark)
	if err != nil {
		return nil, fmt.Errorf("查询证书列表失败: %s", fabric.ExtractErrorMessage(err))
	}
	var queryResult map[string]interface{}
	if err := json.Unmarshal(result, &queryResult); err != nil {
		return nil, fmt.Errorf("解析查询结果失败: %v", err)
	}
	return queryResult, nil
}

func (s *InspectionService) QueryBlockList(pageSize int, pageNum int) (*fabric.BlockQueryResult, error) {
	return fabric.GetBlockListener().GetBlocksByOrg(INSPECTION_ORG, pageSize, pageNum)
}
