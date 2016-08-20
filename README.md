# bookshelf
Программа Bookshelf on GoLang

Одностраничное CRUD приложения - Библиотека книг. Позволяет редактировать список книг. Используется СУБД Postgresql, GoLang, Twitter 
Bootstrap, HTML. Скрипт SQL создания БД находится в файле database.sql. Логин БД и пароль БД = postgres (не стал заморачиваться с 
новым пользователем).

Само приложение состоит из нескольких модулей. Вообще можно было "запихнуть" все в один файл GO, так как это всего лишь пример, но 
решил все-таки сделать более-менее красиво. HTML странички не формируются из программы, а находятся в папке html, откуда вызываются 
приложением в зависимости от произведенных действий и возникших событий.

Использование: 
	1) git clone https://github.com/bar0metr/bookshelf.git 
	2) cd ./bookshelf 
	3) go build *.go 
	4) ./book

З.Ы.: к сожалению, функции отображения и редактирования БИБЛИОТЕК не реализовал, так как время было ограничено не 8 часами, а 3.
