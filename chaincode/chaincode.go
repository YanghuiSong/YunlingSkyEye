package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/v2/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/v2/shim"
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	"github.com/hyperledger/fabric-protos-go-apiv2/peer"
)

// ============================================================
// 智行云岭 - 农产品溯源系统
// 基于 Hyperledger Fabric 联盟链的多合约溯源平台
// 包含：产地合约、产品合约、检测合约、物流合约、溯源合约
// ============================================================

// 文档类型常量（用于创建复合键）
const (
	PREFIX_FARM           = "FARM"    // 农场/产地信息
	PREFIX_PRODUCT        = "PROD"    // 农产品信息
	PREFIX_INSPECTION     = "INSP"    // 检测报告
	PREFIX_LOGISTICS      = "LOG"     // 物流记录
	PREFIX_CERTIFICATE    = "CERT"    // 认证证书
	PREFIX_TRACE_LINK     = "TRACE"   // 溯源链接
	PREFIX_PURCHASE_ORDER = "PO"      // 采购订单
)

// 组织 MSP ID 常量
const (
	ORG_AGRICULTURE_MSPID = "Org1MSP" // 农业局监管部门
	ORG_INSPECTION_MSPID  = "Org2MSP" // 检测认证中心
	ORG_SUPPLYCHAIN_MSPID = "Org3MSP" // 供应链平台
)

// ===================== 状态枚举 =====================

type FarmStatus string

const (
	FARM_ACTIVE    FarmStatus = "ACTIVE"
	FARM_SUSPENDED FarmStatus = "SUSPENDED"
)

type ProductStatus string

const (
	PROD_PLANTED    ProductStatus = "PLANTED"
	PROD_HARVESTED  ProductStatus = "HARVESTED"
	PROD_INSPECTING ProductStatus = "INSPECTING"
	PROD_CERTIFIED  ProductStatus = "CERTIFIED"
	PROD_SHIPPING   ProductStatus = "SHIPPING"
	PROD_SOLD       ProductStatus = "SOLD"
	PROD_RECALLED   ProductStatus = "RECALLED"
)

type InspectionResult string

const (
	INSP_PASS InspectionResult = "PASS"
	INSP_FAIL InspectionResult = "FAIL"
)

type LogisticsStatus string

const (
	LOG_PREPARING  LogisticsStatus = "PREPARING"
	LOG_IN_TRANSIT LogisticsStatus = "IN_TRANSIT"
	LOG_DELIVERED  LogisticsStatus = "DELIVERED"
)

type OrderStatus string

const (
	ORDER_PENDING   OrderStatus = "PENDING"
	ORDER_PAID      OrderStatus = "PAID"
	ORDER_SHIPPED   OrderStatus = "SHIPPED"
	ORDER_CONFIRMED OrderStatus = "CONFIRMED"
	ORDER_CANCELLED OrderStatus = "CANCELLED"
)

// ===================== 数据模型 =====================

// Farm 农场/产地信息
type Farm struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Province   string     `json:"province"`
	City       string     `json:"city"`
	District   string     `json:"district"`
	Address    string     `json:"address"`
	Area       float64    `json:"area"`
	OwnerName  string     `json:"ownerName"`
	OwnerPhone string     `json:"ownerPhone"`
	CertLevel  string     `json:"certLevel"`  // 认证等级: 无公害/绿色/有机
	Status     FarmStatus `json:"status"`
	CreateTime time.Time  `json:"createTime"`
	UpdateTime time.Time  `json:"updateTime"`
}

// Product 农产品信息
type Product struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Category     string        `json:"category"`     // 类别：蔬菜/水果/粮食/肉类/水产等
	FarmID       string        `json:"farmId"`
	FarmName     string        `json:"farmName"`
	BatchNo      string        `json:"batchNo"`      // 批次号
	Quantity     float64       `json:"quantity"`     // 数量
	Unit         string        `json:"unit"`         // 单位：斤/公斤/吨/箱
	PlantingDate time.Time     `json:"plantingDate"` // 种植日期
	HarvestDate  time.Time     `json:"harvestDate"`  // 采收日期
	Status       ProductStatus `json:"status"`
	Description  string        `json:"description"`
	CreateTime   time.Time     `json:"createTime"`
	UpdateTime   time.Time     `json:"updateTime"`
}

// InspectionReport 检测报告
type InspectionReport struct {
	ID               string           `json:"id"`
	ProductID        string           `json:"productId"`
	ProductName      string           `json:"productName"`
	Inspector        string           `json:"inspector"`
	InspectorOrg     string           `json:"inspectorOrg"`
	InspectionDate   time.Time        `json:"inspectionDate"`
	Result           InspectionResult `json:"result"`
	Grade            string           `json:"grade"`            // 等级：特级/一级/二级
	CertNumber       string           `json:"certNumber"`       // 认证编号
	PesticideResidue string           `json:"pesticideResidue"` // 农残检测结果
	HeavyMetal       string           `json:"heavyMetal"`       // 重金属检测结果
	Microorganism    string           `json:"microorganism"`    // 微生物检测结果
	Conclusion       string           `json:"conclusion"`
	InspectorID      string           `json:"inspectorId"`
	CreateTime       time.Time        `json:"createTime"`
	UpdateTime       time.Time        `json:"updateTime"`
}

// LogisticsRecord 物流记录
type LogisticsRecord struct {
	ID            string          `json:"id"`
	ProductID     string          `json:"productId"`
	ProductName   string          `json:"productName"`
	Quantity      float64         `json:"quantity"`
	FromProvince  string          `json:"fromProvince"`
	FromCity      string          `json:"fromCity"`
	ToProvince    string          `json:"toProvince"`
	ToCity        string          `json:"toCity"`
	FromAddress   string          `json:"fromAddress"`
	ToAddress     string          `json:"toAddress"`
	Transporter   string          `json:"transporter"`
	TransportMode string          `json:"transportMode"`   // 运输方式：冷链/常温
	Temperature   string          `json:"temperature"`
	Humidity      string          `json:"humidity"`
	Status        LogisticsStatus `json:"status"`
	StartTime     time.Time       `json:"startTime"`
	EndTime       time.Time       `json:"endTime"`
	CreateTime    time.Time       `json:"createTime"`
	UpdateTime    time.Time       `json:"updateTime"`
}

// PurchaseOrder 采购订单
type PurchaseOrder struct {
	ID          string      `json:"id"`
	ProductID   string      `json:"productId"`
	ProductName string      `json:"productName"`
	Quantity    float64     `json:"quantity"`
	BuyerName   string      `json:"buyerName"`
	BuyerOrg    string      `json:"buyerOrg"`
	SellerName  string      `json:"sellerName"`
	SellerOrg   string      `json:"sellerOrg"`
	TotalPrice  float64     `json:"totalPrice"`
	Status      OrderStatus `json:"status"`
	CreateTime  time.Time   `json:"createTime"`
	UpdateTime  time.Time   `json:"updateTime"`
}

// QueryResult 分页查询结果
type QueryResult struct {
	Records             []interface{} `json:"records"`
	RecordsCount        int32         `json:"recordsCount"`
	Bookmark            string        `json:"bookmark"`
	FetchedRecordsCount int32         `json:"fetchedRecordsCount"`
}

// TraceabilityInfo 完整溯源信息
type TraceabilityInfo struct {
	Product    Product            `json:"product"`
	Farm       Farm               `json:"farm"`
	Inspection *InspectionReport  `json:"inspection"`
	Logistics  []LogisticsRecord  `json:"logistics"`
}

// =================== 基础合约（通用方法） ===================

// BaseContract 提供所有合约共享的通用方法
type BaseContract struct {
	contractapi.Contract
}

func (c *BaseContract) getClientMSPID(ctx contractapi.TransactionContextInterface) (string, error) {
	clientID, err := cid.New(ctx.GetStub())
	if err != nil {
		return "", fmt.Errorf("获取客户端身份失败: %v", err)
	}
	return clientID.GetMSPID()
}

func (c *BaseContract) getCompositeKey(ctx contractapi.TransactionContextInterface, objectType string, attributes []string) (string, error) {
	key, err := ctx.GetStub().CreateCompositeKey(objectType, attributes)
	if err != nil {
		return "", fmt.Errorf("创建复合键失败: %v", err)
	}
	return key, nil
}

func (c *BaseContract) getState(ctx contractapi.TransactionContextInterface, key string, value interface{}) error {
	bytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("读取状态失败: %v", err)
	}
	if bytes == nil {
		return fmt.Errorf("键 %s 不存在", key)
	}
	return json.Unmarshal(bytes, value)
}

func (c *BaseContract) putState(ctx contractapi.TransactionContextInterface, key string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("序列化数据失败: %v", err)
	}
	return ctx.GetStub().PutState(key, bytes)
}

func (c *BaseContract) getStateWithPrefix(ctx contractapi.TransactionContextInterface, prefix string, attributes []string, pageSize int32, bookmark string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error) {
	return ctx.GetStub().GetStateByPartialCompositeKeyWithPagination(prefix, attributes, pageSize, bookmark)
}

func (c *BaseContract) paginateQuery(ctx contractapi.TransactionContextInterface, prefix string, attributes []string, pageSize int32, bookmark string) (*QueryResult, error) {
	iterator, metadata, err := c.getStateWithPrefix(ctx, prefix, attributes, pageSize, bookmark)
	if err != nil {
		return nil, fmt.Errorf("分页查询失败: %v", err)
	}
	defer iterator.Close()

	records := make([]interface{}, 0)
	for iterator.HasNext() {
		response, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("获取记录失败: %v", err)
		}
		var data interface{}
		if err := json.Unmarshal(response.Value, &data); err != nil {
			return nil, fmt.Errorf("解析数据失败: %v", err)
		}
		records = append(records, data)
	}

	return &QueryResult{
		Records:             records,
		RecordsCount:        int32(len(records)),
		Bookmark:            metadata.Bookmark,
		FetchedRecordsCount: metadata.FetchedRecordsCount,
	}, nil
}

// ============================================================
// 合约一：FarmContract - 产地合约（农业局专用）
// 负责农场的注册、更新、查询管理
// ============================================================

type FarmContract struct {
	BaseContract
}

// InitLedger 初始化账本（合约部署时调用）
func (c *FarmContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	log.Println("云岭天眼溯源链码 InitLedger - 初始化演示数据")

	now, _ := time.Parse(time.RFC3339, "2026-06-29T20:00:00+08:00")

	// ========== 1. 注册农场 ==========
	farm := Farm{
		ID: "FARM001", Name: "高原生态苹果基地",
		Province: "云南省", City: "丽江市", District: "玉龙县",
		Address: "拉市镇吉余村", Area: 1000,
		OwnerName: "张建国", OwnerPhone: "13988880001",
		CertLevel: "有机", Status: FARM_ACTIVE,
		CreateTime: now, UpdateTime: now,
	}
	key, _ := c.getCompositeKey(ctx, PREFIX_FARM, []string{string(FARM_ACTIVE), farm.ID})
	if err := c.putState(ctx, key, farm); err != nil {
		return fmt.Errorf("创建演示农场失败: %v", err)
	}

	farm2 := Farm{
		ID: "FARM002", Name: "斗南鲜切花基地",
		Province: "云南省", City: "昆明市", District: "呈贡区",
		Address: "斗南街道", Area: 5000,
		OwnerName: "李华", OwnerPhone: "13988880002",
		CertLevel: "地理标志", Status: FARM_ACTIVE,
		CreateTime: now, UpdateTime: now,
	}
	key2, _ := c.getCompositeKey(ctx, PREFIX_FARM, []string{string(FARM_ACTIVE), farm2.ID})
	if err := c.putState(ctx, key2, farm2); err != nil {
		return fmt.Errorf("创建演示农场2失败: %v", err)
	}

	// ========== 2. 注册产品 ==========
	product1 := Product{
		ID: "PROD001", Name: "高原有机红富士苹果",
		Category: "水果", FarmID: "FARM001", FarmName: "高原生态苹果基地",
		BatchNo: "BATCH20260601", Quantity: 5000, Unit: "斤",
		PlantingDate: now.AddDate(0, -3, 0), HarvestDate: now,
		Status: PROD_HARVESTED, Description: "产自丽江玉龙雪山脚下，海拔2600米高原有机种植",
		CreateTime: now, UpdateTime: now,
	}
	k, _ := c.getCompositeKey(ctx, PREFIX_PRODUCT, []string{string(product1.Status), product1.ID})
	c.putState(ctx, k, product1)

	product2 := Product{
		ID: "PROD002", Name: "昆明绣球花",
		Category: "花卉", FarmID: "FARM002", FarmName: "斗南鲜切花基地",
		BatchNo: "BATCH20260602", Quantity: 2000, Unit: "箱",
		PlantingDate: now.AddDate(0, -2, 0), HarvestDate: now,
		Status: PROD_CERTIFIED, Description: "优质鲜切花，出口级品质",
		CreateTime: now, UpdateTime: now,
	}
	k, _ = c.getCompositeKey(ctx, PREFIX_PRODUCT, []string{string(product2.Status), product2.ID})
	c.putState(ctx, k, product2)

	// ========== 3. 创建检测报告（Org2） ==========
	inspection := InspectionReport{
		ID: "INSP001", ProductID: "PROD002", ProductName: "昆明绣球花",
		Inspector: "王检测", InspectorOrg: "云南省质量检测中心",
		InspectionDate: now, Result: INSP_PASS,
		Grade: "特级", CertNumber: "YN2026CERT0001",
		PesticideResidue: "未检出", HeavyMetal: "低于国标限值",
		Microorganism: "合格", Conclusion: "经检测，产品符合有机认证标准",
		InspectorID: "INSPECTOR001",
		CreateTime: now, UpdateTime: now,
	}
	key, _ = c.getCompositeKey(ctx, PREFIX_INSPECTION, []string{string(inspection.Result), inspection.ID})
	c.putState(ctx, key, inspection)

	// ========== 4. 创建物流记录（Org3） ==========
	logistics := LogisticsRecord{
		ID: "LOG001", ProductID: "PROD002", ProductName: "昆明绣球花",
		Quantity: 500, Transporter: "顺丰冷链",
		FromProvince: "云南省", FromCity: "昆明市",
		ToProvince: "上海市", ToCity: "浦东新区",
		TransportMode: "冷链", Temperature: "2-8℃",
		StartTime: now, EndTime: now.AddDate(0, 0, 2),
		Status: LOG_PREPARING,
		CreateTime: now, UpdateTime: now,
	}
	key, _ = c.getCompositeKey(ctx, PREFIX_LOGISTICS, []string{string(logistics.Status), logistics.ID})
	c.putState(ctx, key, logistics)

	// ========== 5. 创建溯源链接 ==========
	traceKey, _ := c.getCompositeKey(ctx, PREFIX_TRACE_LINK, []string{product2.ID})
	traceData := map[string]interface{}{
		"productId":   product2.ID,
		"logisticsId": logistics.ID,
		"timestamp":   now,
	}
	c.putState(ctx, traceKey, traceData)

	// ========== 6. 创建采购订单（Org3） ==========
	order := PurchaseOrder{
		ID: "ORDER001", ProductID: "PROD002", ProductName: "昆明绣球花",
		Quantity: 500, TotalPrice: 60000,
		BuyerName: "上海花卉贸易有限公司", BuyerOrg: "上海花卉",
		SellerName: "斗南鲜切花基地", SellerOrg: "斗南花卉",
		Status: ORDER_PENDING,
		CreateTime: now, UpdateTime: now,
	}
	key, _ = c.getCompositeKey(ctx, PREFIX_PURCHASE_ORDER, []string{string(ORDER_PENDING), order.ID})
	c.putState(ctx, key, order)

	log.Println("演示数据初始化完成")
	return nil
}

// RegisterFarm 注册农场（仅农业局可调用）
func (c *FarmContract) RegisterFarm(ctx contractapi.TransactionContextInterface,
	id, name, province, city, district, address string, area float64,
	ownerName, ownerPhone, certLevel string, createTime time.Time) error {

	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}
	if mspID != ORG_AGRICULTURE_MSPID {
		return fmt.Errorf("仅农业局监管机构可以注册农场")
	}
	if len(id) == 0 || len(name) == 0 || len(province) == 0 {
		return fmt.Errorf("农场ID、名称和省份不能为空")
	}
	if area <= 0 {
		return fmt.Errorf("面积必须大于0")
	}

	// 检查是否已存在
	key, err := c.getCompositeKey(ctx, PREFIX_FARM, []string{string(FARM_ACTIVE), id})
	if err != nil {
		return err
	}
	exists, err := ctx.GetStub().GetState(key)
	if err != nil {
		return err
	}
	if exists != nil {
		return fmt.Errorf("农场ID %s 已存在", id)
	}

	farm := Farm{
		ID:         id,
		Name:       name,
		Province:   province,
		City:       city,
		District:   district,
		Address:    address,
		Area:       area,
		OwnerName:  ownerName,
		OwnerPhone: ownerPhone,
		CertLevel:  certLevel,
		Status:     FARM_ACTIVE,
		CreateTime: createTime,
		UpdateTime: createTime,
	}

	key, err = c.getCompositeKey(ctx, PREFIX_FARM, []string{string(FARM_ACTIVE), id})
	if err != nil {
		return err
	}
	return c.putState(ctx, key, farm)
}

// UpdateFarm 更新农场信息
func (c *FarmContract) UpdateFarm(ctx contractapi.TransactionContextInterface,
	id, name, province, city, district, address string, area float64,
	ownerName, ownerPhone, certLevel string, updateTime time.Time) error {

	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}
	if mspID != ORG_AGRICULTURE_MSPID {
		return fmt.Errorf("仅农业局监管机构可以更新农场信息")
	}

	key, err := c.getCompositeKey(ctx, PREFIX_FARM, []string{string(FARM_ACTIVE), id})
	if err != nil {
		return err
	}
	var farm Farm
	if err := c.getState(ctx, key, &farm); err != nil {
		return fmt.Errorf("农场不存在: %v", err)
	}

	farm.Name = name
	farm.Province = province
	farm.City = city
	farm.District = district
	farm.Address = address
	farm.Area = area
	farm.OwnerName = ownerName
	farm.OwnerPhone = ownerPhone
	farm.CertLevel = certLevel
	farm.UpdateTime = updateTime

	return c.putState(ctx, key, farm)
}

// SuspendFarm 暂停农场经营
func (c *FarmContract) SuspendFarm(ctx contractapi.TransactionContextInterface, id string, updateTime time.Time) error {
	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}
	if mspID != ORG_AGRICULTURE_MSPID {
		return fmt.Errorf("仅农业局监管机构可以暂停农场")
	}

	activeKey, err := c.getCompositeKey(ctx, PREFIX_FARM, []string{string(FARM_ACTIVE), id})
	if err != nil {
		return err
	}
	var farm Farm
	if err := c.getState(ctx, activeKey, &farm); err != nil {
		return err
	}
	farm.Status = FARM_SUSPENDED
	farm.UpdateTime = updateTime

	if err := ctx.GetStub().DelState(activeKey); err != nil {
		return fmt.Errorf("删除旧状态失败: %v", err)
	}

	newKey, err := c.getCompositeKey(ctx, PREFIX_FARM, []string{string(FARM_SUSPENDED), id})
	if err != nil {
		return err
	}
	return c.putState(ctx, newKey, farm)
}

// QueryFarm 查询农场信息
func (c *FarmContract) QueryFarm(ctx contractapi.TransactionContextInterface, id string) (*Farm, error) {
	for _, status := range []FarmStatus{FARM_ACTIVE, FARM_SUSPENDED} {
		key, err := c.getCompositeKey(ctx, PREFIX_FARM, []string{string(status), id})
		if err != nil {
			return nil, err
		}
		bytes, err := ctx.GetStub().GetState(key)
		if err != nil {
			return nil, err
		}
		if bytes != nil {
			var farm Farm
			if err := json.Unmarshal(bytes, &farm); err != nil {
				return nil, err
			}
			return &farm, nil
		}
	}
	return nil, fmt.Errorf("农场ID %s 不存在", id)
}

// QueryFarmList 分页查询农场列表
func (c *FarmContract) QueryFarmList(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string, status string) (*QueryResult, error) {
	attrs := []string{}
	if status != "" {
		attrs = []string{status}
	}
	return c.paginateQuery(ctx, PREFIX_FARM, attrs, pageSize, bookmark)
}

// ============================================================
// 合约二：ProductContract - 产品合约（农业局+供应链平台）
// 负责农产品信息的注册和生命周期管理
// ============================================================

type ProductContract struct {
	BaseContract
}

// RegisterProduct 注册农产品
func (c *ProductContract) RegisterProduct(ctx contractapi.TransactionContextInterface,
	id, name, category, farmId, farmName, batchNo string, quantity float64, unit string,
	plantingDate time.Time, description string, createTime time.Time) error {

	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}
	if mspID != ORG_AGRICULTURE_MSPID && mspID != ORG_SUPPLYCHAIN_MSPID {
		return fmt.Errorf("仅农业局或供应链平台可以注册产品")
	}
	if len(id) == 0 || len(name) == 0 || len(farmId) == 0 {
		return fmt.Errorf("产品ID、名称和产地ID不能为空")
	}
	if quantity <= 0 {
		return fmt.Errorf("数量必须大于0")
	}

	// 验证农场是否存在
	farmContract := &FarmContract{BaseContract: c.BaseContract}
	_, err = farmContract.QueryFarm(ctx, farmId)
	if err != nil {
		return fmt.Errorf("关联农场不存在: %v", err)
	}

	key, err := c.getCompositeKey(ctx, PREFIX_PRODUCT, []string{string(PROD_PLANTED), id})
	if err != nil {
		return err
	}

	product := Product{
		ID:           id,
		Name:         name,
		Category:     category,
		FarmID:       farmId,
		FarmName:     farmName,
		BatchNo:      batchNo,
		Quantity:     quantity,
		Unit:         unit,
		PlantingDate: plantingDate,
		Status:       PROD_PLANTED,
		Description:  description,
		CreateTime:   createTime,
		UpdateTime:   createTime,
	}

	return c.putState(ctx, key, product)
}

// UpdateProductStatus 更新产品状态（产品生命周期管理）
func (c *ProductContract) UpdateProductStatus(ctx contractapi.TransactionContextInterface, id string, newStatus string, updateTime time.Time) error {
	status := ProductStatus(newStatus)
	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}

	var product *Product
	for _, status := range []ProductStatus{
		PROD_PLANTED, PROD_HARVESTED, PROD_INSPECTING, PROD_CERTIFIED, PROD_SHIPPING, PROD_SOLD, PROD_RECALLED,
	} {
		key, err := c.getCompositeKey(ctx, PREFIX_PRODUCT, []string{string(status), id})
		if err != nil {
			return err
		}
		bytes, err := ctx.GetStub().GetState(key)
		if err != nil {
			return err
		}
		if bytes != nil {
			var p Product
			if err := json.Unmarshal(bytes, &p); err != nil {
				return err
			}
			product = &p
			break
		}
	}
	if product == nil {
		return fmt.Errorf("产品ID %s 不存在", id)
	}

	if mspID != ORG_AGRICULTURE_MSPID && mspID != ORG_SUPPLYCHAIN_MSPID {
		return fmt.Errorf("无权限更新产品状态")
	}

	oldKey, _ := c.getCompositeKey(ctx, PREFIX_PRODUCT, []string{string(product.Status), id})
	newKey, err := c.getCompositeKey(ctx, PREFIX_PRODUCT, []string{string(status), id})
	if err != nil {
		return err
	}

	if err := ctx.GetStub().DelState(oldKey); err != nil {
		return fmt.Errorf("删除旧状态失败: %v", err)
	}

	product.Status = status
	product.UpdateTime = updateTime
	if status == PROD_HARVESTED {
		product.HarvestDate = updateTime
	}

	return c.putState(ctx, newKey, product)
}

// RecordHarvest 记录采收
func (c *ProductContract) RecordHarvest(ctx contractapi.TransactionContextInterface, id string, harvestDate time.Time) error {
	return c.UpdateProductStatus(ctx, id, string(PROD_HARVESTED), harvestDate)
}

// QueryProduct 查询产品信息
func (c *ProductContract) QueryProduct(ctx contractapi.TransactionContextInterface, id string) (*Product, error) {
	for _, status := range []ProductStatus{
		PROD_PLANTED, PROD_HARVESTED, PROD_INSPECTING, PROD_CERTIFIED, PROD_SHIPPING, PROD_SOLD, PROD_RECALLED,
	} {
		key, err := c.getCompositeKey(ctx, PREFIX_PRODUCT, []string{string(status), id})
		if err != nil {
			return nil, err
		}
		bytes, err := ctx.GetStub().GetState(key)
		if err != nil {
			return nil, err
		}
		if bytes != nil {
			var product Product
			if err := json.Unmarshal(bytes, &product); err != nil {
				return nil, err
			}
			return &product, nil
		}
	}
	return nil, fmt.Errorf("产品ID %s 不存在", id)
}

// QueryProductList 分页查询产品列表
func (c *ProductContract) QueryProductList(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string, status string) (*QueryResult, error) {
	attrs := []string{}
	if status != "" {
		attrs = []string{status}
	}
	return c.paginateQuery(ctx, PREFIX_PRODUCT, attrs, pageSize, bookmark)
}

// QueryProductsByFarm 查询某农场的所有产品
func (c *ProductContract) QueryProductsByFarm(ctx contractapi.TransactionContextInterface, farmId string) ([]Product, error) {
	query := fmt.Sprintf(`{"selector":{"farmId":"%s"}}`, farmId)
	iterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	defer iterator.Close()

	var products []Product
	for iterator.HasNext() {
		response, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		var product Product
		if err := json.Unmarshal(response.Value, &product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// ============================================================
// 合约三：InspectionContract - 检测合约（检测认证中心专用）
// 负责农产品质量检测、认证和评级
// ============================================================

type InspectionContract struct {
	BaseContract
}

// CreateInspection 创建检测报告（仅检测认证中心可调用）
func (c *InspectionContract) CreateInspection(ctx contractapi.TransactionContextInterface,
	id, productId, productName, inspector, inspectorOrg string, inspectionDate time.Time,
	pesticideResidue, heavyMetal, microorganism, conclusion string,
	certNumber string, updateTime time.Time) error {

	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}
	if mspID != ORG_INSPECTION_MSPID {
		return fmt.Errorf("仅检测认证中心可以创建检测报告")
	}
	if len(id) == 0 || len(productId) == 0 {
		return fmt.Errorf("检测报告ID和产品ID不能为空")
	}

	prodContract := &ProductContract{BaseContract: c.BaseContract}
	product, err := prodContract.QueryProduct(ctx, productId)
	if err != nil {
		return fmt.Errorf("关联产品不存在: %v", err)
	}
	if pesticideResidue == "" && heavyMetal == "" && microorganism == "" {
		return fmt.Errorf("检测数据不能为空")
	}

	grade := "一级"
	if certNumber != "" {
		grade = "特级"
	}

	report := InspectionReport{
		ID:               id,
		ProductID:        productId,
		ProductName:      productName,
		Inspector:        inspector,
		InspectorOrg:     inspectorOrg,
		InspectionDate:   inspectionDate,
		Result:           INSP_PASS,
		Grade:            grade,
		CertNumber:       certNumber,
		PesticideResidue: pesticideResidue,
		HeavyMetal:       heavyMetal,
		Microorganism:    microorganism,
		Conclusion:       conclusion,
		CreateTime:       updateTime,
		UpdateTime:       updateTime,
	}

	key, err := c.getCompositeKey(ctx, PREFIX_INSPECTION, []string{string(INSP_PASS), id})
	if err != nil {
		return err
	}
	if err := c.putState(ctx, key, report); err != nil {
		return err
	}

	if product.Status == PROD_INSPECTING || product.Status == PROD_HARVESTED || product.Status == PROD_PLANTED {
		prodContract.UpdateProductStatus(ctx, productId, string(PROD_CERTIFIED), updateTime)
	}

	if certNumber != "" {
		certKey, err := c.getCompositeKey(ctx, PREFIX_CERTIFICATE, []string{certNumber})
		if err != nil {
			return err
		}
		certData := map[string]interface{}{
			"certNumber":  certNumber,
			"productId":   productId,
			"productName": productName,
			"reportId":    id,
			"grade":       grade,
			"issueDate":   updateTime,
		}
		if err := c.putState(ctx, certKey, certData); err != nil {
			return err
		}
	}

	return nil
}

// FailInspection 检测不通过
func (c *InspectionContract) FailInspection(ctx contractapi.TransactionContextInterface,
	id, productId, productName, inspector, inspectorOrg string, inspectionDate time.Time,
	pesticideResidue, heavyMetal, microorganism, conclusion string, updateTime time.Time) error {

	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}
	if mspID != ORG_INSPECTION_MSPID {
		return fmt.Errorf("仅检测认证中心可以提交检测报告")
	}

	report := InspectionReport{
		ID:               id,
		ProductID:        productId,
		ProductName:      productName,
		Inspector:        inspector,
		InspectorOrg:     inspectorOrg,
		InspectionDate:   inspectionDate,
		Result:           INSP_FAIL,
		Grade:            "不合格",
		PesticideResidue: pesticideResidue,
		HeavyMetal:       heavyMetal,
		Microorganism:    microorganism,
		Conclusion:       conclusion,
		CreateTime:       updateTime,
		UpdateTime:       updateTime,
	}

	key, err := c.getCompositeKey(ctx, PREFIX_INSPECTION, []string{string(INSP_FAIL), id})
	if err != nil {
		return err
	}
	return c.putState(ctx, key, report)
}

// QueryInspection 查询检测报告
func (c *InspectionContract) QueryInspection(ctx contractapi.TransactionContextInterface, id string) (*InspectionReport, error) {
	for _, result := range []InspectionResult{INSP_PASS, INSP_FAIL} {
		key, err := c.getCompositeKey(ctx, PREFIX_INSPECTION, []string{string(result), id})
		if err != nil {
			return nil, err
		}
		bytes, err := ctx.GetStub().GetState(key)
		if err != nil {
			return nil, err
		}
		if bytes != nil {
			var report InspectionReport
			if err := json.Unmarshal(bytes, &report); err != nil {
				return nil, err
			}
			return &report, nil
		}
	}
	return nil, fmt.Errorf("检测报告ID %s 不存在", id)
}

// QueryInspectionList 分页查询检测报告
func (c *InspectionContract) QueryInspectionList(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string, result string) (*QueryResult, error) {
	attrs := []string{}
	if result != "" {
		attrs = []string{result}
	}
	return c.paginateQuery(ctx, PREFIX_INSPECTION, attrs, pageSize, bookmark)
}

// QueryCertificates 查询认证证书列表
func (c *InspectionContract) QueryCertificates(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string) (*QueryResult, error) {
	return c.paginateQuery(ctx, PREFIX_CERTIFICATE, []string{}, pageSize, bookmark)
}

// ============================================================
// 合约四：LogisticsContract - 物流合约（供应链平台专用）
// 负责农产品物流运输的全流程追踪
// ============================================================

type LogisticsContract struct {
	BaseContract
}

// CreateLogisticsRecord 创建物流记录
func (c *LogisticsContract) CreateLogisticsRecord(ctx contractapi.TransactionContextInterface,
	id, productId, productName string, quantity float64,
	fromProvince, fromCity, toProvince, toCity, fromAddress, toAddress string,
	transporter, transportMode, temperature, humidity string, startTime time.Time, createTime time.Time) error {

	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}
	if mspID != ORG_SUPPLYCHAIN_MSPID {
		return fmt.Errorf("仅供应链平台可以创建物流记录")
	}
	if len(id) == 0 || len(productId) == 0 {
		return fmt.Errorf("物流记录ID和产品ID不能为空")
	}

	prodContract := &ProductContract{BaseContract: c.BaseContract}
	product, err := prodContract.QueryProduct(ctx, productId)
	if err != nil {
		return fmt.Errorf("关联产品不存在: %v", err)
	}

	record := LogisticsRecord{
		ID:            id,
		ProductID:     productId,
		ProductName:   productName,
		Quantity:      quantity,
		FromProvince:  fromProvince,
		FromCity:      fromCity,
		ToProvince:    toProvince,
		ToCity:        toCity,
		FromAddress:   fromAddress,
		ToAddress:     toAddress,
		Transporter:   transporter,
		TransportMode: transportMode,
		Temperature:   temperature,
		Humidity:      humidity,
		Status:        LOG_PREPARING,
		StartTime:     startTime,
		CreateTime:    createTime,
		UpdateTime:    createTime,
	}

	key, err := c.getCompositeKey(ctx, PREFIX_LOGISTICS, []string{string(LOG_PREPARING), id})
	if err != nil {
		return err
	}
	if err := c.putState(ctx, key, record); err != nil {
		return err
	}

	if product.Status == PROD_CERTIFIED || product.Status == PROD_HARVESTED {
		prodContract.UpdateProductStatus(ctx, productId, string(PROD_SHIPPING), createTime)
	}

	return nil
}

// UpdateLogisticsStatus 更新物流状态
func (c *LogisticsContract) UpdateLogisticsStatus(ctx contractapi.TransactionContextInterface, id string, newStatus string, updateTime time.Time) error {
	status := LogisticsStatus(newStatus)
	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}
	if mspID != ORG_SUPPLYCHAIN_MSPID {
		return fmt.Errorf("仅供应链平台可以更新物流状态")
	}

	var record *LogisticsRecord
	for _, status := range []LogisticsStatus{LOG_PREPARING, LOG_IN_TRANSIT, LOG_DELIVERED} {
		key, err := c.getCompositeKey(ctx, PREFIX_LOGISTICS, []string{string(status), id})
		if err != nil {
			return err
		}
		bytes, err := ctx.GetStub().GetState(key)
		if err != nil {
			return err
		}
		if bytes != nil {
			var r LogisticsRecord
			if err := json.Unmarshal(bytes, &r); err != nil {
				return err
			}
			record = &r
			break
		}
	}
	if record == nil {
		return fmt.Errorf("物流记录ID %s 不存在", id)
	}

	oldKey, _ := c.getCompositeKey(ctx, PREFIX_LOGISTICS, []string{string(record.Status), id})
	newKey, err := c.getCompositeKey(ctx, PREFIX_LOGISTICS, []string{string(status), id})
	if err != nil {
		return err
	}

	if err := ctx.GetStub().DelState(oldKey); err != nil {
		return fmt.Errorf("删除旧状态失败: %v", err)
	}

	record.Status = status
	record.UpdateTime = updateTime
	if status == LOG_IN_TRANSIT {
		traceKey, err := c.getCompositeKey(ctx, PREFIX_TRACE_LINK, []string{record.ProductID})
		if err != nil {
			return err
		}
		traceData := map[string]interface{}{
			"productId":   record.ProductID,
			"logisticsId": id,
			"timestamp":   updateTime,
		}
		c.putState(ctx, traceKey, traceData)
	}
	if status == LOG_DELIVERED {
		record.EndTime = updateTime
		prodContract := &ProductContract{BaseContract: c.BaseContract}
		prodContract.UpdateProductStatus(ctx, record.ProductID, string(PROD_SOLD), updateTime)
	}

	return c.putState(ctx, newKey, record)
}

// QueryLogisticsRecord 查询物流记录
func (c *LogisticsContract) QueryLogisticsRecord(ctx contractapi.TransactionContextInterface, id string) (*LogisticsRecord, error) {
	for _, status := range []LogisticsStatus{LOG_PREPARING, LOG_IN_TRANSIT, LOG_DELIVERED} {
		key, err := c.getCompositeKey(ctx, PREFIX_LOGISTICS, []string{string(status), id})
		if err != nil {
			return nil, err
		}
		bytes, err := ctx.GetStub().GetState(key)
		if err != nil {
			return nil, err
		}
		if bytes != nil {
			var record LogisticsRecord
			if err := json.Unmarshal(bytes, &record); err != nil {
				return nil, err
			}
			return &record, nil
		}
	}
	return nil, fmt.Errorf("物流记录ID %s 不存在", id)
}

// QueryLogisticsList 分页查询物流记录
func (c *LogisticsContract) QueryLogisticsList(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string, status string) (*QueryResult, error) {
	attrs := []string{}
	if status != "" {
		attrs = []string{status}
	}
	return c.paginateQuery(ctx, PREFIX_LOGISTICS, attrs, pageSize, bookmark)
}

// QueryLogisticsByProduct 查询某产品的所有物流记录
func (c *LogisticsContract) QueryLogisticsByProduct(ctx contractapi.TransactionContextInterface, productId string) ([]LogisticsRecord, error) {
	query := fmt.Sprintf(`{"selector":{"productId":"%s"}}`, productId)
	iterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	defer iterator.Close()

	var records []LogisticsRecord
	for iterator.HasNext() {
		response, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		var record LogisticsRecord
		if err := json.Unmarshal(response.Value, &record); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

// ============================================================
// 合约五：TraceContract - 溯源合约（所有组织可用）
// 提供完整的农产品溯源信息查询
// ============================================================

type TraceContract struct {
	BaseContract
}

// GetFullTraceability 获取完整溯源信息（从产地到销售全链路）
func (c *TraceContract) GetFullTraceability(ctx contractapi.TransactionContextInterface, productId string) (*TraceabilityInfo, error) {
	prodContract := &ProductContract{BaseContract: c.BaseContract}
	product, err := prodContract.QueryProduct(ctx, productId)
	if err != nil {
		return nil, fmt.Errorf("产品查询失败: %v", err)
	}

	farmContract := &FarmContract{BaseContract: c.BaseContract}
	farm, err := farmContract.QueryFarm(ctx, product.FarmID)
	if err != nil {
		return nil, fmt.Errorf("产地查询失败: %v", err)
	}

	query := fmt.Sprintf(`{"selector":{"productId":"%s","result":"PASS"}}`, productId)
	iter, err := ctx.GetStub().GetQueryResult(query)
	if err == nil {
		defer iter.Close()
		if iter.HasNext() {
			resp, _ := iter.Next()
			var inspection InspectionReport
			if json.Unmarshal(resp.Value, &inspection) == nil {
				traceInfo := &TraceabilityInfo{
					Product:    *product,
					Farm:       *farm,
					Inspection: &inspection,
				}

				logisticsContract := &LogisticsContract{BaseContract: c.BaseContract}
				logistics, err := logisticsContract.QueryLogisticsByProduct(ctx, productId)
				if err == nil {
					traceInfo.Logistics = logistics
				}

				return traceInfo, nil
			}
		}
	}

	return &TraceabilityInfo{
		Product:    *product,
		Farm:       *farm,
		Inspection: nil,
		Logistics:  []LogisticsRecord{},
	}, nil
}

// VerifyProductByBatch 通过批次号验证产品
func (c *TraceContract) VerifyProductByBatch(ctx contractapi.TransactionContextInterface, batchNo string) ([]Product, error) {
	query := fmt.Sprintf(`{"selector":{"batchNo":"%s"}}`, batchNo)
	iterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	defer iterator.Close()

	var products []Product
	for iterator.HasNext() {
		response, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		var product Product
		if err := json.Unmarshal(response.Value, &product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// ============================================================
// 合约六：TradeContract - 交易合约（供应链平台专用）
// 负责农产品的采购和销售订单管理
// ============================================================

type TradeContract struct {
	BaseContract
}

// CreatePurchaseOrder 创建采购订单
func (c *TradeContract) CreatePurchaseOrder(ctx contractapi.TransactionContextInterface,
	id, productId, productName string, quantity float64,
	buyerName, buyerOrg, sellerName, sellerOrg string,
	totalPrice float64, createTime time.Time) error {

	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}
	if mspID != ORG_SUPPLYCHAIN_MSPID {
		return fmt.Errorf("仅供应链平台可以创建采购订单")
	}
	if len(id) == 0 || len(productId) == 0 {
		return fmt.Errorf("订单ID和产品ID不能为空")
	}

	prodContract := &ProductContract{BaseContract: c.BaseContract}
	if _, err := prodContract.QueryProduct(ctx, productId); err != nil {
		return fmt.Errorf("关联产品不存在: %v", err)
	}

	order := PurchaseOrder{
		ID:          id,
		ProductID:   productId,
		ProductName: productName,
		Quantity:    quantity,
		BuyerName:   buyerName,
		BuyerOrg:    buyerOrg,
		SellerName:  sellerName,
		SellerOrg:   sellerOrg,
		TotalPrice:  totalPrice,
		Status:      ORDER_PENDING,
		CreateTime:  createTime,
		UpdateTime:  createTime,
	}

	key, err := c.getCompositeKey(ctx, PREFIX_PURCHASE_ORDER, []string{string(ORDER_PENDING), id})
	if err != nil {
		return err
	}
	return c.putState(ctx, key, order)
}

// UpdateOrderStatus 更新订单状态
func (c *TradeContract) UpdateOrderStatus(ctx contractapi.TransactionContextInterface, id string, newStatus string, updateTime time.Time) error {
	status := OrderStatus(newStatus)
	mspID, err := c.getClientMSPID(ctx)
	if err != nil {
		return err
	}
	if mspID != ORG_SUPPLYCHAIN_MSPID {
		return fmt.Errorf("仅供应链平台可以更新订单状态")
	}

	var order *PurchaseOrder
	for _, status := range []OrderStatus{ORDER_PENDING, ORDER_PAID, ORDER_SHIPPED, ORDER_CONFIRMED, ORDER_CANCELLED} {
		key, err := c.getCompositeKey(ctx, PREFIX_PURCHASE_ORDER, []string{string(status), id})
		if err != nil {
			return err
		}
		bytes, err := ctx.GetStub().GetState(key)
		if err != nil {
			return err
		}
		if bytes != nil {
			var o PurchaseOrder
			if err := json.Unmarshal(bytes, &o); err != nil {
				return err
			}
			order = &o
			break
		}
	}
	if order == nil {
		return fmt.Errorf("订单ID %s 不存在", id)
	}

	oldKey, _ := c.getCompositeKey(ctx, PREFIX_PURCHASE_ORDER, []string{string(order.Status), id})
	newKey, err := c.getCompositeKey(ctx, PREFIX_PURCHASE_ORDER, []string{string(status), id})
	if err != nil {
		return err
	}
	if err := ctx.GetStub().DelState(oldKey); err != nil {
		return fmt.Errorf("删除旧状态失败: %v", err)
	}
	order.Status = status
	order.UpdateTime = updateTime

	if status == ORDER_CONFIRMED {
		prodContract := &ProductContract{BaseContract: c.BaseContract}
		prodContract.UpdateProductStatus(ctx, order.ProductID, string(PROD_SOLD), updateTime)
	}

	return c.putState(ctx, newKey, order)
}

// QueryPurchaseOrder 查询采购订单
func (c *TradeContract) QueryPurchaseOrder(ctx contractapi.TransactionContextInterface, id string) (*PurchaseOrder, error) {
	for _, status := range []OrderStatus{ORDER_PENDING, ORDER_PAID, ORDER_SHIPPED, ORDER_CONFIRMED, ORDER_CANCELLED} {
		key, err := c.getCompositeKey(ctx, PREFIX_PURCHASE_ORDER, []string{string(status), id})
		if err != nil {
			return nil, err
		}
		bytes, err := ctx.GetStub().GetState(key)
		if err != nil {
			return nil, err
		}
		if bytes != nil {
			var order PurchaseOrder
			if err := json.Unmarshal(bytes, &order); err != nil {
				return nil, err
			}
			return &order, nil
		}
	}
	return nil, fmt.Errorf("订单ID %s 不存在", id)
}

// QueryPurchaseOrderList 分页查询采购订单
func (c *TradeContract) QueryPurchaseOrderList(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string, status string) (*QueryResult, error) {
	attrs := []string{}
	if status != "" {
		attrs = []string{status}
	}
	return c.paginateQuery(ctx, PREFIX_PURCHASE_ORDER, attrs, pageSize, bookmark)
}

// ============================================================
// 合约七：BlockQueryContract - 区块查询合约（所有组织可用）
// 提供通用的区块链数据查询接口
// ============================================================

type BlockQueryContract struct {
	BaseContract
}

// Hello 合约验证
func (c *BlockQueryContract) Hello(ctx contractapi.TransactionContextInterface) (string, error) {
	return "云岭天眼高原特色农产品溯源系统 v1.0", nil
}

// GetContractsInfo 获取系统中所有合约信息
func (c *BlockQueryContract) GetContractsInfo(ctx contractapi.TransactionContextInterface) ([]map[string]string, error) {
	return []map[string]string{
		{"name": "FarmContract", "description": "产地合约 - 高原农场注册与管理", "org": "Org1MSP(农业局)"},
		{"name": "ProductContract", "description": "产品合约 - 高原农产品注册与生命周期", "org": "Org1/Org3(农业局/供应链)"},
		{"name": "InspectionContract", "description": "检测合约 - 质量检测与认证", "org": "Org2MSP(检测中心)"},
		{"name": "LogisticsContract", "description": "物流合约 - 高原冷链物流追踪", "org": "Org3MSP(供应链平台)"},
		{"name": "TraceContract", "description": "溯源合约 - 全链路溯源查询", "org": "所有组织"},
		{"name": "TradeContract", "description": "交易合约 - 采购订单管理", "org": "Org3MSP(供应链平台)"},
	}, nil
}

// ============================================================
// main - 启动智行云岭链码
// ============================================================

func main() {
	chaincode, err := contractapi.NewChaincode(
		&FarmContract{},
		&ProductContract{},
		&InspectionContract{},
		&LogisticsContract{},
		&TraceContract{},
		&TradeContract{},
		&BlockQueryContract{},
	)
	if err != nil {
		log.Panicf("创建智行云岭链码失败: %v", err)
	}

	log.Println("云岭天眼高原特色农产品溯源链码启动中...")
	if err := chaincode.Start(); err != nil {
		log.Panicf("启动智行云岭链码失败: %v", err)
	}
}
