package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementApproveOrRejectPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "privateEndpointConnectionName", armapimanagement.PrivateEndpointConnectionRequest{
		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/privateEndpointConnections/connectionName"),
		Properties: &armapimanagement.PrivateEndpointConnectionRequestProperties{
			PrivateLinkServiceConnectionState: &armapimanagement.PrivateLinkServiceConnectionState{
				Description: to.Ptr("The Private Endpoint Connection is approved."),
				Status:      to.Ptr(armapimanagement.PrivateEndpointServiceConnectionStatusApproved),
			},
		},
	}, nil)
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
	// res = armapimanagement.PrivateEndpointConnectionClientCreateOrUpdateResponse{
	// 	PrivateEndpointConnection: armapimanagement.PrivateEndpointConnection{
	// 		Name: to.Ptr("privateEndpointConnectionName"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/privateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/privateEndpointConnections/privateEndpointConnectionName"),
	// 		Properties: &armapimanagement.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armapimanagement.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armapimanagement.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("The request has been approved."),
	// 				Status: to.Ptr(armapimanagement.PrivateEndpointServiceConnectionStatus("Succeeded")),
	// 			},
	// 			ProvisioningState: to.Ptr(armapimanagement.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
