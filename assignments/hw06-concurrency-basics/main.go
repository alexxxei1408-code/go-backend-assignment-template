package main

// SafeCounter должен быть безопасен для использования из нескольких горутин.
// Реализацию рекомендуется строить на основе map и sync.Mutex или sync.RWMutex.
type SafeCounter struct {
	// TODO: добавьте необходимые поля (map, мьютекс и т.д.)
}

// NewSafeCounter создаёт новый счётчик.
func NewSafeCounter() *SafeCounter {
	// TODO: инициализируйте структуру
	return &SafeCounter{}
}

// Inc увеличивает значение счётчика для ключа key на 1.
func (c *SafeCounter) Inc(key string) {
	// TODO: реализуйте потокобезопасное увеличение счётчика
}

// Value возвращает текущее значение счётчика для ключа key.
// Если ключ ранее не встречался, должно возвращаться 0.
func (c *SafeCounter) Value(key string) int {
	// TODO: реализуйте безопасное чтение значения
	return 0
}

// Total возвращает сумму значений по всем ключам.
func (c *SafeCounter) Total() int {
	// TODO: реализуйте безопасное вычисление суммы
	return 0
}

// SumConcurrently считает сумму элементов слайса nums.
//
// Требования:
//   - Вычисление должно быть корректным независимо от количества workers.
//   - При workers <= 0 следует использовать один worker.
//   - Для эффективности рекомендуется разбивать работу на несколько горутин
//     и синхронизировать их через каналы или sync.WaitGroup.
func SumConcurrently(nums []int, workers int) int {
	// TODO: реализуйте конкурентное суммирование
	return 0
}

func main() {}


