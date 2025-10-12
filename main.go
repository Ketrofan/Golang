package main

//Подробнее разобаться в копиях, что можно поменять. а что нельзя
import "fmt"

func main() {
	//Для функции суммирования двух чисел
	number := sum(1, 2)
	fmt.Printf("Вы получили число: %d\n", number)
	fmt.Println("_______________________")

	//Для функции сложения двух строк
	text1 := "Привет "
	text2 := "Мир!"

	//Условие для проверки диапозона чисел и вывода текста из func sumString
	if number < 10 {
		fmt.Printf("Вы ввели корректный диапозон чисел < 10, число: %d\n", number)
		text3 := sumString(text1, text2)
		fmt.Println(text3)
		fmt.Println("_______________________")

	} else {
		fmt.Println("Вы не попали в дапозон числе < 10")
		fmt.Println("_______________________")

	}

	//Вывод строки из func str
	s := str()
	fmt.Println(s)
	fmt.Println("_______________________")

	//true or false для func subscribed
	sub := subscribed(number)
	fmt.Println(sub)
	fmt.Println("_______________________")

	//фФункциия приветсвия с ранним выходом если false
	greeting("Andrey")
	fmt.Println("_______________________")

	//Использование отложенного вызова функции (defer)/stack
	fmt.Println("Я main я начал свою работу")
	defer func() {
		fmt.Println("Я main я закончил свою работу")
	}()
	fmt.Println("Hello1")
	fmt.Println("Hello2")
	fmt.Println("Hello3")
	fmt.Println("Hello4")

	exampleDefer()
	fmt.Println("Тут я закончил выполнять func exampleDefer ")
	fmt.Println("_______________________")

	//Пример завершения канала dataBase при помощи defer
	dataBase()
	fmt.Println("_______________________")

	//Указатели

	numN := 568

	pointer := &numN
	fmt.Println(pointer)

	//Указатели на типы данных
	examplePointer(pointer)

	//Пример укзтеля nil
	// Разыменовывание указателя nil приводит к фатальной ошибке
	/* numEx := 24

	var ptr *int //Нулевой ууказатель делется исключительно через var переменная *тип_данны
	fmt.Println(numEx)
	fmt.Println(ptr)
	//для разыменовывания указателя nil пделается проверка
	if ptr != nil {
		fmt.Println("Разименуем указатель", *ptr)
	} else {
		fmt.Println("Получен nil ууказатель, разыменование невозможно")
	} */

	/* struct-это логически объедененные переменные, под одной какой-то сущностью
	все переменные в сущности-это поле структуры
	Дейсвия которые мы можем делть с этой струуктурой-это методы */
}

// Функция для суммирования двух чисел
func sum(a, b int) int {
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("_______________________")
	s := a + b
	return s
}

// Функция для конкатенации двух строк
func sumString(s1, s2 string) string {
	s := s1 + s2

	return s
}

// Функция для создания строки
func str() string {
	return "Это моя строка"
}

// Функция для проверки подписки (Условное ветвление)
func subscribed(x int) string {
	if x < 10 {
		return "Вы подписаны на канал"
	}
	return "Вы не подписаны на канал"
}

/* Функция для приветствия пользователя (Ранний выход из функции)
Если функция ничего не возвращает, то можно использовать return без значения для раннего
выхода из функции. */
func greeting(name string) {
	if name == "" {
		fmt.Println("Вы не указали имя")
		return
	}
	fmt.Println("Привет многоуважаемый,", name)
}

// Функция defer будет работать как stack, первую положил, последней первую вытщил
func exampleDefer() {
	fmt.Println("Тут я начал выполнять func exampleDefer ")
	defer func() {
		fmt.Println("Defer1")
	}()
	defer func() {
		fmt.Println("Defer2")
	}()
	defer func() {
		fmt.Println("Defer3")
	}()
}

/* // Пример с базой данных которой надо завершить канал, иначе каналы могут зкончиться */
func dataBase() {
	fmt.Println("Создать подключение")
	defer func() {
		fmt.Println("Закрыть подключение")
	}()

	boolean := false

	if boolean {
		fmt.Println("Работа с базой данных")
		// продолжаем выполнение — не выходим из функции, чтобы выполнить цикл
	} else {
		fmt.Println("Ошибка подключения к базе данных")
		return
	}
	// Ищем первую успешную итерацию (например, когда i чётное).
	// Не печатаем "Ошибка" на каждую неудачную попытку — это делает вывод чище.
	found := false
	for i := 1; i <= 5; i++ {
		// Симуляция работы в итерации
		fmt.Println("Работа с циклом в базе данных, итерация:", i)
		if i%2 == 0 {
			fmt.Println("Продолжить работу с базой данных")
			found = true
			break
		}
		// не печатаем ошибку здесь — просто пробуем следующую итерацию
	}
	if !found {
		fmt.Println("Не удалось найти успешную итерацию в цикле базы данных")
	}

}

//Указатели на типы данных
/* Дай мне адрес переменной делается через &. Когда делаем через * и тип,
пердставленный тип данных является указателем на какую то из переменных  */
func examplePointer(n *int) {
	fmt.Println(n)
	//Для вывода самого резултата по адресу используем * (разыменование)
	fmt.Println(*n)
}
