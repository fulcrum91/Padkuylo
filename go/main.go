package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Customer struct {
	ID_customer uint `gorm:"AUTO_INCREMENT; primary_key"`
	Name string `gorm:"type:varchar(50)"`
	Email string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(50)"`

}

type Order struct{
	ID_order uint `gorm:"AUTO_INCREMENT; primary_key"`
	OrderDate string
	Customer
	ID_customer uint
}

type Order_product struct {
	ID_order_product uint `gorm:"AUTO_INCREMENT; primary_key"`
	ID_Order Order
	ID_Product Product
}

type Product struct {
	ID_product   uint `gorm:"AUTO_INCREMENT; primary_key"`
	Name         string `gorm:"type:varchar(50)"`
	price        float32
	number       int
	ID_manufacturer uint
	//Manufacturer Manufacturer `gorm:"ForeignKey:ID_manufacturer"`
}
type Manufacturer struct {
	ID_manufacturer uint `gorm:"AUTO_INCREMENT; primary_key"`
	Name string `gorm:"type:varchar(50)"`
	products []Product
}


func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=web-shop sslmode=disable password=root")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&Manufacturer{})
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Order_product{})
	db.AutoMigrate(&Product{})
	var customer Customer
	var products Product
	var order = Order{OrderDate:"11.12.1996", ID_customer:1 }
	var manufacturer = Manufacturer{Name: "sd"}
	customer = Customer{  Name:"Sashka Padkuylo", Email:"wcg15@mail.ru", Password:"qwerty"}

	// Create
	db.Create(&customer)
	db.Create(&order)
	db.Create(&manufacturer)


	db.Model(&manufacturer).Related(&products)
	//db.Model(&Manufacturer{}).Related(&product)
	//db.Model(&Manufacturer{}).AddForeignKey("ID__manufacturer", "Product (ID_manufacturer)", "CASCADE", "CASCADE")
	//one to one
	db.Model(&customer).Related(&order)

	// Read
	//var product Product
	//db.First(&product, 1)                   // find product with id 1
	//db.First(&product, "code = ?", "L1212") // find product with code l1212
	//
	//// Update - update product's price to 2000
	//db.Model(&product).Update("Price", 2000)
}
//var tpl *template.Template
//
//func init()  {
//	tpl = template.Must(template.ParseGlob("templates/*"))
//
//}
//
//type person struct {
//	FirstName string
//	LastName string
//}
//
//func main()  {
//	http.HandleFunc("/", foo)
//	http.ListenAndServe(":8080", nil)
//}
//
//
//
//func foo(w http.ResponseWriter, req *http.Request)  {
//	f := req.FormValue("first")
//	l := req.FormValue("last")
//	err := tpl.ExecuteTemplate(w, "index.html", person{f, l})
//	if err != nil{
//		http.Error(w, err.Error(), 500)
//		log.Fatalln(err)
//	}
//}
