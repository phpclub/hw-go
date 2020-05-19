package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in mN goroutines and stops its work when receiving mM errors from tasks.
// Переименовал названия потому что golangci-lint ругался на CAPS.
func Run(tasks []Task, nN int, mM int) error {
	if mM < 0 {
		return ErrErrorsLimitExceeded
	}
	var errCount int32
	var wg sync.WaitGroup
	var ch = make(chan Task, len(tasks))
	atomic.StoreInt32(&errCount, 0)
	// запускаем n горутин ждут задач из канала
	for i := 0; i < nN; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				v, ok := <-ch
				if !ok {
					//Канал закрыт выходим
					break
				}
				err := v()
				if err != nil {
					atomic.AddInt32(&errCount, 1)
				}
				if atomic.LoadInt32(&errCount) >= int32(mM) {
					//Кол-во ошибок >mM  выходим
					break
				}
			}
		}()
	}

	//range задач пулим в буферизированный канал
	for _, task := range tasks {
		ch <- task
	}
	close(ch)
	wg.Wait()
	if atomic.LoadInt32(&errCount) >= int32(mM) {
		//вернем ошибку есть превышен лимит ошибок
		return ErrErrorsLimitExceeded
	}
	return nil
}
