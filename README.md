# Go Project

Простой проект на Go с HTTP сервером и PostgreSQL.

## Требования

- Go 1.21+
- PostgreSQL

## Настройка базы данных

1. Установите PostgreSQL
2. Создайте базу данных:
```sql
CREATE DATABASE myproject;
```

3. Выполните SQL скрипт для создания таблиц:
```bash
psql -d myproject -f scripts/init_db.sql
```

## Настройка переменных окружения

Скопируйте файл конфигурации:
```bash
cp configs/env.example .env
```

Отредактируйте `.env` файл с вашими настройками базы данных.

## Запуск

```bash
go mod tidy
go run main.go
```

Сервер будет доступен по адресу: http://localhost:8080

## Структура проекта

```
.
├── main.go          # Основной файл приложения
├── go.mod           # Файл модуля Go
├── README.md        # Документация
├── cmd/             # Директория для исполняемых файлов
├── internal/        # Внутренние пакеты
├── pkg/             # Публичные пакеты
├── api/             # API определения
├── configs/         # Конфигурационные файлы
├── docs/            # Документация
├── scripts/         # Скрипты
├── test/            # Тестовые файлы
└── vendor/          # Зависимости (если используется)
```

## API Endpoints

### Пользователи
- `GET /` - главная страница
- `GET /api/users` - получить всех пользователей
- `POST /api/users` - создать нового пользователя
- `GET /api/users/get?id=1` - получить пользователя по ID
- `PUT /api/users/update` - обновить пользователя
- `DELETE /api/users/delete?id=1` - удалить пользователя

### Товары
- `GET /api/products` - получить все товары
- `POST /api/products` - создать новый товар
- `GET /api/products/get?id=1` - получить товар по ID
- `GET /api/products/category?category=Electronics` - получить товары по категории
- `PUT /api/products/update` - обновить товар
- `DELETE /api/products/delete?id=1` - удалить товар

### Заказы
- `GET /api/orders` - получить все заказы
- `POST /api/orders` - создать новый заказ
- `GET /api/orders/get?id=1` - получить заказ по ID
- `GET /api/orders/user?user_id=1` - получить заказы пользователя
- `GET /api/orders/status?status=pending` - получить заказы по статусу
- `PUT /api/orders/update-status` - обновить статус заказа
- `DELETE /api/orders/delete?id=1` - удалить заказ

## Команды

- `go run main.go` - запуск приложения
- `go build` - сборка приложения
- `go test ./...` - запуск тестов
- `go mod tidy` - очистка зависимостей

## Примеры запросов

### Пользователи

#### Создание пользователя
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

#### Получение всех пользователей
```bash
curl http://localhost:8080/api/users
```

#### Получение пользователя по ID
```bash
curl http://localhost:8080/api/users/get?id=1
```

### Товары

#### Создание товара
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"name":"iPhone 15","description":"Смартфон Apple","price":89999.99,"category":"Electronics","stock":10}'
```

#### Получение всех товаров
```bash
curl http://localhost:8080/api/products
```

#### Получение товара по ID
```bash
curl http://localhost:8080/api/products/get?id=1
```

#### Получение товаров по категории
```bash
curl http://localhost:8080/api/products/category?category=Electronics
```

### Заказы

#### Создание заказа
```bash
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -d '{"user_id":1,"product_id":1,"quantity":2}'
```

#### Получение всех заказов
```bash
curl http://localhost:8080/api/orders
```

#### Получение заказа по ID
```bash
curl http://localhost:8080/api/orders/get?id=1
```

#### Получение заказов пользователя
```bash
curl http://localhost:8080/api/orders/user?user_id=1
```

#### Получение заказов по статусу
```bash
curl http://localhost:8080/api/orders/status?status=pending
```

#### Обновление статуса заказа
```bash
curl -X PUT http://localhost:8080/api/orders/update-status?id=1 \
  -H "Content-Type: application/json" \
  -d '{"status":"confirmed"}'
``` 