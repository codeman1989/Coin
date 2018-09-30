package main

// CCriticalSection 互斥锁
// 先写上，可能也没用
var cs_main sync.Mutex

// 仅仅保存没有被打包到主链中交易
var mapTransactions map[uint256]CTransaction

// 每次对mapTransactions中交易进行更新，都对该字段进行++ --操作
var uint64 nTransactionsUpdated = 0

//如果对应的区块已经放入到主链中，则对应的区块交易应该要从本节点保存的交易内存池中删除
var mapNextTx map[COutPoint]CInPoint

// 块索引信息：其中key对应的block的hash值
var mapBlockIndex map[uint256]*CBlockIndex
