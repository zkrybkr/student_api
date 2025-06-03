package main

import(
	c "studentApi/config"
	d "studentApi/database"
)

func main() {
	c.RunConfig()
	d.RunDB()
}