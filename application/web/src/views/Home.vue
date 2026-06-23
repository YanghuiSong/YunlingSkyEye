<template>
  <div class="home">
    <!-- 顶部横幅 -->
    <div class="hero-section">
      <div class="hero-particles"></div>
      <div class="hero-badge">
        <EyeOutlined /> Hyperledger Fabric · 联盟链
      </div>
      <h1 class="hero-title">👁️ 云岭天眼</h1>
      <p class="hero-subtitle">聚焦高原特色农业的产品溯源系统</p>
      <p class="hero-description">
        基于 Hyperledger Fabric 联盟链技术，深度融合云南高原特色农业场景<br/>
        六大智能合约驱动 · 产地到餐桌全链路可信追溯
      </p>
      <div class="hero-stats">
        <div class="stat-item">
          <span class="stat-value">6</span>
          <span class="stat-label">智能合约</span>
        </div>
        <div class="stat-item">
          <span class="stat-value">3</span>
          <span class="stat-label">联盟组织</span>
        </div>
        <div class="stat-item">
          <span class="stat-value">全链路</span>
          <span class="stat-label">溯源追踪</span>
        </div>
        <div class="stat-item">
          <span class="stat-value">高原</span>
          <span class="stat-label">特色农业</span>
        </div>
      </div>
    </div>

    <!-- 组织卡片 -->
    <div class="home-content">
      <router-link to="/agriculture" class="card">
        <div class="card-icon">
          <EnvironmentOutlined />
        </div>
        <div class="card-title">🌿 农业局（Org1）</div>
        <div class="card-description">高原农场注册 · 特色产品管理 · 全链溯源</div>
        <a-tag color="green" class="card-tag">FarmContract</a-tag>
        <a-tag color="cyan" class="card-tag">ProductContract</a-tag>
      </router-link>

      <router-link to="/inspection" class="card">
        <div class="card-icon">
          <SafetyCertificateOutlined />
        </div>
        <div class="card-title">🔬 检测认证中心（Org2）</div>
        <div class="card-description">质量检测 · 认证评级 · 证书颁发</div>
        <a-tag color="blue" class="card-tag">InspectionContract</a-tag>
      </router-link>

      <router-link to="/supply-chain" class="card">
        <div class="card-icon">
          <CarOutlined />
        </div>
        <div class="card-title">📦 供应链平台（Org3）</div>
        <div class="card-description">物流追踪 · 采购交易 · 订单管理</div>
        <a-tag color="purple" class="card-tag">LogisticsContract</a-tag>
        <a-tag color="orange" class="card-tag">TradeContract</a-tag>
      </router-link>
    </div>

    <!-- 系统流程图 -->
    <div class="flow-section">
      <h2 class="flow-title">📋 高原特色农产品全流程溯源体系</h2>
      <div class="flow-steps">
        <div class="flow-step">
          <span class="step-icon">🌱</span>
          <span class="step-text">产地种植</span>
        </div>
        <ArrowRightOutlined />
        <div class="flow-step">
          <span class="step-icon">🧪</span>
          <span class="step-text">检测认证</span>
        </div>
        <ArrowRightOutlined />
        <div class="flow-step">
          <span class="step-icon">🚚</span>
          <span class="step-text">冷链物流</span>
        </div>
        <ArrowRightOutlined />
        <div class="flow-step">
          <span class="step-icon">🛒</span>
          <span class="step-text">销售交易</span>
        </div>
        <ArrowRightOutlined />
        <div class="flow-step">
          <span class="step-icon">🔍</span>
          <span class="step-text">溯源查询</span>
        </div>
      </div>
    </div>
    <!-- 可视化入口 - 左下角技术入口 -->
    <router-link to="/blockchain-viz" class="tech-entry">
      <span class="tech-entry-icon">⛓️</span>
      <span class="tech-entry-text">链上架构可视化</span>
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { EnvironmentOutlined, SafetyCertificateOutlined, CarOutlined, EyeOutlined, ArrowRightOutlined } from '@ant-design/icons-vue';
</script>

<style scoped>
.home {
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f7ee 0%, #f0f2f5 100%);
}

.hero-section {
  padding: 60px 24px 36px;
  text-align: center;
  background: linear-gradient(135deg, #0a1628 0%, #1a2744 40%, #0d1f3c 70%, #162040 100%);
  color: #fff;
  position: relative;
  overflow: hidden;
}

.hero-particles {
  position: absolute;
  top: 0; left: 0; right: 0; bottom: 0;
  background:
    radial-gradient(circle at 20% 50%, rgba(64,169,255,0.08) 0%, transparent 50%),
    radial-gradient(circle at 80% 30%, rgba(82,196,26,0.05) 0%, transparent 50%),
    radial-gradient(circle at 50% 80%, rgba(255,215,0,0.04) 0%, transparent 50%);
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
  margin-bottom: 20px;
  position: relative;
}

.hero-title {
  font-size: 52px;
  font-weight: 800;
  margin: 0 0 6px;
  background: linear-gradient(135deg, #ffd700 0%, #ffaa00 50%, #ffd700 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  text-shadow: none;
  letter-spacing: 4px;
  position: relative;
}

.hero-subtitle {
  font-size: 20px;
  opacity: 0.9;
  margin: 0 0 8px;
}

.hero-description {
  font-size: 15px;
  opacity: 0.8;
  margin: 0 0 24px;
  line-height: 1.8;
}

.hero-stats {
  display: flex;
  justify-content: center;
  gap: 48px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stat-value {
  font-size: 28px;
  font-weight: 700;
}

.stat-label {
  font-size: 13px;
  opacity: 0.8;
}

.home-content {
  display: flex;
  gap: 24px;
  padding: 32px 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.card {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 32px 24px;
  border-radius: 12px;
  background: #fff;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.06);
  transition: all 0.3s;
  cursor: pointer;
  text-decoration: none;
  color: rgba(0, 0, 0, 0.85);
  border: 1px solid transparent;
}

.card:hover {
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  transform: translateY(-4px);
  border-color: #52c41a;
}

.card-icon {
  font-size: 48px;
  margin-bottom: 16px;
  color: #52c41a;
}

.card:nth-child(2) .card-icon { color: #1890ff; }
.card:nth-child(3) .card-icon { color: #722ed1; }

.card-title {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 8px;
}

.card-description {
  font-size: 14px;
  color: rgba(0, 0, 0, 0.45);
  margin-bottom: 12px;
}

.card-tag {
  margin: 2px;
}

/* 技术可视化入口 - 左下角悬浮按钮 */
.tech-entry {
  position: fixed;
  bottom: 24px;
  left: 24px;
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 14px;
  background: rgba(10, 22, 40, 0.85);
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
  background: rgba(10, 22, 40, 0.95);
  border-color: rgba(255,215,0,0.4);
  color: #ffd700;
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(0,0,0,0.3);
}
.tech-entry-icon { font-size: 16px; }
.tech-entry-text { font-size: 12px; }

.flow-section {
  padding: 32px 24px 48px;
  text-align: center;
  background: #f5f8fe;
}

.flow-title {
  font-size: 22px;
  color: #1a2744;
  margin-bottom: 32px;
}

.flow-steps {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  flex-wrap: wrap;
}

.flow-step {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.step-icon {
  font-size: 32px;
}

.step-text {
  font-size: 13px;
  color: rgba(0,0,0,0.65);
}

@media (max-width: 768px) {
  .home-content { flex-direction: column; }
  .hero-stats { gap: 24px; }
  .flow-steps { gap: 4px; }
}
</style> 