package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSGetPrivateEndpointConnection.json
func ExampleIotDpsResourceClient_GetPrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIotDpsResourceClient().GetPrivateEndpointConnection(ctx, "myResourceGroup", "myFirstProvisioningService", "myPrivateEndpointConnection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armdeviceprovisioningservices.PrivateEndpointConnection{
	// 	Name: to.Ptr("myPrivateEndpointConnection"),
	// 	Type: to.Ptr("Microsoft.Devices/ProvisioningServices/PrivateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/ProvisioningServices/myFirstProvisioningService/PrivateEndpointConnections/myPrivateEndpointConnection"),
	// 	Properties: &armdeviceprovisioningservices.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armdeviceprovisioningservices.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/a9eba280-4734-4d49-878f-b5549d1d0453/resourceGroups/networkResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armdeviceprovisioningservices.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Please approve my request!"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armdeviceprovisioningservices.PrivateLinkServiceConnectionStatusPending),
	// 		},
	// 	},
	// }
}
