package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
)

// Generated from example definition: 2025-05-01/ManagedHsm_getPrivateEndpointConnection.json
func ExampleMHSMPrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMHSMPrivateEndpointConnectionsClient().Get(ctx, "sample-group", "sample-mhsm", "sample-pec", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkeyvault.MHSMPrivateEndpointConnectionsClientGetResponse{
	// 	MHSMPrivateEndpointConnection: &armkeyvault.MHSMPrivateEndpointConnection{
	// 		Name: to.Ptr("sample-pec"),
	// 		Type: to.Ptr("Microsoft.KeyVault/managedhsms/privateEndpointConnections"),
	// 		Etag: &azcore.ETag(""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedhsms/sample-mhsm/privateEndpointConnections/sample-pec"),
	// 		Properties: &armkeyvault.MHSMPrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armkeyvault.MHSMPrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armkeyvault.MHSMPrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("This was automatically approved by user1234@contoso.com"),
	// 				ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
	// 				Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
