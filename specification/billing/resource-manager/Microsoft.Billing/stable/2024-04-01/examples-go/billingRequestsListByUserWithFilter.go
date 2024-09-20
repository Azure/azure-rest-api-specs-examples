package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRequestsListByUserWithFilter.json
func ExampleRequestsClient_NewListByUserPager_billingRequestsListByUserWithFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRequestsClient().NewListByUserPager(&armbilling.RequestsClientListByUserOptions{Filter: to.Ptr("properties/status eq 'Approved'"),
		OrderBy: nil,
		Top:     nil,
		Skip:    nil,
		Count:   nil,
		Search:  nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.RequestListResult = armbilling.RequestListResult{
		// 	Value: []*armbilling.Request{
		// 		{
		// 			Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 			Type: to.Ptr("Microsoft.Billing/billingRequests"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingRequests/11111111-1111-1111-1111-111111111111"),
		// 			Properties: &armbilling.RequestProperties{
		// 				Type: to.Ptr(armbilling.BillingRequestTypeProvisioningAccess),
		// 				BillingAccountDisplayName: to.Ptr("Contoso"),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				BillingAccountName: to.Ptr("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				CreatedBy: &armbilling.RequestPropertiesCreatedBy{
		// 					Upn: to.Ptr("foo@contoso.com"),
		// 				},
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T17:32:28.000Z"); return t}()),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-15T17:32:28.000Z"); return t}()),
		// 				LastUpdatedBy: &armbilling.RequestPropertiesLastUpdatedBy{
		// 					Upn: to.Ptr("admin@contosofinance.com"),
		// 				},
		// 				LastUpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-02T17:32:28.000Z"); return t}()),
		// 				RequestScope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/associatedTenants/100000000-0000-0000-0000-00000000"),
		// 				ReviewalDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-02T17:32:28.000Z"); return t}()),
		// 				ReviewedBy: &armbilling.RequestPropertiesReviewedBy{
		// 					Upn: to.Ptr("admin@contosofinance.com"),
		// 				},
		// 				Status: to.Ptr(armbilling.BillingRequestStatusApproved),
		// 			},
		// 	}},
		// }
	}
}
