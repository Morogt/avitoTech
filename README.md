# avitoTech
Запуск докера: docker-compose up --build avito-tech


Запуск бд: migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' up


Создаём 2 пользователей: localhost:8000/user/create


Body:

{
    "id": 1,
    "balance": 0
}


Ответ:

{
    "id": 1
}

Body:

{
    "id": 2,
    "balance": 0
}

Ответ:

{
    "id": 2
}

Добавляем пользователю деньги на счёт: localhost:8000/balance/addFunds


Body:

{
    "id": 1,
    "count": 100
}


Ответ: 

{
    "added": 100,
    "balance": 100,
    "id": 1
}


Пользователь выводит средства: localhost:8000/balance/withdrawalOfMoney


Body:

{
    "id": 1,
    "count": 50
}

Ответ:

{
    "balance": 50,
    "id": 1,
    "withdraw": 50
}


Пользотватель переводит средства: localhost:8000/balance/transfer

Body:

{
    "sender" : 1,
    "recipient": 2,
    "count": 25
}

Ответ: 

{
    "recipientBalance": 25,
    "recipientId": 2,
    "senderBalance": 25,
    "senderId": 1
}


Пользотватель покупает что-либо: localhost:8000/balance/pay

Body:

{
    "userId": 2,
    "serviceId": 1,
    "count": 25,
    "description": "Покупка БУ стола"
}

Ответ:

{
    "balance": 0,
    "orderId": 5,
    "userId": 2
}

Получаем отчёт по сервисам: localhost:8000/reports/get

Ответ:

{
    "": [
        {
            "id": 1,
            "amount": 25
        }
    ]
}

Получаем историю пользователя: localhost:8000/balance/history

Body: 

{
    "id": 2
}

Ответ:

{
    "": [
        {
            "OrderId": 4,
            "ServiceId": 0,
            "Amount": 25,
            "Description": "Перевод от пользователя",
            "Refill": true,
            "Time": "2022-11-14T05:59:12.523194Z"
        },
        {
            "OrderId": 5,
            "ServiceId": 1,
            "Amount": 25,
            "Description": "Покупка БУ стола",
            "Refill": false,
            "Time": "2022-11-14T05:59:23.803567Z"
        }
    ]
}

Узнать баланс пользователья: localhost:8000/balance/info

Body:

{
    "id": 1
}

Ответ:

{
    "balance": 25,
    "id": 1
}
