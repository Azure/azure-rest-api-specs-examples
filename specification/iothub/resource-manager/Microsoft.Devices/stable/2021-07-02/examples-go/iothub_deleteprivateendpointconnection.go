package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_deleteprivateendpointconnection.json
func ExamplePrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginDelete(ctx, "myResourceGroup", "testHub", "myPrivateEndpointConnection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armiothub.PrivateEndpointConnection{
	// 	Name: to.Ptr("myPrivateEndpointConnection"),
	// 	Type: to.Ptr("Microsoft.Devices/IotHubs/PrivateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/testHub/PrivateEndpointConnections/myPrivateEndpointConnection"),
	// 	Properties: &armiothub.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armiothub.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/a9eba280-4734-4d49-878f-b5549d1d0453/resourceGroups/networkResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armiothub.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Deleted"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armiothub.PrivateLinkServiceConnectionStatusDisconnected),
	// 		},
	// 	},
	// }
}
