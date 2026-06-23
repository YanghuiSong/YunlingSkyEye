<template>
  <div class="agriculture">
    <div class="app-page-header">
      <a-page-header
        title="🌿 农业局监管平台"
        sub-title="高原农场注册 · 特色产品管理 · 全链溯源 · 智能合约监管 - Org1MSP"
        @back="() => $router.push('/')"
      >
        <template #extra>
          <a-space>
            <a-tooltip title="注册新的农场生产基地">
              <a-button type="primary" @click="showFarmModal = true">
                <template #icon><PlusOutlined /></template>
                注册新农场
              </a-button>
            </a-tooltip>
            <a-tooltip title="注册新的农产品">
              <a-button @click="showProductModal = true">
                <template #icon><AppstoreOutlined /></template>
                注册新产品
              </a-button>
            </a-tooltip>
          </a-space>
        </template>
      </a-page-header>
    </div>

    <div class="app-content">
      <!-- Tabs 切换 -->
      <a-tabs v-model:activeKey="activeTab" @change="handleTabChange">
        <a-tab-pane key="farms" tab="🏘️ 农场管理">
          <a-card :bordered="false">
            <template #extra>
              <div class="card-extra">
                <a-input-search
                  v-model:value="searchFarmId"
                  placeholder="输入农场ID精确查询"
                  style="width: 280px; margin-right: 12px;"
                  @search="handleSearchFarm"
                  @change="handleSearchFarmChange"
                  allow-clear
                />
                <a-radio-group v-model:value="farmStatusFilter" button-style="solid" size="small">
                  <a-radio-button value="">全部</a-radio-button>
                  <a-radio-button value="ACTIVE">正常</a-radio-button>
                  <a-radio-button value="SUSPENDED">暂停</a-radio-button>
                </a-radio-group>
              </div>
            </template>
            <div class="table-container">
              <a-table
                :columns="farmColumns"
                :data-source="farmList"
                :loading="farmLoading"
                :pagination="false"
                :scroll="{ x: 1400, y: 'calc(100vh - 400px)' }"
                row-key="id"
              >
                <template #bodyCell="{ column, record }">
                  <template v-if="column.key === 'id'">
                    <a-tooltip :title="record.id">
                      <span style="margin-right:6px">{{ record.id.slice(0,8) }}...</span>
                      <copy-outlined class="copy-btn" @click.stop="copyToClipboard(record.id)" />
                    </a-tooltip>
                  </template>
                  <template v-else-if="column.key === 'status'">
                    <a-tag :color="getStatusColor(record.status)">{{ getStatusText(record.status) }}</a-tag>
                  </template>
                  <template v-else-if="column.key === 'certLevel'">
                    <a-tag v-if="record.certLevel" color="green">{{ record.certLevel }}</a-tag>
                    <span v-else>-</span>
                  </template>
                  <template v-else-if="column.key === 'createTime'">
                    {{ new Date(record.createTime).toLocaleString() }}
                  </template>
                  <template v-else-if="column.key === 'action'">
                    <a-button type="link" size="small" @click="handleSuspendFarm(record)">暂停</a-button>
                  </template>
                </template>
              </a-table>
              <div class="load-more">
                <a-button :loading="farmLoading" @click="loadMoreFarms" :disabled="!farmBookmark">
                  {{ farmBookmark ? '加载更多' : '没有更多数据' }}
                </a-button>
              </div>
            </div>
          </a-card>
        </a-tab-pane>

        <a-tab-pane key="products" tab="🥬 产品管理">
          <a-card :bordered="false">
            <template #extra>
              <div class="card-extra">
                <a-input-search
                  v-model:value="searchProductId"
                  placeholder="输入产品ID精确查询"
                  style="width: 280px; margin-right: 12px;"
                  @search="handleSearchProduct"
                  @change="handleSearchProductChange"
                  allow-clear
                />
                <a-radio-group v-model:value="productStatusFilter" button-style="solid" size="small">
                  <a-radio-button value="">全部</a-radio-button>
                  <a-radio-button value="PLANTED">已种植</a-radio-button>
                  <a-radio-button value="HARVESTED">已采收</a-radio-button>
                  <a-radio-button value="CERTIFIED">已认证</a-radio-button>
                  <a-radio-button value="SHIPPING">运输中</a-radio-button>
                  <a-radio-button value="SOLD">已销售</a-radio-button>
                </a-radio-group>
              </div>
            </template>
            <div class="table-container">
              <a-table
                :columns="productColumns"
                :data-source="productList"
                :loading="productLoading"
                :pagination="false"
                :scroll="{ x: 1600, y: 'calc(100vh - 400px)' }"
                row-key="id"
              >
                <template #bodyCell="{ column, record }">
                  <template v-if="column.key === 'id'">
                    <a-tooltip :title="record.id">
                      <span style="margin-right:6px">{{ record.id.slice(0,8) }}...</span>
                      <copy-outlined class="copy-btn" @click.stop="copyToClipboard(record.id)" />
                    </a-tooltip>
                  </template>
                  <template v-else-if="column.key === 'status'">
                    <a-tag :color="getStatusColor(record.status)">{{ getStatusText(record.status) }}</a-tag>
                  </template>
                  <template v-else-if="column.key === 'createTime'">
                    {{ new Date(record.createTime).toLocaleString() }}
                  </template>
                  <template v-else-if="column.key === 'action'">
                    <a-space>
                      <a-button type="link" size="small" @click="handleHarvest(record)">采收</a-button>
                      <a-button type="link" size="small" @click="handleTrace(record)">溯源</a-button>
                    </a-space>
                  </template>
                </template>
              </a-table>
              <div class="load-more">
                <a-button :loading="productLoading" @click="loadMoreProducts" :disabled="!productBookmark">
                  {{ productBookmark ? '加载更多' : '没有更多数据' }}
                </a-button>
              </div>
            </div>
          </a-card>
        </a-tab-pane>

        <a-tab-pane key="trace" tab="🔍 溯源查询">
          <a-card :bordered="false">
            <div class="trace-search">
              <a-input-search
                v-model:value="traceProductId"
                placeholder="输入产品ID查询完整溯源信息"
                enter-button="查询溯源"
                size="large"
                @search="handleTraceSearch"
                style="max-width: 500px"
              />
              <a-input-search
                v-model:value="traceBatchNo"
                placeholder="输入批次号验证产品"
                enter-button="批次验证"
                size="large"
                @search="handleBatchSearch"
                style="max-width: 500px"
              />
            </div>

            <div v-if="traceInfo" class="trace-result">
              <a-descriptions title="📋 溯源信息" bordered :column="2">
                <a-descriptions-item label="产品名称" :span="2">{{ traceInfo.product?.name }}</a-descriptions-item>
                <a-descriptions-item label="产品类别">{{ traceInfo.product?.category }}</a-descriptions-item>
                <a-descriptions-item label="批次号">{{ traceInfo.product?.batchNo }}</a-descriptions-item>
                <a-descriptions-item label="产地名称" :span="2">{{ traceInfo.farm?.name }}</a-descriptions-item>
                <a-descriptions-item label="产地地址" :span="2">
                  {{ traceInfo.farm?.province }}{{ traceInfo.farm?.city }}{{ traceInfo.farm?.district }}{{ traceInfo.farm?.address }}
                </a-descriptions-item>
                <a-descriptions-item label="认证等级">{{ traceInfo.farm?.certLevel || '无' }}</a-descriptions-item>
                <a-descriptions-item label="农场主">{{ traceInfo.farm?.ownerName }}</a-descriptions-item>
                <a-descriptions-item label="产品状态">
                  <a-tag :color="getStatusColor(traceInfo.product?.status)">{{ getStatusText(traceInfo.product?.status) }}</a-tag>
                </a-descriptions-item>
              </a-descriptions>

              <a-divider />
              <h3>🧪 检测报告</h3>
              <a-descriptions v-if="traceInfo.inspection" bordered :column="2">
                <a-descriptions-item label="检测结果">
                  <a-tag :color="traceInfo.inspection.result === 'PASS' ? 'green' : 'red'">
                    {{ traceInfo.inspection.result === 'PASS' ? '✓ 通过' : '✗ 不通过' }}
                  </a-tag>
                </a-descriptions-item>
                <a-descriptions-item label="认证编号">{{ traceInfo.inspection.certNumber || '无' }}</a-descriptions-item>
                <a-descriptions-item label="农残检测">{{ traceInfo.inspection.pesticideResidue }}</a-descriptions-item>
                <a-descriptions-item label="重金属">{{ traceInfo.inspection.heavyMetal }}</a-descriptions-item>
                <a-descriptions-item label="微生物">{{ traceInfo.inspection.microorganism }}</a-descriptions-item>
                <a-descriptions-item label="等级">{{ traceInfo.inspection.grade }}</a-descriptions-item>
              </a-descriptions>
              <a-empty v-else description="暂无检测报告" />

              <a-divider />
              <h3>🚚 物流记录</h3>
              <a-list v-if="traceInfo.logistics?.length" :data-source="traceInfo.logistics">
                <template #renderItem="{ item }">
                  <a-list-item>
                    <a-list-item-meta
                      :title="`${item.fromProvince}${item.fromCity} → ${item.toProvince}${item.toCity}`"
                      :description="`运输方式: ${item.transportMode} | 承运方: ${item.transporter} | 温度: ${item.temperature} | 湿度: ${item.humidity}`"
                    >
                      <template #avatar>
                        <a-tag :color="getStatusColor(item.status)">{{ getStatusText(item.status) }}</a-tag>
                      </template>
                    </a-list-item-meta>
                  </a-list-item>
                </template>
              </a-list>
              <a-empty v-else description="暂无物流记录" />
            </div>
          </a-card>
        </a-tab-pane>

        <a-tab-pane key="contracts" tab="⚙️ 合约信息">
          <a-card :bordered="false">
            <a-table
              :columns="contractColumns"
              :data-source="contractList"
              :loading="contractLoading"
              :pagination="false"
              row-key="name"
            />
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
            <template #title>
              <span class="block-number">区块 #{{ block.block_num }}</span>
              <span class="block-time">{{ new Date(block.save_time).toLocaleString() }}</span>
            </template>
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

    <!-- 注册农场对话框 -->
    <a-modal v-model:visible="showFarmModal" title="🌾 注册新农场" @ok="handleFarmOk" @cancel="handleFarmCancel" :confirmLoading="farmModalLoading" width="700">
      <a-form ref="farmFormRef" :model="farmForm" :rules="farmRules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="农场名称" name="name">
              <a-input-group compact>
                <a-input v-model:value="farmForm.name" placeholder="请输入农场名称" style="width: calc(100% - 80px)" />
                <a-button @click="farmForm.name = generateRandomFarmName()">随机</a-button>
              </a-input-group>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="认证等级" name="certLevel">
              <a-select v-model:value="farmForm.certLevel" placeholder="选择认证等级">
                <a-select-option value="无公害">无公害</a-select-option>
                <a-select-option value="绿色食品">绿色食品</a-select-option>
                <a-select-option value="有机食品">有机食品</a-select-option>
                <a-select-option value="地理标志">地理标志</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="8"><a-form-item label="省份" name="province"><a-input v-model:value="farmForm.province" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="城市" name="city"><a-input v-model:value="farmForm.city" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="区县" name="district"><a-input v-model:value="farmForm.district" /></a-form-item></a-col>
        </a-row>
        <a-form-item label="详细地址" name="address"><a-input v-model:value="farmForm.address" /></a-form-item>
        <a-row :gutter="16">
          <a-col :span="8"><a-form-item label="面积(亩)" name="area"><a-input-number v-model:value="farmForm.area" :min="0.1" style="width:100%" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="负责人" name="ownerName"><a-input v-model:value="farmForm.ownerName" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="联系电话" name="ownerPhone"><a-input v-model:value="farmForm.ownerPhone" /></a-form-item></a-col>
        </a-row>
      </a-form>
    </a-modal>

    <!-- 注册产品对话框 -->
    <a-modal v-model:visible="showProductModal" title="🥬 注册新产品" @ok="handleProductOk" @cancel="handleProductCancel" :confirmLoading="productModalLoading" width="700">
      <a-form ref="productFormRef" :model="productForm" :rules="productRules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="产品名称" name="name">
              <a-input v-model:value="productForm.name" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="产品类别" name="category">
              <a-select v-model:value="productForm.category" @change="c => { productForm.name = generateRandomProductName(c) }">
                <a-select-option v-for="c in productCategories" :key="c" :value="c">{{ c }}</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="关联农场" name="farmId">
              <a-select v-model:value="productForm.farmId" show-search :filter-option="(input, option) => option.title?.includes(input)">
                <a-select-option v-for="f in farmList" :key="f.id" :title="f.name" :value="f.id">{{ f.name }} ({{ f.id.slice(0,8) }}...)</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="农场名称" name="farmName">
              <a-input v-model:value="productForm.farmName" disabled />
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="8"><a-form-item label="批次号" name="batchNo"><a-input v-model:value="productForm.batchNo" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="数量" name="quantity"><a-input-number v-model:value="productForm.quantity" :min="1" style="width:100%" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="单位" name="unit">
            <a-select v-model:value="productForm.unit">
              <a-select-option value="斤">斤</a-select-option>
              <a-select-option value="公斤">公斤</a-select-option>
              <a-select-option value="吨">吨</a-select-option>
              <a-select-option value="箱">箱</a-select-option>
            </a-select>
          </a-form-item></a-col>
        </a-row>
        <a-form-item label="描述" name="description"><a-textarea v-model:value="productForm.description" :rows="2" /></a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { message } from 'ant-design-vue';
import { PlusOutlined, AppstoreOutlined, ApartmentOutlined, CopyOutlined } from '@ant-design/icons-vue';
import { agricultureApi } from '../api';
import { ref, reactive, watch, onMounted } from 'vue';
import type { Farm, Product, TraceabilityInfo, ContractInfo, BlockData } from '../types';
import { getStatusText, getStatusColor, generateUUID, generateRandomFarmName, generateRandomFarmAddress, generateRandomProvince, generateRandomCity, generateRandomDistrict, generateRandomName, generateRandomPhone, generateRandomArea, generateRandomCategory, generateRandomProductName, generateBatchNo, copyToClipboard } from '../utils';

const activeTab = ref('farms');

// ====== 农场相关 ======
const farmList = ref<Farm[]>([]);
const farmLoading = ref(false);
const farmBookmark = ref('');
const farmStatusFilter = ref('');
const searchFarmId = ref('');
const showFarmModal = ref(false);
const farmModalLoading = ref(false);
const farmFormRef = ref();
const farmForm = reactive({ name: '', province: '', city: '', district: '', address: '', area: undefined as number | undefined, ownerName: '', ownerPhone: '', certLevel: '' });
const farmRules = { name: [{ required: true, message: '请输入农场名称' }], province: [{ required: true, message: '请输入省份' }], city: [{ required: true, message: '请输入城市' }] };

const farmColumns = [
  { title: '农场ID', dataIndex: 'id', key: 'id', width: 120 },
  { title: '名称', dataIndex: 'name', key: 'name', width: 140 },
  { title: '地址', key: 'address', width: 200, customRender: ({ record }: any) => `${record.province}${record.city}${record.district}${record.address}` },
  { title: '面积(亩)', dataIndex: 'area', key: 'area', width: 90 },
  { title: '负责人', dataIndex: 'ownerName', key: 'ownerName', width: 90 },
  { title: '认证等级', dataIndex: 'certLevel', key: 'certLevel', width: 100 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 80 },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime', width: 180 },
  { title: '操作', key: 'action', width: 80, fixed: 'right' },
];

const loadFarms = async () => {
  try {
    farmLoading.value = true;
    const result = await agricultureApi.getFarmList({ pageSize: 10, bookmark: farmBookmark.value, status: farmStatusFilter.value });
    if (!farmBookmark.value) farmList.value = result.records;
    else farmList.value = [...farmList.value, ...result.records];
    farmBookmark.value = result.bookmark;
  } catch (e: any) { message.error(e.message || '加载失败'); }
  finally { farmLoading.value = false; }
};

const loadMoreFarms = () => loadFarms();

watch(farmStatusFilter, () => { farmList.value = []; farmBookmark.value = ''; loadFarms(); });

const handleSearchFarm = async (v: string) => {
  if (!v) return;
  try { const r = await agricultureApi.getFarm(v); farmList.value = [r]; farmBookmark.value = ''; }
  catch (e: any) { message.error(e.message); farmList.value = []; }
};
const handleSearchFarmChange = (e: any) => { if (!e.target.value) { farmList.value = []; farmBookmark.value = ''; loadFarms(); } };

const handleFarmOk = () => {
  farmFormRef.value?.validate().then(async () => {
    farmModalLoading.value = true;
    try {
      await agricultureApi.registerFarm({ ...farmForm, id: generateUUID(), area: farmForm.area || 0 });
      message.success('农场注册成功');
      showFarmModal.value = false; farmFormRef.value?.resetFields();
      farmList.value = []; farmBookmark.value = ''; loadFarms();
    } catch (e: any) { message.error(e.message); }
    finally { farmModalLoading.value = false; }
  });
};
const handleFarmCancel = () => { showFarmModal.value = false; farmFormRef.value?.resetFields(); };

const handleSuspendFarm = async (record: Farm) => {
  try { await agricultureApi.suspendFarm({ id: record.id }); message.success('农场已暂停'); loadFarms(); }
  catch (e: any) { message.error(e.message); }
};

// ====== 产品相关 ======
const productList = ref<Product[]>([]);
const productLoading = ref(false);
const productBookmark = ref('');
const productStatusFilter = ref('');
const searchProductId = ref('');
const showProductModal = ref(false);
const productModalLoading = ref(false);
const productFormRef = ref();
const productCategories = ['茶叶', '咖啡', '水果', '蔬菜', '粮食', '肉类', '野生菌', '中药材', '花卉', '坚果', '豆类', '乳制品', '调味品', '加工食品'];
const productForm = reactive({ name: '', category: '', farmId: '', farmName: '', batchNo: '', quantity: undefined as number | undefined, unit: '斤', description: '' });
const productRules = { name: [{ required: true }], category: [{ required: true }], farmId: [{ required: true }], quantity: [{ required: true }] };

const productColumns = [
  { title: '产品ID', dataIndex: 'id', key: 'id', width: 120 },
  { title: '名称', dataIndex: 'name', key: 'name', width: 120 },
  { title: '类别', dataIndex: 'category', key: 'category', width: 80 },
  { title: '农场', dataIndex: 'farmName', key: 'farmName', width: 120 },
  { title: '批次号', dataIndex: 'batchNo', key: 'batchNo', width: 150 },
  { title: '数量', key: 'quantity', width: 100, customRender: ({ record }: any) => `${record.quantity}${record.unit}` },
  { title: '状态', dataIndex: 'status', key: 'status', width: 80 },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime', width: 180 },
  { title: '操作', key: 'action', width: 120, fixed: 'right' },
];

const loadProducts = async () => {
  try {
    productLoading.value = true;
    const result = await agricultureApi.getProductList({ pageSize: 10, bookmark: productBookmark.value, status: productStatusFilter.value });
    if (!productBookmark.value) productList.value = result.records;
    else productList.value = [...productList.value, ...result.records];
    productBookmark.value = result.bookmark;
  } catch (e: any) { message.error(e.message); }
  finally { productLoading.value = false; }
};

const loadMoreProducts = () => loadProducts();

watch(productStatusFilter, () => { productList.value = []; productBookmark.value = ''; loadProducts(); });

const handleSearchProduct = async (v: string) => {
  if (!v) return;
  try { const r = await agricultureApi.getProduct(v); productList.value = [r]; productBookmark.value = ''; }
  catch (e: any) { message.error(e.message); productList.value = []; }
};
const handleSearchProductChange = (e: any) => { if (!e.target.value) { productList.value = []; productBookmark.value = ''; loadProducts(); } };

const handleProductOk = () => {
  productFormRef.value?.validate().then(async () => {
    productModalLoading.value = true;
    try {
      await agricultureApi.registerProduct({ ...productForm, id: generateUUID(), quantity: productForm.quantity || 0 });
      message.success('产品注册成功');
      showProductModal.value = false; productFormRef.value?.resetFields();
      productList.value = []; productBookmark.value = ''; loadProducts();
    } catch (e: any) { message.error(e.message); }
    finally { productModalLoading.value = false; }
  });
};
const handleProductCancel = () => { showProductModal.value = false; productFormRef.value?.resetFields(); };

// 选择农场时自动填充
watch(() => productForm.farmId, async (farmId) => {
  if (!farmId) return;
  const farm = farmList.value.find(f => f.id === farmId);
  if (farm) productForm.farmName = farm.name;
});

const handleHarvest = async (record: Product) => {
  try { await agricultureApi.recordHarvest({ productId: record.id }); message.success('采收记录成功'); loadProducts(); }
  catch (e: any) { message.error(e.message); }
};

// ====== 溯源相关 ======
const traceProductId = ref('');
const traceBatchNo = ref('');
const traceInfo = ref<TraceabilityInfo | null>(null);

const handleTraceSearch = async (v: string) => {
  if (!v) return;
  try { traceInfo.value = await agricultureApi.getFullTraceability(v); }
  catch (e: any) { message.error(e.message); traceInfo.value = null; }
};

const handleBatchSearch = async (v: string) => {
  if (!v) return;
  try {
    const products = await agricultureApi.verifyByBatch(v);
    if (products.length > 0) {
      traceInfo.value = await agricultureApi.getFullTraceability(products[0].id);
      message.success(`找到 ${products.length} 个匹配产品`);
    } else { message.warning('未找到该批次产品'); }
  } catch (e: any) { message.error(e.message); }
};

// ====== 合约信息 ======
const contractList = ref<ContractInfo[]>([]);
const contractLoading = ref(false);

const loadContracts = async () => {
  contractLoading.value = true;
  try { contractList.value = await agricultureApi.getContractsInfo(); }
  catch (e: any) { message.error(e.message); }
  finally { contractLoading.value = false; }
};

const contractColumns = [
  { title: '合约名称', dataIndex: 'name', key: 'name', width: 180 },
  { title: '合约描述', dataIndex: 'description', key: 'description' },
  { title: '所属组织', dataIndex: 'org', key: 'org', width: 220 },
];

// ====== 区块 ======
const blockDrawer = ref(false);
const blockList = ref<BlockData[]>([]);
const blockTotal = ref(0);
const blockQuery = reactive({ pageSize: 10, pageNum: 1 });
const openBlockDrawer = async () => { blockDrawer.value = true; await fetchBlockList(); };
const fetchBlockList = async () => {
  try { const r = await agricultureApi.getBlockList({ pageSize: blockQuery.pageSize, pageNum: blockQuery.pageNum }); blockList.value = r.blocks; blockTotal.value = r.total; }
  catch (e) { console.error(e); }
};
const handleBlockPageChange = async (page: number, size: number) => { blockQuery.pageNum = page; blockQuery.pageSize = size; await fetchBlockList(); };

const handleTabChange = (key: string) => {
  if (key === 'contracts' && contractList.value.length === 0) loadContracts();
};

onMounted(() => { loadFarms(); loadProducts(); });
</script>

<style scoped>
.trace-search { display: flex; gap: 16px; margin-bottom: 24px; flex-wrap: wrap; }
.trace-result { margin-top: 24px; }
</style>