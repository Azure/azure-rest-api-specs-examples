package armnetworkfunction_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/527f6d35fb0d85c48210ca0f6f6f42814d63bd33/specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/AzureTrafficCollectorsByResourceGroupList.json
func ExampleAzureTrafficCollectorsByResourceGroupClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkfunction.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureTrafficCollectorsByResourceGroupClient().NewListPager("rg1", nil)
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
		// page.AzureTrafficCollectorListResult = armnetworkfunction.AzureTrafficCollectorListResult{
		// 	Value: []*armnetworkfunction.AzureTrafficCollector{
		// 		{
		// 			Name: to.Ptr("atc"),
		// 			Type: to.Ptr("Microsoft.NetworkFunction/azureTrafficCollectors"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.NetworkFunction/azureTrafficCollectors/atc"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetworkfunction.AzureTrafficCollectorPropertiesFormat{
		// 				CollectorPolicies: []*armnetworkfunction.ResourceReference{
		// 				},
		// 				ProvisioningState: to.Ptr(armnetworkfunction.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
