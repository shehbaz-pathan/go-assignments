## Ruuning Program
**Pre-Requisites**
- docker-compose

**Start DB**
1. Clone the repo and navigate to final_assignment/database folder
2. create a new direcotry
```sh
sudo mkdir -p /mysql/data
```
3. run the docker-compose up command from database folder
```sh
docker-compose up -d
```
**Test the App**
1. navigate to final_assignment folder
2. run the the go test command
```sh
go test -v 
```
```
=== RUN   TestAddBookDB
--- PASS: TestAddBookDB (0.00s)
=== RUN   TestGetBooksDB
--- PASS: TestGetBooksDB (0.00s)
=== RUN   TestGetBookByTitleDB
--- PASS: TestGetBookByTitleDB (0.00s)
=== RUN   TestAddBookHandler
[GIN] 2022/12/05 - 13:49:18 | 201 |     118.751µs |                 | POST     "/addbook"
--- PASS: TestAddBookHandler (0.00s)
=== RUN   TestGetBooksHandler
[GIN] 2022/12/05 - 13:49:18 | 200 |     302.934µs |                 | GET      "/getbooks"
--- PASS: TestGetBooksHandler (0.00s)
=== RUN   TestGetBookByTitleHandler
[GIN] 2022/12/05 - 13:49:18 | 200 |      62.567µs |                 | GET      "/getbookbytitle/The Comedy of Errors"
--- PASS: TestGetBookByTitleHandler (0.00s)
PASS
ok  	github.com/shehbaz-pathan/go-assignments/fina-assignment/book-info	0.006s
```
***Build the App**
1. run the go build command
```sh
go build .
```
**Run the App**
1. book-info executable file will be get generated, execute the file it will start the web-server on port 8888 with all the routes
```sh
./book-info
```
