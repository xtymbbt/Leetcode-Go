package main

import "fmt"

type Trie struct {
	val uint8
	next map[uint8]*Trie
	end bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		val:  0,
		next: make(map[uint8]*Trie, 0),
		end:  false,
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	insertVal(this.next, word, 0)
}

func insertVal(tries map[uint8]*Trie, word string, i int) {
	if i == len(word) {
		return
	}
	if v, ok := tries[word[i]];ok {
		if i == len(word) - 1 {
			tries[word[i]].end = true
		} else {
			insertVal(v.next, word, i+1)
		}
	} else {
		tries[word[i]] = &Trie{
			val:  word[i],
			next: make(map[uint8]*Trie, 0),
			end:  false,
		}
		if i == len(word) - 1 {
			tries[word[i]].end = true
		} else {
			insertVal(tries[word[i]].next, word, i+1)
		}
	}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return searchVal(this.next, word, 0)
}

func searchVal(tries map[uint8]*Trie, word string, i int) bool {
	if v, ok := tries[word[i]];ok {
		if i == len(word) - 1 {
			return v.end
		} else {
			return searchVal(v.next, word, i+1)
		}
	} else {
		return false
	}
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return searchStart(this.next, prefix, 0)
}

func searchStart(tries map[uint8]*Trie, word string, i int) bool {
	if v, ok := tries[word[i]];ok {
		if i == len(word) - 1 {
			return true
		} else {
			return searchStart(v.next, word, i+1)
		}
	} else {
		return false
	}
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor();

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // 返回 true
	fmt.Println(trie.Search("app"))     // 返回 false
	fmt.Println(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	fmt.Println(trie.Search("app"))     // 返回 true

}