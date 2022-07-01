package hashset

type MyHashSet struct {
	Set [10][]int
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func HashKey(key int) int {
	return key % 10
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.Set[HashKey(key)] = append(this.Set[HashKey(key)], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	hash := HashKey(key)
	list := this.Set[hash]
	for i, value := range list {
		if value == key {
			temp := this.Set[hash][i+1:]
			this.Set[hash] = this.Set[hash][0:i]
			this.Set[hash] = append(this.Set[hash], temp...)
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	list := this.Set[HashKey(key)]
	for _, value := range list {
		if value == key {
			return true
		}
	}
	return false
}
