// 随机生成中文姓名
const lastNames = ['张', '王', '李', '赵', '刘', '陈', '杨', '黄', '周', '吴'];
const firstNames = ['伟', '芳', '娜', '秀英', '敏', '静', '丽', '强', '磊', '洋', '艳', '勇', '军', '杰', '娟', '涛', '明', '超', '秀兰', '霞'];

export const generateRandomName = () => {
  const lastName = lastNames[Math.floor(Math.random() * lastNames.length)];
  const firstName = firstNames[Math.floor(Math.random() * firstNames.length)];
  return lastName + firstName;
};

// ====== 云南高原特色产地数据 ======
const farmNames = [
  '普洱茶山生态庄园', '小粒咖啡示范农场', '高原野生菌基地', '昭通苹果种植园',
  '蒙自石榴合作社', '丽江雪桃基地', '宣威火腿原料农场', '文山三七种植园',
  '香格里拉松茸基地', '大理乳扇牧场', '临沧坚果庄园', '滇红茶叶生态园',
  '红河梯田红米基地', '哀牢山紫米农场', '斗南鲜切花基地', '楚雄牛肝菌基地',
  '怒江草果种植园', '德宏咖啡庄园', '西双版纳香蕉园', '曲靖韭菜花基地',
  '通海蔬菜合作社', '元谋蔬菜农场', '保山绿豆种植园', '丽江高原牧场',
  '广南石斛基地', '昆明宝珠梨庄园', '会泽石榴种植园', '迪庆青稞农场'
];
const provinces = ['云南省'];
const cities = [
  '昆明市', '大理市', '丽江市', '普洱市', '保山市', '楚雄市', '红河州',
  '文山州', '西双版纳州', '德宏州', '怒江州', '迪庆州', '临沧市',
  '昭通市', '曲靖市', '玉溪市', '蒙自市', '宣威市', '安宁市'
];
const districts = [
  '古城区', '大理市', '思茅区', '隆阳区', '楚雄市', '蒙自市', '文山市',
  '景洪市', '芒市', '泸水市', '香格里拉市', '临翔区', '昭阳区',
  '麒麟区', '红塔区', '官渡区', '西山区', '五华区', '盘龙区',
  '呈贡区', '安宁市', '宜良县', '石林县', '禄劝县', '寻甸县'
];
const villageNames = [
  '普洱茶村', '古树茶园村', '梯田红米村', '咖啡庄园村', '松茸山村',
  '牛肝菌基地', '三七种植村', '乳扇牧业村', '鲜花基地村',
  '核桃古寨', '苹果种植村', '石榴庄园', '雪桃基地', '火腿原料村',
  '玛咖种植村', '蓝莓庄园', '药材种植村', '蔬菜大棚基地'
];

export const generateRandomFarmAddress = () => {
  const province = provinces[Math.floor(Math.random() * provinces.length)];
  const city = cities[Math.floor(Math.random() * cities.length)];
  const district = districts[Math.floor(Math.random() * districts.length)];
  const village = villageNames[Math.floor(Math.random() * villageNames.length)];
  const group = Math.floor(Math.random() * 10 + 1);
  return `${province}${city}${district}${village}${group}组`;
};

export const generateRandomFarmName = () => {
  return farmNames[Math.floor(Math.random() * farmNames.length)];
};

export const generateRandomProvince = () => {
  return provinces[Math.floor(Math.random() * provinces.length)];
};

export const generateRandomCity = () => {
  return cities[Math.floor(Math.random() * cities.length)];
};

export const generateRandomDistrict = () => {
  return districts[Math.floor(Math.random() * districts.length)];
};

// ====== 云南高原特色农产品数据库 ======
const productCategories = [
  '茶叶', '咖啡', '水果', '蔬菜', '粮食', '肉类', '野生菌',
  '中药材', '花卉', '坚果', '豆类', '乳制品', '调味品', '加工食品'
];

const productNames: Record<string, string[]> = {
  '茶叶': [
    '普洱茶（生茶）', '普洱茶（熟茶）', '滇红茶', '古树普洱茶', '云南白茶',
    '月光白', '云南绿茶', '景迈山普洱茶', '冰岛古树茶', '老班章普洱茶',
    '易武正山茶', '南糯山普洱', '勐海熟茶', '凤庆滇红', '昌宁红茶',
    '大理沱茶', '德昂酸茶', '基诺山凉茶'
  ],
  '咖啡': [
    '云南小粒咖啡', '保山铁毕卡咖啡', '普洱阿拉比卡', '德宏咖啡豆',
    '临沧咖啡', '云南精品咖啡', '大理咖啡', '西双版纳咖啡',
    '云南有机咖啡', '怒江咖啡', '云南冷萃咖啡', '滇西咖啡豆'
  ],
  '水果': [
    '昭通苹果', '蒙自石榴', '丽江雪桃', '德宏芒果', '华宁柑桔',
    '富民杨梅', '大理樱桃', '昆明宝珠梨', '会泽石榴', '元谋番茄',
    '河口香蕉', '西双版纳香蕉', '红河蜜柚', '玉溪冰糖橙',
    '丽江华坪芒果', '石林人参果', '弥勒葡萄', '建水酸石榴',
    '曲靖蓝莓', '大理草莓', '宾川柑橘', '瑞丽菠萝'
  ],
  '蔬菜': [
    '通海蔬菜', '元谋蔬菜', '大理海菜', '昆明叶菜', '曲靖韭菜花',
    '泸西小葱', '弥渡大蒜', '江川花菜', '呈贡白菜',
    '云南高山娃娃菜', '丽江雪莲果', '普洱黄瓜', '保山四季豆',
    '玉溪番茄', '楚雄莲藕', '云南高原芦笋', '文山生姜'
  ],
  '粮食': [
    '八宝贡米', '红河梯田红米', '哀牢山紫米', '墨江紫米', '勐海香米',
    '云粳大米', '文山香米', '丽江高原粳米', '大理蚕豆', '曲靖玉米',
    '楚雄荞麦', '云南苦荞', '昭通土豆', '迪庆青稞', '怒江藜麦'
  ],
  '肉类': [
    '宣威火腿', '诺邓火腿', '撒坝火腿', '云南牛干巴', '傣味烤乳猪',
    '滇西黄牛肉', '云南肉牛', '高原牦牛肉', '大理乳扇乳饼',
    '云南土鸡', '武定壮鸡', '盐津乌骨鸡', '丽江腊排骨',
    '云南香肠', '昭通酱牛肉', '西双版纳冬瓜猪'
  ],
  '野生菌': [
    '香格里拉松茸', '楚雄牛肝菌', '大理鸡枞菌', '丽江羊肚菌',
    '南华野生菌', '云南干巴菌', '青头菌', '见手青',
    '奶浆菌', '松露（块菌）', '竹荪', '虎掌菌',
    '白葱菌', '红葱菌', '谷熟菌', '铜绿菌',
    '鸡油菌', '珊瑚菌', '刷把菌', '一窝菌'
  ],
  '中药材': [
    '文山三七', '昭通天麻', '广南铁皮石斛', '丽江玛咖', '大理滇重楼',
    '楚雄茯苓', '云南黄芪', '怒江草果', '迪贝母', '云南黄连',
    '滇黄精', '滇龙胆', '灯盏花', '云南红豆杉', '砂仁',
    '云木香', '云南金银花', '滇丹参', '白及', '黄草乌'
  ],
  '花卉': [
    '昆明鲜切玫瑰', '斗南百合', '云南康乃馨', '大理山茶花',
    '丽江高山杜鹃', '云南蝴蝶兰', '红河观赏花卉', '云南多肉植物',
    '云南食用玫瑰', '昆明绣球花', '玉溪非洲菊', '云南兰花'
  ],
  '坚果': [
    '临沧澳洲坚果', '云南核桃', '大理核桃', '漾濞核桃',
    '楚雄薄壳核桃', '云南板栗', '昭通核桃', '保山核桃',
    '云南松子', '云南腰果', '滇西杏仁'
  ],
  '豆类': [
    '保山透心绿豆', '云南大白芸豆', '丽江芸豆', '曲靖黄豆',
    '大理蚕豆', '楚雄红豆', '文山黑豆', '云南豌豆',
    '滇西鹰嘴豆', '昭通绿豆', '云南扁豆'
  ],
  '乳制品': [
    '大理乳扇', '云南乳饼', '大理酸奶', '丽江酥油',
    '云南奶粉', '大理奶酪', '高原牦牛奶', '云南乳饮料',
    '大理鲜牛奶', '香格里拉酥油茶'
  ],
  '调味品': [
    '昭通酱', '云南腐乳', '石屏豆腐', '云南小米辣', '大理酸木瓜',
    '云南花椒', '丘北辣椒', '漾濞花椒', '云南小黄姜',
    '德宏涮涮辣', '云南胡椒', '玉溪酱油', '昆明麸醋'
  ],
  '加工食品': [
    '云南鲜花饼', '过桥米线', '云南饵丝', '大理耙肉饵丝', '丽江粑粑',
    '云南米线', '宣威火腿饼', '云南荞饼', '版纳糯玉米',
    '云南果脯', '大理雕梅', '云南红糖', '滇式月饼'
  ],
};

export const generateRandomCategory = () => {
  return productCategories[Math.floor(Math.random() * productCategories.length)];
};

export const generateRandomProductName = (category: string) => {
  const names = productNames[category] || productNames['蔬菜'];
  return names[Math.floor(Math.random() * names.length)];
};

// 随机生成农场面积（10-500亩）
export const generateRandomArea = () => {
  return Number((Math.random() * 490 + 10).toFixed(2));
};

// 随机生成产品数量
export const generateRandomQuantity = (unit: string) => {
  if (unit === '吨') return Number((Math.random() * 90 + 10).toFixed(1));
  if (unit === '公斤') return Math.floor(Math.random() * 9000 + 1000);
  if (unit === '斤') return Math.floor(Math.random() * 18000 + 2000);
  return Math.floor(Math.random() * 900 + 100);
};

// 随机生成单位
export const generateRandomUnit = () => {
  const units = ['斤', '公斤', '吨', '箱'];
  return units[Math.floor(Math.random() * units.length)];
};

// 随机生成批次号
export const generateBatchNo = () => {
  const now = new Date();
  const dateStr = `${now.getFullYear()}${String(now.getMonth()+1).padStart(2,'0')}${String(now.getDate()).padStart(2,'0')}`;
  const seq = String(Math.floor(Math.random() * 9999 + 1)).padStart(4, '0');
  return `ZX-${dateStr}-${seq}`;
};

// 随机生成价格（农产品）
export const generateRandomPrice = () => {
  return Number((Math.random() * 90000 + 1000).toFixed(2));
};

// 随机生成手机号
export const generateRandomPhone = () => {
  const prefixes = ['138', '139', '150', '151', '152', '186', '187', '188'];
  const prefix = prefixes[Math.floor(Math.random() * prefixes.length)];
  const suffix = String(Math.floor(Math.random() * 100000000)).padStart(8, '0');
  return `${prefix}${suffix}`;
};

// 随机生成检测数据
export const generateRandomPesticideResidue = () => {
  const values = ['未检出', '0.01mg/kg', '0.02mg/kg', '0.05mg/kg', '0.1mg/kg'];
  return values[Math.floor(Math.random() * values.length)];
};

export const generateRandomHeavyMetal = () => {
  const values = ['未检出', '0.001mg/kg', '0.005mg/kg', '0.01mg/kg', '0.02mg/kg'];
  return values[Math.floor(Math.random() * values.length)];
};

export const generateRandomMicroorganism = () => {
  const values = ['合格', '菌落总数<10CFU/g', '菌落总数<50CFU/g', '大肠杆菌未检出'];
  return values[Math.floor(Math.random() * values.length)];
};

export const generateCertNumber = () => {
  const year = new Date().getFullYear();
  const seq = String(Math.floor(Math.random() * 99999 + 1)).padStart(5, '0');
  return `ZX-JC-${year}-${seq}`;
}; 