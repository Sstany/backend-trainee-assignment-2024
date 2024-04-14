# Banney | Backend Avito Trainee Assignment 2024

## Развертывание

### Запуск всех нужных контейнеров

```
docker compose up -d
```

### Остановка всех контейнеров

```
docker compose down
```

### Остановка всех контейнеров с удалением данных

```
docker compose down -v
```

## Тестирование

## Endpoints

- [ ] GET /user_banner
- [ ] GET /banner
- [ ] POST /banner
- [ ] PATCH /banner/:id
- [ ] DELETE /banner/:id
- [ ] POST /auth/register
- [ ] POST /auth/login

## Вопросы и решения

#### Может ли быть два баннера с одинаковым feature_id и хотя бы одним одинаковым тегом?

О: Нет, не может быть, т.к. в условиях сказано, что любая пара feature_id и tag_id связана с единственным баннером.

#### Какой токен использовать?

О: Я решили использовать JWT токен, с кастомным Claims, в котором лежит user_id и is_admin для работы с доступами. Они удобные для работы и есть много информации об их использовании. Также использование стандартизированного подхода уменьшит вероятность рисков связанных со взломом.

#### Как сгенирировать токен?

О: Зарегистрировать пользователя отправив POST /auth/register со следующим json в Body:

```json
{
	"login":"user",
	"password":"pass",
	"is_admin":false
}
```

На такой запрос мы получим ответ (200 OK), содержащий токен или ошибку с деталями

```json
{
	"token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo4LCJpc19hZG1pbiI6ZmFsc2V9.b4dyBphPMBvQRqhz92TaANBZl5TFIleR8j79cirZbA8tI00E1tY-2-57V6LQoWpnZ_FN18jBObAWFLNiQc3ogQ"
}
```

Также можно воспользоваться методом /auth/login следующего формата, который возвращает аналогичный ответ, как и предыдущий метод

```json
{
	"login":"admin",
	"password":"securepass"
}
```


#### Может ли в PATCH /banner/:id в body быть не полный JSON баннера?

О: Я допустила, что в body содержится полный json с актуальными данными, которые обновляют уже существующие данные.
