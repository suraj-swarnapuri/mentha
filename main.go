package main

import (
	"context"
	"fmt"
	"os"

	"github.com/plaid/plaid-go/plaid"
)

func main() {

	configuration := plaid.NewConfiguration()
	clientId := os.Getenv("PLAID_CLIENT_ID")
	secret := os.Getenv("PLAID_SECRET")
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", clientId)
	configuration.AddDefaultHeader("PLAID-SECRET", secret)
	configuration.UseEnvironment(plaid.Sandbox)
	client := plaid.NewAPIClient(configuration)
	ctx := context.Background()
	accessToken := "access-sandbox-1830008c-6f4c-4080-b3ce-6642250b9109"
	// Get Accounts
	accountsGetResp, _, err := client.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*plaid.NewAccountsGetRequest(accessToken),
	).Execute()
	if err != nil {
		fmt.Println(err)
	}
	accountResp := accountsGetResp
	for _, x := range accountResp.Accounts {
		fmt.Println(x.Name)
	}

	transactionGetResp, _, err := client.PlaidApi.TransactionsGet(ctx).TransactionsGetRequest(
		*plaid.NewTransactionsGetRequest(
			accessToken,
			"2022-01-01",
			"2022-12-31"),
	).Execute()

	for _, x := range transactionGetResp.GetTransactions() {
		fmt.Println(x.GetAmount())
		break
	}
}
