package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/transactionSummaryGetByInvoice.json
func ExampleTransactionsClient_GetTransactionSummaryByInvoice() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTransactionsClient().GetTransactionSummaryByInvoice(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "G123456789", &armbilling.TransactionsClientGetTransactionSummaryByInvoiceOptions{Filter: nil,
		Search: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TransactionSummary = armbilling.TransactionSummary{
	// 	AzureCreditApplied: to.Ptr[float32](100),
	// 	BillingCurrency: to.Ptr("USD"),
	// 	ConsumptionCommitmentDecremented: to.Ptr[float32](1000),
	// 	SubTotal: to.Ptr[float32](1000),
	// 	Tax: to.Ptr[float32](500),
	// 	Total: to.Ptr[float32](5400),
	// }
}
