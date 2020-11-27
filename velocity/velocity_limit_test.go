package velocity

import (
	"testing"
	"time"
)

func TestLimits_Allowed(t *testing.T) {
	type fields struct {
		userTransactions       map[UniqueTransactionKey]bool
		userDailyTransactions  map[DailyTransactionKey]*BalanceAndCount
		userWeeklyTransactions map[WeeklyTransactionKey]int
	}
	type args struct {
		funds []Funds
	}
	fixedDate := time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "testBasic",
			fields: fields{
				userTransactions:       make(map[UniqueTransactionKey]bool),
				userDailyTransactions:  make(map[DailyTransactionKey]*BalanceAndCount),
				userWeeklyTransactions: make(map[WeeklyTransactionKey]int),
			},
			args: args{
				funds: []Funds{
					{Id: 33, CustomerId: 44, Dollar: Dollar{132512}, Timestamp: time.Now()},
				},
			},
			want: []int{Accept},
		}, {
			name: "testReject4thLoadFund",
			fields: fields{
				userTransactions:       make(map[UniqueTransactionKey]bool),
				userDailyTransactions:  make(map[DailyTransactionKey]*BalanceAndCount),
				userWeeklyTransactions: make(map[WeeklyTransactionKey]int),
			},
			args: args{
				funds: []Funds{
					{Id: 33, CustomerId: 44, Dollar: Dollar{1000}, Timestamp: fixedDate},
					{Id: 34, CustomerId: 44, Dollar: Dollar{1000}, Timestamp: fixedDate},
					{Id: 35, CustomerId: 44, Dollar: Dollar{1000}, Timestamp: fixedDate},
					{Id: 36, CustomerId: 44, Dollar: Dollar{2000}, Timestamp: fixedDate},
				},
			},
			want: []int{Accept, Accept, Accept, Reject},
		},
		//{name: "testSingleFundGreaterThanDailyLimit"},
		//{name: "testMultipleFundAddUpToGreaterThanDailyLimit"},
		//{name: "testSingleFundGreaterThanWeeklyLimit"},
		//{name: "testMultipleFundAddUpToGreaterThanWeeklyLimit"},
		//{name: "testIgnoreDuplicatedIdAndCustomerId"},
		//{name: "testDayEndsAtUtcMidnight"},
		//{name: "testDayEarlyBoundaryIsAtUtcMidnightExclusive"},
		//{name: "testDayLaterBoundaryIsAtUtcMidnightInclusive"},
		//{name: "testWeekStartsOnMondayMorningUtcZeroInclusive"},
		//{name: "testWeekEndsOnSundayUtcMidnightMinusOneSecond"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Limits{
				userTransactions:       tt.fields.userTransactions,
				userDailyTransactions:  tt.fields.userDailyTransactions,
				userWeeklyTransactions: tt.fields.userWeeklyTransactions,
			}
			for i, fund := range tt.args.funds {
				if got := l.Allowed(fund); got != tt.want[i] {
					t.Errorf("Allowed() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
