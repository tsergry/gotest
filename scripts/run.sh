#!/bin/bash

# Orders Project - Скрипт запуска
# Использование: ./scripts/run.sh [backend|android|all]

set -e

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Функция для вывода сообщений
print_message() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_header() {
    echo -e "${BLUE}================================${NC}"
    echo -e "${BLUE} $1${NC}"
    echo -e "${BLUE}================================${NC}"
}

# Проверка наличия необходимых инструментов
check_requirements() {
    print_message "Проверка требований..."
    
    # Проверка Go
    if ! command -v go &> /dev/null; then
        print_error "Go не установлен. Установите Go 1.21+"
        exit 1
    fi
    
    # Проверка PostgreSQL (базовая)
    if ! command -v psql &> /dev/null; then
        print_warning "PostgreSQL клиент не найден. Убедитесь, что PostgreSQL установлен"
    fi
    
    # Проверка Android SDK (если нужно)
    if [ "$1" = "android" ] || [ "$1" = "all" ]; then
        if ! command -v adb &> /dev/null; then
            print_warning "Android SDK не найден. Установите Android Studio"
        fi
    fi
}

# Запуск backend
run_backend() {
    print_header "Запуск Backend (Go)"
    
    cd backend
    
    # Проверка модуля
    print_message "Проверка Go модуля..."
    go mod tidy
    
    # Проверка подключения к БД
    print_message "Проверка подключения к базе данных..."
    if ! go run cmd/init_db/main.go &> /dev/null; then
        print_warning "Не удалось подключиться к базе данных. Проверьте настройки в .env"
    fi
    
    # Запуск сервера
    print_message "Запуск сервера на http://localhost:8080"
    go run main.go
}

# Запуск Android
run_android() {
    print_header "Запуск Android приложения"
    
    cd mobile/android
    
    # Проверка Gradle
    print_message "Проверка Gradle..."
    if [ ! -f "gradlew" ]; then
        print_error "Gradle wrapper не найден"
        exit 1
    fi
    
    # Сборка приложения
    print_message "Сборка Android приложения..."
    ./gradlew assembleDebug
    
    # Установка на устройство (если подключено)
    if adb devices | grep -q "device$"; then
        print_message "Установка приложения на устройство..."
        ./gradlew installDebug
    else
        print_warning "Android устройство не подключено. Запустите эмулятор или подключите устройство"
    fi
}

# Инициализация базы данных
init_database() {
    print_header "Инициализация базы данных"
    
    cd backend/cmd/init_db
    
    print_message "Создание таблиц и тестовых данных..."
    go run main.go
    
    print_message "База данных инициализирована успешно!"
}

# Показать справку
show_help() {
    echo "Использование: $0 [команда]"
    echo ""
    echo "Команды:"
    echo "  backend     - Запустить только backend сервер"
    echo "  android     - Собрать и запустить Android приложение"
    echo "  all         - Запустить backend и Android"
    echo "  init-db     - Инициализировать базу данных"
    echo "  help        - Показать эту справку"
    echo ""
    echo "Примеры:"
    echo "  $0 backend"
    echo "  $0 android"
    echo "  $0 all"
    echo "  $0 init-db"
}

# Основная логика
main() {
    case "${1:-help}" in
        "backend")
            check_requirements backend
            run_backend
            ;;
        "android")
            check_requirements android
            run_android
            ;;
        "all")
            check_requirements all
            run_backend &
            BACKEND_PID=$!
            sleep 3
            run_android
            kill $BACKEND_PID
            ;;
        "init-db")
            check_requirements
            init_database
            ;;
        "help"|*)
            show_help
            ;;
    esac
}

# Запуск скрипта
main "$@" 