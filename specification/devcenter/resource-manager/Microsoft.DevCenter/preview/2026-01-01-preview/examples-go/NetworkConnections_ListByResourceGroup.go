package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/NetworkConnections_ListByResourceGroup.json
func ExampleNetworkConnectionsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkConnectionsClient().NewListByResourceGroupPager("rg1", nil)
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
		// page = armdevcenter.NetworkConnectionsClientListByResourceGroupResponse{
		// 	NetworkConnectionListResult: armdevcenter.NetworkConnectionListResult{
		// 		Value: []*armdevcenter.NetworkConnection{
		// 			{
		// 				Name: to.Ptr("uswest3network"),
		// 				Type: to.Ptr("Microsoft.DevCenter/networkconnections"),
		// 				ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/networkconnections/uswest3network"),
		// 				Location: to.Ptr("centralus"),
		// 				Properties: &armdevcenter.NetworkProperties{
		// 					DomainJoinType: to.Ptr(armdevcenter.DomainJoinTypeHybridAzureADJoin),
		// 					DomainName: to.Ptr("mydomaincontroller.local"),
		// 					DomainUsername: to.Ptr("testuser@mydomaincontroller.local"),
		// 					NetworkingResourceGroupName: to.Ptr("NetworkInterfaces"),
		// 					ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
		// 					SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ExampleRG/providers/Microsoft.Network/virtualNetworks/ExampleVNet/subnets/default"),
		// 				},
		// 				SystemData: &armdevcenter.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user1"),
		// 					LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
