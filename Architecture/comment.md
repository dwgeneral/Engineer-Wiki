## 评论系统架构设计

> 评论系统对于互联网产品来说是一个通用需求, 正好今天看到了一篇关于评论系统设计的分享, 梳理一下和大家分享

架构设计最重要的就是理解整个产品体系在系统中的定位。搞清楚系统背后的业务场景，才能做出最佳的设计和抽象。不要做需求的翻译机，先理解业务背后的本质，事情的初衷。

### 功能概述
评论系统主要包括:
- 发布评论: 支持回复楼层、楼中楼。
- 读取评论: 按照时间、热度排序。
- 删除评论: 用户删除、作者删除。
- 管理评论: 作者置顶、后台运营管理(搜索、删除、审核等)。

评论系统，我们往小做就是视频评论系统，往大做就是评论平台，可以接入各种业务形态。

### 架构设计
![某厂评论系统架构](./assets/comment.png ':size=700')

各模块概述:
- BFF: comment
    复杂评论业务的服务编排，比如访问账号服务进行等级判定，同时需要在 BFF 面向移动端/WEB场景来设计 API，这一层抽象把评论的本身的内容列表处理(加载、分页、排序等)进行了隔离，关注在业务平台化逻辑上。
- Service: comment-sevice
    服务层，去平台业务的逻辑，专注在评论功能的 API 实现上，比如发布、读取、删除等，关注在稳定性、可用性上，这样让上游可以灵活组织逻辑把基础能力和业务能力剥离。
- Job: comment-job
    异步任务, 消息队列的消费端,异步地完成数据写入, 数据同步, 缓存重建等工作.
- Admin: comment-admin
    管理平台，按照安全等级划分服务，尤其划分运营平台，他们会共享服务层的存储层(MySQL、Redis)。运营体系的数据大量都是检索，我们使用 canal 进行同步到 ES 中，整个数据的展示都是通过 ES，再通过业务主键更新业务数据层，这样运营端的查询压力就下方给了独立的 fulltext search 系统。
- Dependency: account-service, filter-service
    整个评论服务还会依赖一些外部 gRPC 服务，统一的平台业务逻辑在 comment BFF 层收敛，这里 account-service 主要是账号服务，filter-service 是敏感词过滤服务。

数据流转:

- 读逻辑:
   
   使用 cache-aside 模式,先读缓存,再读DB,对于重建缓存逻辑, 一般会使用 read ahead 预读, 用户访问了第一页, 很有可能访问第二页, 所以缓存会超前加载, 减少 cache miss. 但某些场景下缓存会发生抖动, 造成cache miss, 这时如果请求量大的话, 会出现惊群现象, 大量触发缓存重建, 又因为使用了预读机制, 容易导致服务OOM.
   
   所以这里使用异步回源策略, 即 comment-service 向 kafka 发送回源指令消息, comment-job消费此类消息并串行进行缓存重建, 消除流量峰值.

- 写逻辑:
    
    为了避免热点事件瞬时大量写入对DB造成压力, 使用消息队列对写操作进行分压. 我们认为刚发布的评论有极短的延迟(通常小于几 ms)对用户可见是可接受的，把对存储的直接冲击下放到消息队列，按照消息反压的思路，即如果存储 latency 升高，消费能力就下降，自然消息容易堆积，系统始终以最大化方式消费。
    
    为了保证数据的一致性, 会先更新DB, 再更新缓存.

- 运营体系:

    mysql binlog 中的数据被 canal 中间件流式消费，获取到业务的原始 CRUD 操作，需要回放录入到 es 中，但是 es 中的数据最终是面向运营体系提供服务能力，需要检索的数据维度比较多，在入 es 前需要做一个异构的 joiner，把单表变宽预处理好 join 逻辑，然后倒入到 es 中。

    一般来说，运营后台的检索条件都是组合的，使用 es 的好处是避免依赖 mysql 来做多条件组合检索，同时 mysql 毕竟是 oltp 面向线上联机事务处理的。通过冗余数据的方式，使用其他引擎来实现。
    es 一般会存储检索、展示、primary key 等数据，当我们操作编辑的时候，找到记录的 primary key，最后交由 comment-admin 进行运营测的 CRUD 操作。

### 存储设计

- 索引内容分离:

    ![comment表结构](./assets/comment-table.png ':size=400')

    comment_index: 评论的索引表, 不存储具体的评论内容, 存储索引数据与动态数据(经常变更的数据)

    comment_content: 评论内容表, 包含具体的评论内容, 并且该表没有id, 是为了减少一次二级索引查询, 直接基于主键(comment_id)检索, 同时 comment_id 在写入时要尽可能的顺序自增

    comment_subject: 评论与主题关系表, 记录该评论所在的主题的相关信息.

    索引,内容分离, 方便 mysql datapage 缓存更多的 row, 如果和 content 耦合, 会导致更大的IO, 长远来看 content 信息可以直接使用 KV storage 存储.

- 缓存设计:

    ![comment缓存](./assets/comment-cache.png ':size=400')

    comment_subject_cache: 对应主题的缓存，value 使用 protobuf 序列化的方式存入。

    comment_index_cache: 使用 redis sortedset 进行索引的缓存，索引即数据的组织顺序，而非数据内容。参考过百度的贴吧，他们使用自己研发的拉链存储来组织索引，我认为 mysql 作为主力存储，利用 redis 来做加速完全足够，因为 cache miss 的构建，我们前面讲过使用 kafka 的消费者中处理，预加载少量数据，通过增量加载的方式逐渐预热填充缓存，而 redis sortedset skiplist 的实现，可以做到 O(logN) + O(M) 的时间复杂度，效率很高。sorted set 是要增量追加的，因此必须判定 key 存在，才能 zdd。

    comment_content_cache: 对应评论内容数据，使用 protobuf 序列化的方式存入.

### 可用性设计

- singleflight

    对于热门的主题，如果存在缓存穿透的情况，会导致大量的同进程、跨进程的数据回源到存储层，可能会引起存储过载的情况，如何只交给同进程内，一个人去做加载存储?

    使用归并回源的思路: https://pkg.go.dev/golang.org/x/sync/singleflight

    同进程只交给一个人去获取 mysql 数据，然后批量返回。同时这个 lease owner 投递一个 kafka 消息，做 index cache 的 recovery 操作。这样可以大大减少 mysql 的压力，以及大量透穿导致的密集写 kafka 的问题。

    更进一步的，后续连续的请求，仍然可能会短时 cache miss，我们可以在进程内设置一个 short-lived flag，标记最近有一个人投递了 cache rebuild 的消息，直接 drop。

- 热点识别

    ![comment热点识别](./assets/comment-hotscan.png ':size=400')

    流量热点是因为突然热门的主题，被高频次的访问，因为底层的 cache 设计，一般是按照主题 key 进行一致性 hash 来进行分片，但是热点 key 一定命中某一个节点，这时候 remote cache 可能会变为瓶颈，因此做 cache 的升级 local cache 是有必要的，我们一般使用单进程自适应发现热点的思路，附加一个短时的 ttl local cache，可以在进程内吞掉大量的读请求。

    在内存中使用 hashmap 统计每个 key 的访问频次，这里可以使用滑动窗口统计，即每个窗口中，维护一个 hashmap，之后统计所有未过去的 bucket，汇总所有 key 的数据。

    之后使用小堆计算 TopK 的数据，自动进行热点识别。

