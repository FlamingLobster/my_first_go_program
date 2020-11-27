package velocity

const (
	DailyDistinctLimit int = 3
	DailyFundLimit     int = 500000
	WeeklyFundLimit    int = 2000000
)

type Limits struct {
	userTransactions       map[UniqueTransactionKey]bool
	userDailyTransactions  map[DailyTransactionKey]*BalanceAndCount
	userWeeklyTransactions map[WeeklyTransactionKey]int
}

type BalanceAndCount struct {
	balance int
	count   int
}

func (b *BalanceAndCount) addBalance(amount int) *BalanceAndCount {
	b.balance += amount
	b.increment()
	return b
}

func (b *BalanceAndCount) increment() {
	b.count = b.count + 1
}

func (l *Limits) Allowed(funds Funds) int {
	if _, present := l.userTransactions[TransactionKey(funds.Id, funds.CustomerId)]; present {
		return Ignore
	}
	l.userTransactions[TransactionKey(funds.Id, funds.CustomerId)] = true

	isAllowedByDailyLimit := l.allowedByDailyLimit(funds)
	isAllowedByWeeklyLimit := l.allowedByWeeklyLimit(funds)

	if isAllowedByDailyLimit && isAllowedByWeeklyLimit {
		l.update(funds)
		return Accept
	} else {
		return Reject
	}
}

func (l *Limits) allowedByDailyLimit(funds Funds) bool {
	if funds.Dollar.Amount > DailyFundLimit {
		return false
	}

	startOfDay := ToStartOfDay(funds.Timestamp)
	if balanceAndCount, present := l.userDailyTransactions[DailyKey(funds.CustomerId, startOfDay)]; present {
		if balanceAndCount.balance+funds.Dollar.Amount > DailyFundLimit ||
			balanceAndCount.count == DailyDistinctLimit {
			return false
		}
	}
	return true
}

func (l *Limits) allowedByWeeklyLimit(funds Funds) bool {
	if funds.Dollar.Amount > WeeklyFundLimit {
		return false
	}

	week := WeeklyKey(funds.CustomerId, ToStartOfWeek(funds.Timestamp))
	if balance, present := l.userWeeklyTransactions[week]; present {
		if balance+funds.Dollar.Amount > WeeklyFundLimit {
			return false
		}
	}
	return true
}

func (l *Limits) update(funds Funds) {
	startOfDay := ToStartOfDay(funds.Timestamp)
	if balance, present := l.userDailyTransactions[DailyKey(funds.CustomerId, startOfDay)]; present {
		l.userDailyTransactions[DailyKey(funds.CustomerId, startOfDay)] = balance.addBalance(funds.Dollar.Amount)
	} else {
		l.userDailyTransactions[DailyKey(funds.CustomerId, startOfDay)] = &BalanceAndCount{funds.Dollar.Amount, 1}
	}

	week := WeeklyKey(funds.CustomerId, ToStartOfWeek(funds.Timestamp))
	if balance, present := l.userWeeklyTransactions[week]; present {
		l.userWeeklyTransactions[week] = balance + funds.Dollar.Amount
	} else {
		l.userWeeklyTransactions[week] = funds.Dollar.Amount
	}
}

func NewLimit() *Limits {
	limits := Limits{
		userTransactions:       make(map[UniqueTransactionKey]bool),
		userDailyTransactions:  make(map[DailyTransactionKey]*BalanceAndCount),
		userWeeklyTransactions: make(map[WeeklyTransactionKey]int),
	}
	return &limits
}
