# Database API Gen 

### Goals 

- Generate a basic REST CRUD service and OpenAPI spec from an existing MYSQL database schema
- Load sample data into the database and provide a local test environment using Docker compose

### Dependencies 

- [Entgo](https://entgo.io/)
- MySQL 8.x
- OpenAPI 3.x
- Docker

### Data

Leveraged data from [https://github.com/datacharmer/test_db](https://github.com/datacharmer/test_db), but made a few changes to table names to default to singular naming and also added `id` columns to every table especially those with composite primary keys because the Entgo framework does not seem to support this quite yet (see: [https://github.com/ent/ent/issues/400(https://github.com/ent/ent/issues/400)]).

### Running

```bash
# bootstrap the DB and schema in ./schemas
make database-start 

# generate Ent schema/model from db (Docker required)
# generate openapi spec and service from schema 
make build

# run service on localhost:3000
make run

# stop db
make database-stop
```

Generated spec in: [./ent/openapi.json](./ent/openapi.json)

#### Examples

Example:

```bash
❯ curl -s -X 'GET' \
  "http://localhost:3000/employees" \
  -H 'accept: application/json'|jq
[
  {
    "id": "100000",
    "birth_date": "1956-01-11T00:00:00Z",
    "first_name": "Hiroyasu",
    "last_name": "Emden",
    "gender": "M",
    "hire_date": "1991-07-02T00:00:00Z",
    "created_at": "2024-05-02T18:32:55Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  {
    "id": "100001",
    "birth_date": "1953-02-07T00:00:00Z",
    "first_name": "Jasminko",
    "last_name": "Antonakopoulos",
    "gender": "M",
    "hire_date": "1994-12-25T00:00:00Z",
    "created_at": "2024-05-02T18:32:55Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
...
]
```

### References

- https://entgo.io/blog/2021/10/11/generating-ent-schemas-from-existing-sql-databases
- https://entgo.io/blog/2022/02/15/generate-rest-crud-with-ent-and-ogen/
