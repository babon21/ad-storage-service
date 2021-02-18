# Сервис для хранения и подачи объявлений

<!-- ToC start -->
# Обзор

1. [Запуск приложения](#Запуск-приложения)
1. [Запуск линтера](#Запуск-линтера)
1. [Реализация](#Реализация)
1. [Усложнения](#Реализация)
1. [Используемые библиотеки](#Используемые-библиотеки)
1. [API](#API)
<!-- ToC end -->

# Запуск приложения
```
docker-compose up
``` 

# Запуск линтера
Предварительно необходимо установить линтер в локальный репозиторий командой:
```
make lint-prepare
```
Запуск линтера:
```
make lint
```

# Юнит-тесты
Запуск юнит-тестов:
```
make test
```

# Реализация
- Язык программирования Go
- База данных PostgreSQL
- HTTP JSON API
# Усложнения
- Следование принципам подхода "Чистая архитектура"
- Покрытие кода Юнит тестами (покрытие 73% всего кода)
- Использование индексов БД (создание в файле init.sql, с помощью EXPLAIN было проверено, что созданные индексы на деле используются)
- Докеризация приложения и запуск с помощью `docker-compose`
- Многоэтапная сборка Docker-образа сервиса
- Использование переменных окружения для конфигурации приложения (One of The Twelve-Factor App)
- Использование линтера `golangci-lint`
- Логирование HTTP запросов с помощью middleware
- Валидация запросов
# Используемые библиотеки
- Веб-фреймворк `echo`
- Конфигурирование приложения - библиотека `viper`
- `sqlx` для работы с БД
- Генерация mock-объектов для unit-тестирования - `mockery`
- Валидация запросов - `go-playground/validator`

# API
### POST /ads
**Создание объявления.**
Принимает поля: название, описание, несколько ссылок на фотографии, цена;

```
{
    "title": "some_title",
    "description": "some_description",
    "price": 200,
    "photos": [
        "https://rozetked.me/images/uploads/dwoilp3BVjlE.jpg",
        "https://cdn23.img.ria.ru/images/148839/96/1488399659_0:0:960:960_600x0_80_0_1_e38b72053fffa5d3d7e82d2fe116f0b3.jpg"
    ]
}
```
Возвращает ID созданного объявления и код результата (ошибка или успех).

Запрос:
```
curl --request POST 'localhost:8080/ads' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "some_title",
    "description": "some_description",
    "price": 200,
    "photos": [
        "https://rozetked.me/images/uploads/dwoilp3BVjlE.jpg",
        "https://cdn23.img.ria.ru/images/148839/96/1488399659_0:0:960:960_600x0_80_0_1_e38b72053fffa5d3d7e82d2fe116f0b3.jpg"
    ]
}'
```
Ответ: 200 Status
```
{
  "id": 1
}
```

### GET /ads?page=[число]&sort=[(-)поле]
**Получение списка объявлений.** 
- Пагинация: на одной странице должно присутствовать 10 объявлений;
- Cортировки: по цене (возрастание/убывание) и по дате создания (возрастание/убывание);
- Поля в ответе: название объявления, ссылка на главное фото (первое в списке), цена.

#### Пример запроса и ответа
Запрос:
```
curl --request GET 'localhost:8080/ads?page=1&sort=-price'
```
Ответ:
```
{
    "ads": [
        {
            "title": "some_title100",
            "price": 100,
            "main_photo": "https://rozetked.me/images/uploads/dwoilp3BVjlE.jpg"
        },
        {
            "title": "some_title20",
            "price": 20,
            "main_photo": "https://rozetked.me/images/uploads/dwoilp3BVjlE.jpg"
        }
    ]
}
```

### GET /ads?fields=[description,photos] по любым комбинациям полей
**Получение конкретного объявления.**
- Обязательные поля в ответе: название объявления, цена, ссылка на главное фото;
- Опциональные поля (можно запросить, передав параметр fields): описание, ссылки на все фото.


#### Пример запроса и ответа
Запрос:
```
curl --request GET 'localhost:8080/ads/1?fields=description,photos'
```
Ответ:
```
{
    "title": "some_title100",
    "description": "some_description100",
    "price": 100,
    "main_photo": "link3",
    "photos": [
        "https://rozetked.me/images/uploads/dwoilp3BVjlE.jpg"
    ]
}
```
