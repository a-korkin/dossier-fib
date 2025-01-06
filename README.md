## Web api written on fiber framework

### Request examples
```http
# getting all persons
GET localhost:8000/persons 

# getting person by id
GET localhost:8000/persons/1

# creating person
POST localhost:8000/persons
Content-Type: application/json

{
    "last_name": "Bickle",
    "first_name": "Travis",
    "middle_name": "Unknown",
    "age": 23
}

# updating person
PUT localhost:8000/persons/1
Content-Type: application/json

{
    "last_name": "BickleUpd",
    "first_name": "TravisUpd",
    "middle_name": "UnknownUpd",
    "age": 24
}

# deleting person
DELETE localhost:8000/persons/1

# creating payment
POST localhost:8000/payments/1
Content-Type: application/json

{
    "sum": 43.03
}

# getting payment by person
GET localhost:8000/payments/1
```
