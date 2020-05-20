package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	// Добавить значение в кэш по ключу
	Set(key Key, value interface{}) bool
	// Получить значение из кэша по ключу
	Get(key Key) (interface{}, bool)
	// Очистить кэш
	Clear()
}

type ListItemValue struct {
	iKey   Key
	iValue interface{}
}

type lruCache struct {
	//ёмкость (количество сохраняемых в кэше элементов)
	capacity int
	//очередь [последних используемых элементов] на основе двусвязного списка
	queue List
	//словарь, отображающий ключ (строка) на элемент очереди
	items map[Key]*ListItem
}

// при добавлении элемента:
// если элемент присутствует в словаре, то обновить его значение и переместить элемент в начало очереди;
// если элемента нет в словаре, то добавить в словарь и в начало очереди
// (при этом, если размер очереди больше ёмкости кэша, то необходимо удалить
// последний элемент из очереди и его значение из словаря);
// возвращаемое значение - флаг, присутствовал ли элемент в кэше.
func (l *lruCache) Set(key Key, value interface{}) bool {
	item := ListItem{Value: ListItemValue{iKey: key, iValue: value}}
	_, bExists := l.Get(key)
	if bExists {
		l.updateItemValue(key, value)
		return bExists
	}
	l.items[key] = l.queue.PushFront(&item)
	if l.capacity < l.queue.Len() {
		removeItem := l.queue.Back()
		if removeItem != nil {
			l.queue.Remove(removeItem)
			delete(l.items, removeItem.Value.(*ListItem).Value.(ListItemValue).iKey)
		}
	}
	return bExists
}

// Обновление значения по ключу.
func (l *lruCache) updateItemValue(key Key, value interface{}) bool {
	item, ok := l.items[key]
	if ok {
		item.Value.(*ListItem).Value = ListItemValue{iKey: key, iValue: value}
		return true
	}
	return false
}

// при получении элемента:
// если элемент присутствует в словаре, то переместить элемент в начало очереди и вернуть его значение и true;
// если элемента нет в словаре, то вернуть nil и false (работа с кешом похожа на работу с map).
func (l *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(item)
		return item.Value.(*ListItem).Value.(ListItemValue).iValue, true
	}
	return nil, false
}

// для очистки текущего кеша надо передавать
// ссылку иначе мы ничего не очистим.
func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
