package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/IpamPools_ListAssociatedResources.json
func ExampleIpamPoolsClient_NewListAssociatedResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIpamPoolsClient().NewListAssociatedResourcesPager("rg1", "TestNetworkManager", "TestPool", nil)
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
		// page = armnetwork.IpamPoolsClientListAssociatedResourcesResponse{
		// 	PoolAssociationList: armnetwork.PoolAssociationList{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/TestNetworkManager/ipamPools/TestPool/listAssociatedResources?api-version=2025-07-01"),
		// 		Value: []*armnetwork.PoolAssociation{
		// 			{
		// 				Description: to.Ptr(""),
		// 				AddressPrefixes: []*string{
		// 					to.Ptr("10.0.0.0/24"),
		// 				},
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T20:00:00.0000000Z"); return t}()),
		// 				NumberOfReservedIPAddresses: to.Ptr("0"),
		// 				PoolID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/TestNetworkManager/ipamPools/TestPool"),
		// 				ReservedPrefixes: []*string{
		// 				},
		// 				ResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/testVnet"),
		// 				TotalNumberOfIPAddresses: to.Ptr("256"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
