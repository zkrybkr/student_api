package main

import (
	c "studentApi/config"
	d "studentApi/database"
	r "studentApi/routes"
)

func main() {
	c.RunConfig()
	d.RunDB()
	r.RunService()
}
