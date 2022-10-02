package main

func main() {
	database.start_db()
	r := router.start_app()
	r.Run(":8080")
}
