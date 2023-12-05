## 核心思想
- 以空间换时间，一次缓存，多次复用 
- 每次向操作系统申请内存会多申请一些，以备后用(申请单位为 heapArena 64MB)

## mheap/mcentral/mcache 数据结构
- Golang运行时的堆 mheap 正是基于该思想产生的数据结构
- 对操作系统而言，这是一个用户进程申请了一坨内存
- 对于Go进程内部，堆是所有对象的内存起源，全局唯一对象
- 为了提高内存分配效率，采用多级缓存，实现无/细锁化
    - mheap 是 Go runtime 中最大的临界共享资源，这意味着每次存取都要加锁，在性能层面是一件很可怕的事情
    - 所以，Go在 mheap 之上，依次细化粒度，建立了 mcentral, mcache 的模型
    - mcentral: 每种对象大小规格(全局共划分为68种)对应的缓存，锁的粒度也仅限于同一种规格以内
    - mcache: 每个P(GMP)持有的一份内存缓存，访问时无锁

## 分页管理内存

### mspan
- Golang借鉴操作系统分页管理的思想，每个最小的存储单元也称之为页(page)，但大小为8KB
- 最小的内存管理单元 mspan，大小为page的整数倍，从 8B 到 80KB 共67种规格，分配对象时，会根据大小映射到对应规格的 mspan, 从中获取空间

### pageAlloc
- 而 pageAlloc 是在Golang堆内存中，用来加快寻找空闲页的一种索引结构,底层实现是基数树(Radix Tree 前缀树)
- 每棵基数树聚合了16GB内存空间中各页使用情况的索引信息，用于帮助 mheap 快速找到指定长度的连续空闲页的所在位置
- mheap 堆内存的上限是 256TB，所以其持有 2^14 棵基数树，因此索引能够覆盖到所有堆内存

### heapArena
- heapArena 是 mheap 向操作系统申请内存的单位（64MB）
- 记录了 page 与 mspan 的映射，因为GC时，通过地址偏移找到页很容易，但找到其所属的mspan不容易，因此需要通过这个映射信息辅助查找

## 对象内存分配流程
- 依据Object大小，划分为三种对象类型
    - Tiny Obj(0-16B)
        1. 从P专属的 mcache 的 tiny 分配器取内存(无锁)
        2. 如果没有，则根据所属的 spanClass, 从 P 专属 mcache 的 mspan 中取(无锁)
        3. 如果没有，则根据所属的 spanClass, 从对应的 mcentral 中取 mspan 填充到 mcache,然后从 mspan 中取(spanClass粒度锁)
        4. 如果没有，则根据所属的 spanClass, 从 mheap 的页分配器 pageAlloc 找出空闲页组装成 mspan 填充到 mcache ,然后从 mspan 中取(全局锁)
        5. 如果没有，mheap 向操作系统申请内存，更新 pageAlloc 索引信息，然后重复 step 4
    - Small Obj(16B, 32K)
        - 跳过1，执行2-5
    - Large Obj(>32K)
        - 跳过1-3，执行4-5

- 无论是 new, make, &T{}, 分配内存最终由 mallocgc 方法负责
    - 因为内存分配本身也是触发GC的一个入口，当发现mcache/mcentral中的内存不够用了，
会将 shouldhelpgc 置为true，发起GC
- 综上，总结来看，核心流程类似读多级缓存的过程，由上而下，每一步成功则返回，不成功继续下层处理