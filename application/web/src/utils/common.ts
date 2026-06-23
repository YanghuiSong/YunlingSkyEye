// 生成UUID
export const generateUUID = () => {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
    const r = Math.random() * 16 | 0;
    const v = c === 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
};

// 格式化状态文本（农产品溯源）
export const getStatusText = (status: string) => {
  const statusMap: Record<string, string> = {
    'ACTIVE': '正常经营',
    'SUSPENDED': '已暂停',
    'PLANTED': '已种植',
    'HARVESTED': '已采收',
    'INSPECTING': '检测中',
    'CERTIFIED': '已认证',
    'SHIPPING': '运输中',
    'SOLD': '已销售',
    'RECALLED': '已召回',
    'PASS': '检测通过',
    'FAIL': '检测不通过',
    'PREPARING': '准备发货',
    'IN_TRANSIT': '运输中',
    'DELIVERED': '已送达',
    'PENDING': '待处理',
    'PAID': '已付款',
    'SHIPPED': '已发货',
    'CONFIRMED': '已确认',
    'CANCELLED': '已取消',
  };
  return statusMap[status] || '未知';
};

// 获取状态对应的颜色
export const getStatusColor = (status: string) => {
  const colorMap: Record<string, string> = {
    'ACTIVE': 'green',
    'SUSPENDED': 'red',
    'PLANTED': 'cyan',
    'HARVESTED': 'orange',
    'INSPECTING': 'blue',
    'CERTIFIED': 'green',
    'SHIPPING': 'purple',
    'SOLD': 'green',
    'RECALLED': 'red',
    'PASS': 'green',
    'FAIL': 'red',
    'PREPARING': 'blue',
    'IN_TRANSIT': 'purple',
    'DELIVERED': 'green',
    'PENDING': 'blue',
    'PAID': 'cyan',
    'SHIPPED': 'purple',
    'CONFIRMED': 'green',
    'CANCELLED': 'red',
  };
  return colorMap[status] || 'default';
};

// 格式化金额显示
export const formatPrice = (price: number) => {
  return `¥ ${price}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',');
}; 