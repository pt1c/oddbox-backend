# Odd Box
Сервераная (backend) часть приложения для организации вещей.  
Frontend будет доступен в другом репозитории.

## Проблема:
Часто случается так, что ты покупаешь вещи, которые тебе нужны немедлено (отвертка, таблетка от головы, батарейка, etc), но проблема в том, что дома (на даче\кладовке\гараже\ящике) эта вещь уже есть. Ты просто забыл!

Чтобы избежать покупки вещей которые у тебя уже есть создается приложение, которое ведет (категоризирует) весь этот хлам.

## Пример:
Врач выписал антибиотик 7 таблеток. Открываешь приложение и смотришь, что дома в аптечке лежит как раз 7 таблеток с нормальным сроком годности.

Или надо починить протекающую трубу, - ищешь в приложении и понимаешь, что есть нужный инструмен и лежит он в гараже на 2-ой полке.

Или нужно каталогизировать электронные компоненты и их количество.

## Как запустить
Для запуска нужен `docker`. 
Переименовать файл `.env.sample` в `.env`
В консоле `docker-compose up`
При успешном старте Backend будет доступен по адресу `http://localhost:8080`

## Seed
Для сидирования БД нужно запустить файл сидирования командой `go run ./cmd/seedUsers.go`  
Если приложение работает в `docker` подключиться к контейнеру и выполнить `go run /app/cmd/seedUsers.go`

## Схема таблиц
Схема таблиц веденся в приложении [dbdiagram.io](https://dbdiagram.io/)  
Схемма доступна по [ссылке](https://dbdiagram.io/d/64abd94802bd1c4a5ecb3624)  
Экспортированный файл в формате [DBML](https://dbml.dbdiagram.io/docs) со схемой `dbscheme.dbml`

## Приложение использует модули
- [Fiber v2](https://gofiber.io/) ([Fiber Github](https://github.com/gofiber/fiber))
- [Fiber JWT](https://github.com/gofiber/contrib/tree/main/jwt)
- [Validator v10](https://pkg.go.dev/github.com/go-playground/validator/v10)
- [GORM](https://gorm.io/) + [MySQL Driver](https://github.com/go-gorm/mysql)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [GoDotENV](https://github.com/joho/godotenv)
- [Faker v4](github.com/bxcodec/faker)

## Референс
- [Go Ambassador](https://github.com/antoniopapa/go-ambassador) by AntonioPapa


