package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

// функция для получения данных и перекладывания в dispatcherCh.
func dispatcherStage(in In, done In, dispatcherCh Bi) {
	//При выходе закроем dispatcherCh
	defer close(dispatcherCh)
	for {
		select {
		case value, ok := <-in:
			// Ждем значения
			if !ok {
				// Выходим если канал закрыт
				return
			}
			dispatcherCh <- value
			// Передадим полученное значение в dispatcherCh
		case <-done:
			// Завершим если получили сигнал done
			return
		}
	}
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		// Выполним все stage
		dispatcherCh := make(Bi)
		// Создадим канал для обработки stage
		go dispatcherStage(in, done, dispatcherCh)
		// Запустим горутину которая слушает in и прислушивается к done ;-)
		in = stage(dispatcherCh)
		// Выполним stage
	}
	return in
}
