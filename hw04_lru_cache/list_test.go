package hw04_lru_cache //nolint:golint,stylecheck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, l.Len(), 0)
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, l.Len(), 3)

		middle := l.Back().Next // 20
		l.Remove(middle)        // [10, 30]
		require.Equal(t, l.Len(), 2)

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, l.Len(), 7)
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Back(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{50, 30, 10, 40, 60, 80, 70}, elems)
	})
}

func internalTestMoveFront(t *testing.T, call string) {
	// Инит
	l := NewList()

	// Тестовые данные
	l.PushFront(10) // [10]
	l.PushBack(20)  // [10, 20]
	l.PushBack(30)  // [10, 20, 30]
	l.PushBack(40)  // [10, 20, 30, 40]
	l.PushBack(50)  // [10, 20, 30, 40, 50]

	// Отладка для проверки
	for i := l.Back(); i != nil; i = i.Next {
		println(fmt.Sprintf("Pointers %d: %p %v", i.Value.(int), i, i))

	}
	println("==^ Before ^ ==")
	middle := l.Back().Next // 40
	switch call {
	case "MoveToFront":
		l.MoveToFront(middle)
	case "MoveToFrontBad":
		l.MoveToFrontBad(middle)
	default:
		panic("Unknow call")
	}
	elems := make([]int, 0, l.Len())
	for i := l.Back(); i != nil; i = i.Next {
		println(fmt.Sprintf("Pointers %d: %p %v", i.Value.(int), i, i))
		elems = append(elems, i.Value.(int))
	}
	println("==^ After ^ ==")
	//Типа все хорошо, значения ожидаемы
	require.Equal(t, []int{50, 30, 20, 10, 40}, elems)

	switch call {
	case "MoveToFront":
		require.Truef(t,
			*middle == *l.Front(),
			fmt.Sprintf("Pointers not equal %p != %p", &*middle, &*l.Front()))
	case "MoveToFrontBad":
		//Но на самом деле мы создали новый элемент, а не переместили его
		//Проверяем что указатели не равны
		require.Falsef(t,
			*middle == *l.Front(),
			fmt.Sprintf("Pointers equal %p == %p", &*middle, &*l.Front()))
	}
}

// Написать тест, что MoveFront перемещает, а не пересоздает элемент.
// Убедиться, что тест падает.
// Пофиксить код и убедиться, что тест проходит.
// https://github.com/phpclub/hw-go/pull/5#pullrequestreview-395906126
// Решил оставить оба теста для наглядности
func TestMoveFrontBad(t *testing.T) {
	internalTestMoveFront(t, "MoveToFrontBad")
}

func TestMoveFront(t *testing.T) {
	internalTestMoveFront(t, "MoveToFront")
}
