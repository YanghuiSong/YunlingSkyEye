<template>
  <div class="supply-chain">
    <div class="app-page-header">
      <a-page-header
        title="📦 供应链平台"
        sub-title="高原冷链物流追踪 · 特色农产品采购 · 订单管理 - Org3MSP"
        @back="() => $router.push('/')"
      >
        <template #extra>
          <a-space>
            <a-tooltip title="创建新的物流记录">
              <a-button type="primary" @click="showLogisticsModal = true">
                <template #icon><PlusOutlined /></template>
                创建物流单
              </a-button>
            </a-tooltip>
            <a-tooltip title="创建采购订单">
              <a-button @click="showOrderModal = true">
                <template #icon><ShoppingCartOutlined /></template>
                创建采购订单
              </a-button>
            </a-tooltip>
          </a-space>
        </template>
      </a-page-header>
    </div>

    <div class="app-content">
      <a-tabs v-model:activeKey="activeTab">
        <!-- 物流管理 -->
        <a-tab-pane key="logistics" tab="🚚 物流管理">
          <a-card :bordered="false">
            <template #extra>
              <div class="card-extra">
                <a-input-search v-model:value="searchLogId" placeholder="输入物流单号查询" style="width: 260px; margin-right: 12px;"
                  @search="handleSearchLog" @change="handleSearchLogChange" allow-clear />
                <a-radio-group v-model:value="logStatusFilter" button-style="solid" size="small">
                  <a-radio-button value="">全部</a-radio-button>
                  <a-radio-button value="PREPARING">待发货</a-radio-button>
                  <a-radio-button value="IN_TRANSIT">运输中</a-radio-button>
                  <a-radio-button value="DELIVERED">已送达</a-radio-button>
                </a-radio-group>
              </div>
            </template>
            <div class="table-container">
              <a-table :columns="logColumns" :data-source="logisticsList" :loading="logLoading" :pagination="false"
                :scroll="{ x: 1800, y: 'calc(100vh - 350px)' }" row-key="id">
                <template #bodyCell="{ column, record }">
                  <template v-if="column.key === 'id'">
                    <a-tooltip :title="record.id">
                      <span style="margin-right:6px">{{ record.id.slice(0,8) }}...</span>
                      <copy-outlined class="copy-btn" @click.stop="copyToClipboard(record.id)" />
                    </a-tooltip>
                  </template>
                  <template v-else-if="column.key === 'status'"><a-tag :color="getStatusColor(record.status)">{{ getStatusText(record.status) }}</a-tag></template>
                  <template v-else-if="column.key === 'transportMode'"><a-tag :color="record.transportMode === '冷链' ? 'blue' : 'default'">{{ record.transportMode }}</a-tag></template>
                  <template v-else-if="column.key === 'createTime'">{{ new Date(record.createTime).toLocaleString() }}</template>
                  <template v-else-if="column.key === 'action'">
                    <a-select v-model:value="record._newStatus" style="width: 110px" @change="v => handleUpdateLogStatus(record, v)" placeholder="更新状态">
                      <a-select-option value="PREPARING">待发货</a-select-option>
                      <a-select-option value="IN_TRANSIT">运输中</a-select-option>
                      <a-select-option value="DELIVERED">已送达</a-select-option>
                    </a-select>
                  </template>
                </template>
              </a-table>
              <div class="load-more">
                <a-button :loading="logLoading" @click="loadMoreLogs" :disabled="!logBookmark">
                  {{ logBookmark ? '加载更多' : '没有更多数据' }}
                </a-button>
              </div>
            </div>
          </a-card>
        </a-tab-pane>

        <!-- 采购订单 -->
        <a-tab-pane key="orders" tab="📄 采购订单">
          <a-card :bordered="false">
            <template #extra>
              <div class="card-extra">
                <a-input-search v-model:value="searchOrderId" placeholder="输入订单号查询" style="width: 260px; margin-right: 12px;"
                  @search="handleSearchOrder" @change="handleSearchOrderChange" allow-clear />
                <a-radio-group v-model:value="orderStatusFilter" button-style="solid" size="small">
                  <a-radio-button value="">全部</a-radio-button>
                  <a-radio-button value="PENDING">待处理</a-radio-button>
                  <a-radio-button value="PAID">已付款</a-radio-button>
                  <a-radio-button value="SHIPPED">已发货</a-radio-button>
                  <a-radio-button value="CONFIRMED">已完成</a-radio-button>
                </a-radio-group>
              </div>
            </template>
            <div class="table-container">
              <a-table :columns="orderColumns" :data-source="orderList" :loading="orderLoading" :pagination="false"
                :scroll="{ x: 1700, y: 'calc(100vh - 350px)' }" row-key="id">
                <template #bodyCell="{ column, record }">
                  <template v-if="column.key === 'id'">
                    <a-tooltip :title="record.id">
                      <span style="margin-right:6px">{{ record.id.slice(0,8) }}...</span>
                      <copy-outlined class="copy-btn" @click.stop="copyToClipboard(record.id)" />
                    </a-tooltip>
                  </template>
                  <template v-else-if="column.key === 'status'"><a-tag :color="getStatusColor(record.status)">{{ getStatusText(record.status) }}</a-tag></template>
                  <template v-else-if="column.key === 'totalPrice'"><span style="color:#cf1322;font-weight:500;">¥{{ record.totalPrice?.toLocaleString() }}</span></template>
                  <template v-else-if="column.key === 'createTime'">{{ new Date(record.createTime).toLocaleString() }}</template>
                  <template v-else-if="column.key === 'action'">
                    <a-select v-model:value="record._newStatus" style="width: 110px" @change="v => handleUpdateOrderStatus(record, v)" placeholder="更新状态">
                      <a-select-option value="PENDING">待处理</a-select-option>
                      <a-select-option value="PAID">已付款</a-select-option>
                      <a-select-option value="SHIPPED">已发货</a-select-option>
                      <a-select-option value="CONFIRMED">已完成</a-select-option>
                    </a-select>
                  </template>
                </template>
              </a-table>
              <div class="load-more">
                <a-button :loading="orderLoading" @click="loadMoreOrders" :disabled="!orderBookmark">
                  {{ orderBookmark ? '加载更多' : '没有更多数据' }}
                </a-button>
              </div>
            </div>
          </a-card>
        </a-tab-pane>

        <!-- 产品查询 -->
        <a-tab-pane key="products" tab="🥬 产品查询">
          <a-card :bordered="false">
            <template #extra>
              <a-input-search v-model:value="searchSCProductId" placeholder="输入产品ID查询" style="width: 280px"
                @search="handleSearchSCProduct" @change="handleSearchSCProductChange" allow-clear />
            </template>
            <a-table :columns="scProductColumns" :data-source="scProductList" :loading="scProductLoading" :pagination="false"
              :scroll="{ x: 1500, y: 'calc(100vh - 350px)' }" row-key="id">
              <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'id'">
                    <a-tooltip :title="record.id">
                      <span style="margin-right:6px">{{ record.id.slice(0,8) }}...</span>
                      <copy-outlined class="copy-btn" @click.stop="copyToClipboard(record.id)" />
                    </a-tooltip>
                  </template>
                <template v-else-if="column.key === 'status'"><a-tag :color="getStatusColor(record.status)">{{ getStatusText(record.status) }}</a-tag></template>
                <template v-else-if="column.key === 'quantity'">{{ record.quantity }}{{ record.unit }}</template>
                <template v-else-if="column.key === 'createTime'">{{ new Date(record.createTime).toLocaleString() }}</template>
                <template v-else-if="column.key === 'action'">
                  <a-button type="link" size="small" @click="handleTraceFromSC(record)">溯源</a-button>
                </template>
              </template>
            </a-table>
          </a-card>
        </a-tab-pane>

        <!-- 溯源查询 -->
        <a-tab-pane key="trace" tab="🔍 溯源码查询">
          <a-card :bordered="false">
            <a-input-search v-model:value="traceSCProductId" placeholder="输入产品ID查询完整溯源信息"
              enter-button="查询溯源" size="large" @search="handleTraceSCSearch" style="max-width: 500px" />
            <div v-if="traceInfo" class="trace-result">
              <a-descriptions title="📋 溯源信息" bordered :column="2">
                <a-descriptions-item label="产品名称" :span="2">{{ traceInfo.product?.name }}</a-descriptions-item>
                <a-descriptions-item label="类别">{{ traceInfo.product?.category }}</a-descriptions-item>
                <a-descriptions-item label="批次号">{{ traceInfo.product?.batchNo }}</a-descriptions-item>
                <a-descriptions-item label="产地" :span="2">{{ traceInfo.farm?.name }} - {{ traceInfo.farm?.province }}{{ traceInfo.farm?.city }}</a-descriptions-item>
                <a-descriptions-item label="认证等级">{{ traceInfo.farm?.certLevel || '无' }}</a-descriptions-item>
                <a-descriptions-item label="状态"><a-tag :color="getStatusColor(traceInfo.product?.status)">{{ getStatusText(traceInfo.product?.status) }}</a-tag></a-descriptions-item>
              </a-descriptions>
              <a-divider /><h3>🧪 检测报告</h3>
              <a-descriptions v-if="traceInfo.inspection" bordered :column="2">
                <a-descriptions-item label="结果"><a-tag :color="traceInfo.inspection.result === 'PASS' ? 'green' : 'red'">{{ traceInfo.inspection.result === 'PASS' ? '✓ 通过' : '✗ 不通过' }}</a-tag></a-descriptions-item>
                <a-descriptions-item label="认证编号">{{ traceInfo.inspection.certNumber }}</a-descriptions-item>
                <a-descriptions-item label="农残">{{ traceInfo.inspection.pesticideResidue }}</a-descriptions-item>
                <a-descriptions-item label="重金属">{{ traceInfo.inspection.heavyMetal }}</a-descriptions-item>
              </a-descriptions>
              <a-empty v-else description="暂无检测报告" />
              <a-divider /><h3>🚚 物流记录</h3>
              <a-timeline v-if="traceInfo.logistics?.length">
                <a-timeline-item v-for="item in traceInfo.logistics" :key="item.id" :color="getStatusColor(item.status)">
                  <p><strong>{{ item.fromProvince }}{{ item.fromCity }} → {{ item.toProvince }}{{ item.toCity }}</strong></p>
                  <p>承运方: {{ item.transporter }} | 方式: {{ item.transportMode }} | 温度: {{ item.temperature }} | 状态: <a-tag :color="getStatusColor(item.status)">{{ getStatusText(item.status) }}</a-tag></p>
                </a-timeline-item>
              </a-timeline>
              <a-empty v-else description="暂无物流记录" />
            </div>
          </a-card>
        </a-tab-pane>
      </a-tabs>
    </div>

    <!-- 区块抽屉 -->
    <div class="block-icon" @click="openBlockDrawer"><ApartmentOutlined /></div>
    <a-drawer v-model:visible="blockDrawer" title="📦 区块信息" placement="right" :width="960">
      <div class="block-container">
        <div class="block-header">
          <h3>区块列表</h3>
          <a-pagination v-model:current="blockQuery.pageNum" v-model:pageSize="blockQuery.pageSize"
            :total="blockTotal" :show-total="t => `共 ${t} 条`"
            :page-size-options="['5','10','20','50']" show-size-changer @change="handleBlockPageChange" />
        </div>
        <div class="block-list">
          <a-card v-for="block in blockList" :key="block.block_num" class="block-item">
            <template #title><span class="block-number">区块 #{{ block.block_num }}</span><span class="block-time">{{ new Date(block.save_time).toLocaleString() }}</span></template>
            <div class="block-item-content">
              <div class="block-field"><span class="field-label">区块哈希:</span><span class="field-value hash">{{ block.block_hash }}</span></div>
              <div class="block-field"><span class="field-label">数据哈希:</span><span class="field-value hash">{{ block.data_hash }}</span></div>
              <div class="block-field"><span class="field-label">前块哈希:</span><span class="field-value hash">{{ block.prev_hash }}</span></div>
              <div class="block-field"><span class="field-label">交易数量:</span><span class="field-value">{{ block.tx_count }}</span></div>
            </div>
          </a-card>
        </div>
      </div>
    </a-drawer>

    <!-- 创建物流单对话框 -->
    <a-modal v-model:visible="showLogisticsModal" title="🚚 创建物流单" @ok="handleLogisticsOk" @cancel="handleLogisticsCancel"
      :confirmLoading="logModalLoading" width="800">
      <a-form ref="logFormRef" :model="logForm" :rules="logRules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12"><a-form-item label="产品ID" name="productId"><a-input v-model:value="logForm.productId" @change="handleLogProductChange" /></a-form-item></a-col>
          <a-col :span="12"><a-form-item label="产品名称" name="productName"><a-input v-model:value="logForm.productName" /></a-form-item></a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="8"><a-form-item label="数量" name="quantity"><a-input-number v-model:value="logForm.quantity" :min="1" style="width:100%" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="承运方" name="transporter"><a-input v-model:value="logForm.transporter" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="运输方式" name="transportMode">
            <a-select v-model:value="logForm.transportMode">
              <a-select-option value="冷链">冷链运输</a-select-option>
              <a-select-option value="常温">常温运输</a-select-option>
              <a-select-option value="快递">快递物流</a-select-option>
            </a-select>
          </a-form-item></a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="6"><a-form-item label="出发省份" name="fromProvince"><a-input v-model:value="logForm.fromProvince" /></a-form-item></a-col>
          <a-col :span="6"><a-form-item label="出发城市" name="fromCity"><a-input v-model:value="logForm.fromCity" /></a-form-item></a-col>
          <a-col :span="6"><a-form-item label="目的省份" name="toProvince"><a-input v-model:value="logForm.toProvince" /></a-form-item></a-col>
          <a-col :span="6"><a-form-item label="目的城市" name="toCity"><a-input v-model:value="logForm.toCity" /></a-form-item></a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12"><a-form-item label="出发地址" name="fromAddress"><a-input v-model:value="logForm.fromAddress" /></a-form-item></a-col>
          <a-col :span="12"><a-form-item label="目的地址" name="toAddress"><a-input v-model:value="logForm.toAddress" /></a-form-item></a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12"><a-form-item label="温度记录" name="temperature" extra="如: 2-8°C"><a-input v-model:value="logForm.temperature" /></a-form-item></a-col>
          <a-col :span="12"><a-form-item label="湿度记录" name="humidity" extra="如: 60-70%"><a-input v-model:value="logForm.humidity" /></a-form-item></a-col>
        </a-row>
      </a-form>
    </a-modal>

    <!-- 创建采购订单对话框 -->
    <a-modal v-model:visible="showOrderModal" title="📄 创建采购订单" @ok="handleOrderOk" @cancel="handleOrderCancel"
      :confirmLoading="orderModalLoading" width="700">
      <a-form ref="orderFormRef" :model="orderForm" :rules="orderRules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12"><a-form-item label="产品ID" name="productId"><a-input v-model:value="orderForm.productId" @change="handleOrderProductChange" /></a-form-item></a-col>
          <a-col :span="12"><a-form-item label="产品名称" name="productName"><a-input v-model:value="orderForm.productName" /></a-form-item></a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="8"><a-form-item label="数量" name="quantity"><a-input-number v-model:value="orderForm.quantity" :min="1" style="width:100%" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="单价(元)" name="unitPrice"><a-input-number v-model:value="orderForm.unitPrice" :min="0.01" style="width:100%" @change="calcTotalPrice" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="总价"><span style="color:#cf1322;font-weight:700;font-size:18px">¥{{ totalPrice }}</span></a-form-item></a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12"><a-form-item label="采购方" name="buyerName"><a-input v-model:value="orderForm.buyerName" /></a-form-item></a-col>
          <a-col :span="12"><a-form-item label="采购单位" name="buyerOrg"><a-input v-model:value="orderForm.buyerOrg" /></a-form-item></a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12"><a-form-item label="供应方" name="sellerName"><a-input v-model:value="orderForm.sellerName" /></a-form-item></a-col>
          <a-col :span="12"><a-form-item label="供应单位" name="sellerOrg"><a-input v-model:value="orderForm.sellerOrg" /></a-form-item></a-col>
        </a-row>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { message } from 'ant-design-vue';
import { PlusOutlined, ShoppingCartOutlined, ApartmentOutlined, CopyOutlined } from '@ant-design/icons-vue';
import { supplyChainApi } from '../api';
import { ref, reactive, watch, onMounted, computed } from 'vue';
import type { LogisticsRecord, PurchaseOrder, Product, TraceabilityInfo, BlockData } from '../types';
import { getStatusText, getStatusColor, generateUUID, generateRandomProvince, generateRandomCity, generateRandomName, copyToClipboard } from '../utils';

const activeTab = ref('logistics');

// ====== 物流管理 ======
const logisticsList = ref<LogisticsRecord[]>([]);
const logLoading = ref(false);
const logBookmark = ref('');
const logStatusFilter = ref('');
const searchLogId = ref('');

const logColumns = [
  { title: '物流单号', dataIndex: 'id', key: 'id', width: 120 },
  { title: '产品', dataIndex: 'productName', key: 'productName', width: 110 },
  { title: '出发地', key: 'from', width: 160, customRender: ({ record }: any) => `${record.fromProvince}${record.fromCity}` },
  { title: '目的地', key: 'to', width: 160, customRender: ({ record }: any) => `${record.toProvince}${record.toCity}` },
  { title: '承运方', dataIndex: 'transporter', key: 'transporter', width: 100 },
  { title: '方式', dataIndex: 'transportMode', key: 'transportMode', width: 80 },
  { title: '温度', dataIndex: 'temperature', key: 'temperature', width: 80 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 90 },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime', width: 170 },
  { title: '操作', key: 'action', width: 130, fixed: 'right' },
];

const loadLogistics = async () => {
  try {
    logLoading.value = true;
    const result = await supplyChainApi.getLogisticsList({ pageSize: 10, bookmark: logBookmark.value, status: logStatusFilter.value });
    if (!logBookmark.value) logisticsList.value = result.records;
    else logisticsList.value = [...logisticsList.value, ...result.records];
    logBookmark.value = result.bookmark;
  } catch (e: any) { message.error(e.message); }
  finally { logLoading.value = false; }
};

const loadMoreLogs = () => loadLogistics();
watch(logStatusFilter, () => { logisticsList.value = []; logBookmark.value = ''; loadLogistics(); });

const handleSearchLog = async (v: string) => {
  if (!v) return;
  try { const r = await supplyChainApi.getLogisticsRecord(v); logisticsList.value = [r]; logBookmark.value = ''; }
  catch (e: any) { message.error(e.message); logisticsList.value = []; }
};
const handleSearchLogChange = (e: any) => { if (!e.target.value) { logisticsList.value = []; logBookmark.value = ''; loadLogistics(); } };

const handleUpdateLogStatus = async (record: any, status: string) => {
  try { await supplyChainApi.updateLogisticsStatus({ id: record.id, status }); message.success('物流状态已更新'); loadLogistics(); }
  catch (e: any) { message.error(e.message); }
};

// 创建物流单
const showLogisticsModal = ref(false);
const logModalLoading = ref(false);
const logFormRef = ref();
const logForm = reactive({
  productId: '', productName: '', quantity: undefined as number | undefined,
  fromProvince: '', fromCity: '', toProvince: '', toCity: '',
  fromAddress: '', toAddress: '',
  transporter: '', transportMode: '冷链', temperature: '', humidity: '',
});
const logRules = { productId: [{ required: true }], transporter: [{ required: true }] };

const handleLogProductChange = async (e: any) => {
  const id = e.target.value;
  if (!id) return;
  try { const p = await supplyChainApi.getProduct(id); logForm.productName = p.name; }
  catch { logForm.productName = ''; }
};

const handleLogisticsOk = () => {
  logFormRef.value?.validate().then(async () => {
    logModalLoading.value = true;
    try {
      await supplyChainApi.createLogistics({ ...logForm, id: generateUUID(), quantity: logForm.quantity || 0 });
      message.success('物流单创建成功');
      showLogisticsModal.value = false; logFormRef.value?.resetFields();
      logisticsList.value = []; logBookmark.value = ''; loadLogistics();
    } catch (e: any) { message.error(e.message); }
    finally { logModalLoading.value = false; }
  });
};
const handleLogisticsCancel = () => { showLogisticsModal.value = false; logFormRef.value?.resetFields(); };

// ====== 采购订单 ======
const orderList = ref<PurchaseOrder[]>([]);
const orderLoading = ref(false);
const orderBookmark = ref('');
const orderStatusFilter = ref('');
const searchOrderId = ref('');

const orderColumns = [
  { title: '订单号', dataIndex: 'id', key: 'id', width: 120 },
  { title: '产品', dataIndex: 'productName', key: 'productName', width: 110 },
  { title: '采购方', dataIndex: 'buyerName', key: 'buyerName', width: 100 },
  { title: '供应方', dataIndex: 'sellerName', key: 'sellerName', width: 100 },
  { title: '数量', key: 'quantity', width: 80, customRender: ({ record }: any) => record.quantity },
  { title: '总价', dataIndex: 'totalPrice', key: 'totalPrice', width: 110 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 90 },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime', width: 170 },
  { title: '操作', key: 'action', width: 130, fixed: 'right' },
];

const loadOrders = async () => {
  try {
    orderLoading.value = true;
    const result = await supplyChainApi.getOrderList({ pageSize: 10, bookmark: orderBookmark.value, status: orderStatusFilter.value });
    if (!orderBookmark.value) orderList.value = result.records;
    else orderList.value = [...orderList.value, ...result.records];
    orderBookmark.value = result.bookmark;
  } catch (e: any) { message.error(e.message); }
  finally { orderLoading.value = false; }
};

const loadMoreOrders = () => loadOrders();
watch(orderStatusFilter, () => { orderList.value = []; orderBookmark.value = ''; loadOrders(); });

const handleSearchOrder = async (v: string) => {
  if (!v) return;
  try { const r = await supplyChainApi.getOrder(v); orderList.value = [r]; orderBookmark.value = ''; }
  catch (e: any) { message.error(e.message); orderList.value = []; }
};
const handleSearchOrderChange = (e: any) => { if (!e.target.value) { orderList.value = []; orderBookmark.value = ''; loadOrders(); } };

const handleUpdateOrderStatus = async (record: any, status: string) => {
  try { await supplyChainApi.updateOrderStatus({ id: record.id, status }); message.success('订单状态已更新'); loadOrders(); }
  catch (e: any) { message.error(e.message); }
};

// 创建采购订单
const showOrderModal = ref(false);
const orderModalLoading = ref(false);
const orderFormRef = ref();
const orderForm = reactive({
  productId: '', productName: '', quantity: undefined as number | undefined,
  unitPrice: undefined as number | undefined,
  buyerName: '', buyerOrg: '', sellerName: '', sellerOrg: '',
});
const totalPrice = computed(() => {
  if (orderForm.quantity && orderForm.unitPrice) return (orderForm.quantity * orderForm.unitPrice).toLocaleString();
  return '0';
});
const orderRules = { productId: [{ required: true }], quantity: [{ required: true }], buyerName: [{ required: true }], sellerName: [{ required: true }] };

const handleOrderProductChange = async (e: any) => {
  const id = e.target.value;
  if (!id) return;
  try { const p = await supplyChainApi.getProduct(id); orderForm.productName = p.name; }
  catch { orderForm.productName = ''; }
};

const calcTotalPrice = () => { /* computed handles this */ };

const handleOrderOk = () => {
  orderFormRef.value?.validate().then(async () => {
    orderModalLoading.value = true;
    try {
      const total = (orderForm.quantity || 0) * (orderForm.unitPrice || 0);
      await supplyChainApi.createOrder({
        id: generateUUID(), productId: orderForm.productId, productName: orderForm.productName,
        quantity: orderForm.quantity || 0, buyerName: orderForm.buyerName, buyerOrg: orderForm.buyerOrg || '采购公司',
        sellerName: orderForm.sellerName, sellerOrg: orderForm.sellerOrg || '供应公司', totalPrice: total,
      });
      message.success('采购订单创建成功');
      showOrderModal.value = false; orderFormRef.value?.resetFields();
      orderList.value = []; orderBookmark.value = ''; loadOrders();
    } catch (e: any) { message.error(e.message); }
    finally { orderModalLoading.value = false; }
  });
};
const handleOrderCancel = () => { showOrderModal.value = false; orderFormRef.value?.resetFields(); };

// ====== 供应链产品查询 ======
const scProductList = ref<Product[]>([]);
const scProductLoading = ref(false);
const searchSCProductId = ref('');

const scProductColumns = [
  { title: '产品ID', dataIndex: 'id', key: 'id', width: 120 },
  { title: '名称', dataIndex: 'name', key: 'name', width: 110 },
  { title: '类别', dataIndex: 'category', key: 'category', width: 70 },
  { title: '产地', dataIndex: 'farmName', key: 'farmName', width: 120 },
  { title: '批次号', dataIndex: 'batchNo', key: 'batchNo', width: 150 },
  { title: '数量', dataIndex: 'quantity', key: 'quantity', width: 90 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 80 },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime', width: 170 },
  { title: '操作', key: 'action', width: 70, fixed: 'right' },
];

const loadSCProducts = async () => {
  scProductLoading.value = true;
  try { const r = await supplyChainApi.getProductList({ pageSize: 20, bookmark: '' }); scProductList.value = r.records; }
  catch (e: any) { message.error(e.message); }
  finally { scProductLoading.value = false; }
};

const handleSearchSCProduct = async (v: string) => {
  if (!v) return;
  try { const r = await supplyChainApi.getProduct(v); scProductList.value = [r]; }
  catch (e: any) { message.error(e.message); scProductList.value = []; }
};
const handleSearchSCProductChange = (e: any) => { if (!e.target.value) loadSCProducts(); };

// ====== 溯源 ======
const traceSCProductId = ref('');
const traceInfo = ref<TraceabilityInfo | null>(null);

const handleTraceSCSearch = async (v: string) => {
  if (!v) return;
  try { traceInfo.value = await supplyChainApi.getFullTraceability(v); }
  catch (e: any) { message.error(e.message); traceInfo.value = null; }
};

const handleTraceFromSC = async (record: Product) => {
  traceSCProductId.value = record.id;
  activeTab.value = 'trace';
  try { traceInfo.value = await supplyChainApi.getFullTraceability(record.id); }
  catch (e: any) { message.error(e.message); }
};

// ====== 区块 ======
const blockDrawer = ref(false);
const blockList = ref<BlockData[]>([]);
const blockTotal = ref(0);
const blockQuery = reactive({ pageSize: 10, pageNum: 1 });

const openBlockDrawer = async () => { blockDrawer.value = true; await fetchBlockList(); };
const fetchBlockList = async () => {
  try { const r = await supplyChainApi.getBlockList({}); blockList.value = r.blocks; blockTotal.value = r.total; }
  catch (e) { console.error(e); }
};
const handleBlockPageChange = async (page: number, size: number) => { blockQuery.pageNum = page; blockQuery.pageSize = size; await fetchBlockList(); };

watch(activeTab, (key) => { if (key === 'products' && scProductList.value.length === 0) loadSCProducts(); });

onMounted(() => { loadLogistics(); loadOrders(); });
</script>