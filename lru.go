package main

import (
	"errors"
	"sync"
)

type Node struct {
	head bool
	key int
	value int
	Pre *Node
	Next *Node
}

type LRU struct {
	Cap int32
	Len int32
	Head *Node
	Tail *Node
	RelationMap map[int]*Node
	Lock sync.RWMutex
}

var Lru *LRU

func init() {
	Lru.RelationMap = make(map[int]*Node, 0)
	Lru.Head = &Node{true,-1, -1, nil, nil}
}

func (l *LRU) Set(key int,value int) error {
	l.Lock.Lock()
	defer l.Lock.Unlock()

	if l.Len >= l.Cap {
		deleteKey := l.Tail.key
		l.Tail.Pre = nil
		l.Tail.Next = nil
		delete(l.RelationMap, deleteKey)
		l.Len--
	}
	node := &Node{false, key, value,nil,nil}
	second := l.Head.Next
	l.Head.Next = node
	node.Pre = l.Head
	node.Next = second
	second.Pre = node

	l.RelationMap[key] = node
	l.Len++
	if l.Len == 1 {
		l.Tail = node
	}
	return nil
}

func (l *LRU) Get(key int) (int, error) {
	l.Lock.RLock()
	defer l.Lock.RUnlock()

	if value, ok := l.RelationMap[key]; ok {
		if l.Len >1 && !value.Pre.head {
			// 交接
			pre := value.Pre
			next := value.Next
			pre.Next = next
			if next != nil {
				next.Pre = pre
			} else {
				l.Tail = pre
			}

			// 加入
			second := l.Head.Next
			l.Head.Next = value
			value.Pre = l.Head
			value.Next = second
			second.Pre = value
		}
	}

	return 0, errors.New("not found")
}