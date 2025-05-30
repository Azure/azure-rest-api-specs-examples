package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ManagedEnvironmentUsages_List.json
func ExampleManagedEnvironmentUsagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedEnvironmentUsagesClient().NewListPager("examplerg", "jlaw-demo1", nil)
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
		// 				LocalizedValue: to.Ptr("Managed Environment Consumption Cores"),
		// 				Value: to.Ptr("ManagedEnvironmentConsumptionCores"),
		// 			},
		// 			CurrentValue: to.Ptr[float32](0.5),
		// 			Limit: to.Ptr[float32](10),
		// 			Unit: to.Ptr("Count"),
		// 	}},
		// }
	}
}
