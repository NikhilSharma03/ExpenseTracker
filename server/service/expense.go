package service

import (
	"context"
	"strconv"

	"github.com/NikhilSharma03/expensetracker/server/db"
	"github.com/NikhilSharma03/expensetracker/server/expensepb"
	"github.com/go-redis/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	redisClient := db.GetRedisClient()
	val, err := redisClient.Get("AMOUNT").Result()
	switch {
	case err == redis.Nil:
		_, err := redisClient.Set("AMOUNT", 0, 0).Result()
		if err != nil {
			return &expensepb.User{}, status.Errorf(codes.Internal, err.Error())
		}
	case err != nil:
		return &expensepb.User{}, status.Errorf(codes.Internal, err.Error())
	case val == "":
		_, err := redisClient.Set("AMOUNT", 0, 0).Result()
		if err != nil {
			return &expensepb.User{}, status.Errorf(codes.Internal, err.Error())
		}
	}
	val, _ = redisClient.Get("AMOUNT").Result()
	fVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return &expensepb.User{}, status.Errorf(codes.Internal, err.Error())
	}
	return &expensepb.User{Balance: fVal}, nil
}

func (*ExpenseServer) AddExpense(ctx context.Context, req *expensepb.Transaction) (*expensepb.Transaction, error) {
	transactionType := req.Type
	transactionAmount := req.Amount

	return &expensepb.Transaction{Amount: transactionAmount, Type: transactionType}, nil
}
