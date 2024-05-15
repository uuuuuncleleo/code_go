package main

type List interface {
	Add(index int, val any)
	Append(val any) error
	Delete(index int) error
}

type LinkedList struct {
	head node
	Head node
}

func (l *LinkedList) Add() {
	// 实现这个方法
}

type node struct {
}
