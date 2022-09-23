package dto

import (
	"log"
	"time"
)

type Response[T interface{}] struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type Record struct {
	Id         int64  `json:"id,omitempty"`
	UserId     int64  `json:"userId,omitempty"`
	Price      int64  `json:"price"`
	Memo       string `json:"memo"`
	Date       string `json:"date"`
	CategoryId int64  `json:"categoryId"`
	PaymentId  int64  `json:"paymentId,omitepmty"`
	ReceiptId  int64  `json:"receiptId,omitepmty"`
	IsExpense  bool   `json:"isExpense"`
}

type Expense struct {
	Id         int64     `json:"id,omitempty"`
	UserId     int64     `json:"userId"`
	Price      int64     `json:"price"`
	Memo       string    `json:"memo"`
	Date       time.Time `json:"date"`
	CategoryId int64     `json:"categoryId"`
	PaymentId  int64     `json:"paymentId"`
}

func NewExpense(r Record) *Expense {
	date, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		log.Println(err)
		return nil
	}
	if !r.IsExpense || r.ReceiptId != 0 {
		log.Printf("atrribute error: IsExpense: %v, r.ReceiptId: %v\n", r.IsExpense, r.ReceiptId)
		return nil
	}
	return &Expense{
		Id:         r.Id,
		UserId:     r.UserId,
		Price:      r.Price,
		Memo:       r.Memo,
		Date:       date,
		CategoryId: r.CategoryId,
		PaymentId:  r.PaymentId,
	}
}

func (ex *Expense) NewRecord() *Record {
	return &Record{
		Id:         ex.Id,
		UserId:     ex.UserId,
		Price:      ex.Price,
		Memo:       ex.Memo,
		Date:       ex.Date.String(),
		CategoryId: ex.CategoryId,
		PaymentId:  ex.PaymentId,
		IsExpense:  true,
	}
}

type Income struct {
	Id         int64     `json:"id,omitempty"`
	UserId     int64     `json:"userId"`
	Price      int64     `json:"price"`
	Memo       string    `json:"memo"`
	Date       time.Time `json:"date"`
	CategoryId int64     `json:"categoryId"`
	ReceiptId  int64     `json:"receiptId"`
}

func NewIncome(r Record) *Income {
	date, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		log.Println(err)
		return nil
	}
	if r.IsExpense || (r.PaymentId != 0) {
		log.Printf("validation error: IsExpense: %v, r.ReceiptId: %v\n", r.IsExpense, r.ReceiptId)
		return nil
	}
	return &Income{
		Id:         r.Id,
		UserId:     r.UserId,
		Price:      r.Price,
		Memo:       r.Memo,
		Date:       date,
		CategoryId: r.CategoryId,
		ReceiptId:  r.ReceiptId,
	}
}

func (in *Income) NewRecord() *Record {
	return &Record{
		Id:         in.Id,
		UserId:     in.UserId,
		Price:      in.Price,
		Memo:       in.Memo,
		Date:       in.Date.String(),
		CategoryId: in.CategoryId,
		ReceiptId:  in.ReceiptId,
		IsExpense:  false,
	}
}
