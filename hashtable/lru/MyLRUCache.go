package lru

type Node struct {
	Key, Value int
	Pre, Next  *Node
}

func NewNode(key, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
		Pre:   nil,
		Next:  nil,
	}
}

type LinkedList struct {
	Capacity   int
	Head, Tail *Node
}

func NewLinkedList() *LinkedList {
	head := NewNode(0, 0)
	tail := NewNode(0, 0)
	head.Next = tail
	tail.Pre = head
	return &LinkedList{
		Capacity: 0,
		Head:     head,
		Tail:     tail,
	}
}

func (linkedList *LinkedList) AddTail(node *Node) {
	linkedList.Tail.Pre.Next = node
	node.Pre = linkedList.Tail.Pre
	node.Next = linkedList.Tail
	linkedList.Tail.Pre = node
	linkedList.Capacity++
}

func (linkedList *LinkedList) DeleteNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	linkedList.Capacity--
}

func (linkedList *LinkedList) DeleteFirst() *Node {
	if linkedList.Head.Next == linkedList.Tail {
		return nil
	}
	node := linkedList.Head.Next
	linkedList.DeleteNode(node)
	return node
}

type LRUCache struct {
	Capacity int
	List     *LinkedList
	Cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		List:     NewLinkedList(),
		Cache:    map[int]*Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Cache[key]; !ok {
		return -1
	}
	node := this.MakeRecent(key)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Cache[key]; ok {
		node := this.DeleteNode(key)
		this.AddRecent(node.Key, node.Value)
	} else {
		//不存在
		if this.List.Capacity == this.Capacity {
			this.DeleteFirst()
		}
		this.AddRecent(key, value)
	}
}

func (this *LRUCache) MakeRecent(key int) *Node {
	node := this.Cache[key]
	this.List.DeleteNode(node)
	this.List.AddTail(node)
	return node
}

func (this *LRUCache) AddRecent(key, value int) {
	node := NewNode(key, value)
	this.List.AddTail(node)
	this.Cache[key] = node
}

func (this *LRUCache) DeleteNode(key int) *Node {
	node := this.Cache[key]
	this.List.DeleteNode(node)
	delete(this.Cache, key)
	return node
}

func (this *LRUCache) DeleteFirst() {
	node := this.List.DeleteFirst()
	if node != nil {
		delete(this.Cache, node.Key)
	}
}
