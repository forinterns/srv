Написать код для
1) подключения к БД (Mysql/Postgres)
2) в для хэндлеров add добавить код который позволит записать данные и вернет ID, get отдаст данные по ID
структура запроса add
 ```curl --location --request POST 'http://localhost:8080/employee/add' \
 --header 'Content-Type: application/json' \
 --data-raw '{
   "name": "name",
   "surname": "surname",
   "patronymic": "patronymic",
   "skills": {
     "soft": [
       "languages",
       "motivation"
     ],
     "hard": [
       "golang",
       "math"
     ]
   },
   "salary": {
     "main": 1000,
     "bonus": 500.15
   }
 }'
 ```
запрос get
 ```
curl --location --request POST 'http://localhost:8080/employee/get' \
--header 'Content-Type: application/json' \
--data-raw '{
  "employeeId": 123
}'
 ```
в ответе приложить sql для создания схемы в БД

