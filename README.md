# Merch-Shop API

## Описание
Merch-Shop — это сервис для перевода монет между пользователями, покупки товаров за монеты и получения информации о пользователях и их инвентаре.

## Установка и запуск

1. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/gfarida/merch-shop.git
    cd merch-shop
    ```

2. Установите зависимости:
    ```sh
    go mod tidy
    ```

3. Создайте файл `.env` в корне проекта с переменной окружения:
    ```env
    JWT_SECRET_KEY=your_secure_jwt_secret_key_here
    ```

4. Запустите проект с использованием Docker:
    ```sh
    docker-compose up --build
    ```

5. Приложение будет доступно по адресу: `http://localhost:8080`.

---

## API

### 1. Логин пользователя
- **POST** `/api/auth/login`
- **Тело запроса:**
    ```json
    {
        "username": "your_username",
        "password": "your_password"
    }
    ```

- **Ответ:**
    ```json
    {
        "token": "your_jwt_token_here"
    }
    ```

### 2. Получение информации о пользователе
- **GET** `/api/info`
- **Заголовки:**
    ```http
    Authorization: Bearer your_jwt_token_here
    ```

- **Ответ:**
    ```json
    {
        "coins": 1000,
        "inventory": [
            {"type": "t-shirt", "quantity": 2},
            {"type": "book", "quantity": 1}
        ],
        "coinHistory": {
            "received": [{"fromUser": "user1", "amount": 50}],
            "sent": [{"toUser": "user2", "amount": 30}]
        }
    }
    ```

### 3. Перевод монет
- **POST** `/api/sendCoin`
- **Заголовки:**
    ```http
    Authorization: Bearer your_jwt_token_here
    ```

- **Тело запроса:**
    ```json
    {
        "toUser": "user2",
        "amount": 50
    }
    ```

### 4. Покупка товара
- **GET** `/api/buy/{item}`
- **Заголовки:**
    ```http
    Authorization: Bearer your_jwt_token_here
    ```

- **Ответ:**
    ```json
    {
        "message": "Покупка успешна"
    }
    ```

---

## Сценарии использования

1. **Авторизация:**  
    Пользователь отправляет запрос с логином и паролем, получает JWT и использует его для доступа к защищённым маршрутам.

2. **Перевод монет:**  
    Пользователь может отправлять монеты другим пользователям, если у него достаточно средств.

3. **Покупка товара:**  
    Пользователь может купить товар, если у него хватает монет.

4. **Получение информации:**  
    Пользователь может запросить информацию о своём балансе, инвентаре и истории транзакций.