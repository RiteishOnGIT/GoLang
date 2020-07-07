package main


import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type book struct{
	isbn string
	title string
	author string
	price float32
}
func main(){
	db,err:= sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic (err)
	}
	fmt.Println("You are connected to your database")

	rows, err := db.Query("SELECT  * FROM books")
	if err != nil{
		panic(err)
	}
	defer rows.Close()
	bks := make([]book, 0)
	for rows.Next(){
		bk := book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil{
			panic(err)
		}
		bks = append(bks, bk)
	}
	for _, bk := range bks {
		// fmt.Println(bk.isbn, bk.title, bk.author, bk.price)
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}

}































