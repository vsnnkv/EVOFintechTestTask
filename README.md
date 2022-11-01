# EVOFintechTestTask

Запустити проект:
1. docker-compose build
2. docker-compose up

Команди:
1. Зберігти csv файл, данні з нього в БД (GET): http://localhost:8080/saveTransactions
Відповідь:
- 200, {"status": "success"}
- 500, {"error": Error}

2. Отримати данні з БД (GET): http://localhost:8080/transactions
Можливі фільтри:
- transaction_id
- terminal_id
- status
- payment_type
- from
- to
- payment_narrative

Відповіді:
- 200, [{Transaction}]
- 500, {"error": Error}
