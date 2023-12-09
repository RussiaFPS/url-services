## url-services
# Краткое описание
Какое хранилище использовать, указывается параметром в dockerfile: ENTRYPOINT
Выбор:
- pst = PostgresSQL
- memory = in-memory

Основные данные из *docker-compose.yml*
- Пользователь для БД: postgres
- Пароль для БД: qwer1234
- Порт для БД: 5430

Основные данные для подключения к БД прописанны в env файле в папке envs, необходима
база данных __url-services__ она создается при выполнении команды make up автоматически.

> При первом запуске использовать __make up__, где включена сборка докер контейнеров,
создание бд, миграции табличек. Возможно нужно будет поменять username в makefile и данные от БД.

> При повторном запуске используйте команду __make start__, где выполнится только сборка докер контейнеров.

Так же можно запустить вручную через __docker-compose up -d__

Сервис работает через GRPC, proto файл находится в ./internal/controller/proto

В директории ./pkg/postgres/migrations/schema отображена схема БД urls, которую можно создать вручную

Примеры запросов:

 ```
 localhost:8080 url:Get
 request:
{
   "shortUrl": "7dzQKRRdZk"
}

response:
{
    "originalUrl": "test"
}
 ```

 ```
 localhost:8080 url:Save
 request:
{
   "originalUrl": "test"
}

response:
{
    "shortUrl": "7dzQKRRdZk"
}
 ```