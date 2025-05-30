package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Usages_List.json
func ExampleUsagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListPager("westus", nil)
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
		// page.ListUsagesResult = armappcontainers.ListUsagesResult{
		// 	Value: []*armappcontainers.Usage{
		// 		{
		// 			Name: &armappcontainers.UsageName{
		// 				LocalizedValue: to.Ptr("ManagedEnvironmentCount"),
		// 				Value: to.Ptr("ManagedEnvironmentCount"),
		// 			},
		// 			CurrentValue: to.Ptr[float32](5),
		// 			Limit: to.Ptr[float32](10),
		// 			Unit: to.Ptr("Count"),
		// 		},
		// 		{
		// 			Name: &armappcontainers.UsageName{
		// 				LocalizedValue: to.Ptr("ManagedEnvironmentCores"),
		// 				Value: to.Ptr("ManagedEnvironmentCores"),
		// 			},
		// 			CurrentValue: to.Ptr[float32](3),
		// 			Limit: to.Ptr[float32](20),
		// 			Unit: to.Ptr("Count"),
		// 	}},
		// }
	}
}
