import request from '../utils/request';
import type {
  FarmPageResult, ProductPageResult, InspectionPageResult,
  LogisticsPageResult, OrderPageResult, BlockQueryResult,
  Farm, Product, InspectionReport, LogisticsRecord, PurchaseOrder,
  TraceabilityInfo, ContractInfo
} from '../types';

// ==================== 农业局 API ====================
export const agricultureApi = {
  // 农场管理
  registerFarm: (data: {
    id: string; name: string; province: string; city: string;
    district: string; address: string; area: number;
    ownerName: string; ownerPhone: string; certLevel: string;
  }) => request.post<never, void>('/agriculture/farm/register', data),

  updateFarm: (data: {
    id: string; name: string; province: string; city: string;
    district: string; address: string; area: number;
    ownerName: string; ownerPhone: string; certLevel: string;
  }) => request.post<never, void>('/agriculture/farm/update', data),

  suspendFarm: (data: { id: string }) =>
    request.post<never, void>('/agriculture/farm/suspend', data),

  getFarm: (id: string) => request.get<never, Farm>(`/agriculture/farm/${id}`),

  getFarmList: (params: { pageSize: number; bookmark: string; status?: string }) =>
    request.get<never, FarmPageResult>('/agriculture/farm/list', { params }),

  // 产品管理
  registerProduct: (data: {
    id: string; name: string; category: string;
    farmId: string; farmName: string; batchNo: string;
    quantity: number; unit: string; description: string;
  }) => request.post<never, void>('/agriculture/product/register', data),

  recordHarvest: (data: { productId: string }) =>
    request.post<never, void>('/agriculture/product/harvest', data),

  getProduct: (id: string) => request.get<never, Product>(`/agriculture/product/${id}`),

  getProductList: (params: { pageSize: number; bookmark: string; status?: string }) =>
    request.get<never, ProductPageResult>('/agriculture/product/list', { params }),

  getProductsByFarm: (farmId: string) =>
    request.get<never, Product[]>(`/agriculture/product/by-farm/${farmId}`),

  // 溯源查询
  getFullTraceability: (productId: string) =>
    request.get<never, TraceabilityInfo>(`/agriculture/trace/${productId}`),

  verifyByBatch: (batchNo: string) =>
    request.get<never, Product[]>(`/agriculture/trace/batch/${batchNo}`),

  // 区块和合约
  getBlockList: (params: { pageSize?: number; pageNum?: number }) =>
    request.get<never, BlockQueryResult>('/agriculture/block/list', { params }),

  getContractsInfo: () =>
    request.get<never, ContractInfo[]>('/agriculture/contracts/info'),
};

// ==================== 检测认证中心 API ====================
export const inspectionApi = {
  createInspection: (data: {
    id: string; productId: string; productName: string;
    inspector: string; inspectorOrg: string;
    pesticideResidue: string; heavyMetal: string;
    microorganism: string; conclusion: string; certNumber: string;
  }) => request.post<never, void>('/inspection/report/create', data),

  failInspection: (data: {
    id: string; productId: string; productName: string;
    inspector: string; inspectorOrg: string;
    pesticideResidue: string; heavyMetal: string;
    microorganism: string; conclusion: string;
  }) => request.post<never, void>('/inspection/report/fail', data),

  getInspection: (id: string) =>
    request.get<never, InspectionReport>(`/inspection/report/${id}`),

  getInspectionList: (params: { pageSize: number; bookmark: string; result?: string }) =>
    request.get<never, InspectionPageResult>('/inspection/report/list', { params }),

  getCertificates: (params: { pageSize?: number; bookmark?: string }) =>
    request.get<never, any>('/inspection/certificate/list', { params }),

  getBlockList: (params: { pageSize?: number; pageNum?: number }) =>
    request.get<never, BlockQueryResult>('/inspection/block/list', { params }),
};

// ==================== 供应链平台 API ====================
export const supplyChainApi = {
  // 物流管理
  createLogistics: (data: {
    id: string; productId: string; productName: string; quantity: number;
    fromProvince: string; fromCity: string; toProvince: string; toCity: string;
    fromAddress: string; toAddress: string; transporter: string;
    transportMode: string; temperature: string; humidity: string;
  }) => request.post<never, void>('/supply-chain/logistics/create', data),

  updateLogisticsStatus: (data: { id: string; status: string }) =>
    request.post<never, void>('/supply-chain/logistics/update-status', data),

  getLogisticsRecord: (id: string) =>
    request.get<never, LogisticsRecord>(`/supply-chain/logistics/${id}`),

  getLogisticsList: (params: { pageSize: number; bookmark: string; status?: string }) =>
    request.get<never, LogisticsPageResult>('/supply-chain/logistics/list', { params }),

  getLogisticsByProduct: (productId: string) =>
    request.get<never, LogisticsRecord[]>(`/supply-chain/logistics/by-product/${productId}`),

  // 产品查询
  getProduct: (id: string) => request.get<never, Product>(`/supply-chain/product/${id}`),

  getProductList: (params: { pageSize: number; bookmark: string; status?: string }) =>
    request.get<never, ProductPageResult>('/supply-chain/product/list', { params }),

  // 采购订单
  createOrder: (data: {
    id: string; productId: string; productName: string; quantity: number;
    buyerName: string; buyerOrg: string; sellerName: string; sellerOrg: string; totalPrice: number;
  }) => request.post<never, void>('/supply-chain/order/create', data),

  updateOrderStatus: (data: { id: string; status: string }) =>
    request.post<never, void>('/supply-chain/order/update-status', data),

  getOrder: (id: string) => request.get<never, PurchaseOrder>(`/supply-chain/order/${id}`),

  getOrderList: (params: { pageSize: number; bookmark: string; status?: string }) =>
    request.get<never, OrderPageResult>('/supply-chain/order/list', { params }),

  // 溯源查询
  getFullTraceability: (productId: string) =>
    request.get<never, TraceabilityInfo>(`/supply-chain/trace/${productId}`),

  // 区块查询
  getBlockList: (params: { pageSize?: number; pageNum?: number }) =>
    request.get<never, BlockQueryResult>('/supply-chain/block/list', { params }),
};
