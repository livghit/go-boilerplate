package routes

import (
 db "github.com/livghit/go-boilerplate/database"
	"fmt"
)

type Routes struct {
	path string
}

func LoadRoutes() {
  databaseConstant := db.Three; 
	fmt.Printf("Hello %v from Routes folder", databaseConstant);
}
