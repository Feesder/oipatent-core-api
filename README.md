# OiPatent Backend Application

Этот проект представляет собой REST API для работы с патентами. На текущий момент реализованы JWT авторизация пользователей и CRUD-операции над патентами. Приложение разработано на Go c использованием библиотеки Gin для отработки HTTP-запросов.

## Установка

1. Клонирование репозиторий: 
```
git clone https://github.com/Feesder/oipatent-core-api.git
cd oipatent-core-api
```
2. Настройка конфигурации
Создайте конфигурационный файл на основе шаблона и укажите необходимые параметры:
```
cp config/local.example.yaml config/local.yaml
```
3. Применение миграций базы данных
Установите утилиту: https://github.com/golang-migrate/migrate
После установки выполните миграцию:
```
migrate -path ./schema -database <DATABASE_URL> up
```
4. Запуск приложения
```
CONFIG_PATH=./config/local.yaml go run cmd/main.go 
```

## Структура проекта
Проект следует классической многослойной архитектуре
```
oipatent-core-api/
├── cmd/
│   └── main. go                    # Точка входа приложения
│
├── config/
│   └── local.example.yaml         # Пример конфигурации для локальной разработки
│
├── internal/
│   ├── common/                    # Общие компоненты и утилиты
│   │   ├── dto/                   # Data Transfer Objects (DTO) для передачи данных
│   │   ├── entity/                # Модели данных (Entity)
│   │   ├── lib/                   # Общие библиотеки и утилиты
│   │   └── mapper/                # Мапперы для преобразования данных между слоями
│   │
│   ├── config/
│   │   └── config. go              # Загрузка и управление конфигурацией
│   │
│   └── modules/                   # Бизнес-логика по модулям
│       ├── handler/               # HTTP обработчики (контроллеры)
│       ├── repository/            # Доступ к данным (DAL - Data Access Layer)
│       └── service/               # Бизнес-логика (Business Logic Layer)
│
├── schema/                        # Миграции базы данных
│   ├── 000001_init. up.sql        # Миграция создания таблиц
│   └── 000001_init.down.sql      # Откат миграции
│
├── go.mod                         # Определение модуля и зависимости
├── go.sum                         # Хеши зависимостей
└── . gitignore                     # Игнорируемые файлы
```