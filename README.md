# Cloud Native Go

Learning Go by building simple ReST APIs. This code was written by following the Cloud Native Go course on LinkedIn Learning.

## Build the image
```
docker-compose build
```
## Start the application
```
docker-compose -d up
```
## Stop the application
```
docker-compose down
```
## Accessing the APIs
Get all the books
```
curl -H "Content-Type: application/json" -X GET http://localhost:8000/api/books
```
Get a book by id
```
curl -H "Content-Type: application/json" -X GET http://localhost:8000/api/books/:bookId
```
Add a new book
```
curl -H "Content-Type: application/json" -X POST http://localhost:8000/api/books -d '{
    "title": "Book Title",
    "author": "Book Author",
    "isbn": "Book ISBN",
    "description": "Book Description"
}'
```
Update a book
```
curl -H "Content-Type: application/json" -X PUT http://localhost:8000/api/books/:bookId -d '{
    "title": "Book Title",
    "author": "Book Author",
    "isbn": "Book ISBN",
    "description": "New Description"
}'
```
