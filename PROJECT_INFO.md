# OrdersTest - Android проект

## 📱 Описание проекта

OrdersTest - это Android приложение на Java для работы с REST API заказов. Приложение позволяет просматривать, создавать и управлять заказами через удобный пользовательский интерфейс.

## 🏗️ Структура проекта

```
OrdersTest/
├── app/
│   ├── src/main/
│   │   ├── java/com/example/orderstest/
│   │   │   ├── adapter/          # Адаптеры для RecyclerView
│   │   │   ├── api/              # API интерфейсы и клиенты
│   │   │   ├── model/            # Модели данных
│   │   │   ├── MainActivity.java # Главная активность
│   │   │   ├── CreateOrderActivity.java
│   │   │   └── OrderDetailActivity.java
│   │   ├── res/
│   │   │   ├── layout/           # XML layouts
│   │   │   ├── values/           # Ресурсы (строки, цвета, темы)
│   │   │   └── xml/              # XML конфигурации
│   │   └── AndroidManifest.xml
│   └── build.gradle
├── build.gradle                  # Корневая конфигурация
├── settings.gradle              # Настройки проекта
├── gradle.properties            # Свойства Gradle
├── .gitignore                   # Исключения Git
└── README.md                    # Документация
```

## ✨ Основные возможности

- 📋 **Просмотр списка заказов** - отображение всех заказов с детальной информацией
- ➕ **Создание новых заказов** - добавление заказов с указанием пользователя, товара и количества
- 👁️ **Детальный просмотр заказа** - полная информация о заказе включая статус и даты
- 🔄 **Обновление статуса заказа** - изменение статуса (pending, confirmed, shipped, delivered, cancelled)
- 🔄 **Pull-to-refresh** - обновление списка заказов свайпом вниз
- 🎨 **Цветовая индикация статусов** - разные цвета для разных статусов заказов

## 🛠️ Технологии

- **Java** - основной язык программирования
- **Retrofit2** - для HTTP запросов к API
- **Gson** - для JSON парсинга
- **RecyclerView** - для отображения списка заказов
- **CardView** - для красивого отображения элементов
- **Material Design** - современный дизайн интерфейса
- **SwipeRefreshLayout** - для обновления данных
- **FloatingActionButton** - для быстрого создания заказов

## 🔌 API Endpoints

Приложение использует следующие эндпоинты:

- `GET /api/orders` - получение всех заказов
- `GET /api/orders/get?id={id}` - получение заказа по ID
- `POST /api/orders` - создание нового заказа
- `PUT /api/orders/update-status` - обновление статуса заказа

## 📊 Статусы заказов

- 🔶 **pending** - ожидает подтверждения
- 🔵 **confirmed** - подтвержден
- 🟣 **shipped** - отправлен
- 🟢 **delivered** - доставлен
- 🔴 **cancelled** - отменен

## 🚀 Запуск проекта

1. **Откройте в Android Studio:**
   - Запустите Android Studio
   - Выберите "Open an existing project"
   - Укажите папку OrdersTest

2. **Синхронизируйте Gradle:**
   - Дождитесь завершения синхронизации зависимостей

3. **Запустите приложение:**
   - Подключите Android устройство или запустите эмулятор
   - Нажмите "Run" (зеленая кнопка)

## 📝 Git статус

- ✅ Репозиторий инициализирован
- ✅ Базовая структура проекта создана
- ✅ Основные конфигурационные файлы добавлены
- ✅ .gitignore настроен для Android проекта
- ✅ Первоначальный коммит создан

## 🔧 Следующие шаги

1. **Создать недостающие Java файлы:**
   - ApiService.java
   - RetrofitClient.java
   - OrderAdapter.java
   - CreateOrderActivity.java
   - OrderDetailActivity.java

2. **Создать Layout файлы:**
   - activity_main.xml
   - item_order.xml
   - activity_create_order.xml
   - activity_order_detail.xml

3. **Добавить ресурсы:**
   - strings.xml
   - colors.xml
   - themes.xml

4. **Настроить зависимости в build.gradle**

5. **Протестировать приложение**

## 📱 Совместимость

- **Минимальная версия Android:** API 24 (Android 7.0)
- **Целевая версия Android:** API 34 (Android 14)
- **Язык программирования:** Java 8+
- **Среда разработки:** Android Studio Arctic Fox+

## 📄 Лицензия

MIT License

---

**OrdersTest** - Android приложение для управления заказами 