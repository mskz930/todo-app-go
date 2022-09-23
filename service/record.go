package service

// サービス用パッケージ

import (
	"database/sql"
	"todo-app/model/dto"
	"todo-app/repository"
)

type RecordService interface {
	NewExpense(record *dto.Expense)
}

type recordService struct {
	rr repository.RecordRepository
}

func NewRecordService(db sql.DB) *recordService {
	recordRepository := repository.NewRecordRepository(db)
	return &recordService{rr: recordRepository}
}

// 支出レコードの保存
func (rs *recordService) CreateExpense(expense dto.Expense) bool {
	return rs.rr.SaveExpense(expense)
}

// 収入レコードの保存
func (rs *recordService) CreateIncome(income dto.Income) bool {
	return rs.rr.SaveIncome(income)
}
