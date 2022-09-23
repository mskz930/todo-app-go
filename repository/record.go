package repository

import (
	"database/sql"
	"log"
	"todo-app/model/dto"
)

type RecordRepository interface {
	SaveExpense(dto.Expense) bool
	SaveIncome(dto.Income) bool
}

type recordRepository struct {
	db *sql.DB
}

func NewRecordRepository(db sql.DB) *recordRepository {
	return &recordRepository{db: &db}
}

func (rr *recordRepository) SaveExpense(ex dto.Expense) bool {
	sql := `insert into expenses(user_id, price, date, memo, category_id, payment_id, created_at, updated_at) 
	values (?, ?, ?, ?, ?, ?, now(), now());
	`
	_, err := rr.db.Exec(sql, ex.UserId, ex.Price, ex.Date, ex.Memo, ex.CategoryId, ex.PaymentId)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (rr *recordRepository) SaveIncome(inc dto.Income) bool {
	sql := `insert into incomes(user_id, price, date, memo, category_id, receipt_id, created_at, updated_at)
	values (?, ?, ?, ?, ?, ?, now(), now())`
	_, err := rr.db.Exec(sql, inc.UserId, inc.Price, inc.Date, inc.Memo, inc.CategoryId, inc.ReceiptId)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (rr *recordRepository) FindAllExpenses(userId int64) []dto.Expense {
	sql := `select price, date, memo, category_id, payment_id from expenses where user_id = ?`
	rows, err := rr.db.Query(sql, userId)
	if err != nil {
		log.Println(err)
		return nil
	}
	expenses := []dto.Expense{}
	for rows.Next() {
		var ex dto.Expense
		rows.Scan(ex.Price, ex.Date, ex.Memo, ex.CategoryId, ex.PaymentId)
		expenses = append(expenses, ex)
	}
	return expenses
}

func (rr *recordRepository) FindAllIncomes(userId int64) []dto.Income {
	sql := `select price, date, memo, category_id, payment_id from expenses where user_id = ?`
	rows, err := rr.db.Query(sql, userId)
	if err != nil {
		log.Println(err)
		return nil
	}
	incomes := []dto.Income{}
	for rows.Next() {
		var inc dto.Income
		rows.Scan(inc.Price, inc.Date, inc.Memo, inc.CategoryId, inc.ReceiptId)
		incomes = append(incomes, inc)
	}
	return incomes
}
