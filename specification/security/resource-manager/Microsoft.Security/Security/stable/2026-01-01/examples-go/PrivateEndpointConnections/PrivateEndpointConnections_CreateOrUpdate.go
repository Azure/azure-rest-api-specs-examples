package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2026-01-01/PrivateEndpointConnections/PrivateEndpointConnections_CreateOrUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "rg", "pe", armsecurity.PrivateEndpointConnection{
		Properties: &armsecurity.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armsecurity.PrivateLinkServiceConnectionState{
				Description:     to.Ptr("Approved by administrator"),
				ActionsRequired: to.Ptr("None"),
				Status:          to.Ptr(armsecurity.PrivateEndpointServiceConnectionStatusApproved),
			},
		},
	}, armsecurity.PrivateLinkParameters{PrivateLinkName: "pls"}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.PrivateEndpointConnectionsClientCreateOrUpdateResponse{
	// 	PrivateEndpointConnection: armsecurity.PrivateEndpointConnection{
	// 		Name: to.Ptr("pe"),
	// 		Type: to.Ptr("Microsoft.Security/privateLinks/privateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Security/privateLinks/pls/privateEndpointConnections/pe"),
	// 		Properties: &armsecurity.PrivateEndpointConnectionProperties{
	// 			GroupIDs: []*string{
	// 				to.Ptr("containers"),
	// 			},
	// 			PrivateEndpoint: &armsecurity.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Network/privateEndpoints/pe"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armsecurity.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("Approved by administrator"),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armsecurity.PrivateEndpointServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armsecurity.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
