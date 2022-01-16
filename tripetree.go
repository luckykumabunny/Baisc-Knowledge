type Trie struct {
    isEnd bool
    next [26] *Trie
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
	for _,char := range word{
		if this.next[char-'a'] == nil {
			this.next[char-'a'] = new(Trie)
		}
		this = this.next[char-'a']
	}
	// 别忘了将结尾置为true
	this.isEnd = true
}

// search 方法要考虑结尾字母是否为true
func (this *Trie) Search(word string) bool {
	for _,char := range word{
		if this.next[char-'a'] == nil{
			return false
		}
		this = this.next[char-'a']
	}
	return this.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
	for _,char := range prefix{
		if this.next[char-'a'] == nil{
			return false
		}
		this = this.next[char-'a']
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */