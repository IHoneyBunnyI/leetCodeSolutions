package main

import "fmt"

//*
//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func (l *ListNode) GetNode(index int) *ListNode {
	if index < 0 {
		return nil
	}
	cur := l

	for cur != nil && index > 0 {
		cur = cur.Next
		index--
	}
	return cur
}

func PrintList(l *ListNode) {
	cur := l
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	fmt.Print("nil\n")
}

func (l *ListNode) DeleteNode(index int) *ListNode {
	if index < 0 { // Не валидный индекс
		fmt.Print("Invalid index")
		return nil
	}
	cur := l
	var prev *ListNode

	for cur != nil && index > 0 {
		prev = cur
		cur = cur.Next
		index--
	}
	if prev == nil { //Если 0 индекс
		l = l.Next // Заменяем первый элемент на следующий
		return l
	}
	if cur == nil { //Если данного индекса нет, не удаляем
		return l
	}
	prev.Next = cur.Next
	return l
}

func (l *ListNode) AddElem(index int, node *ListNode) *ListNode {
	if index < 0 { // Не валидный индекс
		fmt.Print("Invalid index")
		return nil
	}
	cur := l
	var prev *ListNode

	for cur != nil && index > 0 {
		prev = cur
		cur = cur.Next
		index--
	}
	if prev == nil { //Если 0 индекс
		node.Next = l
		l = node // Заменяем первый элемент на тот что пришел
		return l
	}
	if cur == nil { //Если данного индекса нет, не удаляем
		return l
	}
	prev.Next = node
	node.Next = cur
	return l
}

func main() {
	head := new(ListNode)
	head.Val = 1

	n2 := &ListNode{2, nil}
	head.Next = n2
	n3 := &ListNode{3, nil}
	n2.Next = n3
	n4 := &ListNode{4, nil}
	n3.Next = n4
	n5 := &ListNode{5, nil}
	n4.Next = n5

	PrintList(head)
	//newL := head.DeleteNode(0)
	//head.DeleteNode(10)
	node := &ListNode{100, nil}
	newL := head.AddElem(200, node)
	PrintList(newL)
}
