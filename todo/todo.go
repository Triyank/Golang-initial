package main

fmt (
	"fmt"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBconfig struct{
	Host string
	Port int
	User string
	DBname string
	Password string
}

func BuildDbConfig() *BuildDbConfig{
	dbConfig:= DBconfig{
		Host: "0.0.0.0",
		Port: 3306,
		User: "Triyank",
		DBname: "todos",
		Password: "mypassword",
	}
	return &dbConfig
}

func DbUrl( dbConfig *DBconfig) string {
	return fmt.Sprint(
		"%S:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBname,
	)

}
func main (){
	fmt.Println("ok")
}