package service

import (
	"context"

	"github.com/NikhilSharma03/expensetracker/server/expensepb"
)

type ExpenseServer struct {
	expensepb.UnimplementedExpenseServiceServer
}

func (*ExpenseServer) GetExpenseHistory(req *expensepb.Empty, stream expensepb.ExpenseService_GetExpenseHistoryServer) error {
	data := []*expensepb.Transaction{{Type: "credit", Amount: 1000}, {Type: "credit", Amount: 2000}, {Type: "debit", Amount: 3000}, {Type: "debit", Amount: 3400.50}, {Type: "credit", Amount: 1000}}

	for _, item := range data {
		stream.Send(item)
	}

	return nil
}

func (*ExpenseServer) GetBalance(ctx context.Context, req *expensepb.Empty) (*expensepb.User, error) {
	return &expensepb.User{Balance: 100000.212}, nil
}

func (*ExpenseServer) AddExpense(ctx context.Context, req *expensepb.Transaction) (*expensepb.Transaction, error) {
	transactionType := req.Type
	transactionAmount := req.Amount

	return &expensepb.Transaction{Amount: transactionAmount, Type: transactionType}, nil
}
