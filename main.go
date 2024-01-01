package main

import(
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/routers"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/database"
)

func main(){
	database.InitDB()
	database.MigrateDB()
	r := routers.InitRouter()
	r.Run(":8080")
}