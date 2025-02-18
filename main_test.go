package main

import (
	"sync"
	"testing"
	"time"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	res := MaxInt(a, b)

	if res != b {
		t.Errorf("expected %d, got %d", b, res)
	}
}

func TestMain(m *testing.M) {
	var wg sync.WaitGroup
	numGoroutines := 100 // Увеличим количество горутин
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ { // Увеличим количество итераций
				main()                       // Запускаем main() несколько раз
				time.Sleep(time.Millisecond) // Небольшая задержка для усиления гонки
			}
		}()
	}

	wg.Wait()
}
