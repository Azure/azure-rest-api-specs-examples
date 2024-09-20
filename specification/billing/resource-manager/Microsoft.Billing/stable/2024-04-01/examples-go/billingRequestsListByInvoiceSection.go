package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRequestsListByInvoiceSection.json
func ExampleRequestsClient_NewListByInvoiceSectionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRequestsClient().NewListByInvoiceSectionPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "yyyy-yyyy-yyy-yyy", &armbilling.RequestsClientListByInvoiceSectionOptions{Filter: nil,
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
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingRequests"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingRequests/00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.RequestProperties{
		// 				Type: to.Ptr(armbilling.BillingRequestTypeRoleAssignment),
		// 				AdditionalInformation: map[string]*string{
		// 					"RoleId": to.Ptr("30000000-aaaa-bbbb-cccc-200000000003"),
		// 				},
		// 				BillingAccountDisplayName: to.Ptr("Contoso"),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				BillingAccountName: to.Ptr("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 				BillingProfileName: to.Ptr("xxxx-xxxx-xxx-xxx"),
		// 				CreatedBy: &armbilling.RequestPropertiesCreatedBy{
		// 					Upn: to.Ptr("foo@contoso.com"),
		// 				},
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-01T17:32:28.000Z"); return t}()),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-15T17:32:28.000Z"); return t}()),
		// 				InvoiceSectionDisplayName: to.Ptr("Department 1"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
		// 				InvoiceSectionName: to.Ptr("yyyy-yyyy-yyy-yyy"),
		// 				LastUpdatedBy: &armbilling.RequestPropertiesLastUpdatedBy{
		// 					Upn: to.Ptr("foo@contoso.com"),
		// 				},
		// 				LastUpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-01T17:32:28.000Z"); return t}()),
		// 				RequestScope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
		// 				Status: to.Ptr(armbilling.BillingRequestStatusPending),
		// 			},
		// 	}},
		// }
	}
}
