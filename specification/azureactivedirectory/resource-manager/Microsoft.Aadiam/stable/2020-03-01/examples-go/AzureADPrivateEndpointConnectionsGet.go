package armaad_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateEndpointConnectionsGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armaad.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "myResourceGroup", "example-policy-5849", "{privateEndpointConnection name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armaad.PrivateEndpointConnection{
	// 	Name: to.Ptr("{privateEndpointConnection name}"),
	// 	Type: to.Ptr("microsoft.aadiam/privateLinkForAzureAD/privateEndpointConnections"),
	// 	ID: to.Ptr("subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroup/providers/microsoft.aadiam/privateLinkForAzureAD/example-policy-5849/privateLinkConnections/{privateEndpointConnection name}"),
	// 	Properties: &armaad.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armaad.PrivateEndpoint{
	// 			ID: to.Ptr("subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroup/providers/microsoft.aadiam/privateLinkForAzureAD/example-policy-5849/privateLinkConnections/{privateEndpointConnection name}"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armaad.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("approve please"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armaad.PrivateEndpointServiceConnectionStatusPending),
	// 		},
	// 		ProvisioningState: to.Ptr(armaad.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
