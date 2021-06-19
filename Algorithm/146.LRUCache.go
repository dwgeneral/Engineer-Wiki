/*
 * 使用map存储key与链表节点的对应关系, 定位节点的效率 O(1)
 * 使用双端链表记录最近使用顺序, 头部添加与尾部删除的效率也是 O(1), 将某个节点添加到头部的效率也是 O(1)
 */
type LRUCache struct {
	size, cap  int
	hashmap    map[int]*Node
	head, tail *Node
}

type Node struct {
	key, val   int
	prev, next *Node
}

func initNode(key, val int) *Node {
	return &Node{
		key: key,
		val: val,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cap:     capacity,
		hashmap: make(map[int]*Node),
		head:    initNode(0, 0),
		tail:    initNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hashmap[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.hashmap[key]
	if ok {
		node.val = value
		this.moveToHead(node)
		return
	}
	newHead := initNode(key, value)
	this.addHead(newHead)
	this.hashmap[key] = newHead
	this.size++
	if this.size > this.cap {
		removed := this.removeTail()
		delete(this.hashmap, removed.key)
		this.size--
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addHead(node)
}

func (this *LRUCache) addHead(node *Node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}