package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type student struct {
    Id int
    fullName string
}

func Add(db *sql.DB,id int, fullname string) {
    _, err := db.Exec(`insert into students(id, fullname) values (?, ?) `, id, fullname)
    if err != nil {
        log.Fatalf("couldn't add: %v", err)
    }
    fmt.Printf("added student with id= %d and name = %s\n", id, fullname)
}

func Delete(db *sql.DB, id int){
    result, err := db.Exec(`delete from students where id = ?`, id)
    if err != nil {
        log.Fatal(err)
    }
    i,err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    if i == 0 {
        fmt.Printf("no change to database\n")
        return
    }
}

func Update(db *sql.DB, id int, fullname string){
    _ , err := db.Exec(`update students set fullname = ? where id = ? `, fullname, id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("update the student with id=%d new name:%s\n", id, fullname)
}

func GetAll(db *sql.DB) {
    rows, err := db.Query(`select * from students`)
    if err != nil {
        log.Fatal(err)
    }
    var students []student
    for rows.Next() {
        var student student
        err := rows.Scan(&student.Id, &student.fullName)
        if err != nil {
            log.Fatal(err)
        }
        students = append(students, student)
    }

    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    for _,s := range students {
        fmt.Printf("id = %d, value = %v\n", s.Id, s.fullName)
    }

}

func Get(db *sql.DB, id int ) {
    var s student
    err := db.QueryRow(`select * from students where id = ?`, id).Scan(&s.Id,&s.fullName)
    if err != nil {
        fmt.Printf("couldnt get student with id %d: %v", id, err)
    }
    fmt.Printf("student with id = %d: %s\n", s.Id, s.fullName)
}

func Connect() (db *sql.DB,error error){
    db,err := sql.Open("sqlite3", "./data.db")
    if err != nil {
        return nil, fmt.Errorf("couldnt open connection to db:%v", err)
    }

    statement := `create table if not exists students(
    id int not null,
    fullname string not null)`

    _, err = db.Exec(statement)
    if err != nil {
        log.Fatalf("couldn't execute query:%v", err)
    }

    return db, nil
}
