# MongoDB

## MongoDB 与关系型数据库的区别

1. 数据模型

- MongoDB 使用文档存储，每个文档采用 BSON 格式，能够嵌套结构化数组和对象，无需预先定义固定模式
- 关系型数据库采用表格结构，数据以行和列形式存储，通常需要预定义格式

2. 查询语言

- MongoDB 使用基于 JSON 的查询语言，采用聚合管道实现复杂数据转化，查询灵活性高
- 关系型数据库使用结构化查询语言 SQL， 但复杂查询可能影响性能

3. 扩展性

- MongoDB 天生支持水平扩展，可以跨多态服务器分布存储数据，适合大规模数据和高并发场景。
- 关系型数据库通常依赖垂直扩展，不如 MongoDB 灵活

4. 一致性与事务

- MongoDB 在单文档操作上提供原子性,多文档事务从 V4.0 开始支持，但与关系型数据库相比仍有限。
- 关系型数据库全面支持 ACID 事务，适合对数据一致性要求极高的场景。

## MongoDB 的事务

- 性能： 强一致性事务会影响数据库的性能，而 mongoDB 的设计目标是高性能。
- 数据模型： 文档型数据库的数据模型相对灵活，但同时也增加了实现事务的复杂性。
- 使用场景： MongoDB 更适合用于存储非结构化数据和高并发读写场景，这些场景对事务的要求相对较低。
- 一致性： mongoDB 为了性能采用的是最终一致性，需要一段时间才能被所有节点看到，关系型数据库是强一致性。

在设计数据模型时，仍建议尽量利用单文档操作的原子性，只有在必要时才引入跨多文档的事务。

## MongoDB 的索引类型有哪些？各自的优缺点是什么？

1. 单字段索引(single Field)

- 优点： 简单高效，适用精确匹配/排序
- 缺点： 只优化单字段查询

2. 复合索引(Compound Index)

- 优点： 支持多字段联合查询/排序
- 缺点： 仅最左前缀生效，占用更多内存

3. 多键索引(MultiKey)

- 优点： 自动为数组字段每个元素创建索引
- 缺点： 大规模数组影响写入性能

4. 文本索引(Text)

- 优点： 支持全文搜索(含语言分词)
- 缺点： 仅一个文本索引/权重管理复杂

5. 地理空间索引(Geospatial)

- 优点： 高效处理位置查询
- 缺点： 仅适用地理坐标数据

6. 哈希索引(Hashed)

- 优点： 均匀分片支持， 快速等值查询
- 缺点： 无法范围查询

7. 通配符索引(Wildcard)

- 优点： 动态字段查询优化
- 缺点： 索引尺寸较大

## MongoDB 的查询分析计划

`db.collection.find().explain("executionStats") // 执行统计`

executionStats 实际执行性能指标

```json
"executionStats": {
    "nReturned" : 5,        // 返回文档数
    "executionTimeMillis" : 2, // 执行时间(ms)
    "totalKeysExamined" : 5,  // 索引扫描次数
    "totalDocsExamined" : 5   // 文档扫描次数
  }
```

理想情况下 `nReturned ≈ totalKeysExamined ≈ totalDocsExamined`

COLLSCAN（全表扫描） vs IXSCAN（索引扫描）  
 ▶︎ 应尽量避免出现 COLLSCAN

## 如何保证 MongoDB 的数据一致性？

1. 写入层保障

```js
// 强一致性写入配置
db.orders.insertOne({...}, {
  writeConcern: {
    w: "majority", // 确保数据写入大多数节点
    j: true,       // 写入journal日志保证持久性
    wtimeout: 5000 // 超时设置（毫秒）
  }
})
```

2. 读取层控制

```js
// 读取已提交的数据
db.products.find().readConcern("majority");
```
