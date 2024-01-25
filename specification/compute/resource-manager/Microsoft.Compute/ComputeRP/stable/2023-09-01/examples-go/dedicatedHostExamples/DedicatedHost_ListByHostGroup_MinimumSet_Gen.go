package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/dedicatedHostExamples/DedicatedHost_ListByHostGroup_MinimumSet_Gen.json
func ExampleDedicatedHostsClient_NewListByHostGroupPager_dedicatedHostListByHostGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostsClient().NewListByHostGroupPager("rgcompute", "aaaa", nil)
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
		// page.DedicatedHostListResult = armcompute.DedicatedHostListResult{
		// 	Value: []*armcompute.DedicatedHost{
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myHost"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armcompute.SKU{
		// 			},
		// 	}},
		// }
	}
}
