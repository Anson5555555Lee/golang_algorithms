package ht

import list "github.com/Anson5555555Lee/golang_algorithms/linked-list"

type HashTable struct {
	Table    map[int]*list.LinkedList
	Size     int
	Capacity int
}

type item struct {
	key   string
	value interface{}
}

func New(cap int) *HashTable {
	table := make(map[int]*list.LinkedList, cap)
	return &HashTable{table, 0, cap}
}

func (h *HashTable) Set(key, val string) {
	index := hash
}

// For simplicity reasons we will just use character codes sum of all characters of the key
// to calculate the hash.
//
// But you may also use more sophisticated approaches like polynomial string hash to reduce the
// number of collisions:
//
// hash = charCodeAt(0) * PRIME^(n-1) + charCodeAt(1) * PRIME^(n-2) + ... + charCodeAt(n-1)
//
// where charCodeAt(i) is the i-th character code of the key, n is the length of the key and
// PRIME is just any prime number like 31.

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum
}
