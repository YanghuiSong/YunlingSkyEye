// API 响应基础结构
export interface ApiResponse<T> {
  code: number;
  message: string;
  data: T;
}

// 区块数据类型
export interface BlockData {
  block_num: number;
  block_hash: string;
  data_hash: string;
  prev_hash: string;
  tx_count: number;
  save_time: string;
}

// 区块查询结果类型
export interface BlockQueryResult {
  blocks: BlockData[];
  total: number;
  page_size: number;
  page_num: number;
  has_more: boolean;
}

// 分页查询结果
export interface PageResult<T> {
  bookmark: string;
  fetchedRecordsCount: number;
  records: T[];
  recordsCount: number;
}

// 农场信息
export interface Farm {
  id: string;
  name: string;
  province: string;
  city: string;
  district: string;
  address: string;
  area: number;
  ownerName: string;
  ownerPhone: string;
  certLevel: string;
  status: 'ACTIVE' | 'SUSPENDED';
  createTime: string;
  updateTime: string;
}

// 农产品信息
export interface Product {
  id: string;
  name: string;
  category: string;
  farmId: string;
  farmName: string;
  batchNo: string;
  quantity: number;
  unit: string;
  plantingDate: string;
  harvestDate: string;
  status: 'PLANTED' | 'HARVESTED' | 'INSPECTING' | 'CERTIFIED' | 'SHIPPING' | 'SOLD' | 'RECALLED';
  description: string;
  createTime: string;
  updateTime: string;
}

// 检测报告
export interface InspectionReport {
  id: string;
  productId: string;
  productName: string;
  inspector: string;
  inspectorOrg: string;
  inspectionDate: string;
  result: 'PASS' | 'FAIL';
  grade: string;
  certNumber: string;
  pesticideResidue: string;
  heavyMetal: string;
  microorganism: string;
  conclusion: string;
  createTime: string;
  updateTime: string;
}

// 物流记录
export interface LogisticsRecord {
  id: string;
  productId: string;
  productName: string;
  quantity: number;
  fromProvince: string;
  fromCity: string;
  toProvince: string;
  toCity: string;
  fromAddress: string;
  toAddress: string;
  transporter: string;
  transportMode: string;
  temperature: string;
  humidity: string;
  status: 'PREPARING' | 'IN_TRANSIT' | 'DELIVERED';
  startTime: string;
  endTime: string;
  createTime: string;
  updateTime: string;
}

// 采购订单
export interface PurchaseOrder {
  id: string;
  productId: string;
  productName: string;
  quantity: number;
  buyerName: string;
  buyerOrg: string;
  sellerName: string;
  sellerOrg: string;
  totalPrice: number;
  status: 'PENDING' | 'PAID' | 'SHIPPED' | 'CONFIRMED' | 'CANCELLED';
  createTime: string;
  updateTime: string;
}

// 溯源信息
export interface TraceabilityInfo {
  product: Product;
  farm: Farm;
  inspection?: InspectionReport;
  logistics?: LogisticsRecord[];
}

// 合约信息
export interface ContractInfo {
  name: string;
  description: string;
  org: string;
}

// 分页结果类型
export type FarmPageResult = PageResult<Farm>;
export type ProductPageResult = PageResult<Product>;
export type InspectionPageResult = PageResult<InspectionReport>;
export type LogisticsPageResult = PageResult<LogisticsRecord>;
export type OrderPageResult = PageResult<PurchaseOrder>; 