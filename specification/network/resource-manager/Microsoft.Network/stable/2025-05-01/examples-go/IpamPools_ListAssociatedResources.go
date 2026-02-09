package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/IpamPools_ListAssociatedResources.json
func ExampleIpamPoolsClient_NewListAssociatedResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.PoolAssociationList = armnetwork.PoolAssociationList{
		// 	Value: []*armnetwork.PoolAssociation{
		// 		{
		// 			Description: to.Ptr(""),
		// 			AddressPrefixes: []*string{
		// 				to.Ptr("10.0.0.0/24")},
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T20:00:00.000Z"); return t}()),
		// 				NumberOfReservedIPAddresses: to.Ptr("0"),
		// 				PoolID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/TestNetworkManager/ipamPools/TestPool"),
		// 				ReservedPrefixes: []*string{
		// 				},
		// 				ResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/testVnet"),
		// 				TotalNumberOfIPAddresses: to.Ptr("256"),
		// 		}},
		// 	}
	}
}
