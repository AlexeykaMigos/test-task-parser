# test-task-parser

## Запуск в Docker

1. При необходимости создайте `.env` в корне проекта (или задайте переменные окружения):

```env
PROXY_URL=
CATEGORIES=1,2
STORE_ID=1052
```

2. Запустите парсер:

```bash
docker compose up --build
```

3. Результаты выгрузки будут сохранены в директорию `./output` на хосте.
