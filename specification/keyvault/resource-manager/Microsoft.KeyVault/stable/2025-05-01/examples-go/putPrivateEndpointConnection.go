package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
)

// Generated from example definition: 2025-05-01/putPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Put(ctx, "sample-group", "sample-vault", "sample-pec", armkeyvault.PrivateEndpointConnection{
		Etag: nil,
		Properties: &armkeyvault.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armkeyvault.PrivateLinkServiceConnectionState{
				Description: to.Ptr("My name is Joe and I'm approving this."),
				Status:      to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkeyvault.PrivateEndpointConnectionsClientPutResponse{
	// 	PrivateEndpointConnection: &armkeyvault.PrivateEndpointConnection{
	// 		Name: to.Ptr("sample-pec"),
	// 		Type: to.Ptr("Microsoft.KeyVault/vaults/privateEndpointConnections"),
	// 		Etag: &azcore.ETag(""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault/privateEndpointConnections/sample-pec"),
	// 		Properties: &armkeyvault.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armkeyvault.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armkeyvault.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("My name is Joe and I'm approving this."),
	// 				ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
	// 				Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
