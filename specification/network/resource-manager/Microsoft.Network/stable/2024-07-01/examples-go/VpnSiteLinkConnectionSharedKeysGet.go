package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VpnSiteLinkConnectionSharedKeysGet.json
func ExampleVPNLinkConnectionsClient_NewGetAllSharedKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVPNLinkConnectionsClient().NewGetAllSharedKeysPager("rg1", "gateway1", "vpnConnection1", "Connection-Link1", nil)
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
		// page.ConnectionSharedKeyResultList = armnetwork.ConnectionSharedKeyResultList{
		// 	Value: []*armnetwork.ConnectionSharedKeyResult{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/vpnLinkConnections/Connection-Link1/sharedKeys/default"),
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Network/vpnGateways/vpnConnections/vpnLinkConnections/sharedKeys"),
		// 			Properties: &armnetwork.SharedKeyProperties{
		// 				SharedKeyLength: to.Ptr[int32](16),
		// 			},
		// 	}},
		// }
	}
}
