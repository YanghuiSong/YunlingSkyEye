package service

import (
	"application/pkg/fabric"
	"encoding/json"
	"fmt"
	"time"
)

type AgricultureService struct{}

const AGRICULTURE_ORG = "org1" // 农业局

// ==================== 农场管理 ====================

func (s *AgricultureService) RegisterFarm(id, name, province, city, district, address string, area float64, ownerName, ownerPhone, certLevel string) error {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("FarmContract:RegisterFarm", id, name, province, city, district, address, fmt.Sprintf("%f", area), ownerName, ownerPhone, certLevel, now)
	if err != nil {
		return fmt.Errorf("注册农场失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

func (s *AgricultureService) UpdateFarm(id, name, province, city, district, address string, area float64, ownerName, ownerPhone, certLevel string) error {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("FarmContract:UpdateFarm", id, name, province, city, district, address, fmt.Sprintf("%f", area), ownerName, ownerPhone, certLevel, now)
	if err != nil {
		return fmt.Errorf("更新农场失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

func (s *AgricultureService) SuspendFarm(id string) error {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("FarmContract:SuspendFarm", id, now)
	if err != nil {
		return fmt.Errorf("暂停农场失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

func (s *AgricultureService) QueryFarm(id string) (map[string]interface{}, error) {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	result, err := contract.EvaluateTransaction("FarmContract:QueryFarm", id)
	if err != nil {
		return nil, fmt.Errorf("查询农场失败: %s", fabric.ExtractErrorMessage(err))
	}
	var farm map[string]interface{}
	if err := json.Unmarshal(result, &farm); err != nil {
		return nil, fmt.Errorf("解析农场数据失败: %v", err)
	}
	return farm, nil
}

func (s *AgricultureService) QueryFarmList(pageSize int32, bookmark string, status string) (map[string]interface{}, error) {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	result, err := contract.EvaluateTransaction("FarmContract:QueryFarmList", fmt.Sprintf("%d", pageSize), bookmark, status)
	if err != nil {
		return nil, fmt.Errorf("查询农场列表失败: %s", fabric.ExtractErrorMessage(err))
	}
	var queryResult map[string]interface{}
	if err := json.Unmarshal(result, &queryResult); err != nil {
		return nil, fmt.Errorf("解析查询结果失败: %v", err)
	}
	return queryResult, nil
}

// ==================== 产品管理 ====================

func (s *AgricultureService) RegisterProduct(id, name, category, farmId, farmName, batchNo string, quantity float64, unit, description string) error {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("ProductContract:RegisterProduct", id, name, category, farmId, farmName, batchNo, fmt.Sprintf("%f", quantity), unit, now, description, now)
	if err != nil {
		return fmt.Errorf("注册产品失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

func (s *AgricultureService) RecordHarvest(productId string) error {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("ProductContract:RecordHarvest", productId, now)
	if err != nil {
		return fmt.Errorf("记录采收失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

func (s *AgricultureService) QueryProduct(id string) (map[string]interface{}, error) {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	result, err := contract.EvaluateTransaction("ProductContract:QueryProduct", id)
	if err != nil {
		return nil, fmt.Errorf("查询产品失败: %s", fabric.ExtractErrorMessage(err))
	}
	var product map[string]interface{}
	if err := json.Unmarshal(result, &product); err != nil {
		return nil, fmt.Errorf("解析产品数据失败: %v", err)
	}
	return product, nil
}

func (s *AgricultureService) QueryProductList(pageSize int32, bookmark string, status string) (map[string]interface{}, error) {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	result, err := contract.EvaluateTransaction("ProductContract:QueryProductList", fmt.Sprintf("%d", pageSize), bookmark, status)
	if err != nil {
		return nil, fmt.Errorf("查询产品列表失败: %s", fabric.ExtractErrorMessage(err))
	}
	var queryResult map[string]interface{}
	if err := json.Unmarshal(result, &queryResult); err != nil {
		return nil, fmt.Errorf("解析查询结果失败: %v", err)
	}
	return queryResult, nil
}

func (s *AgricultureService) QueryProductsByFarm(farmId string) (interface{}, error) {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	result, err := contract.EvaluateTransaction("ProductContract:QueryProductsByFarm", farmId)
	if err != nil {
		return nil, fmt.Errorf("查询农场产品失败: %s", fabric.ExtractErrorMessage(err))
	}
	var products interface{}
	if err := json.Unmarshal(result, &products); err != nil {
		return nil, fmt.Errorf("解析产品列表失败: %v", err)
	}
	return products, nil
}

// ==================== 溯源查询 ====================

func (s *AgricultureService) GetFullTraceability(productId string) (map[string]interface{}, error) {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	result, err := contract.EvaluateTransaction("TraceContract:GetFullTraceability", productId)
	if err != nil {
		return nil, fmt.Errorf("溯源查询失败: %s", fabric.ExtractErrorMessage(err))
	}
	var traceInfo map[string]interface{}
	if err := json.Unmarshal(result, &traceInfo); err != nil {
		return nil, fmt.Errorf("解析溯源数据失败: %v", err)
	}
	return traceInfo, nil
}

func (s *AgricultureService) VerifyProductByBatch(batchNo string) (interface{}, error) {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	result, err := contract.EvaluateTransaction("TraceContract:VerifyProductByBatch", batchNo)
	if err != nil {
		return nil, fmt.Errorf("批次验证失败: %s", fabric.ExtractErrorMessage(err))
	}
	var products interface{}
	if err := json.Unmarshal(result, &products); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	return products, nil
}

// ==================== 区块和合约 ====================

func (s *AgricultureService) QueryBlockList(pageSize int, pageNum int) (*fabric.BlockQueryResult, error) {
	return fabric.GetBlockListener().GetBlocksByOrg(AGRICULTURE_ORG, pageSize, pageNum)
}

func (s *AgricultureService) GetContractsInfo() (interface{}, error) {
	contract := fabric.GetContract(AGRICULTURE_ORG)
	result, err := contract.EvaluateTransaction("BlockQueryContract:GetContractsInfo")
	if err != nil {
		return nil, fmt.Errorf("获取合约信息失败: %s", fabric.ExtractErrorMessage(err))
	}
	var info interface{}
	if err := json.Unmarshal(result, &info); err != nil {
		return nil, fmt.Errorf("解析合约信息失败: %v", err)
	}
	return info, nil
}
