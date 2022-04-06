package main

import "errors"

//链表节点定义
type LinkTableNode struct {
	Next *LinkTableNode
	Data interface{}
}

//链表定义
type LinkTable struct {
	Head   *LinkTableNode
	length int
}

//创建链表,空链表有一个虚拟头节点
func InitLinkTable() *LinkTable {
	pLinkTable := new(LinkTable)
	pLinkTable.length = 0
	pLinkTable.Head = &LinkTableNode{
		Next: nil,
		Data: nil,
	}
	return pLinkTable
}

//获取链表头节点
func (l LinkTable) GetHeadNode() *LinkTableNode {
	return l.Head.Next
}

//获取链表长度
func (l LinkTable) GetLength() int {
	return l.length
}

//获取index位置的结点
func (l LinkTable) GetLinkTableNode(index int) (*LinkTableNode, error) {
	if l.length < index {
		return nil, errors.New("beyond max length")
	}
	temp := l.Head
	for index > 1 {
		temp = temp.Next
	}
	return temp.Next, nil
}

//查找值是否存在
func (l LinkTable) IsExist(v interface{}) (bool, error) {
	temp := l.Head
	for temp.Next != nil {
		if temp.Data == v {
			return true, nil
		}
		temp = temp.Next
	}
	return false, errors.New("not exist")
}

//添加节点
func (l *LinkTable) AddLinkTableNode(Node *LinkTableNode) {
	temp := l.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = Node
}

//删除指定位置(从1开始)节点,返回被删除的结点
func (l *LinkTable) DeleteLinkTableNode(index int) (*LinkTableNode, error) {
	if l.length < index {
		return nil, errors.New("beyond max length")
	}
	temp := l.Head
	for index > 1 {
		temp = temp.Next
	}
	remove := temp.Next
	temp.Next = temp.Next.Next
	return remove, nil
}
