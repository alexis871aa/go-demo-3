package main

import "fmt"

/*
	Создать приложение, которое сначала выдаёт меню:

- 1. Посмотреть закладки
- 2. Добавить закладку
- 3. Удалить закладку
- 4. Выход
При 1 - Выводит закладки
При 2 - 2 поля ввода названия и адреса и после добавление
При 3 - Ввод названия и удаление по нему
При 4 - Завершение
*/
type bookmarksMap = map[string]string

func main() {
	bookmarks := make(bookmarksMap)
	fmt.Println("Приложение для закладок")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmarks(bookmarks)
		case 3:
			deleteBookmarks(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	fmt.Println("Выберите вариант")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	var variant int
	fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks bookmarksMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookmarks {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func addBookmarks(bookmarks bookmarksMap) {
	var newBookmarkKey, newBookmarkValue string
	fmt.Print("Введите название: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Введите ссылку: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
}

func deleteBookmarks(bookmarks bookmarksMap) {
	var bookmarkKeyToDelete string
	fmt.Print("Введите название: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
}
