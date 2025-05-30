package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/AvailableDelegationsResourceGroupGet.json
func ExampleAvailableResourceGroupDelegationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailableResourceGroupDelegationsClient().NewListPager("westcentralus", "rg1", nil)
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
		// page.AvailableDelegationsResult = armnetwork.AvailableDelegationsResult{
		// 	Value: []*armnetwork.AvailableDelegation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Provider.resourceType"),
		// 			Type: to.Ptr("Microsoft.Network/availableDelegations"),
		// 			Actions: []*string{
		// 				to.Ptr("Microsoft.Network/resource/action")},
		// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/availableDelegations/Microsoft.Provider.resourceType"),
		// 				ServiceName: to.Ptr("Microsoft.Provider/resourceType"),
		// 		}},
		// 	}
	}
}
