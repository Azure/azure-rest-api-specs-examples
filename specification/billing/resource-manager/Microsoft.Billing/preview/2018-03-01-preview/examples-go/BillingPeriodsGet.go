package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/BillingPeriodsGet.json
func ExamplePeriodsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPeriodsClient().Get(ctx, "201702-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Period = armbilling.Period{
	// 	Name: to.Ptr("201702-1"),
	// 	Type: to.Ptr("Microsoft.Billing/billingPeriods"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Billing/billingPeriods/201702-1"),
	// 	Properties: &armbilling.PeriodProperties{
	// 		BillingPeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2017-02-01"); return t}()),
	// 		BillingPeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2017-01-01"); return t}()),
	// 		InvoiceIDs: []*string{
	// 			to.Ptr("/subscriptions/subid/providers/Microsoft.Billing/invoices/2017-02-05-123456789"),
	// 			to.Ptr("/subscriptions/subid/providers/Microsoft.Billing/invoices/2017-01-05-987654321")},
	// 		},
	// 	}
}
