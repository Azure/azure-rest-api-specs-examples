package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/MultipleModernInvoiceDownload.json
func ExampleInvoicesClient_BeginDownloadMultipleBillingProfileInvoices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInvoicesClient().BeginDownloadMultipleBillingProfileInvoices(ctx, "{billingAccountName}", []*string{
		to.Ptr("https://management.azure.com/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}/download?downloadToken={downloadToken}&useCache=True&api-version=2020-05-01"),
		to.Ptr("https://management.azure.com/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}/download?downloadToken={downloadToken}&useCache=True&api-version=2020-05-01"),
		to.Ptr("https://management.azure.com/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}/download?downloadToken={downloadToken}&useCache=True&api-version=2020-05-01")}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DownloadURL = armbilling.DownloadURL{
	// 	URL: to.Ptr("https://myaccount.blob.core.windows.net/invoices/1383724.pdf?sv=2019-02-02&sr=b&sp=r"),
	// }
}
