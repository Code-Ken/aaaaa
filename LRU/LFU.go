package LRU

type LFUCache struct {
	cache map[int]*Node
	freq map[int]*DoubleList
	cap, size, minFreq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache {
		cache: make(map[int]*Node),
		freq: make(map[int]*DoubleList),
		cap: capacity,
	}
}

func (lfu *LFUCache) Get(key int) int {
	if node, ok := lfu.cache[key]; ok {
		lfu.IncFreq(node)
		return node.val
	}
	return -1
}

func (lfu *LFUCache) Put(key, value int) {
	if lfu.cap == 0 {
		return
	}
	if node, ok := lfu.cache[key]; ok {
		node.val = value
		lfu.IncFreq(node)
	} else {
		if lfu.size >= lfu.cap {
			node := lfu.freq[lfu.minFreq].RemoveLast()
			delete(lfu.cache, node.key)
			lfu.size--
		}
		x := &Node{key: key, val: value, freq: 1}
		lfu.cache[key] = x
		if lfu.freq[1] == nil {
			lfu.freq[1] = CreateDL()
		}
		lfu.freq[1].AddFirst(x)
		lfu.minFreq = 1
		lfu.size++
	}
}

func (lfu *LFUCache) IncFreq(node *Node) {
	_freq := node.freq
	lfu.freq[_freq].Remove(node)
	if lfu.minFreq == _freq && lfu.freq[_freq].IsEmpty() {
		lfu.minFreq++
		delete(lfu.freq, _freq)
	}

	node.freq++
	if lfu.freq[node.freq] == nil {
		lfu.freq[node.freq] = CreateDL()
	}
	lfu.freq[node.freq].AddFirst(node)
}

type DoubleList struct {
	head, tail *Node
}

type Node struct {
	prev, next *Node
	key, val, freq int
}

func CreateDL() *DoubleList {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
	return &DoubleList {
		head: head,
		tail: tail,
	}
}

func (dl *DoubleList) AddFirst(node *Node) {
	node.next = dl.head.next
	node.prev = dl.head

	dl.head.next.prev = node
	dl.head.next = node
}

func (dl *DoubleList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = nil
	node.prev = nil
}

func (dl *DoubleList) RemoveLast() *Node {
	if dl.IsEmpty() {
		return nil
	}

	last := dl.tail.prev
	dl.Remove(last)

	return last
}

func (dl *DoubleList) IsEmpty() bool {
	return dl.head.next == dl.tail
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

