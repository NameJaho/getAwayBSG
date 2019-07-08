// 房价均价查询语句

db.lianjia.aggregate([
  { '$match': { "address.0": { $exists: true } } },
  {
    $group: {
      _id: { "$arrayElemAt": ["$address", 0] },
      count: { $sum: 1 },
      avg_UnitPrice: { $avg: "$UnitPrice" },
      std: { $stdDevPop: "$UnitPrice" },
    }
  },
  {
    $project:
        {
          count: 1,        //总数
          avg_UnitPrice: 1, //每平米均价
          std: 1,   //标准差
          ratio: { $divide: ["$std", "$avg_UnitPrice"] } //标准差与均价的比值
        }
  },
  {
    '$sort': { count: -1 }
  }
]);