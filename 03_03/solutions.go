package solutions

/* Дальше мы планируем использовать созданную нами сегодня библиотеку. Поэтому
давайте ее немного улучшим.
1. Переименуйте функцию ContainsInt на InSliceInt.
2. Так как это изменение названия функции не обратно совместимо, примите
необходимые меры по правильному версионированию библиотеки.
3. В ответе к заданию скиньте ссылку на итоговый репозиторий в github-е. */

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
func InSliceInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}