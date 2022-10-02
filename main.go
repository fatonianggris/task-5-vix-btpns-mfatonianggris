package main

import (
	"rakaminbtpn/database"
	"rakaminbtpn/router"
)

func main() {
	database.Start_db()
	r := router.Start_app()
	r.Run(":8080")
}
