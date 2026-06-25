<template>
  <div class="home">
    <!-- ====== 顶部横幅 ====== -->
    <div class="hero-section">
      <div class="hero-particles" ref="particlesRef"></div>
      <div class="hero-glow"></div>
      <div class="hero-badge fade-in">
        <EyeOutlined /> Hyperledger Fabric · 联盟链
      </div>
      <h1 class="hero-title slide-up">👁️ 云岭天眼</h1>
      <p class="hero-subtitle slide-up-delay">聚焦高原特色农业的产品溯源系统</p>
      <p class="hero-description slide-up-delay2">
        基于 Hyperledger Fabric 联盟链技术，深度融合云南高原特色农业场景<br/>
        六大智能合约驱动 · 产地到餐桌全链路可信追溯
      </p>
      <div class="hero-stats slide-up-delay2">
        <div class="stat-item" v-for="s in heroStats" :key="s.label">
          <span class="stat-value" ref="statRefs">{{ s.value }}</span>
          <span class="stat-label">{{ s.label }}</span>
        </div>
      </div>
      <div class="hero-cta slide-up-delay2">
        <a-button type="primary" size="large" class="cta-btn" @click="scrollToCards">
          🚀 开始体验
          <ArrowDownOutlined />
        </a-button>
      </div>
    </div>

    <!-- ====== 组织卡片 ====== -->
    <div class="home-content" id="org-cards">
      <div class="section-label fade-in-up">
        <span class="label-line"></span>
        <span class="label-text">选择组织身份进入系统</span>
        <span class="label-line"></span>
      </div>
      <div class="cards-row">
        <router-link to="/agriculture" class="card card-org1 fade-in-up-card">
          <div class="card-glow"></div>
          <div class="card-icon-wrap"><EnvironmentOutlined /></div>
          <div class="card-title">🌿 农业局</div>
          <div class="card-subtitle">Org1MSP · 监管机构</div>
          <div class="card-desc">高原农场注册 · 特色产品管理 · 全链溯源</div>
          <div class="card-tags">
            <span class="tag">FarmContract</span>
            <span class="tag">ProductContract</span>
          </div>
          <div class="card-action">进入面板 →</div>
        </router-link>

        <router-link to="/inspection" class="card card-org2 fade-in-up-card">
          <div class="card-glow"></div>
          <div class="card-icon-wrap"><SafetyCertificateOutlined /></div>
          <div class="card-title">🔬 检测认证中心</div>
          <div class="card-subtitle">Org2MSP · 检测机构</div>
          <div class="card-desc">质量检测 · 认证评级 · 证书颁发</div>
          <div class="card-tags">
            <span class="tag">InspectionContract</span>
          </div>
          <div class="card-action">进入面板 →</div>
        </router-link>

        <router-link to="/supply-chain" class="card card-org3 fade-in-up-card">
          <div class="card-glow"></div>
          <div class="card-icon-wrap"><CarOutlined /></div>
          <div class="card-title">📦 供应链平台</div>
          <div class="card-subtitle">Org3MSP · 流通枢纽</div>
          <div class="card-desc">物流追踪 · 采购交易 · 订单管理</div>
          <div class="card-tags">
            <span class="tag">LogisticsContract</span>
            <span class="tag">TradeContract</span>
          </div>
          <div class="card-action">进入面板 →</div>
        </router-link>
      </div>
    </div>

    <!-- ====== 溯源流程 ====== -->
    <div class="flow-section">
      <div class="section-label fade-in-up">
        <span class="label-line"></span>
        <span class="label-text">📋 高原特色农产品全流程溯源体系</span>
        <span class="label-line"></span>
      </div>
      <div class="flow-steps">
        <div class="flow-step" v-for="(step, i) in flowSteps" :key="i" :style="{ animationDelay: i * 0.15 + 's' }">
          <div class="step-icon-wrap">
            <span class="step-icon">{{ step.icon }}</span>
            <div class="step-number">{{ i + 1 }}</div>
          </div>
          <span class="step-text">{{ step.text }}</span>
          <span class="step-desc">{{ step.desc }}</span>
        </div>
      </div>
    </div>

    <!-- ====== 技术栈 ====== -->
    <div class="tech-section fade-in-up">
      <div class="section-label">
        <span class="label-line"></span>
        <span class="label-text">⚙️ 技术栈</span>
        <span class="label-line"></span>
      </div>
      <div class="tech-badges">
        <div class="tech-badge" v-for="t in techStack" :key="t.name">
          <span class="tech-icon">{{ t.icon }}</span>
          <span class="tech-name">{{ t.name }}</span>
          <span class="tech-version">{{ t.version }}</span>
        </div>
      </div>
    </div>

    <!-- 可视化入口 -->
    <router-link to="/blockchain-viz" class="tech-entry">
      <span class="tech-entry-icon">⛓️</span>
      <span class="tech-entry-text">链上架构可视化</span>
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { EnvironmentOutlined, SafetyCertificateOutlined, CarOutlined, EyeOutlined, ArrowDownOutlined } from '@ant-design/icons-vue';

const heroStats = [
  { value: '6', label: '智能合约' },
  { value: '3', label: '联盟组织' },
  { value: '49+', label: '特色产品' },
  { value: '全链路', label: '溯源追踪' },
];

const flowSteps = [
  { icon: '🌱', text: '产地种植', desc: 'FarmContract' },
  { icon: '🧪', text: '检测认证', desc: 'InspectionContract' },
  { icon: '🚚', text: '冷链物流', desc: 'LogisticsContract' },
  { icon: '🛒', text: '销售交易', desc: 'TradeContract' },
  { icon: '🔍', text: '溯源查询', desc: 'TraceContract' },
];

const techStack = [
  { icon: '🔗', name: 'Hyperledger Fabric', version: 'v2.5.10' },
  { icon: '🦫', name: 'Go + Gin', version: 'v1.10' },
  { icon: '⚡', name: 'Vue 3 + Vite', version: 'v4.5' },
  { icon: '🎨', name: 'Ant Design Vue', version: 'v3.2' },
  { icon: '🗄️', name: 'BoltDB + LevelDB', version: '' },
  { icon: '📦', name: 'Docker Compose', version: 'v2.x' },
];

function scrollToCards() {
  document.getElementById('org-cards')?.scrollIntoView({ behavior: 'smooth' });
}

onMounted(() => {
  // Intersection Observer for scroll animations
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('visible');
      }
    });
  }, { threshold: 0.1 });

  document.querySelectorAll('.fade-in-up, .fade-in-up-card').forEach(el => observer.observe(el));
});
</script>

<style scoped>
.home {
  min-height: 100vh;
  background: linear-gradient(180deg, #f0f7ee 0%, #f5f8fe 50%, #f0f2f5 100%);
  overflow-x: hidden;
}

/* ====== Hero 区域 ====== */
.hero-section {
  padding: 80px 24px 50px;
  text-align: center;
  background: linear-gradient(135deg, #0a1628 0%, #1a2744 40%, #0d1f3c 70%, #162040 100%);
  color: #fff;
  position: relative;
  overflow: hidden;
  min-height: 520px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.hero-glow {
  position: absolute;
  top: -30%;
  left: 50%;
  transform: translateX(-50%);
  width: 600px;
  height: 600px;
  background: radial-gradient(circle, rgba(255,215,0,0.06) 0%, transparent 70%);
  pointer-events: none;
}

.hero-particles {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(1px 1px at 10% 20%, rgba(255,215,0,0.3) 0%, transparent 100%),
    radial-gradient(1px 1px at 30% 60%, rgba(82,196,26,0.2) 0%, transparent 100%),
    radial-gradient(1.5px 1.5px at 50% 80%, rgba(255,215,0,0.3) 0%, transparent 100%),
    radial-gradient(1px 1px at 70% 30%, rgba(24,144,255,0.2) 0%, transparent 100%),
    radial-gradient(1px 1px at 90% 70%, rgba(255,215,0,0.2) 0%, transparent 100%),
    radial-gradient(1.5px 1.5px at 20% 90%, rgba(82,196,26,0.2) 0%, transparent 100%),
    radial-gradient(1px 1px at 60% 15%, rgba(255,215,0,0.2) 0%, transparent 100%),
    radial-gradient(1px 1px at 80% 55%, rgba(24,144,255,0.2) 0%, transparent 100%);
  pointer-events: none;
}

.hero-badge {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 20px;
  background: rgba(255,215,0,0.1);
  border: 1px solid rgba(255,215,0,0.2);
  border-radius: 20px;
  font-size: 13px;
  color: #ffd700;
  margin-bottom: 24px;
  position: relative;
}

.hero-title {
  font-size: 56px;
  font-weight: 800;
  margin: 0 0 8px;
  background: linear-gradient(135deg, #ffd700 0%, #ffaa00 50%, #ffd700 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  letter-spacing: 4px;
  line-height: 1.2;
}

.hero-subtitle {
  font-size: 22px;
  opacity: 0.9;
  margin: 0 0 10px;
  font-weight: 500;
}

.hero-description {
  font-size: 15px;
  opacity: 0.7;
  margin: 0 0 30px;
  line-height: 1.8;
  max-width: 600px;
}

.hero-stats {
  display: flex;
  gap: 48px;
  margin-bottom: 28px;
}

.stat-item { display: flex; flex-direction: column; align-items: center; }
.stat-value { font-size: 30px; font-weight: 700; color: #ffd700; }
.stat-label { font-size: 13px; opacity: 0.7; margin-top: 4px; }

.hero-cta { margin-top: 4px; }
.cta-btn {
  height: 48px;
  padding: 0 32px;
  font-size: 16px;
  border-radius: 24px;
  background: linear-gradient(135deg, #ffd700, #ffaa00);
  border: none;
  color: #1a2744;
  font-weight: 600;
  box-shadow: 0 4px 20px rgba(255,215,0,0.3);
  transition: all 0.3s;
}
.cta-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 30px rgba(255,215,0,0.4);
  color: #1a2744;
}

/* ====== 动画 ====== */
.fade-in { animation: fadeIn 0.8s ease-out; }
.slide-up { animation: slideUp 0.8s ease-out; }
.slide-up-delay { animation: slideUp 0.8s ease-out 0.15s both; }
.slide-up-delay2 { animation: slideUp 0.8s ease-out 0.3s both; }
@keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }
@keyframes slideUp {
  from { opacity: 0; transform: translateY(30px); }
  to { opacity: 1; transform: translateY(0); }
}

/* ====== 组织卡片 ====== */
.home-content {
  padding: 48px 24px 32px;
  max-width: 1100px;
  margin: 0 auto;
}

.section-label {
  display: flex;
  align-items: center;
  gap: 16px;
  justify-content: center;
  margin-bottom: 32px;
}
.label-line {
  width: 40px; height: 2px;
  background: linear-gradient(90deg, transparent, #52c41a, transparent);
}
.label-text { font-size: 20px; font-weight: 600; color: #1a2744; white-space: nowrap; }

.cards-row {
  display: flex;
  gap: 24px;
}

.card {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px 24px 32px;
  border-radius: 16px;
  background: #fff;
  box-shadow: 0 2px 16px rgba(0,0,0,0.04);
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
  cursor: pointer;
  text-decoration: none;
  color: rgba(0,0,0,0.85);
  position: relative;
  overflow: hidden;
  border: 1px solid rgba(0,0,0,0.04);
}

.card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 16px 48px rgba(0,0,0,0.1);
}

.card-glow {
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle at 50% 0%, rgba(82,196,26,0.06), transparent 60%);
  opacity: 0;
  transition: opacity 0.4s;
  pointer-events: none;
}
.card-org2 .card-glow { background: radial-gradient(circle at 50% 0%, rgba(24,144,255,0.06), transparent 60%); }
.card-org3 .card-glow { background: radial-gradient(circle at 50% 0%, rgba(114,46,209,0.06), transparent 60%); }
.card:hover .card-glow { opacity: 1; }

.card-icon-wrap {
  width: 72px; height: 72px;
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32px;
  margin-bottom: 16px;
  transition: all 0.4s;
}
.card-org1 .card-icon-wrap { background: rgba(82,196,26,0.1); color: #52c41a; }
.card-org2 .card-icon-wrap { background: rgba(24,144,255,0.1); color: #1890ff; }
.card-org3 .card-icon-wrap { background: rgba(114,46,209,0.1); color: #722ed1; }
.card:hover .card-icon-wrap { transform: scale(1.1) rotate(-5deg); }

.card-title { font-size: 17px; font-weight: 600; margin-bottom: 4px; }
.card-subtitle { font-size: 12px; color: rgba(0,0,0,0.35); margin-bottom: 10px; }
.card-desc { font-size: 13px; color: rgba(0,0,0,0.55); text-align: center; margin-bottom: 14px; line-height: 1.5; }

.card-tags { display: flex; gap: 6px; flex-wrap: wrap; justify-content: center; margin-bottom: 16px; }
.tag {
  font-size: 10px;
  font-family: monospace;
  padding: 3px 10px;
  border-radius: 6px;
  background: rgba(82,196,26,0.08);
  color: #52c41a;
  border: 1px solid rgba(82,196,26,0.12);
}
.card-org2 .tag { background: rgba(24,144,255,0.08); color: #1890ff; border-color: rgba(24,144,255,0.12); }
.card-org3 .tag { background: rgba(114,46,209,0.08); color: #722ed1; border-color: rgba(114,46,209,0.12); }

.card-action {
  font-size: 13px;
  font-weight: 500;
  opacity: 0.5;
  transition: all 0.3s;
}
.card:hover .card-action { opacity: 1; }
.card-org1:hover .card-action { color: #52c41a; }
.card-org2:hover .card-action { color: #1890ff; }
.card-org3:hover .card-action { color: #722ed1; }

/* ====== 溯源流程 ====== */
.flow-section {
  padding: 48px 24px 56px;
  background: #fff;
}

.flow-steps {
  display: flex;
  align-items: flex-start;
  justify-content: center;
  gap: 0;
  max-width: 900px;
  margin: 0 auto;
  position: relative;
}

.flow-step {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  flex: 1;
  position: relative;
  animation: slideUp 0.6s ease-out both;
}

.step-icon-wrap {
  width: 72px; height: 72px;
  border-radius: 50%;
  background: linear-gradient(135deg, #f0f7ee, #e8f5e9);
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  border: 2px solid #e0e0e0;
  transition: all 0.3s;
}
.flow-step:hover .step-icon-wrap {
  border-color: #52c41a;
  transform: scale(1.1);
  box-shadow: 0 4px 20px rgba(82,196,26,0.15);
}
.step-icon { font-size: 28px; }
.step-number {
  position: absolute;
  top: -6px;
  right: -6px;
  width: 22px; height: 22px;
  border-radius: 50%;
  background: #52c41a;
  color: #fff;
  font-size: 11px;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: center;
}

.step-text { font-size: 14px; font-weight: 500; color: #333; }
.step-desc {
  font-size: 10px;
  color: #999;
  font-family: monospace;
  padding: 2px 8px;
  background: #f5f5f5;
  border-radius: 4px;
}

/* 步骤之间的连接线 */
.flow-step:not(:last-child)::after {
  content: '';
  position: absolute;
  top: 36px;
  right: -50%;
  width: 100%;
  height: 2px;
  background: linear-gradient(90deg, #52c41a, #e0e0e0);
  z-index: 0;
}

/* ====== 技术栈 ====== */
.tech-section {
  padding: 48px 24px;
  max-width: 900px;
  margin: 0 auto;
}

.tech-badges {
  display: flex;
  justify-content: center;
  gap: 12px;
  flex-wrap: wrap;
}
.tech-badge {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 18px;
  background: #fff;
  border: 1px solid #eee;
  border-radius: 10px;
  transition: all 0.3s;
}
.tech-badge:hover {
  border-color: #52c41a;
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(0,0,0,0.06);
}
.tech-icon { font-size: 18px; }
.tech-name { font-size: 13px; font-weight: 500; color: #333; }
.tech-version { font-size: 11px; color: #999; }

/* ====== 滚动动画 ====== */
.fade-in-up { opacity: 0; transform: translateY(30px); transition: all 0.6s ease-out; }
.fade-in-up.visible { opacity: 1; transform: translateY(0); }
.fade-in-up-card { opacity: 0; transform: translateY(30px); transition: all 0.6s ease-out; }
.fade-in-up-card:nth-child(1) { transition-delay: 0s; }
.fade-in-up-card:nth-child(2) { transition-delay: 0.1s; }
.fade-in-up-card:nth-child(3) { transition-delay: 0.2s; }
.fade-in-up-card.visible { opacity: 1; transform: translateY(0); }

/* ====== 可视化入口 ====== */
.tech-entry {
  position: fixed;
  bottom: 24px;
  left: 24px;
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 14px;
  background: rgba(10,22,40,0.85);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(255,215,0,0.15);
  border-radius: 20px;
  color: #888;
  font-size: 12px;
  text-decoration: none;
  z-index: 999;
  transition: all 0.3s;
  cursor: pointer;
}
.tech-entry:hover {
  background: rgba(10,22,40,0.95);
  border-color: rgba(255,215,0,0.4);
  color: #ffd700;
  transform: translateY(-2px);
}

/* ====== 响应式 ====== */
@media (max-width: 768px) {
  .hero-title { font-size: 36px; }
  .hero-stats { gap: 24px; flex-wrap: wrap; justify-content: center; }
  .cards-row { flex-direction: column; }
  .flow-steps { flex-direction: column; align-items: center; gap: 16px; }
  .flow-step:not(:last-child)::after { display: none; }
  .tech-badges { gap: 8px; }
}
</style>