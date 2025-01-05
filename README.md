## Web api written on fiber framework

### Request examples
```http
GET localhost:8000/persons 

GET localhost:8000/persons/1

POST localhost:8000/persons
Content-Type: application/json

{
    "last_name": "Bickle",
    "first_name": "Travis",
    "middle_name": "Unknown",
    "age": 23
}

PUT localhost:8000/persons/1
Content-Type: application/json

{
    "last_name": "BickleUpd",
    "first_name": "TravisUpd",
    "middle_name": "UnknownUpd",
    "age": 24
}

DELETE localhost:8000/persons/1

POST localhost:8000/payments/1
Content-Type: application/json

{
    "sum": 43.03
}
```
