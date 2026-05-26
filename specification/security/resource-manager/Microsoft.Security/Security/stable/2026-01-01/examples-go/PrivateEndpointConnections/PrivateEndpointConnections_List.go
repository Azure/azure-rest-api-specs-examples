package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2026-01-01/PrivateEndpointConnections/PrivateEndpointConnections_List.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("rg", armsecurity.PrivateLinkParameters{PrivateLinkName: "pls"}, nil)
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
		// page = armsecurity.PrivateEndpointConnectionsClientListResponse{
		// 	PrivateEndpointConnectionListResult: armsecurity.PrivateEndpointConnectionListResult{
		// 		Value: []*armsecurity.PrivateEndpointConnection{
		// 			{
		// 				Name: to.Ptr("pe"),
		// 				Type: to.Ptr("Microsoft.Security/privateLinks/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Security/privateLinks/pls/privateEndpointConnections/pe"),
		// 				Properties: &armsecurity.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("containers"),
		// 					},
		// 					PrivateEndpoint: &armsecurity.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Network/privateEndpoints/pe"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armsecurity.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armsecurity.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armsecurity.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("pe2"),
		// 				Type: to.Ptr("Microsoft.Security/privateLinks/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Security/privateLinks/pls/privateEndpointConnections/pe2"),
		// 				Properties: &armsecurity.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("containers"),
		// 					},
		// 					PrivateEndpoint: &armsecurity.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Network/privateEndpoints/pe2"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armsecurity.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Waiting for approval"),
		// 						ActionsRequired: to.Ptr("Manual approval required"),
		// 						Status: to.Ptr(armsecurity.PrivateEndpointServiceConnectionStatusPending),
		// 					},
		// 					ProvisioningState: to.Ptr(armsecurity.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
