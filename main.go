package main

import (
	"errors"
	"fmt"
)

type Trie struct {
	Children []*Trie
	Ch       byte
}

func NewTrie() (trie *Trie) {
	trie = &Trie{}
	trie.Children = make([]*Trie, 36)
	return
}

func (m *Trie) Put(word string) (err error) {

	cur := m
	pos := '0'

	for _, ch := range word {
		//只接受小写字母和数字
		if ch >= '0' && ch <= '9' {
			pos = ch - '0'
		} else if ch >= 'a' && ch <= 'z' {
			pos = ch - 'a' + 10
		} else {
			return errors.New("char error")
		}

		//查询当前字母位置是否有树
		if cur.Children[pos] == nil {
			cur.Children[pos] = NewTrie()
			cur.Children[pos].Ch = byte(ch)
		}

		cur = cur.Children[pos]
	}

	return

}

func (m *Trie) Search(word string) (isExists bool) {

	cur := m
	pos := '0'
	for _, ch := range word {

		//只接受小写字母和数字
		if ch >= '0' && ch <= '9' {
			pos = ch - '0'
		} else if ch >= 'a' && ch <= 'z' {
			pos = ch - 'a' + 10
		} else {
			return false
		}

		if cur.Children[pos] == nil {
			return false
		}

		cur = cur.Children[pos]
	}

	return true

}

func main() {

	trie := NewTrie()
	trie.Put("you")
	trie.Put("me")
	trie.Put("he123")

	fmt.Println(trie.Search("he123"))

}
