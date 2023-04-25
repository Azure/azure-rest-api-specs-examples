package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/getPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "sample-group", "sample-vault", "sample-pec", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armkeyvault.PrivateEndpointConnection{
	// 	Name: to.Ptr("sample-pec"),
	// 	Type: to.Ptr("Microsoft.KeyVault/vaults/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault/privateEndpointConnections/sample-pec"),
	// 	Etag: to.Ptr(""),
	// 	Properties: &armkeyvault.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armkeyvault.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armkeyvault.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("This was automatically approved by user1234@contoso.com"),
	// 			ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
	// 			Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
