### consul
`consul agent -dev -bind=127.0.0.1`

### test
* model `SERVICE_ENV=test MYSQL_USERNAME=mysql_username MYSQL_PASSWORD=mysql_pwd MYSQL_DB_NAME=db_name go test -v ./model`
* service `SERVICE_ENV=test MYSQL_USERNAME=mysql_username MYSQL_PASSWORD=mysql_pwd MYSQL_DB_NAME=db_name REGISTRY_SERVER_HOST=registry_host REGISTRY_SERVER_PORT=registry_port go test -v ./service`

### run
`MYSQL_USERNAME=mysql_username MYSQL_PASSWORD=mysql_pwd MYSQL_DB_NAME=db_name go run main.go`


#### user
* create `micro query user Service.CreateUser '{"user": {"id": 1, "name": "zouqilin", "Phone": "12345678", "Email": "mr.zouqilin@gmail.com", "Role": 0}}'`
* update `micro query user Service.UpdateUser '{"user": {"id": 1, "name": "zouqilin", "Phone": "87654321", "Email": "mr.zouqilin@gmail.com", "Role": 0}}'`
* get one `micro query user Service.GetUser '{"user": {"id": 1}}'`
* get batch `micro query user Service.GetUsers '{"ids": 1}'`
* delete `micro query user Service.DeleteUser '{"user": {"id": 1}}'`
