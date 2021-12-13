package service

import (
	"context"
	"fmt"
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
	redisClient := db.GetRedisClient()
	data, err := redisClient.LRange("TRANSACTION", 0, -1).Result()
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, item := range data {
		tranType := ""
		if string(item[0]) == "+" {
			tranType = "credit"
		} else if string(item[0]) == "-" {
			tranType = "debit"
		}
		fVal, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

		stream.Send(&expensepb.Transaction{Amount: fVal, Type: tranType})
	}

	return nil
}

func (*ExpenseServer) GetBalance(ctx context.Context, req *expensepb.Empty) (*expensepb.User, error) {
	redisClient := db.GetRedisClient()

	// Get Amount
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

	// Get Amount again (in case a new value is initialized)
	val, _ = redisClient.Get("AMOUNT").Result()
	fVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return &expensepb.User{}, status.Errorf(codes.Internal, err.Error())
	}
	return &expensepb.User{Balance: fVal}, nil
}

func (*ExpenseServer) AddExpense(ctx context.Context, req *expensepb.Transaction) (*expensepb.Transaction, error) {
	redisClient := db.GetRedisClient()
	transactionType := req.Type
	transactionAmount := req.Amount

	// Get Amount
	val, err := redisClient.Get("AMOUNT").Result()
	switch {
	case err == redis.Nil:
		_, err := redisClient.Set("AMOUNT", 0, 0).Result()
		if err != nil {
			return &expensepb.Transaction{}, status.Errorf(codes.Internal, err.Error())
		}
	case err != nil:
		return &expensepb.Transaction{}, status.Errorf(codes.Internal, err.Error())
	case val == "":
		_, err := redisClient.Set("AMOUNT", 0, 0).Result()
		if err != nil {
			return &expensepb.Transaction{}, status.Errorf(codes.Internal, err.Error())
		}
	}

	// Get Amount again (in case a new value is initialized)
	val, _ = redisClient.Get("AMOUNT").Result()
	fVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return &expensepb.Transaction{}, status.Errorf(codes.Internal, err.Error())
	}

	// Do Expense Calculation
	if transactionType == "credit" {
		_, err := redisClient.LPush("TRANSACTION", "+"+fmt.Sprint(transactionAmount)).Result()
		if err != nil {
			return &expensepb.Transaction{}, status.Errorf(codes.Internal, err.Error())
		}

		latestAmount := fVal + transactionAmount

		_, erro := redisClient.Set("AMOUNT", latestAmount, 0).Result()
		if erro != nil {
			return &expensepb.Transaction{}, status.Errorf(codes.Internal, err.Error())
		}
	} else if transactionType == "debit" {
		_, err := redisClient.LPush("TRANSACTION", "-"+fmt.Sprint(transactionAmount)).Result()
		if err != nil {
			return &expensepb.Transaction{}, status.Errorf(codes.Internal, err.Error())
		}

		latestAmount := fVal - transactionAmount

		_, erro := redisClient.Set("AMOUNT", latestAmount, 0).Result()
		if erro != nil {
			return &expensepb.Transaction{}, status.Errorf(codes.Internal, err.Error())
		}
	} else {
		return &expensepb.Transaction{}, status.Errorf(codes.Internal, "Invalid transaction type")
	}

	return &expensepb.Transaction{Amount: transactionAmount, Type: transactionType}, nil
}
