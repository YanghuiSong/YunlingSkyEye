<template>
  <div class="inspection">
    <div class="app-page-header">
      <a-page-header
        title="🔬 检测认证中心"
        sub-title="高原农产品质量检测 · 农药残留检测 · 重金属检测 · 认证评级 - Org2MSP"
        @back="() => $router.push('/')"
      >
        <template #extra>
          <a-tooltip title="创建新的检测报告">
            <a-button type="primary" @click="showInspectionModal = true">
              <template #icon><PlusOutlined /></template>
              创建检测报告
            </a-button>
          </a-tooltip>
        </template>
      </a-page-header>
    </div>

    <div class="app-content">
      <a-tabs v-model:activeKey="activeTab">
        <a-tab-pane key="reports" tab="📋 检测报告">
          <a-card :bordered="false">
            <template #extra>
              <div class="card-extra">
                <a-input-search v-model:value="searchId" placeholder="输入报告ID精确查询" style="width: 280px; margin-right: 12px;"
                  @search="handleSearch" @change="handleSearchChange" allow-clear />
                <a-radio-group v-model:value="resultFilter" button-style="solid" size="small">
                  <a-radio-button value="">全部</a-radio-button>
                  <a-radio-button value="PASS">通过</a-radio-button>
                  <a-radio-button value="FAIL">不通过</a-radio-button>
                </a-radio-group>
              </div>
            </template>
            <div class="table-container">
              <a-table :columns="columns" :data-source="reportList" :loading="loading" :pagination="false"
                :scroll="{ x: 1600, y: 'calc(100vh - 350px)' }" row-key="id">
                <template #bodyCell="{ column, record }">
                  <template v-if="column.key === 'id'">
                    <a-tooltip :title="record.id">
                      <span style="margin-right:6px">{{ record.id.slice(0,8) }}...</span>
                      <copy-outlined class="copy-btn" @click.stop="copyToClipboard(record.id)" />
                    </a-tooltip>
                  </template>
                  <template v-else-if="column.key === 'result'">
                    <a-tag :color="record.result === 'PASS' ? 'green' : 'red'">
                      {{ record.result === 'PASS' ? '✓ 通过' : '✗ 不通过' }}
                    </a-tag>
                  </template>
                  <template v-else-if="column.key === 'grade'"><a-tag color="blue">{{ record.grade }}</a-tag></template>
                  <template v-else-if="column.key === 'certNumber'"><a-tag color="green">{{ record.certNumber || '-' }}</a-tag></template>
                  <template v-else-if="column.key === 'createTime'">{{ new Date(record.createTime).toLocaleString() }}</template>
                </template>
              </a-table>
              <div class="load-more">
                <a-button :loading="loading" @click="loadMore" :disabled="!bookmark">
                  {{ bookmark ? '加载更多' : '没有更多数据' }}
                </a-button>
              </div>
            </div>
          </a-card>
        </a-tab-pane>

        <a-tab-pane key="certificates" tab="🏅 认证证书">
          <a-card :bordered="false">
            <a-table :columns="certColumns" :data-source="certList" :loading="certLoading" :pagination="false" row-key="certNumber">
              <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'issueDate'">{{ new Date(record.issueDate).toLocaleString() }}</template>
              </template>
            </a-table>
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

    <!-- 创建检测报告对话框 -->
    <a-modal v-model:visible="showInspectionModal" title="🧪 创建检测报告" @ok="handleInspectionOk" @cancel="handleInspectionCancel"
      :confirmLoading="modalLoading" width="720">
      <a-form ref="formRef" :model="formState" :rules="rules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="产品ID" name="productId">
              <a-input v-model:value="formState.productId" placeholder="输入产品ID" @change="handleProductIdChange" />
            </a-form-item>
          </a-col>
          <a-col :span="12"><a-form-item label="产品名称" name="productName"><a-input v-model:value="formState.productName" disabled /></a-form-item></a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12"><a-form-item label="检测人员" name="inspector"><a-input v-model:value="formState.inspector" /></a-form-item></a-col>
          <a-col :span="12"><a-form-item label="检测机构" name="inspectorOrg"><a-input v-model:value="formState.inspectorOrg" /></a-form-item></a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="8"><a-form-item label="农残检测结果" name="pesticideResidue"><a-input v-model:value="formState.pesticideResidue" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="重金属检测" name="heavyMetal"><a-input v-model:value="formState.heavyMetal" /></a-form-item></a-col>
          <a-col :span="8"><a-form-item label="微生物检测" name="microorganism"><a-input v-model:value="formState.microorganism" /></a-form-item></a-col>
        </a-row>
        <a-form-item label="检测结论" name="conclusion"><a-textarea v-model:value="formState.conclusion" :rows="2" /></a-form-item>
        <a-form-item label="认证编号（通过时填写）" name="certNumber"><a-input v-model:value="formState.certNumber" placeholder="如: ZX-JC-2024-00001" /></a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { message } from 'ant-design-vue';
import { PlusOutlined, ApartmentOutlined, CopyOutlined } from '@ant-design/icons-vue';
import { inspectionApi } from '../api';
import { ref, reactive, watch, onMounted } from 'vue';
import type { InspectionReport, BlockData } from '../types';
import { getStatusText, getStatusColor, generateUUID, generateRandomName, generateCertNumber, generateRandomPesticideResidue, generateRandomHeavyMetal, generateRandomMicroorganism, copyToClipboard } from '../utils';

const activeTab = ref('reports');
const showInspectionModal = ref(false);
const modalLoading = ref(false);
const formRef = ref();

const formState = reactive({
  productId: '', productName: '', inspector: '', inspectorOrg: '',
  pesticideResidue: '', heavyMetal: '', microorganism: '', conclusion: '', certNumber: '',
});
const rules = {
  productId: [{ required: true, message: '请输入产品ID' }],
  inspector: [{ required: true, message: '请输入检测人员' }],
  pesticideResidue: [{ required: true, message: '请输入农残检测结果' }],
};

const handleProductIdChange = async (e: any) => {
  const id = e.target.value;
  if (!id) { formState.productName = ''; return; }
  try {
    const { agricultureApi } = await import('../api');
    const product = await agricultureApi.getProduct(id);
    formState.productName = product.name;
  } catch { formState.productName = '（未找到产品）'; }
};

const handleInspectionOk = () => {
  formRef.value?.validate().then(async () => {
    modalLoading.value = true;
    try {
      await inspectionApi.createInspection({ ...formState, id: generateUUID(), inspectorOrg: formState.inspectorOrg || '省级检测认证中心' });
      message.success('检测报告创建成功');
      showInspectionModal.value = false; formRef.value?.resetFields();
      reportList.value = []; bookmark.value = ''; loadReports();
    } catch (e: any) { message.error(e.message); }
    finally { modalLoading.value = false; }
  });
};
const handleInspectionCancel = () => { showInspectionModal.value = false; formRef.value?.resetFields(); };

// ====== 报告列表 ======
const reportList = ref<InspectionReport[]>([]);
const loading = ref(false);
const bookmark = ref('');
const resultFilter = ref('');
const searchId = ref('');

const columns = [
  { title: '报告ID', dataIndex: 'id', key: 'id', width: 120 },
  { title: '产品名称', dataIndex: 'productName', key: 'productName', width: 120 },
  { title: '检测人员', dataIndex: 'inspector', key: 'inspector', width: 90 },
  { title: '农残', dataIndex: 'pesticideResidue', key: 'pesticideResidue', width: 100 },
  { title: '重金属', dataIndex: 'heavyMetal', key: 'heavyMetal', width: 100 },
  { title: '检测结果', dataIndex: 'result', key: 'result', width: 90 },
  { title: '等级', dataIndex: 'grade', key: 'grade', width: 80 },
  { title: '认证编号', dataIndex: 'certNumber', key: 'certNumber', width: 160 },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime', width: 180 },
];

const loadReports = async () => {
  try {
    loading.value = true;
    const result = await inspectionApi.getInspectionList({ pageSize: 10, bookmark: bookmark.value, result: resultFilter.value });
    if (!bookmark.value) reportList.value = result.records;
    else reportList.value = [...reportList.value, ...result.records];
    bookmark.value = result.bookmark;
  } catch (e: any) { message.error(e.message); }
  finally { loading.value = false; }
};

const loadMore = () => loadReports();

watch(resultFilter, () => { reportList.value = []; bookmark.value = ''; loadReports(); });

const handleSearch = async (v: string) => {
  if (!v) return;
  try { const r = await inspectionApi.getInspection(v); reportList.value = [r]; bookmark.value = ''; }
  catch (e: any) { message.error(e.message); reportList.value = []; }
};
const handleSearchChange = (e: any) => { if (!e.target.value) { reportList.value = []; bookmark.value = ''; loadReports(); } };

// ====== 证书列表 ======
const certList = ref<any[]>([]);
const certLoading = ref(false);
const certColumns = [
  { title: '证书编号', dataIndex: 'certNumber', key: 'certNumber', width: 180 },
  { title: '产品ID', dataIndex: 'productId', key: 'productId', width: 120 },
  { title: '产品名称', dataIndex: 'productName', key: 'productName', width: 120 },
  { title: '等级', dataIndex: 'grade', key: 'grade', width: 80 },
  { title: '颁发时间', dataIndex: 'issueDate', key: 'issueDate', width: 180 },
];

watch(activeTab, async (key) => {
  if (key === 'certificates' && certList.value.length === 0) {
    certLoading.value = true;
    try { const r = await inspectionApi.getCertificates({}); certList.value = r.records || []; }
    catch (e: any) { message.error(e.message); }
    finally { certLoading.value = false; }
  }
});

// ====== 区块 ======
const blockDrawer = ref(false);
const blockList = ref<BlockData[]>([]);
const blockTotal = ref(0);
const blockQuery = reactive({ pageSize: 10, pageNum: 1 });
const openBlockDrawer = async () => { blockDrawer.value = true; await fetchBlockList(); };
const fetchBlockList = async () => {
  try { const r = await inspectionApi.getBlockList({}); blockList.value = r.blocks; blockTotal.value = r.total; }
  catch (e) { console.error(e); }
};
const handleBlockPageChange = async (page: number, size: number) => { blockQuery.pageNum = page; blockQuery.pageSize = size; await fetchBlockList(); };

onMounted(() => loadReports());
</script>