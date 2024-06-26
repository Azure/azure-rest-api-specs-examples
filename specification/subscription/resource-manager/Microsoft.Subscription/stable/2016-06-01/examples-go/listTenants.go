package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/subscription/resource-manager/Microsoft.Subscription/stable/2016-06-01/examples/listTenants.json
func ExampleTenantsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTenantsClient().NewListPager(nil)
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
		// page.TenantListResult = armsubscription.TenantListResult{
		// 	Value: []*armsubscription.TenantIDDescription{
		// 		{
		// 			ID: to.Ptr("/tenants/72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 		},
		// 		{
		// 			ID: to.Ptr("/tenants/33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
		// 			TenantID: to.Ptr("33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
		// 	}},
		// }
	}
}
