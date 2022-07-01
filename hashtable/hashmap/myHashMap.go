package hashmap

type MyHashMap struct {
	Set [10][]Node
}

type Node struct {
	key   int
	Value int
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func HashKey(key int) int {
	return key % 10
}

func (this *MyHashMap) Put(key int, value int) {
	hash := HashKey(key)

	exist := false
	for i, v := range this.Set[hash] {
		if v.key == key {
			changeValue(&this.Set[hash][i], value)
			exist = true
		}
	}

	if !exist {
		newNode := Node{
			key:   key,
			Value: value,
		}
		this.Set[hash] = append(this.Set[hash], newNode)
	}
}

func changeValue(node *Node, value int) {
	node.Value = value
}

func (this *MyHashMap) Get(key int) int {
	hash := HashKey(key)

	for _, v := range this.Set[hash] {
		if v.key == key {
			return v.Value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	hash := HashKey(key)
	for i, v := range this.Set[hash] {
		if v.key == key {
			temp := this.Set[hash][i+1:]
			this.Set[hash] = this.Set[hash][0:i]
			this.Set[hash] = append(this.Set[hash], temp...)
		}
	}
}
