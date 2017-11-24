### consul
`consul agent -dev -bind=127.0.0.1`

### test
* model `MYSQL_USERNAME=mysql_username MYSQL_PASSWORD=mysql_pwd MYSQL_DB_NAME=db_name go test -v ./model`
* service `go test -v ./service`

### run
`MYSQL_USERNAME=mysql_username MYSQL_PASSWORD=mysql_pwd MYSQL_DB_NAME=db_name go run main.go`
