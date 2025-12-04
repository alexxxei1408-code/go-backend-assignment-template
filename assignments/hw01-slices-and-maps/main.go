package main

// Purchase описывает одну покупку пользователя.
type Purchase struct {
	UserID   int64
	Category string
	Amount   int64 // сумма в условных минимальных единицах (например, центы)
}

// UserStats агрегирует статистику по пользователю.
type UserStats struct {
	UserID      int64
	TotalAmount int64 // суммарные траты пользователя
	Categories  int   // количество различных категорий покупок
}

// TopUsersBySpending возвращает слайс из не более чем topN пользователей
// с максимальными суммарными тратами.
//
// Требования:
//   - Суммарные траты пользователя считаются как сумма Amount по всем его покупкам.
//   - Поле Categories должно содержать количество различных категорий для пользователя.
//   - Результат должен быть отсортирован:
//   - по TotalAmount по убыванию,
//   - при равенстве TotalAmount — по UserID по возрастанию.
//   - Исходный слайс purchases изменять нельзя.
//   - При topN <= 0 функция должна вернуть пустой слайс.
func TopUsersBySpending(purchases []Purchase, topN int) []UserStats {
	// TODO: реализуйте функцию, используя слайсы и map для агрегации
	return nil
}

// CategoryTotals возвращает сумму трат по каждой категории.
//
// Ключом map является название категории, значением — суммарная Amount по всем покупкам.
// Если purchases пустой, должна возвращаться пустая (но не nil) map.
func CategoryTotals(purchases []Purchase) map[string]int64 {
	// TODO: реализуйте функцию, аккуратно работая с map и обработкой пустых случаев
	return nil
}

func main() {}
