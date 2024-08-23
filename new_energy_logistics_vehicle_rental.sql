-- 删除现有的 `companies` 表（如果存在）
DROP TABLE IF EXISTS `companies`;

-- 创建 `companies` 表，存储公司及子公司信息
CREATE TABLE companies
(
    id                INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自动递增',
    parent_company_id INT      DEFAULT NULL COMMENT '父公司ID (NULL表示顶级公司)',
    company_name      VARCHAR(100) UNIQUE NOT NULL COMMENT '公司名称，唯一',
    address           VARCHAR(255) COMMENT '公司地址',
    contact_number    VARCHAR(20) COMMENT '联系电话',
    email             VARCHAR(100) COMMENT '公司邮箱',
    description       TEXT COMMENT '公司描述（可选）',
    `create_by`       int      DEFAULT NULL COMMENT '创建者',
    `update_by`       int      DEFAULT NULL COMMENT '更新者',
    created_at        DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at        DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (parent_company_id) REFERENCES companies (id) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT = '存储公司及子公司信息的表';

-- 删除现有的 `fleets` 表（如果存在）
DROP TABLE IF EXISTS `fleets`;

-- 创建 `fleets` 表，存储车队信息
CREATE TABLE fleets
(
    id             INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自动递增',
    fleet_name     VARCHAR(100) NOT NULL COMMENT '车队名称',
    company_id     INT          NOT NULL COMMENT '关联的公司ID (参考companies表的主键)',
    description    TEXT COMMENT '车队描述（可选）',
    contact_number VARCHAR(20) COMMENT '联系电话',
    email          VARCHAR(100) COMMENT '负责人邮箱',
    manager_name   VARCHAR(100) COMMENT '负责人',
    `create_by`    int      DEFAULT NULL COMMENT '创建者',
    `update_by`    int      DEFAULT NULL COMMENT '更新者',
    created_at     DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at     DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '车队表';

-- 删除现有的 `vehicle_types` 表（如果存在）
DROP TABLE IF EXISTS `vehicle_types`;

-- 创建 `vehicle_types` 表，存储车辆种类信息
CREATE TABLE vehicle_types
(
    id          INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自动递增',
    type_name   VARCHAR(50) UNIQUE NOT NULL COMMENT '车辆种类名称 (如“小型汽车”)',
    `create_by` int      DEFAULT NULL COMMENT '创建者',
    `update_by` int      DEFAULT NULL COMMENT '更新者',
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '车辆种类表';

-- 删除现有的 `vehicle_brands` 表（如果存在）
DROP TABLE IF EXISTS `vehicle_brands`;

-- 创建 `vehicle_brands` 表，存储车辆品牌信息
CREATE TABLE vehicle_brands
(
    id          INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自动递增',
    brand_name  VARCHAR(100) UNIQUE NOT NULL COMMENT '车辆品牌名称 (如“晶马”)',
    `create_by` int      DEFAULT NULL COMMENT '创建者',
    `update_by` int      DEFAULT NULL COMMENT '更新者',
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '车辆品牌表';

-- 删除现有的 `vehicle_models` 表（如果存在）
DROP TABLE IF EXISTS `vehicle_models`;

-- 创建 `vehicle_models` 表，存储车辆型号信息
CREATE TABLE vehicle_models
(
    id          INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自动递增',
    model_name  VARCHAR(100) NOT NULL COMMENT '车型名称',
    description TEXT COMMENT '车型描述（可选）',
    brand_id    INT          NOT NULL COMMENT '关联的品牌ID (参考vehicle_brands表的主键)',
    type_id     INT          NOT NULL COMMENT '关联的车辆种类ID (参考vehicle_types表的主键)',
    `create_by` int      DEFAULT NULL COMMENT '创建者',
    `update_by` int      DEFAULT NULL COMMENT '更新者',
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (brand_id) REFERENCES vehicle_brands (id) ON DELETE CASCADE,
    FOREIGN KEY (type_id) REFERENCES vehicle_types (id) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '车辆型号表';

-- 删除现有的 `pickup_locations` 表（如果存在）
DROP TABLE IF EXISTS `pickup_locations`;

-- 创建 `pickup_locations` 表，存储提车网点信息
CREATE TABLE pickup_locations
(
    id             INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自动递增',
    company_id     INT            NOT NULL COMMENT '关联的公司ID (参考companies表的主键)',
    location_name  VARCHAR(100)   NOT NULL COMMENT '提车点名称',
    description    TEXT COMMENT '车型描述（可选）',
    address        VARCHAR(255) COMMENT '提车点地址',
    contact_number VARCHAR(20) COMMENT '提车点联系电话',
    manager_name   VARCHAR(100) COMMENT '提车点负责人',
    latitude       DECIMAL(10, 7) NOT NULL COMMENT '提车点纬度',
    longitude      DECIMAL(10, 7) NOT NULL COMMENT '提车点经度',
    `create_by`    int      DEFAULT NULL COMMENT '创建者',
    `update_by`    int      DEFAULT NULL COMMENT '更新者',
    created_at     DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at     DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT = '存储提车网点信息的表';

-- 删除现有的 `vehicles` 表（如果存在）
DROP TABLE IF EXISTS `vehicles`;

-- 创建 `vehicles` 表，存储车辆信息
CREATE TABLE vehicles
(
    id                          INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自动递增',
    company_id                  INT NOT NULL COMMENT '关联的公司ID (参考companies表的主键)',
    pickup_location_id          INT NOT NULL COMMENT '关联的提车网点ID (参考pickup_locations表的主键)',
    fleet_id                    INT NOT NULL COMMENT '关联的车队ID (参考fleets表的主键)',
    license_plate               VARCHAR(20) COMMENT '车牌号码',
    color                       VARCHAR(50) COMMENT '车辆颜色',
    model_id                    INT NOT NULL COMMENT '车型ID (关联到vehicle_models表)',
    vin                         VARCHAR(50) COMMENT '车架号 (Vehicle Identification Number)',
    engine_number               VARCHAR(50) COMMENT '发动机号',
    maintenance_interval_km     INT COMMENT '保养周期 (公里)',
    total_mileage_km            INT COMMENT '总里程 (公里)',
    next_maintenance_mileage_km INT COMMENT '下次保养里程 (公里)',
    emission_standard           VARCHAR(50) COMMENT '环保标准',
    vehicle_owner               VARCHAR(100) COMMENT '车辆所属',
    affiliated_company          VARCHAR(100) COMMENT '挂靠公司',
    transaction_mode            VARCHAR(50) COMMENT '交易方式',
    idle_reason                 TEXT COMMENT '闲置原因',
    remarks                     TEXT COMMENT '备注',
    vehicle_status              VARCHAR(50) COMMENT '车辆状态',
    is_disposed                 BOOLEAN COMMENT '是否处置',
    traffic_management_account  VARCHAR(100) COMMENT '所属的交管账号',
    `create_by`                 int DEFAULT NULL COMMENT '创建者',
    `update_by`                 int DEFAULT NULL COMMENT '更新者',
    created_at                  DATETIME COMMENT '创建日期',
    updated_at                  DATETIME COMMENT '更新日期',
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (fleet_id) REFERENCES fleets (id) ON DELETE CASCADE,
    FOREIGN KEY (model_id) REFERENCES vehicle_models (id) ON DELETE RESTRICT,
    FOREIGN KEY (pickup_location_id) REFERENCES pickup_locations (id) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '车辆表';

-- 删除现有的 `vehicle_photos` 表（如果存在）
DROP TABLE IF EXISTS `vehicle_photos`;

-- 创建 `vehicle_photos` 表，存储车辆图片信息
CREATE TABLE vehicle_photos
(
    id          INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自动递增',
    vehicle_id  INT  NOT NULL COMMENT '关联的车辆ID (参考vehicles表的主键)',
    photo       BLOB NOT NULL COMMENT '车辆图片数据',
    description VARCHAR(255) COMMENT '图片描述（可选）',
    `create_by` int      DEFAULT NULL COMMENT '创建者',
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '图片创建时间',
    FOREIGN KEY (vehicle_id) REFERENCES vehicles (id) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '车辆图片表';

DROP TABLE IF EXISTS `vehicle_photos`;
DROP TABLE IF EXISTS `vehicles`;
DROP TABLE IF EXISTS `pickup_locations`;
-- DROP TABLE IF EXISTS `vehicle_models`;
-- DROP TABLE IF EXISTS `vehicle_brands`;
-- DROP TABLE IF EXISTS `vehicle_types`;
-- DROP TABLE IF EXISTS `fleets`;
-- DROP TABLE IF EXISTS `companies`;

