## 2. Distributed config

### (Тестовое задание в рамках GoCloudCamp)

---

### Технические требования:

- реализация задания может быть выполнена на любом языке программирования
- формат схемы конфига на выбор: json/protobuf
- сервис должен обеспечивать персистентность данных
- сервис должен поддерживать все CRUD операции по работе с конфигом
- должно поддерживаться версионирование конфига при его изменении
- удалять конфиг допускается только если он не используется никаким приложением

---

### Пример использования сервиса:

Создание конфига:

curl -d "@data.json" -H "Content-Type: application/json" -X POST http://localhost:8080/config

    {
        "service": "managed-k8s",
        "data": [
            {"key1": "value1"},
            {"key2": "value2"}
        ]
    }

Получение конфига:

curl http://localhost:8080/config?service=managed-k8s

    {"key1": "value1", "key2": "value2"}

---

## API:

Добавление нового конфига:

curl -d "@data.json" -H "Content-Type: application/json" -X POST http://localhost:8080/config

---

Получение последнего конфига:

curl http://localhost:8080/config?service=managed-k8s

---

Изменение последнего конфига с добавлением подверсии:

curl -d "@data.json" -H "Content-Type: application/json" -X PUT http://localhost:8080/config

---

Удаление неиспользуемого системой конфига:

curl -X DELETE http://localhost:8080/config?service=managed-k8s&version=2

---

Текущая версия конфига:

curl http://localhost:8080/version?service=managed-k8s

---

## Описание работы сервиса:

Сервис является прослойкой между сообщением сервера и базы данных.

На старте сервиса устанавливается последняя версия конфига, при создании или изменении данных ,конфигурация сервиса не изменяется.

Для того чтобы в реальном времени изменить конфигурацию на последнюю версию нужно произвести GET запрос:

curl http://localhost:8080/config?service= \*

с указанием названия сервиса в качестве параметра "service"

### Настройка postgresql:

Запуск Postgresql:

    $ sudo service postgresql start

Смена пользователя:

    $ su - postgres

Создание базы данных:

    $ createdb cloud

Переход в интерактивный терминал Postgresql

    $ psql cloud

Создание таблицы:

    SQL: CREATE TABLE config(version serial,service text,data jsonb);
    ALTER TABLE config
    ALTER COLUMN version
    TYPE numeric;

Установка пароля для подключения к базе данных(в тесте используется пароль "cloud"):

    $ \password
