# Orders

Мультиплатформенное приложение для управления заказами с Go backend и Android клиентом.

## 🏗️ Структура проекта

```
├── backend/           # Go backend сервер
│   ├── cmd/          # Команды (init_db)
│   ├── internal/     # Внутренние пакеты
│   │   ├── database/ # Работа с БД
│   │   ├── handlers/ # HTTP обработчики
│   │   └── models/   # Модели данных
│   ├── configs/      # Конфигурации
│   └── main.go       # Точка входа
├── mobile/           # Мобильные приложения
│   └── android/     # Android приложение
│       ├── app/     # Основное приложение
│       └── OrdersTest/ # Тестовое приложение
├── shared/           # Общие ресурсы
│   ├── api/         # API спецификации
│   └── docs/        # Общая документация
├── docs/            # Документация проекта
└── scripts/         # Общие скрипты
```

## 🚀 Быстрый старт

### Backend (Go)

```bash
# Переход в директорию backend
cd backend

# Установка зависимостей
go mod tidy

# Инициализация базы данных
cd cmd/init_db
go run main.go

# Запуск сервера
cd ../..
go run main.go
```

Сервер запустится на `http://localhost:8080`

### Android

```bash
# Переход в директорию Android
cd mobile/android

# Сборка приложения
./gradlew assembleDebug

# Или открытие в Android Studio
open app/
```

## 📋 Требования

- **Backend**: Go 1.21+, PostgreSQL 12+
- **Android**: Android Studio, Java 11+
- **База данных**: PostgreSQL

## 🔧 Конфигурация

### Переменные окружения

Создайте файл `.env` в директории `backend/`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=orders_db
```

### База данных

Проект использует PostgreSQL. Для инициализации:

```bash
cd backend/cmd/init_db
go run main.go
```

## 📚 API Документация

Полная спецификация API доступна в `shared/api/openapi.yaml`

### Основные endpoints:

- **Users**: `/api/users` - управление пользователями
- **Products**: `/api/products` - управление товарами  
- **Orders**: `/api/orders` - управление заказами

## 🧪 Тестирование

```bash
# Backend тесты
cd backend
go test ./...

# Android тесты
cd mobile/android
./gradlew test
```

## 📦 Развертывание

### Backend

```bash
cd backend
go build -o orders-server main.go
./orders-server
```

### Android

```bash
cd mobile/android
./gradlew assembleRelease
```

## 🤝 Разработка

### Рекомендации

1. **API First**: Изменения API сначала в `shared/api/openapi.yaml`
2. **Документация**: Обновляйте docs при изменении API
3. **Тестирование**: Пишите тесты для всех изменений
4. **Версионирование**: Используйте семантическое версионирование

### Структура коммитов

```
feat: добавить новый endpoint для заказов
fix: исправить ошибку в обработке JSON
docs: обновить API документацию
test: добавить тесты для ProductHandler
```

## 📊 Мониторинг

Рекомендуется настроить:

- Логирование запросов
- Метрики производительности  
- Мониторинг ошибок
- Health checks

## 📄 Лицензия

MIT License

## 🤝 Вклад в проект

1. Fork репозитория
2. Создайте feature branch
3. Внесите изменения
4. Добавьте тесты
5. Создайте Pull Request

## 📞 Поддержка

Для вопросов и предложений создавайте Issues в репозитории. 