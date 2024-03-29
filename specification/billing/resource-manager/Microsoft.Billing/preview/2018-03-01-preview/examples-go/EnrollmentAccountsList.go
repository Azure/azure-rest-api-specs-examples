package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/EnrollmentAccountsList.json
func ExampleEnrollmentAccountsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnrollmentAccountsClient().NewListPager(nil)
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
		// page.EnrollmentAccountListResult = armbilling.EnrollmentAccountListResult{
		// 	Value: []*armbilling.EnrollmentAccountSummary{
		// 		{
		// 			Name: to.Ptr("e1bf1c8c-5ac6-44a0-bdcd-aa7c1cf60556"),
		// 			Type: to.Ptr("Microsoft.Billing/enrollmentAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/enrollmentAccounts/e1bf1c8c-5ac6-44a0-bdcd-aa7c1cf60556"),
		// 			Properties: &armbilling.EnrollmentAccountSummaryProperties{
		// 				PrincipalName: to.Ptr("kathy@contoso.com"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("edd24053-07cd-4ed4-aa5b-326160a6680d"),
		// 			Type: to.Ptr("Microsoft.Billing/enrollmentAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/enrollmentAccounts/edd24053-07cd-4ed4-aa5b-326160a6680d"),
		// 			Properties: &armbilling.EnrollmentAccountSummaryProperties{
		// 				PrincipalName: to.Ptr("dan@contoso.com"),
		// 			},
		// 	}},
		// }
	}
}
