package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type employee_info struct {
	id                 int64
	name, position     string
	experience, salary float64
}

func CreateTable() {
	db, err := sql.Open("mysql", "root:mukheshM1@25-03@/")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec("USE test")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("database selected successfully")
	}
	table, err := db.Prepare("CREATE Table employee(emp_id int NOT NULL AUTO_INCRIMENT,emp_name varchar(50),emp_position varchar(50),emp_experience float,emp_salary float,PRIMARY KEY (emp_id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = table.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("table created successfully")
	}

}

func PostData(emp employee_info) (employee_info, error) {
	//CreateTable()
	result := emp
	db, err := sql.Open("mysql", "root:mukheshM1@25-03@/test")
	if err != nil {
		return result, errors.New("cannot open database")
	}
	if emp.id <= 0 {
		return result, errors.New("invalid id")
	}

	rs, err := db.Exec("INSERT INTO employee (emp_id, emp_name, emp_position, emp_exp, emp_salary) VALUES (?,? ,?, ?, ?)", emp.id, emp.name, emp.position, emp.experience, emp.salary)

	if err != nil {
		return result, errors.New("key already present")
	}
	emp.id, err = rs.LastInsertId()
	if err != nil {
		fmt.Errorf("Error 2: %v", err)
	}
	return emp, nil

}

func GetData(id int64) (employee_info, error) {
	emp := employee_info{}

	db, err := sql.Open("mysql", "root:mukheshM1@25-03@/test")

	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	row := db.QueryRow("SELECT * FROM employee WHERE emp_id = ?", id)

	if err := row.Scan(&emp.id, &emp.name, &emp.position, &emp.experience, &emp.salary); err != nil {
		if err == sql.ErrNoRows {
			return emp, fmt.Errorf("GetId %d: NO id found", id)
		}
		return emp, fmt.Errorf("GetId %d: %v", id, err)
	}
	return emp, nil

}
