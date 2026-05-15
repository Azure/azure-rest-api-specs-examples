package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/invoicesDownloadByBillingAccount.json
func ExampleInvoicesClient_BeginDownloadByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInvoicesClient().BeginDownloadByBillingAccount(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "G123456789", &armbilling.InvoicesClientBeginDownloadByBillingAccountOptions{
		DocumentName: to.Ptr("12345678")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.InvoicesClientDownloadByBillingAccountResponse{
	// 	DocumentDownloadResult: armbilling.DocumentDownloadResult{
	// 		ExpiryTime: to.Ptr("2023-02-16T17:32:28Z"),
	// 		URL: to.Ptr("https://myaccount.blob.core.windows.net/invoices/1383724.pdf?sv=2019-02-02&sr=b&sp=r"),
	// 	},
	// }
}
