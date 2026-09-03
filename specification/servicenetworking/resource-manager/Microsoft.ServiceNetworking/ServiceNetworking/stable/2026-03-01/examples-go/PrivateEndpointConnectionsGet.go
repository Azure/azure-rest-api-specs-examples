package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: 2026-03-01/PrivateEndpointConnectionsGet.json
func ExamplePrivateEndpointConnectionsInterfaceClient_NewListByTrafficControllerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsInterfaceClient().NewListByTrafficControllerPager("rg1", "tc1", nil)
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
		// page = armservicenetworking.PrivateEndpointConnectionsInterfaceClientListByTrafficControllerResponse{
		// 	PrivateEndpointConnectionListResult: armservicenetworking.PrivateEndpointConnectionListResult{
		// 		Value: []*armservicenetworking.PrivateEndpointConnection{
		// 			{
		// 				Name: to.Ptr("pec1"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/privateEndpointConnections/pec1"),
		// 				Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/privateEndpointConnections"),
		// 				Properties: &armservicenetworking.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armservicenetworking.PrivateEndpointReference{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/pe1"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armservicenetworking.PrivateLinkServiceConnectionState{
		// 						Status: to.Ptr(armservicenetworking.PrivateLinkServiceConnectionStatusApproved),
		// 						Description: to.Ptr("Approved by admin"),
		// 					},
		// 					ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
