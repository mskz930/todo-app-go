package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"todo-app/model/dto"
	"todo-app/repository"

	_ "github.com/go-sql-driver/mysql"
)

var recordRepository repository.RecordRepository

// var loginService service.LoginService
// var recordService service.RecordService

func main() {
	db, err := sql.Open("mysql", "root@(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}

	recordRepository = repository.NewRecordRepository(*db)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/record/new", newRecord)
	fmt.Println("listen and serve...")
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		response := &dto.Response[any]{Status: 200, Message: "Ok"}
		json.NewEncoder(w).Encode(response)
	}
}

// 新規レコードの登録
func newRecord(w http.ResponseWriter, r *http.Request) {
	var res *dto.Response[any]
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var record dto.Record
		err := json.NewDecoder(r.Body).Decode(&record)
		if err != nil {
			fmt.Println(err)
			return
		}
		var result bool
		if record.IsExpense {
			expense := dto.NewExpense(record)
			fmt.Printf("%v", expense)
			result = recordRepository.SaveExpense(*expense)
		} else {
			income := dto.NewIncome(record)
			result = recordRepository.SaveIncome(*income)
		}
		if result {
			res = &dto.Response[any]{Status: 200, Message: "ok"}
		} else {
			res = &dto.Response[any]{Status: 500, Message: "error"}
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	res = &dto.Response[any]{Status: 400, Message: "bad request"}
	json.NewEncoder(w).Encode(res)
	return
}
