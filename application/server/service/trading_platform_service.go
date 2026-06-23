package service

import (
	"application/pkg/fabric"
	"encoding/json"
	"fmt"
	"time"
)

type SupplyChainService struct{}

const SUPPLYCHAIN_ORG = "org3" // 供应链平台

// ==================== 物流管理 ====================

func (s *SupplyChainService) CreateLogisticsRecord(id, productId, productName string, quantity float64,
	fromProvince, fromCity, toProvince, toCity, fromAddress, toAddress string,
	transporter, transportMode, temperature, humidity string) error {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("LogisticsContract:CreateLogisticsRecord",
		id, productId, productName, fmt.Sprintf("%f", quantity),
		fromProvince, fromCity, toProvince, toCity, fromAddress, toAddress,
		transporter, transportMode, temperature, humidity, now, now)
	if err != nil {
		return fmt.Errorf("创建物流记录失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

func (s *SupplyChainService) UpdateLogisticsStatus(id, status string) error {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("LogisticsContract:UpdateLogisticsStatus", id, status, now)
	if err != nil {
		return fmt.Errorf("更新物流状态失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

func (s *SupplyChainService) QueryLogisticsRecord(id string) (map[string]interface{}, error) {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
	result, err := contract.EvaluateTransaction("LogisticsContract:QueryLogisticsRecord", id)
	if err != nil {
		return nil, fmt.Errorf("查询物流记录失败: %s", fabric.ExtractErrorMessage(err))
	}
	var record map[string]interface{}
	if err := json.Unmarshal(result, &record); err != nil {
		return nil, fmt.Errorf("解析物流数据失败: %v", err)
	}
	return record, nil
}

func (s *SupplyChainService) QueryLogisticsList(pageSize int32, bookmark string, status string) (map[string]interface{}, error) {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
	result, err := contract.EvaluateTransaction("LogisticsContract:QueryLogisticsList",
		fmt.Sprintf("%d", pageSize), bookmark, status)
	if err != nil {
		return nil, fmt.Errorf("查询物流列表失败: %s", fabric.ExtractErrorMessage(err))
	}
	var queryResult map[string]interface{}
	if err := json.Unmarshal(result, &queryResult); err != nil {
		return nil, fmt.Errorf("解析查询结果失败: %v", err)
	}
	return queryResult, nil
}

func (s *SupplyChainService) QueryLogisticsByProduct(productId string) (interface{}, error) {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
	result, err := contract.EvaluateTransaction("LogisticsContract:QueryLogisticsByProduct", productId)
	if err != nil {
		return nil, fmt.Errorf("查询产品物流失败: %s", fabric.ExtractErrorMessage(err))
	}
	var records interface{}
	if err := json.Unmarshal(result, &records); err != nil {
		return nil, fmt.Errorf("解析物流列表失败: %v", err)
	}
	return records, nil
}

// ==================== 产品查询 ====================

func (s *SupplyChainService) QueryProduct(id string) (map[string]interface{}, error) {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
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

func (s *SupplyChainService) QueryProductList(pageSize int32, bookmark string, status string) (map[string]interface{}, error) {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
	result, err := contract.EvaluateTransaction("ProductContract:QueryProductList",
		fmt.Sprintf("%d", pageSize), bookmark, status)
	if err != nil {
		return nil, fmt.Errorf("查询产品列表失败: %s", fabric.ExtractErrorMessage(err))
	}
	var queryResult map[string]interface{}
	if err := json.Unmarshal(result, &queryResult); err != nil {
		return nil, fmt.Errorf("解析查询结果失败: %v", err)
	}
	return queryResult, nil
}

// ==================== 采购订单管理 ====================

func (s *SupplyChainService) CreatePurchaseOrder(id, productId, productName string, quantity float64,
	buyerName, buyerOrg, sellerName, sellerOrg string, totalPrice float64) error {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("TradeContract:CreatePurchaseOrder",
		id, productId, productName, fmt.Sprintf("%f", quantity),
		buyerName, buyerOrg, sellerName, sellerOrg,
		fmt.Sprintf("%f", totalPrice), now)
	if err != nil {
		return fmt.Errorf("创建采购订单失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

func (s *SupplyChainService) UpdateOrderStatus(id, status string) error {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
	now := time.Now().Format(time.RFC3339)
	_, err := contract.SubmitTransaction("TradeContract:UpdateOrderStatus", id, status, now)
	if err != nil {
		return fmt.Errorf("更新订单状态失败: %s", fabric.ExtractErrorMessage(err))
	}
	return nil
}

func (s *SupplyChainService) QueryPurchaseOrder(id string) (map[string]interface{}, error) {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
	result, err := contract.EvaluateTransaction("TradeContract:QueryPurchaseOrder", id)
	if err != nil {
		return nil, fmt.Errorf("查询采购订单失败: %s", fabric.ExtractErrorMessage(err))
	}
	var order map[string]interface{}
	if err := json.Unmarshal(result, &order); err != nil {
		return nil, fmt.Errorf("解析订单数据失败: %v", err)
	}
	return order, nil
}

func (s *SupplyChainService) QueryPurchaseOrderList(pageSize int32, bookmark string, status string) (map[string]interface{}, error) {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
	result, err := contract.EvaluateTransaction("TradeContract:QueryPurchaseOrderList",
		fmt.Sprintf("%d", pageSize), bookmark, status)
	if err != nil {
		return nil, fmt.Errorf("查询订单列表失败: %s", fabric.ExtractErrorMessage(err))
	}
	var queryResult map[string]interface{}
	if err := json.Unmarshal(result, &queryResult); err != nil {
		return nil, fmt.Errorf("解析查询结果失败: %v", err)
	}
	return queryResult, nil
}

// ==================== 溯源查询 ====================

func (s *SupplyChainService) GetFullTraceability(productId string) (map[string]interface{}, error) {
	contract := fabric.GetContract(SUPPLYCHAIN_ORG)
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

// ==================== 区块查询 ====================

func (s *SupplyChainService) QueryBlockList(pageSize int, pageNum int) (*fabric.BlockQueryResult, error) {
	return fabric.GetBlockListener().GetBlocksByOrg(SUPPLYCHAIN_ORG, pageSize, pageNum)
}
