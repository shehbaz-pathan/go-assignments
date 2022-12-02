package main

import (
        "log"
        "time"
        "context"
        "regexp"
        "github.com/magiconair/properties"
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
)

func DBConnection() (*sql.DB, error){
	p:=properties.MustLoadFile("./config/dbconfig.properties",properties.UTF8)
	re:=regexp.MustCompile("[\"]")
	host:= re.ReplaceAllString(p.MustGetString("host"),"")
	dbname:= re.ReplaceAllString(p.MustGetString("dbname"),"")
	username:= re.ReplaceAllString(p.MustGetString("username"),"")
	password:= re.ReplaceAllString(p.MustGetString("password"),"")
	connString:=username+":"+password+"@tcp("+host+")/"+dbname
	db,err:=sql.Open("mysql",connString)
	if err!=nil {
		log.Printf("Error %s while connecting to DB",err)
		return nil,err
	}
	ctx, cancelfunc:= context.WithTimeout(context.Background(), 5*time.Second)
        defer cancelfunc()
	err = db.PingContext(ctx)
        if err != nil {
           log.Printf("Errors %s pinging DB", err)
           return nil,err
        }
	return db, nil

}

func (dbh *DBHandler)AddBook(b Book) error {
	query:= "insert into book_info(Isbn,Title,Synopsis,AuthorName,Price) values(?,?,?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancelfunc()
	stmt, err:= dbh.db.PrepareContext(ctx,query)
	if err!= nil {
		log.Printf("Error %s while preparing sql statement",err)
		return err
	}
	defer stmt.Close()
	res,err:=stmt.ExecContext(ctx,b.Isbn,b.Title,b.Synopsis,b.AuthorName,b.Price)
	if err!=nil {
	        log.Printf("Error %s while inserting data into book_info table",err)
                return err
	}
	_,err=res.RowsAffected()
	if err!=nil {
		log.Printf("Error %s while checking for rows affected")
                return err
	}
	return nil
}
func (dbh *DBHandler)GetBooks() ([]Book,error) {
	query:= "select * from book_info"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancelfunc()
	stmt,err:= dbh.db.PrepareContext(ctx,query)
	if err!=nil {
		log.Printf("%s while preparing query",err)
		return []Book{},err
	}
	rows,err:=stmt.QueryContext(ctx)
	if err!=nil {
		log.Printf("%s while reading data from table",err)
		return []Book{},err
	}
	defer rows.Close()
	var books = []Book{}
	for rows.Next() {
		var book Book
		if err = rows.Scan(&book.Isbn,&book.Title,&book.Synopsis,&book.AuthorName,&book.Price);err!=nil {
			log.Printf("%s while reading rows",err)
			return []Book{},err
		}
		books=append(books,book)
	}
	if err = rows.Err();err!=nil {
		log.Printf("%s while reading data from table",err)
		return []Book{},err
	}
	return books,nil

}

func (dbh *DBHandler)GetBookByTitle(s string) ([]Book,error) {
	query:= "select * from book_info where Title = ?"
	ctx,cancelfunc:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancelfunc()
	stmt,err:= dbh.db.PrepareContext(ctx,query)
	if err!=nil {
		log.Printf("%s while preparing query",err)
		return []Book{},nil
	}
	rows,err:= stmt.QueryContext(ctx,s)
	if err!=nil {
		log.Printf("%s while reading data",err)
		return []Book{},err
	}
	var books = []Book{}
	for rows.Next() {
		var book = Book{}
		if err = rows.Scan(&book.Isbn,&book.Title,&book.Synopsis,&book.AuthorName,&book.Price);err!=nil {
                        log.Printf("%s while reading rows",err)
                        return []Book{},err
                }
		books=append(books,book)
	}
	if err = rows.Err();err!=nil {
                log.Printf("%s while reading data from table",err)
                return []Book{},err
        }
        return books,nil
}
