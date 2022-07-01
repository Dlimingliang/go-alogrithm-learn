package hashset

type MyHashSet struct {
	Set []int
}

func Constructor() MyHashSet {
	return MyHashSet{Set: []int{}}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.Set = append(this.Set, key)
	}
}

func (this *MyHashSet) Remove(key int) {
	for i, value := range this.Set {
		if value == key {
			temp := this.Set[i+1:]
			this.Set = this.Set[0:i]
			this.Set = append(this.Set, temp...)
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, value := range this.Set {
		if value == key {
			return true
		}
	}
	return false
}
