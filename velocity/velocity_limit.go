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

func (l Limits) Allowed(funds Funds) int {
	if _, present := l.userTransactions[KeyOf(funds.Id, funds.CustomerId)]; present {
		return Ignore
	}
	l.userTransactions[KeyOf(funds.Id, funds.CustomerId)] = true

	isAllowedByDailyLimit := l.allowedByDailyLimit(funds)
	isAllowedByWeeklyLimit := l.allowedByWeeklyLimit(funds)

	if isAllowedByDailyLimit && isAllowedByWeeklyLimit {
		l.update(funds)
		return Accept
	} else {
		return Reject
	}
}

func (l Limits) allowedByDailyLimit(funds Funds) bool {
	if funds.Dollar.Amount > DailyFundLimit {
		return false
	}

	startOfDay := ToStartOfDay(funds.Timestamp)
	if balance, present := l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)]; present {
		if balance.balance+funds.Dollar.Amount > DailyFundLimit || balance.count == DailyDistinctLimit {
			return false
		}
	}
	return true
}

func (l Limits) allowedByWeeklyLimit(funds Funds) bool {
	if funds.Dollar.Amount > WeeklyFundLimit {
		return false
	}

	week := WeekKeyOf(funds.CustomerId, ToStartOfWeek(funds.Timestamp))
	if balance, present := l.userWeeklyTransactions[week]; present {
		if balance+funds.Dollar.Amount > WeeklyFundLimit {
			return false
		}
	}
	return true
}

func (l Limits) update(funds Funds) {
	startOfDay := ToStartOfDay(funds.Timestamp)
	if balance, present := l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)]; present {
		l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)] = balance.addBalance(funds.Dollar.Amount)
	} else {
		l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)] = &BalanceAndCount{funds.Dollar.Amount, 1}
	}

	week := WeekKeyOf(funds.CustomerId, ToStartOfWeek(funds.Timestamp))
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
