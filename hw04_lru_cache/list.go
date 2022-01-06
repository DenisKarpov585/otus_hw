package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.first
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	elem := &ListItem{v, nil, nil}
	if l.len == 0 {
		l.first = elem
		l.last = elem
	} else {
		elem.Next = l.first
		l.first.Prev = elem
		l.first = elem
		if l.last.Prev == nil {
			l.last = l.first
		}
	}
	l.len++
	return l.first
}

func (l *list) PushBack(v interface{}) *ListItem {
	elem := &ListItem{v, nil, nil}
	if l.len == 0 {
		l.first = elem
		l.last = elem
	} else {
		elem.Prev = l.last
		l.last.Next = elem
		l.last = elem
		if l.first.Next == nil {
			l.first = l.last
		}
	}
	l.len++
	return l.last
}

func (l *list) Remove(elem *ListItem) {
	if l.Check(elem) {
		if elem.Prev == nil {
			l.first = elem.Next
		} else {
			elem.Prev.Next = elem.Next
		}
		if elem.Next == nil {
			l.last = elem.Prev
		} else {
			elem.Next.Prev = elem.Prev
		}
		l.len--
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if l.Check(i) {
		l.Remove(i)
		l.PushFront(i.Value)
	}
}

func (l *list) Check(i *ListItem) bool {
	ends := l.first == i || l.last == i
	if !ends {
		if i.Next == nil || i.Prev == nil {
			return false
		}
		return !(i.Next.Prev.Value != i.Value || i.Next.Prev.Prev != i.Prev || i.Prev.Next.Next != i.Next)
	}
	return true
}
