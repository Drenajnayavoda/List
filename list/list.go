package list

import "fmt"

// List представляет односвязный список
type List struct {
	len       int64
	firstNode *node
}

// NewList создает новый список
func NewList() *List {
	return &List{}
}

// Len возвращает длину списка
func (l *List) Len() int64 {
	return l.len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) int64 {
	newNode := &node{value: value}
	if l.firstNode == nil {
		l.firstNode = newNode
	} else {
		cur := l.firstNode
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = newNode
	}
	l.len++
	return l.len - 1
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if id < 0 || id >= l.len {
		return
	}
	if id == 0 {
		l.firstNode = l.firstNode.next
	} else {
		prev := l.firstNode
		for i := int64(0); i < id-1; i++ {
			prev = prev.next
		}
		prev.next = prev.next.next
	}
	l.len--
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	if l.firstNode != nil && l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
		return
	}
	cur := l.firstNode
	for cur != nil && cur.next != nil {
		if cur.next.value == value {
			cur.next = cur.next.next
			l.len--
			break
		} else {
			cur = cur.next
		}
	}
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	if l.firstNode != nil && l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
	}
	cur := l.firstNode
	for cur != nil {
		if cur.next != nil && cur.next.value == value {
			cur.next = cur.next.next
			l.len--
		} else {
			cur = cur.next
		}
	}
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (int64, bool) {
	if id < 0 || id >= l.len {
		return 0, false
	}
	cur := l.firstNode
	for i := int64(0); i < id; i++ {
		cur = cur.next
	}
	return cur.value, true
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (int64, bool) {
	cur := l.firstNode
	index := int64(0)
	for cur != nil {
		if cur.value == value {
			return index, true
		}
		cur = cur.next
		index++
	}
	return 0, false
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) ([]int64, bool) {
	var indxs []int64
	cur := l.firstNode
	index := int64(0)

	for cur != nil {
		if cur.value == value {
			indxs = append(indxs, index)
		}
		cur = cur.next
		index++
	}
	if len(indxs) > 0 {
		return indxs, true
	}
	return nil, false
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() ([]int64, bool) {
	if l.len == 0 {
		return nil, false
	}
	values := make([]int64, l.len)
	cur := l.firstNode
	index := int64(0)
	for cur != nil {
		values[index] = cur.value
		cur = cur.next
		index++
	}
	return values, true
}

// Clear очищает список
func (l *List) Clear() {
	l.len = 0
	l.firstNode = nil
}

// Print выводит список в консоль
func (l *List) Print() {
	cur := l.firstNode
	for cur != nil {
		fmt.Printf("%d ", cur.value)
		cur = cur.next
	}
	fmt.Println()
}
