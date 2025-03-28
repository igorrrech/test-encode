# enCode test task
## API
Функции реализованного api:
- **GET** ```/person/?limit=<uint>&offset=<uint>&search<string>``` - Данные всех лиц. 
Параметры: *limit*, *offset*, *search*; реализуют выборку записей с лимитом по количеству, смещение относительно *id*
и поиск по *first_name* соответственно.
- **GET**```/person/:id``` - Данные одного человека по *:id*
- **POST**```/person/``` - Добавление человека в систему.\
Request body:
```json
"person":{
    "id":1,//"не используется"
    "email":"example@example.com",
    "phone":"1234567890",
    "first-name":"Foo",
    "last-name":"Bar"
}
```
- **DELETE**```/person/:id``` - Удаление человека по *:id*
- **UPDATE**```/person/:id``` - Обновлние данных о человеке по *:id*\
Request body:
```json
"person":{
    "id":1,//"не используется"
    "email":"example@example.com",
    "phone":"1234567890",
    "first-name":"Foo",
    "last-name":"Bar"
}
```
## Организация кода
*./app* - Файлы моделей\
*./persondb* - Файлы репозитория и провайдера соединений с бд\
*./internal/*:
- */config* Файл подгрузки конфигурации из *config.json* и *app.env* (переменные перечислены в .env)
- */http/* 
    - */handlers* Все обработчики запросов
    - */middleware* Создание сессий, логирование запросов и ошибок
- */logic* Бизнес логика, моки, интерфейсы usecase и репозитория
- *docker-compose.yaml* - Конфигурация контейнеров postgresql, pgadmin и приложения
- *init.sql* - Первичная инициализация бд.
- *Makefile* - Скрипты для быстрого запуска и остановки контейнеров
## Немного о использовании
Запуск осуществляется через подъем всех сервисов в compose.\
Разработка велась на базе wsl (Linux).\
При запуска менеджер бд доступен по адресу: [pgadmin](http://localhost:5050/browser/)
- Имя хоста - контейнер базы данных (test-db-pg)
- Имя базы - из .env файла (POSTGRES_DB)
- Имя пользователя - из .env файла (POSTGRES_USER)
- Пароль - из .env файла (POSTGRES_PASSWORD)


Задаются при первом запуске, вместе со скриптами init.sql.
Для очистки нужно усыпить *docker-compose volumes*
