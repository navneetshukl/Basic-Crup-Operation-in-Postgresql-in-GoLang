package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	//connect to a database

	conn, err := sql.Open("pgx","host= port= dbname= user= password=")
	if err!=nil{
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n",err))
	}

	defer conn.Close()

	log.Println("Connected to database")

	//test my connection

	err=conn.Ping()
	if err!=nil{
		log.Fatal("Cannot Ping the database")
	}
	log.Println("pinged database")

	//get rows from table

	err=getAllRows(conn)

	if err!=nil{
		log.Fatal(err)
	}

	//insert a row

	query:=`insert into users(first_name,last_name) values($1,$2)`

	_,err =conn.Exec(query,"John","Doe")

	if err!=nil{
		log.Fatal(err)
	}

	log.Println("Inserted a row")

	//get rows from table again

	err=getAllRows(conn)

	if err!=nil{
		log.Fatal(err)
	}

	//update a row

	stmt:=`update users set first_name=$1 where id=$2`
	_,err=conn.Exec(stmt,"Jackie",5)

	if err!=nil{
		log.Fatal(err)
	}

	log.Println("Updated one or more row")


	//get row from table again

	err=getAllRows(conn)

	if err!=nil{
		log.Fatal(err)
	}

	//get one row by id

	query=`select id,first_name,last_name from users where id=$1`
	
	var firstname,lastname string
	var id int
	
	row:=conn.QueryRow(query, 1)
	err=row.Scan(&id,&firstname,&lastname)
	if err!=nil{
		log.Fatal(err)
	}

	log.Println("QueryRow returns",id,firstname,lastname)

	//deleta a row

	query=`delete from users where id=$1`
	_,err=conn.Exec(query,6)
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("Deleted a row")

	//get rows again

	err=getAllRows(conn)

	if err!=nil{
		log.Fatal(err)
	}
}


func getAllRows(conn *sql.DB) error {

	rows,err:=conn.Query("select id,first_name,last_name from users")
	if err!=nil{
		log.Println(err)
		return err
	}
	defer rows.Close()
	var firstname,lastname string
	var id int
	for rows.Next(){
		err:=rows.Scan(&id,&firstname,&lastname)
		if err!=nil{
			log.Println(err)
			return err
		}
		fmt.Println("Record is",id,firstname,lastname)
	}
	if err=rows.Err(); err!=nil{
		log.Fatal("Error Scanning rows",err)
	}

	fmt.Println("----------------------------")

	return nil
}