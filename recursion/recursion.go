package main

func main() {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[right], s[left] = s[left], s[right]
		left++
		right--
	}
}
