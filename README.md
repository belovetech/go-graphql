# go-graphql

# Set up gqlgen

1. go run github.com/99designs/gqlgen init
2. go run github.com/99designs/gqlgen generate

# MYSQL

```bash
    go get -u github.com/go-sql-driver/mysql
    go build -tags 'mysql' -ldflags="-X main.Version=1.0.0" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/
    cd internal/pkg/db/migrations/
    migrate create -ext sql -dir mysql -seq create_users_table
    migrate create -ext sql -dir mysql -seq create_links_table
    migrate -database mysql://root:dbpass@/hackernews -path internal/pkg/db/migrations/mysql up
```
