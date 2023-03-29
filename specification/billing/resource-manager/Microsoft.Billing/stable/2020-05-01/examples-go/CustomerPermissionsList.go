package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomerPermissionsList.json
func ExamplePermissionsClient_NewListByCustomerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPermissionsClient().NewListByCustomerPager("{billingAccountName}", "{customerName}", nil)
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
		// page.PermissionsListResult = armbilling.PermissionsListResult{
		// 	Value: []*armbilling.PermissionsProperties{
		// 		{
		// 			Actions: []*string{
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000001")},
		// 				NotActions: []*string{
		// 				},
		// 		}},
		// 	}
	}
}
