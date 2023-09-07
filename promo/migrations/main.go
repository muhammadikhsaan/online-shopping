package main

import (
	"context"
	"fmt"

	"pensiel.com/material/src/client/postgresql"
	"pensiel.com/material/src/client/postgresql/migrate"
)

func main() {
	ctx := context.Background()
	c, _ := postgresql.NewClient()

	dbi := c.Cnx(ctx).(*postgresql.Connection).Conn

	if err := migrate.Promo(dbi).Executor(); err != nil {
		fmt.Println(err)
	}

}
