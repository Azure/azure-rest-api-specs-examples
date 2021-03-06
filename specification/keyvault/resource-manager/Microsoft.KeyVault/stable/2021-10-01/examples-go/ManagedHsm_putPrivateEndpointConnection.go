package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/ManagedHsm_putPrivateEndpointConnection.json
func ExampleMHSMPrivateEndpointConnectionsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkeyvault.NewMHSMPrivateEndpointConnectionsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Put(ctx,
		"sample-group",
		"sample-mhsm",
		"sample-pec",
		armkeyvault.MHSMPrivateEndpointConnection{
			Properties: &armkeyvault.MHSMPrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armkeyvault.MHSMPrivateLinkServiceConnectionState{
					Description: to.Ptr("My name is Joe and I'm approving this."),
					Status:      to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
