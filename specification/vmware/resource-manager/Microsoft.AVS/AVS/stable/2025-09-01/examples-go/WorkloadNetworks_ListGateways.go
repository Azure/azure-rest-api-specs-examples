package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/WorkloadNetworks_ListGateways.json
func ExampleWorkloadNetworksClient_NewListGatewaysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkloadNetworksClient().NewListGatewaysPager("group1", "cloud1", nil)
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
		// page = armavs.WorkloadNetworksClientListGatewaysResponse{
		// 	WorkloadNetworkGatewayList: armavs.WorkloadNetworkGatewayList{
		// 		Value: []*armavs.WorkloadNetworkGateway{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/gateways/gateway1"),
		// 				Name: to.Ptr("gateway1"),
		// 				Properties: &armavs.WorkloadNetworkGatewayProperties{
		// 					DisplayName: to.Ptr("gateway1"),
		// 					Path: to.Ptr("/infra/tier-1s/gateway1"),
		// 				},
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/segments"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
