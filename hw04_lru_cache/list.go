package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	// длина списка
	Len() int
	// первый Item
	Front() *ListItem
	// последний Item
	Back() *ListItem
	// добавить значение в начало
	PushFront(v interface{}) *ListItem
	// добавить значение в конец
	PushBack(v interface{}) *ListItem
	// удалить элемент
	Remove(i *ListItem)
	// переместить элемент в начало
	MoveToFront(i *ListItem)
}

type ListItem struct {
	// значение
	Value interface{}
	// следующий элемент
	Next *ListItem
	// предыдущий элемент
	Prev *ListItem
	// привязка к списку
	//list *list
}

type list struct {
	head   *ListItem
	tail   *ListItem
	length int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	if l.length == 0 {
		return nil
	}
	return l.tail
}

func (l *list) Back() *ListItem {
	if l.length == 0 {
		return nil
	}
	return l.head
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if l.head == nil {
		l.head = item
		l.tail = item
	} else {
		fi := l.tail
		l.tail = item
		fi.Next = item
		item.Prev = fi
	}
	l.length++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if l.tail == nil {
		l.tail = item
		l.head = item
	} else {
		ti := l.head
		l.head = item
		ti.Prev = item
		item.Next = ti
	}
	l.length++
	return item
}

func (l *list) Remove(i *ListItem) {
	hi := i.Prev // 10
	ti := i.Next //30
	switch {
	case hi == nil:
		ti.Prev = i.Prev
		l.head = ti
	case ti == nil:
		hi.Next = i.Next
		l.tail = hi
	default:
		hi.Next = i.Next
		ti.Prev = i.Prev
	}
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i.Value)
	l.Remove(i)
}

func NewList() List {
	return &list{}
}
