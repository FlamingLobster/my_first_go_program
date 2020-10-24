package main

import (
	"bufio"
	"koho/velocity"
	"os"
	"strconv"
	"testing"
)

/**
The provided input file does not cover this requirement:
- A maximum of 3 loads can be performed per day, regardless of amount

only 4 patients have more than 3 loads a day and the above requirement is always hidden by this requirement:
- A maximum of $5,000 can be loaded per day

you can confirm this by checking out the commit that this comment was check in on and run the acceptance_test
On this commit, there is no logic checking for the 3 load per day limit. But there will be no failures
and everything will pass.

I've not modified the provided input.txt nor made a new file. Such a file should be easy to construct. And a
test in limit_test covers this scenario

In sequence:
Customer ID: 732
{"id":"6091","customer_id":"732","load_amount":"$4665.30","time":"2000-02-09T03:55:58Z"}
{"id":"25064","customer_id":"732","load_amount":"$475.00","time":"2000-02-09T11:05:32Z"}
{"id":"15357","customer_id":"732","load_amount":"$5198.50","time":"2000-02-09T13:08:16Z"}
{"id":"6947","customer_id":"732","load_amount":"$3838.06","time":"2000-02-09T15:11:00Z"}
-------------------
Customer ID: 392
{"id":"27017","customer_id":"392","load_amount":"$1697.59","time":"2000-01-07T03:16:48Z"}
{"id":"10894","customer_id":"392","load_amount":"$5039.22","time":"2000-01-07T06:20:54Z"}
{"id":"7650","customer_id":"392","load_amount":"$5009.09","time":"2000-01-07T09:25:00Z"}
{"id":"21596","customer_id":"392","load_amount":"$5194.66","time":"2000-01-07T17:35:56Z"}
-------------------
Customer ID: 426
{"id":"12377","customer_id":"426","load_amount":"$3412.62","time":"2000-01-07T02:15:26Z"}
{"id":"12110","customer_id":"426","load_amount":"$2868.72","time":"2000-01-07T18:37:18Z"}
{"id":"29446","customer_id":"426","load_amount":"$99.21","time":"2000-01-07T20:40:02Z"}
{"id":"3026","customer_id":"426","load_amount":"$4692.19","time":"2000-01-07T23:44:08Z"}
-------------------
Customer ID: 222
{"id":"1045","customer_id":"222","load_amount":"$3005.41","time":"2000-01-08T03:49:36Z"}
{"id":"18516","customer_id":"222","load_amount":"$2833.22","time":"2000-01-08T09:57:48Z"}
{"id":"8217","customer_id":"222","load_amount":"$4107.58","time":"2000-01-08T12:00:32Z"}
{"id":"5401","customer_id":"222","load_amount":"$3073.28","time":"2000-01-08T23:15:34Z"}
*/
func TestProvidedInputFile(t *testing.T) {
	userTransactions := make(map[velocity.DailyTransactionKey][]string)

	if inputFile, err := os.Open("input.txt"); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			line := scanner.Text()

			loadFund, err := unmarhshalFunds(line)
			if err != nil {
				t.Error("Could not unmarshal input json")
				continue
			}
			startOfDay := velocity.ToStartOfDay(loadFund.Timestamp)
			key := velocity.TimeKeyOf(loadFund.CustomerId, startOfDay)
			if transactions, present := userTransactions[key]; present {
				transactions = append(transactions, line)
				userTransactions[key] = transactions
			} else {
				transactions = make([]string, 1)
				transactions[0] = line
				userTransactions[key] = transactions
			}
		}
	}

	for key, transactions := range userTransactions {
		if len(transactions) > 3 {
			println("-------------------")
			println("Customer ID: " + strconv.Itoa(key.CustomerId))
			for _, transaction := range transactions {
				println(transaction)
			}
		}
	}
}
