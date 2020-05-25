package LRU

type (
	LRUCache struct {
		Capacity int
		HashMap  map[int]*Node
		Head     *Node
		Last     *Node
	}
	
	Node struct {
		Val  int
		Key  int
		Pre  *Node
		Next *Node
	}
)


func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Capacity: capacity,
		HashMap:  make(map[int]*Node, capacity),
		Head:     &Node{},
		Last:     &Node{},
	}
	cache.Head.Next = cache.Last
	cache.Last.Pre = cache.Head
	return cache
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.HashMap[key]
	if !ok {
		return -1
	}
	lru.remove(node)
	lru.setHead(node)
	return node.Val
}

func (lru *LRUCache) Put(key int, value int) {
	node, ok := lru.HashMap[key]
	if ok {
		node.Val = value
		lru.remove(node)
	} else {
		if len(lru.HashMap) == lru.Capacity {
			delete(lru.HashMap, lru.Last.Pre.Key)
			lru.remove(lru.Last.Pre)
		}
		node = &Node{
			Val:  value,
			Key:  key,
			Pre:  nil,
			Next: nil,
		}
		lru.HashMap[node.Key] = node
	}
	lru.setHead(node)

}

func (lru *LRUCache) setHead(node *Node) {
	lru.Head.Next.Pre = node
	node.Next = lru.Head.Next
	lru.Head.Next = node
	node.Pre = lru.Head
}

func (lru *LRUCache) remove(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

/*============================*/

/*
使用方法
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
*/