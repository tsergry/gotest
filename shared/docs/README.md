# Orders - Общая документация

Мультиплатформенное приложение для управления заказами с Go backend и Android клиентом.

## Структура проекта

```
├── backend/           # Go backend сервер
├── mobile/           # Мобильные приложения
│   └── android/     # Android приложение
├── shared/           # Общие ресурсы
│   ├── api/         # API спецификации
│   └── docs/        # Общая документация
├── docs/            # Документация проекта
└── scripts/         # Общие скрипты
```

## API Endpoints

### Users (Пользователи)
- `GET /api/users` - Получить всех пользователей
- `POST /api/users` - Создать пользователя
- `GET /api/users/get` - Получить пользователя по ID
- `PUT /api/users/update` - Обновить пользователя
- `DELETE /api/users/delete` - Удалить пользователя

### Products (Товары)
- `GET /api/products` - Получить все товары
- `POST /api/products` - Создать товар
- `GET /api/products/get` - Получить товар по ID
- `GET /api/products/category` - Получить товары по категории
- `PUT /api/products/update` - Обновить товар
- `DELETE /api/products/delete` - Удалить товар

### Orders (Заказы)
- `GET /api/orders` - Получить все заказы
- `POST /api/orders` - Создать заказ
- `GET /api/orders/get` - Получить заказ по ID
- `GET /api/orders/user` - Получить заказы пользователя
- `GET /api/orders/status` - Получить заказы по статусу
- `PUT /api/orders/update-status` - Обновить статус заказа
- `DELETE /api/orders/delete` - Удалить заказ

## Модели данных

### User (Пользователь)
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Product (Товар)
```json
{
  "id": 1,
  "name": "iPhone 15",
  "description": "Смартфон Apple iPhone 15 128GB",
  "price": 89999.99,
  "category": "Electronics",
  "stock": 10,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Order (Заказ)
```json
{
  "id": 1,
  "user_id": 1,
  "product_id": 1,
  "quantity": 1,
  "total_price": 89999.99,
  "status": "pending",
  "order_date": "2024-01-01T00:00:00Z",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## Статусы заказов

- `pending` - Ожидает подтверждения
- `confirmed` - Подтвержден
- `shipped` - Отправлен
- `delivered` - Доставлен
- `cancelled` - Отменен

## Запуск проекта

### Backend
```bash
cd backend
go mod tidy
go run main.go
```

Сервер запустится на `http://localhost:8080`

### Android
```bash
cd mobile/android
./gradlew assembleDebug
```

## База данных

Проект использует PostgreSQL. Для инициализации базы данных:

```bash
cd backend/cmd/init_db
go run main.go
```

## Переменные окружения

Создайте файл `.env` в директории `backend/`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=orders_db
```

## Разработка

### Требования
- Go 1.21+
- PostgreSQL 12+
- Android Studio (для Android разработки)
- Java 11+ (для Android)

### Рекомендации по разработке

1. **API First подход**: Все изменения API должны сначала отражаться в `shared/api/openapi.yaml`
2. **Документация**: Обновляйте документацию при изменении API
3. **Тестирование**: Пишите тесты для backend и Android приложения
4. **Версионирование**: Используйте семантическое версионирование для API

## CI/CD

Рекомендуется настроить CI/CD для:
- Автоматического тестирования backend
- Сборки Android приложения
- Деплоя на staging/production серверы

## Мониторинг

Рекомендуется добавить:
- Логирование запросов
- Метрики производительности
- Мониторинг ошибок
- Health checks для API 