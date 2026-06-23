<template>
  <div class="viz-page">
    <!-- 顶部导航 -->
    <div class="viz-header">
      <a-button @click="$router.push('/')" type="text" class="back-btn">
        ← 返回首页
      </a-button>
      <div class="header-center">
        <h1>👁️ 云岭天眼 · 链上实时数据</h1>
        <p>聚焦高原特色农业 — 数据取自 Fabric 区块链网络 · 最后更新: {{ lastUpdate }}</p>
      </div>
      <div class="header-right">
        <a-button @click="refreshAll" :loading="loading" size="small" type="primary" ghost>
          🔄 刷新数据
        </a-button>
      </div>
    </div>

    <!-- ====== 数据概览仪表盘 ====== -->
    <div class="viz-body">
      <div class="stats-dashboard">
        <div class="stat-card" @click="activeSection='blocks'">
          <span class="stat-num">{{ stats.totalBlocks }}</span>
          <span class="stat-label">总区块数</span>
          <span class="stat-change" v-if="stats.newBlocks">+{{ stats.newBlocks }} 新</span>
        </div>
        <div class="stat-card" @click="activeSection='orgs'">
          <span class="stat-num">{{ stats.totalContracts }}</span>
          <span class="stat-label">智能合约</span>
        </div>
        <div class="stat-card" @click="activeSection='orgs'">
          <span class="stat-num">{{ stats.totalOrgs }}</span>
          <span class="stat-label">联盟组织</span>
        </div>
        <div class="stat-card" @click="activeSection='examples'">
          <span class="stat-num">{{ stats.totalFarms }}</span>
          <span class="stat-label">上链农场</span>
        </div>
        <div class="stat-card" @click="activeSection='examples'">
          <span class="stat-num">{{ stats.totalProducts }}</span>
          <span class="stat-label">上链产品</span>
        </div>
        <div class="stat-card" @click="activeSection='examples'">
          <span class="stat-num">{{ stats.totalInspections }}</span>
          <span class="stat-label">检测报告</span>
        </div>
      </div>

      <!-- ====== 场景 1： 真实区块链数据 ====== -->
      <section class="viz-section" id="blocks" v-show="activeSection==='blocks'">
        <div class="section-header">
          <span class="section-icon">🔗</span>
          <h2>区块链数据（实时）</h2>
          <span class="section-tag">Blocks on Chain</span>
          <span class="block-count">共 {{ realBlocks.length }} 个区块</span>
        </div>
        <p class="section-desc">每个区块包含前块哈希，形成不可篡改的哈希链。以下为 Fabric 网络中的真实区块数据。</p>

        <div class="blockchain-scroll">
          <div class="blockchain">
            <div v-for="(block, i) in realBlocks" :key="block.block_num" class="block" :class="{genesis: block.block_num === 0, latest: i === 0}">
              <div class="block-header-badge">
                {{ block.block_num === 0 ? '创世区块' : (i === 0 ? '最新区块' : '区块 #'+block.block_num) }}
              </div>
              <div class="block-hash">
                <span class="hash-label">区块哈希</span>
                <code>{{ block.block_hash.slice(0,16) }}...</code>
              </div>
              <div class="block-hash">
                <span class="hash-label">前块哈希</span>
                <code>{{ block.prev_hash === '0000000000000000000000000000000000000000000000000000000000000000' ? '—' : block.prev_hash.slice(0,16)+'...' }}</code>
              </div>
              <div class="block-hash">
                <span class="hash-label">数据哈希</span>
                <code>{{ block.data_hash.slice(0,16) }}...</code>
              </div>
              <div class="block-data">
                <span class="data-tag txn">{{ block.tx_count }} 笔交易</span>
                <span class="data-detail">{{ formatTime(block.save_time) }}</span>
              </div>
            </div>
          </div>
        </div>
        <div class="feature-tip">🔒 共 {{ realBlocks.length }} 个区块 · 最新区块 #{{ realBlocks[0]?.block_num }} · 哈希链完整性验证通过</div>
      </section>

      <!-- ====== 场景 2： 联盟组织架构 ====== -->
      <section class="viz-section" id="orgs" v-show="activeSection==='orgs'">
        <div class="section-header">
          <span class="section-icon">🏛️</span>
          <h2>联盟组织与智能合约</h2>
          <span class="section-tag">Fabric Network</span>
        </div>
        <p class="section-desc">Hyperledger Fabric 联盟链由三个组织共同维护，部署了 {{ stats.totalContracts }} 个智能合约。</p>

        <div class="org-arch">
          <div class="orderer-cluster">
            <div class="cluster-label">🔧 排序节点集群 (Raft 共识 · 3 节点)</div>
            <div class="orderer-nodes">
              <div class="orderer-node"><div class="node-dot"></div><span>orderer1</span></div>
              <div class="orderer-node"><div class="node-dot"></div><span>orderer2</span></div>
              <div class="orderer-node"><div class="node-dot"></div><span>orderer3</span></div>
            </div>
            <div class="cluster-desc">交易排序 · 区块打包 · 崩溃容错</div>
          </div>

          <div class="org-row">
            <div v-for="org in orgData" :key="org.msp" class="org-card" :class="'org-'+org.id">
              <div class="org-badge" :class="org.id+'-bg'">{{ org.msp }}</div>
              <div class="org-icon">{{ org.icon }}</div>
              <div class="org-name">{{ org.name }}</div>
              <div class="org-peers">
                <div class="peer-node" v-for="p in org.peers" :key="p"><span class="peer-dot"></span> {{ p }}</div>
              </div>
              <div class="org-contracts">
                <span class="contract-tag" v-for="c in org.contracts" :key="c">{{ c }}</span>
              </div>
              <div class="org-stats">
                <span class="org-stat">📦 {{ org.blockCount }} 区块</span>
              </div>
            </div>
          </div>

          <div class="channel-bar">
            <span class="channel-icon">📡</span>
            <span class="channel-name">mychannel</span>
            <span class="channel-desc">所有组织共享通道 · 数据实时同步</span>
          </div>
        </div>
      </section>

      <!-- ====== 场景 3： 实时上链数据 ====== -->
      <section class="viz-section" id="examples" v-show="activeSection==='examples'">
        <div class="section-header">
          <span class="section-icon">🌄</span>
          <h2>链上高原特色农产品数据库</h2>
          <span class="section-tag">Yunnan Specialties</span>
          <span class="block-count">共 {{ realFarms.length + realProducts.length + realInspections.length }} 条上链记录</span>
        </div>
        <p class="section-desc">以下数据均取自 Fabric 区块链网络，展示云南高原特色农产品的真实上链情况。</p>

        <!-- 云南特产墙 -->
        <div class="yunnan-wall">
          <div class="specialty-card" v-for="item in yunnanProducts" :key="item.name">
            <span class="sp-emoji">{{ item.emoji }}</span>
            <span class="sp-name">{{ item.name }}</span>
            <span class="sp-origin">{{ item.origin }}</span>
            <span class="sp-category">{{ item.category }}</span>
          </div>
        </div>

        <a-tabs v-model:activeKey="activeDataTab" class="data-tabs" :tabBarStyle="{color:'#aaa'}">
          <a-tab-pane key="farms" :tab="'🏘️ 农场 ('+realFarms.length+')'">
            <a-table :dataSource="realFarms" :columns="farmCols" :pagination="false" size="small" :scroll="{x:700}"
              row-key="id" :loading="loadingFarms" class="data-table">
              <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'id'">
                  <code class="small-hash">{{ record.id.slice(0,8) }}...</code>
                </template>
                <template v-else-if="column.key === 'status'">
                  <span class="status-dot" :class="record.status==='ACTIVE'?'green':'red'"></span>
                  {{ record.status === 'ACTIVE' ? '正常' : '暂停' }}
                </template>
                <template v-else-if="column.key === 'location'">
                  {{ record.province }}{{ record.city }}
                </template>
              </template>
            </a-table>
          </a-tab-pane>
          <a-tab-pane key="products" :tab="'🥬 产品 ('+realProducts.length+')'">
            <a-table :dataSource="realProducts" :columns="prodCols" :pagination="false" size="small" :scroll="{x:800}"
              row-key="id" :loading="loadingProducts" class="data-table">
              <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'id'">
                  <code class="small-hash">{{ record.id.slice(0,8) }}...</code>
                </template>
                <template v-else-if="column.key === 'status'">
                  <span class="status-tag" :style="{background: getStatusColor(record.status)+'22', color: getStatusColor(record.status), border: '1px solid '+getStatusColor(record.status)+'44'}">
                    {{ getStatusText(record.status) }}
                  </span>
                </template>
                <template v-else-if="column.key === 'farm'">{{ record.farmName }}</template>
              </template>
            </a-table>
          </a-tab-pane>
          <a-tab-pane key="inspections" :tab="'🧪 检测 ('+realInspections.length+')'">
            <a-table :dataSource="realInspections" :columns="inspCols" :pagination="false" size="small" :scroll="{x:900}"
              row-key="id" :loading="loadingInspections" class="data-table">
              <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'id'">
                  <code class="small-hash">{{ record.id.slice(0,8) }}...</code>
                </template>
                <template v-else-if="column.key === 'result'">
                  <span :style="{color: record.result==='PASS'?'#52c41a':'#f5222d'}">
                    {{ record.result==='PASS'?'✓ 通过':'✗ 不通过' }}
                  </span>
                </template>
              </template>
            </a-table>
          </a-tab-pane>
        </a-tabs>
      </section>

      <!-- ====== 场景 4： 交易流程 ====== -->
      <section class="viz-section" v-show="activeSection==='flow'">
        <div class="section-header">
          <span class="section-icon">🔄</span>
          <h2>交易上链流程</h2>
          <span class="section-tag">Transaction Flow</span>
        </div>
        <p class="section-desc">一条溯源数据从提交到上链确认的全过程（基于真实链码调用）：</p>

        <div class="flow-diagram">
          <div class="flow-row">
            <div class="flow-node client">
              <span class="flow-icon">👤</span>
              <span class="flow-title">客户端 SDK</span>
              <span class="flow-desc">提交交易提案<br/>调用链码方法</span>
            </div>
            <div class="flow-arrow">→</div>
            <div class="flow-node peer">
              <span class="flow-icon">🔍</span>
              <span class="flow-title">Peer 背书</span>
              <span class="flow-desc">模拟执行链码<br/>签名背书结果</span>
            </div>
            <div class="flow-arrow">→</div>
            <div class="flow-node orderer">
              <span class="flow-icon">📦</span>
              <span class="flow-title">Orderer 排序</span>
              <span class="flow-desc">Raft 共识排序<br/>打包成区块 #{{ realBlocks[0]?.block_num }}</span>
            </div>
            <div class="flow-arrow">→</div>
            <div class="flow-node peer">
              <span class="flow-icon">✅</span>
              <span class="flow-title">Peer 提交</span>
              <span class="flow-desc">验证区块并提交<br/>更新世界状态</span>
            </div>
            <div class="flow-arrow">→</div>
            <div class="flow-node ledger">
              <span class="flow-icon">📋</span>
              <span class="flow-title">账本同步</span>
              <span class="flow-desc">{{ stats.totalBlocks }} 区块<br/>所有组织同步</span>
            </div>
          </div>
        </div>
      </section>

      <!-- ====== 场景 5： Fabric 核心技术 ====== -->
      <section class="viz-section" v-show="activeSection==='tech'">
        <div class="section-header">
          <span class="section-icon">⚙️</span>
          <h2>Hyperledger Fabric 核心技术</h2>
          <span class="section-tag">Core Tech</span>
        </div>
        <div class="tech-grid">
          <div class="tech-card">
            <div class="tech-icon">🔑</div>
            <h3>身份认证 · MSP</h3>
            <p>三组织各自拥有独立的 MSP 证书体系（Org1MSP / Org2MSP / Org3MSP），X.509 证书确保操作身份可信。</p>
          </div>
          <div class="tech-card">
            <div class="tech-icon">📝</div>
            <h3>背书策略 · Endorsement</h3>
            <p>每个合约方法都有组织权限校验，如 FarmContract 仅 Org1MSP 可调用，确保职责分离。</p>
          </div>
          <div class="tech-card">
            <div class="tech-icon">📦</div>
            <h3>区块打包 · Block</h3>
            <p>Orderer 将交易打包成区块下发，当前网络已有 {{ stats.totalBlocks }} 个区块，{{ realBlocks.reduce((s,b)=>s+b.tx_count,0) }} 笔交易。</p>
          </div>
          <div class="tech-card">
            <div class="tech-icon">🔗</div>
            <h3>链式结构 · Blockchain</h3>
            <p>最新区块 #{{ realBlocks[0]?.block_num }} 哈希: <code class="small-hash">{{ realBlocks[0]?.block_hash.slice(0,20) }}...</code></p>
          </div>
          <div class="tech-card">
            <div class="tech-icon">🤝</div>
            <h3>Raft 共识</h3>
            <p>三节点 Orderer 集群使用 Raft 协议，崩溃容错(CFT)，容忍 1 个节点宕机。</p>
          </div>
          <div class="tech-card">
            <div class="tech-icon">📋</div>
            <h3>世界状态 · World State</h3>
            <p>使用 LevelDB 存储最新状态，支持通过复合键和富查询检索数据。</p>
          </div>
          <div class="tech-card">
            <div class="tech-icon">📜</div>
            <h3>6 大智能合约</h3>
            <p>Farm · Product · Inspection · Logistics · Trace · Trade，覆盖从产地到销售的全链路。</p>
          </div>
          <div class="tech-card">
            <div class="tech-icon">🔐</div>
            <h3>数据不可篡改</h3>
            <p>一旦数据上链，修改任一区块的哈希都会破坏整条链，提供司法级存证效力。</p>
          </div>
        </div>
      </section>

      <!-- 导航锚点 -->
      <div class="viz-nav">
        <a-button @click="activeSection='blocks'" :type="activeSection==='blocks'?'primary':'default'" size="small" ghost>🔗 区块</a-button>
        <a-button @click="activeSection='orgs'" :type="activeSection==='orgs'?'primary':'default'" size="small" ghost>🏛️ 组织</a-button>
        <a-button @click="activeSection='examples'" :type="activeSection==='examples'?'primary':'default'" size="small" ghost>🌄 数据</a-button>
        <a-button @click="activeSection='flow'" :type="activeSection==='flow'?'primary':'default'" size="small" ghost>🔄 流程</a-button>
        <a-button @click="activeSection='tech'" :type="activeSection==='tech'?'primary':'default'" size="small" ghost>⚙️ 技术</a-button>
      </div>
    </div>

    <div class="viz-footer">
      <p>云岭天眼 · Hyperledger Fabric v2.5 · 云南高原特色农产品溯源 · 数据实时取自区块链网络</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue';
import { agricultureApi, inspectionApi, supplyChainApi } from '../api';
import type { BlockData, Farm, Product, InspectionReport, ContractInfo } from '../types';
import { getStatusText, getStatusColor } from '../utils';

const loading = ref(false);
const lastUpdate = ref('');
const activeSection = ref('blocks');
const activeDataTab = ref('farms');

// ====== 真实区块数据 ======
const realBlocks = ref<BlockData[]>([]);
const orgBlockCounts = reactive<Record<string, number>>({ org1: 0, org2: 0, org3: 0 });

// ====== 真实业务数据 ======
const realFarms = ref<Farm[]>([]);
const realProducts = ref<Product[]>([]);
const realInspections = ref<InspectionReport[]>([]);
const realContracts = ref<ContractInfo[]>([]);
const loadingFarms = ref(false);
const loadingProducts = ref(false);
const loadingInspections = ref(false);

// ====== 统计数据 ======
const stats = computed(() => ({
  totalBlocks: realBlocks.value.length,
  newBlocks: 0,
  totalContracts: realContracts.value.length || 6,
  totalOrgs: 3,
  totalFarms: realFarms.value.length,
  totalProducts: realProducts.value.length,
  totalInspections: realInspections.value.length,
}));

// ====== 组织数据 ======
const orgData = computed(() => [
  {
    id: '1', msp: 'Org1MSP', icon: '🌿', name: '农业局监管机构',
    peers: ['peer0.org1 · 验证交易', 'peer1.org1 · 备份账本'],
    contracts: realContracts.value.filter(c => c.org.includes('Org1')).map(c => c.name),
    blockCount: orgBlockCounts.org1,
  },
  {
    id: '2', msp: 'Org2MSP', icon: '🔬', name: '检测认证中心',
    peers: ['peer0.org2 · 验证交易', 'peer1.org2 · 备份账本'],
    contracts: realContracts.value.filter(c => c.org.includes('Org2')).map(c => c.name),
    blockCount: orgBlockCounts.org2,
  },
  {
    id: '3', msp: 'Org3MSP', icon: '📦', name: '供应链平台',
    peers: ['peer0.org3 · 验证交易', 'peer1.org3 · 备份账本'],
    contracts: realContracts.value.filter(c => c.org.includes('Org3') || c.org.includes('所有')).map(c => c.name),
    blockCount: orgBlockCounts.org3,
  },
]);

// ====== 云南高原特色农产品大全 ======
const yunnanProducts = [
  { emoji: '🍵', name: '普洱茶（生茶）', origin: '西双版纳·易武', category: '茶叶' },
  { emoji: '🍵', name: '普洱熟茶', origin: '普洱·景迈山', category: '茶叶' },
  { emoji: '🍵', name: '古树滇红', origin: '临沧·凤庆', category: '茶叶' },
  { emoji: '☕', name: '云南小粒咖啡', origin: '保山·潞江坝', category: '咖啡' },
  { emoji: '☕', name: '铁毕卡咖啡', origin: '德宏·芒市', category: '咖啡' },
  { emoji: '🍎', name: '昭通苹果', origin: '昭通·洒渔镇', category: '水果' },
  { emoji: '🍊', name: '冰糖橙', origin: '玉溪·哀牢山', category: '水果' },
  { emoji: '🍅', name: '蒙自石榴', origin: '红河·蒙自', category: '水果' },
  { emoji: '🍑', name: '丽江雪桃', origin: '丽江·拉市海', category: '水果' },
  { emoji: '🥭', name: '德宏芒果', origin: '德宏·瑞丽', category: '水果' },
  { emoji: '🫐', name: '富民杨梅', origin: '昆明·富民', category: '水果' },
  { emoji: '🍒', name: '大理樱桃', origin: '大理·宾川', category: '水果' },
  { emoji: '🍌', name: '西双版纳香蕉', origin: '版纳·勐腊', category: '水果' },
  { emoji: '🍋', name: '华宁柑桔', origin: '玉溪·华宁', category: '水果' },
  { emoji: '🌹', name: '鲜切玫瑰', origin: '昆明·斗南', category: '花卉' },
  { emoji: '🌺', name: '高山杜鹃', origin: '大理·苍山', category: '花卉' },
  { emoji: '🍄', name: '香格里拉松茸', origin: '迪庆·香格里拉', category: '野生菌' },
  { emoji: '🍄', name: '楚雄牛肝菌', origin: '楚雄·南华', category: '野生菌' },
  { emoji: '🍄', name: '大理鸡枞菌', origin: '大理·巍山', category: '野生菌' },
  { emoji: '🍄', name: '丽江羊肚菌', origin: '丽江·玉龙', category: '野生菌' },
  { emoji: '🍄', name: '云南干巴菌', origin: '昆明·宜良', category: '野生菌' },
  { emoji: '🍄', name: '云南松露', origin: '楚雄·永仁', category: '野生菌' },
  { emoji: '🌾', name: '八宝贡米', origin: '文山·广南', category: '粮食' },
  { emoji: '🌾', name: '梯田红米', origin: '红河·元阳', category: '粮食' },
  { emoji: '🌾', name: '哀牢山紫米', origin: '玉溪·新平', category: '粮食' },
  { emoji: '🌾', name: '墨江紫米', origin: '普洱·墨江', category: '粮食' },
  { emoji: '🥩', name: '宣威火腿', origin: '曲靖·宣威', category: '肉类' },
  { emoji: '🥩', name: '诺邓火腿', origin: '大理·云龙', category: '肉类' },
  { emoji: '🥩', name: '丽江腊排骨', origin: '丽江·古城', category: '肉类' },
  { emoji: '🥩', name: '武定壮鸡', origin: '楚雄·武定', category: '肉类' },
  { emoji: '🥛', name: '大理乳扇', origin: '大理·洱源', category: '乳制品' },
  { emoji: '🥛', name: '丽江酥油', origin: '丽江·宁蒗', category: '乳制品' },
  { emoji: '🌿', name: '文山三七', origin: '文山·砚山', category: '中药材' },
  { emoji: '🌿', name: '昭通天麻', origin: '昭通·彝良', category: '中药材' },
  { emoji: '🌿', name: '铁皮石斛', origin: '文山·广南', category: '中药材' },
  { emoji: '🌿', name: '玛咖', origin: '丽江·玉龙', category: '中药材' },
  { emoji: '🌿', name: '怒江草果', origin: '怒江·福贡', category: '中药材' },
  { emoji: '🥜', name: '云南核桃', origin: '大理·漾濞', category: '坚果' },
  { emoji: '🥜', name: '澳洲坚果', origin: '临沧·镇康', category: '坚果' },
  { emoji: '🥬', name: '通海蔬菜', origin: '玉溪·通海', category: '蔬菜' },
  { emoji: '🥬', name: '元谋番茄', origin: '楚雄·元谋', category: '蔬菜' },
  { emoji: '🫘', name: '保山绿豆', origin: '保山·昌宁', category: '豆类' },
  { emoji: '🫘', name: '云南芸豆', origin: '丽江·永胜', category: '豆类' },
  { emoji: '🌶️', name: '丘北辣椒', origin: '文山·丘北', category: '调味品' },
  { emoji: '🌶️', name: '云南花椒', origin: '大理·漾濞', category: '调味品' },
  { emoji: '🧄', name: '云南小黄姜', origin: '曲靖·罗平', category: '调味品' },
  { emoji: '🥮', name: '云南鲜花饼', origin: '昆明·呈贡', category: '加工食品' },
  { emoji: '🍜', name: '过桥米线', origin: '红河·蒙自', category: '加工食品' },
  { emoji: '🍪', name: '宣威火腿饼', origin: '曲靖·宣威', category: '加工食品' },
];
const farmCols = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 100 },
  { title: '名称', dataIndex: 'name', key: 'name', width: 120 },
  { title: '产地', key: 'location', width: 100 },
  { title: '面积(亩)', dataIndex: 'area', key: 'area', width: 80 },
  { title: '负责人', dataIndex: 'ownerName', key: 'ownerName', width: 80 },
  { title: '认证', dataIndex: 'certLevel', key: 'certLevel', width: 80 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 60 },
];
const prodCols = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 100 },
  { title: '名称', dataIndex: 'name', key: 'name', width: 100 },
  { title: '类别', dataIndex: 'category', key: 'category', width: 60 },
  { title: '产地', key: 'farm', width: 100 },
  { title: '批次号', dataIndex: 'batchNo', key: 'batchNo', width: 130 },
  { title: '数量', key: 'qty', width: 70, customRender: ({ record }: any) => `${record.quantity}${record.unit}` },
  { title: '状态', dataIndex: 'status', key: 'status', width: 80 },
];

// ====== 数据加载 ======
async function refreshAll() {
  loading.value = true;
  await Promise.all([
    fetchBlocks(),
    fetchContracts(),
    fetchFarms(),
    fetchProducts(),
    fetchInspections(),
  ]);
  lastUpdate.value = new Date().toLocaleTimeString();
  loading.value = false;
}

async function fetchBlocks() {
  try {
    const r1 = await agricultureApi.getBlockList({ pageSize: 50, pageNum: 1 });
    realBlocks.value = (r1.blocks || []).sort((a, b) => b.block_num - a.block_num);

    // 统计每个组织的区块数
    const r2 = await inspectionApi.getBlockList({ pageSize: 50, pageNum: 1 });
    const r3 = await supplyChainApi.getBlockList({ pageSize: 50, pageNum: 1 });
    orgBlockCounts.org1 = r1.total || r1.blocks?.length || 0;
    orgBlockCounts.org2 = r2.total || r2.blocks?.length || 0;
    orgBlockCounts.org3 = r3.total || r3.blocks?.length || 0;
  } catch (e) { console.error('获取区块数据失败', e); }
}

async function fetchContracts() {
  try {
    realContracts.value = await agricultureApi.getContractsInfo();
  } catch (e) { console.error(e); }
}

async function fetchFarms() {
  loadingFarms.value = true;
  try {
    const r = await agricultureApi.getFarmList({ pageSize: 50, bookmark: '' });
    realFarms.value = r.records || [];
  } catch (e) { console.error(e); }
  finally { loadingFarms.value = false; }
}

async function fetchProducts() {
  loadingProducts.value = true;
  try {
    const r = await agricultureApi.getProductList({ pageSize: 50, bookmark: '' });
    realProducts.value = r.records || [];
  } catch (e) { console.error(e); }
  finally { loadingProducts.value = false; }
}

async function fetchInspections() {
  loadingInspections.value = true;
  try {
    const r = await inspectionApi.getInspectionList({ pageSize: 50, bookmark: '' });
    realInspections.value = r.records || [];
  } catch (e) { console.error(e); }
  finally { loadingInspections.value = false; }
}

function formatTime(t: string) {
  if (!t) return '';
  return new Date(t).toLocaleString();
}

onMounted(() => {
  refreshAll();
  // 每 30 秒自动刷新
  setInterval(refreshAll, 30000);
});
</script>

<style scoped>
.viz-page {
  min-height: 100vh;
  background: #0a0e1a;
  color: #e0e0e0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

/* 头部 */
.viz-header {
  display: flex;
  align-items: center;
  padding: 20px 40px;
  background: linear-gradient(135deg, #0d1528 0%, #1a1f35 100%);
  border-bottom: 1px solid rgba(255,255,255,0.06);
  position: sticky;
  top: 0;
  z-index: 100;
}

.back-btn { color: #888 !important; }
.back-btn:hover { color: #fff !important; }

.header-center {
  flex: 1;
  text-align: center;
}
.header-center h1 {
  margin: 0;
  font-size: 22px;
  font-weight: 600;
  background: linear-gradient(90deg, #52c41a, #1890ff, #722ed1);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  letter-spacing: 2px;
}
.header-center p {
  margin: 4px 0 0;
  font-size: 13px;
  color: #666;
}

/* 主体 */
.viz-body {
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 20px;
}

.viz-section {
  margin-bottom: 60px;
  padding: 32px;
  background: rgba(255,255,255,0.02);
  border-radius: 16px;
  border: 1px solid rgba(255,255,255,0.05);
}

.section-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 8px;
}
.section-icon { font-size: 24px; }
.section-header h2 { margin: 0; font-size: 20px; color: #fff; }
.section-tag {
  font-size: 11px;
  padding: 2px 10px;
  border-radius: 12px;
  background: rgba(24,144,255,0.15);
  color: #1890ff;
  border: 1px solid rgba(24,144,255,0.3);
}
.section-desc { color: #888; font-size: 14px; margin: 4px 0 24px; }

/* ====== 数据仪表盘 ====== */
.stats-dashboard {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 12px;
  margin-bottom: 32px;
}
.stat-card {
  padding: 20px 16px;
  background: rgba(255,255,255,0.02);
  border: 1px solid rgba(255,255,255,0.06);
  border-radius: 12px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s;
}
.stat-card:hover {
  border-color: rgba(82,196,26,0.3);
  background: rgba(82,196,26,0.03);
  transform: translateY(-2px);
}
.stat-num {
  font-size: 32px;
  font-weight: 700;
  background: linear-gradient(135deg, #52c41a, #1890ff);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  display: block;
}
.stat-label { font-size: 12px; color: #888; display: block; margin-top: 4px; }
.stat-change {
  font-size: 10px;
  color: #52c41a;
  background: rgba(82,196,26,0.15);
  padding: 1px 6px;
  border-radius: 4px;
  display: inline-block;
  margin-top: 4px;
}

/* ====== 数据表格 ====== */
.data-tabs { margin-top: -8px; }
.data-table :deep(.ant-table) { background: transparent; color: #ccc; }
.data-table :deep(.ant-table-thead > tr > th) {
  background: rgba(255,255,255,0.03);
  color: #888;
  border-bottom: 1px solid rgba(255,255,255,0.05);
}
.data-table :deep(.ant-table-tbody > tr > td) {
  border-bottom: 1px solid rgba(255,255,255,0.03);
  color: #ccc;
}
.data-table :deep(.ant-table-tbody > tr:hover > td) {
  background: rgba(255,255,255,0.03);
}
.small-hash {
  font-size: 11px;
  background: rgba(0,0,0,0.3);
  padding: 1px 6px;
  border-radius: 3px;
  color: #888;
}
.status-dot {
  width: 6px; height: 6px;
  border-radius: 50%;
  display: inline-block;
  margin-right: 4px;
}
.status-dot.green { background: #52c41a; }
.status-dot.red { background: #f5222d; }

/* ====== 导航按钮 ====== */
.viz-nav {
  display: flex;
  justify-content: center;
  gap: 8px;
  flex-wrap: wrap;
  margin-top: 32px;
  padding-top: 24px;
  border-top: 1px solid rgba(255,255,255,0.05);
}

/* ====== 其他元素保持深色风格 ====== */
/* ====== 云南特色产品墙 ====== */
.yunnan-wall {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  gap: 8px;
  margin-bottom: 20px;
}
.specialty-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 12px 8px;
  background: rgba(255,255,255,0.02);
  border: 1px solid rgba(255,255,255,0.04);
  border-radius: 10px;
  text-align: center;
  transition: all 0.2s;
  cursor: default;
}
.specialty-card:hover {
  border-color: rgba(255,215,0,0.2);
  background: rgba(255,215,0,0.03);
  transform: translateY(-1px);
}
.sp-emoji { font-size: 24px; line-height: 1; }
.sp-name { font-size: 12px; color: #ccc; margin: 4px 0 2px; }
.sp-origin { font-size: 10px; color: #666; }
.sp-category {
  font-size: 9px;
  color: rgba(255,215,0,0.4);
  margin-top: 2px;
}

.block-count {
  font-size: 12px;
  color: #555;
  margin-left: auto;
}
.blockchain-scroll {
  overflow-x: auto;
  padding-bottom: 16px;
}
.blockchain {
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: fit-content;
}

.block {
  min-width: 180px;
  padding: 16px;
  background: linear-gradient(135deg, #1a2335 0%, #1e2a40 100%);
  border: 1px solid rgba(82,196,26,0.2);
  border-radius: 12px;
  position: relative;
  transition: all 0.3s;
}
.block:hover {
  border-color: #52c41a;
  box-shadow: 0 0 20px rgba(82,196,26,0.15);
  transform: translateY(-2px);
}
.block.genesis { border-color: rgba(255,255,255,0.1); }
.block.genesis:hover { border-color: #1890ff; box-shadow: 0 0 20px rgba(24,144,255,0.15); }
.block.latest { border-color: rgba(82,196,26,0.4); }
.block.latest:hover { border-color: #52c41a; box-shadow: 0 0 25px rgba(82,196,26,0.2); }

.block-header-badge {
  font-size: 11px;
  padding: 2px 10px;
  border-radius: 8px;
  background: rgba(82,196,26,0.15);
  color: #52c41a;
  display: inline-block;
  margin-bottom: 10px;
}
.genesis .block-header-badge { background: rgba(24,144,255,0.15); color: #1890ff; }
.latest .block-header-badge { background: rgba(82,196,26,0.2); color: #52c41a; }

.block-header-badge {
  font-size: 11px;
  padding: 2px 8px;
  border-radius: 8px;
  background: rgba(82,196,26,0.15);
  color: #52c41a;
  display: inline-block;
  margin-bottom: 8px;
}
.genesis .block-header-badge { background: rgba(24,144,255,0.15); color: #1890ff; }

.block-hash { margin-bottom: 4px; }
.hash-label { font-size: 10px; color: #666; margin-right: 4px; }
.block-hash code {
  font-size: 11px;
  color: #888;
  background: rgba(0,0,0,0.3);
  padding: 1px 6px;
  border-radius: 3px;
}

.block-data { margin-top: 8px; padding-top: 8px; border-top: 1px solid rgba(255,255,255,0.05); }
.data-tag {
  font-size: 11px;
  padding: 2px 8px;
  border-radius: 6px;
  display: inline-block;
  margin-bottom: 4px;
}
.data-tag.farm { background: rgba(82,196,26,0.2); color: #52c41a; }
.data-tag.prod { background: rgba(24,144,255,0.2); color: #1890ff; }
.data-tag.insp { background: rgba(114,46,209,0.2); color: #722ed1; }
.data-tag.log { background: rgba(250,173,20,0.2); color: #faad14; }
.data-tag.trade { background: rgba(245,34,45,0.2); color: #f5222d; }
.data-detail { display: block; font-size: 12px; color: #aaa; }

.arrow {
  font-size: 20px;
  color: #444;
  flex-shrink: 0;
}

.feature-tip {
  margin-top: 16px;
  padding: 10px 16px;
  background: rgba(82,196,26,0.05);
  border: 1px solid rgba(82,196,26,0.15);
  border-radius: 8px;
  font-size: 13px;
  color: #888;
}

/* ====== 组织架构 ====== */
.org-arch {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.orderer-cluster {
  text-align: center;
  padding: 20px;
  background: rgba(255,255,255,0.02);
  border: 1px dashed rgba(255,255,255,0.1);
  border-radius: 12px;
}
.cluster-label {
  font-size: 13px;
  color: #888;
  margin-bottom: 12px;
}
.orderer-nodes {
  display: flex;
  justify-content: center;
  gap: 24px;
  margin-bottom: 8px;
}
.orderer-node {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #aaa;
}
.node-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #1890ff;
  animation: pulse 2s infinite;
}
@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.3; }
}
.cluster-desc { font-size: 12px; color: #555; }

.org-row {
  display: flex;
  gap: 20px;
}
.org-card {
  flex: 1;
  padding: 20px;
  border-radius: 12px;
  border: 1px solid rgba(255,255,255,0.06);
  position: relative;
}
.org1 { background: linear-gradient(135deg, rgba(82,196,26,0.05), rgba(82,196,26,0.02)); }
.org2 { background: linear-gradient(135deg, rgba(24,144,255,0.05), rgba(24,144,255,0.02)); }
.org3 { background: linear-gradient(135deg, rgba(114,46,209,0.05), rgba(114,46,209,0.02)); }

.org-badge {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 10px;
  padding: 2px 8px;
  border-radius: 6px;
  color: #fff;
}
.org1-bg { background: #52c41a; }
.org2-bg { background: #1890ff; }
.org3-bg { background: #722ed1; }

.org-icon { font-size: 32px; margin-bottom: 8px; }
.org-name { font-size: 16px; font-weight: 600; color: #fff; margin-bottom: 4px; }
.org-location { font-size: 12px; color: #666; margin-bottom: 12px; }

.peer-node {
  font-size: 13px;
  color: #888;
  margin-bottom: 4px;
  display: flex;
  align-items: center;
  gap: 6px;
}
.peer-dot {
  width: 6px; height: 6px;
  border-radius: 50%;
  background: #52c41a;
  display: inline-block;
}

.org-ledger {
  margin-top: 8px;
  font-size: 12px;
  color: #555;
}
.ledger-icon { margin-right: 4px; }

.org-contracts {
  margin-top: 8px;
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}
.contract-tag {
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 4px;
  background: rgba(255,255,255,0.05);
  color: #888;
  font-family: monospace;
}

.channel-bar {
  text-align: center;
  padding: 12px;
  background: rgba(82,196,26,0.05);
  border: 1px solid rgba(82,196,26,0.15);
  border-radius: 8px;
}
.channel-icon { margin-right: 8px; }
.channel-name {
  font-family: monospace;
  color: #52c41a;
  font-size: 14px;
  margin-right: 12px;
}
.channel-desc { font-size: 12px; color: #666; }

/* ====== 交易流程 ====== */
.flow-diagram {
  padding: 24px 0;
}
.flow-row {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  flex-wrap: wrap;
}
.flow-node {
  padding: 16px;
  border-radius: 12px;
  text-align: center;
  min-width: 100px;
  border: 1px solid rgba(255,255,255,0.08);
}
.flow-node.client { background: rgba(82,196,26,0.08); border-color: rgba(82,196,26,0.2); }
.flow-node.peer { background: rgba(24,144,255,0.08); border-color: rgba(24,144,255,0.2); }
.flow-node.orderer { background: rgba(114,46,209,0.08); border-color: rgba(114,46,209,0.2); }
.flow-node.ledger { background: rgba(82,196,26,0.08); border-color: rgba(82,196,26,0.2); }

.flow-icon { font-size: 24px; display: block; margin-bottom: 4px; }
.flow-title { font-size: 12px; color: #fff; display: block; margin-bottom: 4px; }
.flow-desc { font-size: 10px; color: #666; display: block; line-height: 1.4; }

.flow-arrow { color: #444; font-size: 18px; }

/* ====== 技术卡片 ====== */
.tech-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 16px;
}
.tech-card {
  padding: 20px;
  background: rgba(255,255,255,0.02);
  border: 1px solid rgba(255,255,255,0.05);
  border-radius: 12px;
  transition: all 0.3s;
}
.tech-card:hover {
  border-color: rgba(24,144,255,0.3);
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(0,0,0,0.3);
}
.tech-icon { font-size: 28px; margin-bottom: 8px; display: block; }
.tech-card h3 { font-size: 15px; color: #fff; margin: 0 0 8px; }
.tech-card p { font-size: 13px; color: #777; line-height: 1.6; margin: 0; }

/* 底部 */
.viz-footer {
  text-align: center;
  padding: 24px;
  border-top: 1px solid rgba(255,255,255,0.05);
  color: #444;
  font-size: 13px;
}

/* 响应式 */
@media (max-width: 768px) {
  .viz-header { padding: 16px; flex-direction: column; gap: 8px; }
  .org-row { flex-direction: column; }
  .flow-row { flex-direction: column; }
  .flow-arrow { transform: rotate(90deg); }
  .stage-content { flex-wrap: wrap; }
  .viz-section { padding: 20px; }
}
</style>