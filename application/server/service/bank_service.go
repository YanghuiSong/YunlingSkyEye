package service

import (
	"application/pkg/fabric"
	"encoding/json"
	"fmt"
	"time"
)

type InspectionService struct{}

const INSPECTION_ORG = "org2" // 检测认证中心

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
