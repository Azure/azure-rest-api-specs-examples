package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_Get.json
func ExamplePrivateLinkScopesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkScopesClient().Get(ctx, "my-resource-group", "my-privatelinkscope", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkScope = armhybridcompute.PrivateLinkScope{
	// 	Name: to.Ptr("my-privatelinkscope"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes"),
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-privatelinkscope"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armhybridcompute.PrivateLinkScopeProperties{
	// 		PrivateEndpointConnections: []*armhybridcompute.PrivateEndpointConnectionDataModel{
	// 			{
	// 				Name: to.Ptr("private-endpoint-connection-name"),
	// 				Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
	// 				Properties: &armhybridcompute.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armhybridcompute.PrivateEndpointProperty{
	// 						ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armhybridcompute.PrivateLinkServiceConnectionStateProperty{
	// 						Description: to.Ptr("Auto-approved"),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr("Approved"),
	// 					},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 				},
	// 		}},
	// 		PrivateLinkScopeID: to.Ptr("f5dc51d3-92ed-4d7e-947a-775ea79b4919"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
	// 	},
	// }
}
