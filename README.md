# Azure Go Postgres and MySQL Quick Test
Quick test to try out building a connection to Azure PostgreSQL and MySQL

## How to Run
Replace `{{var name}}` with the appropriate variable from the Azure portal. Upon creation of a PostgreSQL or a MySQL
instance, you should quickly be able find username, password, the server name under the `Connection strings` tab.
- `go get github.com/devigned/azure-golang-pg-mysql`
- `cd $GOPATH/src/github.com/devigned/azure-golang-pg-mysql`
- `go build -o azsql`
- `./azsql mysql -n ‘{{username}}:{{password}}@tcp({{server_name}}.mysql.database.azure.com:3306)/mysql?tls=skip-verify&allowNativePasswords=true’`
- `./azsql pg -n 'user={{username}} password={{password}} host={{server_name}}.postgres.database.azure.com dbname=postgres port=5432'`

