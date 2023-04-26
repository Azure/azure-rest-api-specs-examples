package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/NameSpaces/PrivateEndPointConnectionGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "SDK-EventHub-4794", "sdk-Namespace-5828", "privateEndpointConnectionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armeventhub.PrivateEndpointConnection{
	// 	Name: to.Ptr("privateEndpointConnectionName"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/PrivateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-EventHub-4794/providers/Microsoft.EventHub/namespaces/sdk-Namespace-5828/privateEndpointConnections/privateEndpointConnectionName"),
	// 	Properties: &armeventhub.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armeventhub.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-EventHub-4794/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-5828"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armeventhub.ConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			Status: to.Ptr(armeventhub.PrivateLinkConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armeventhub.EndPointProvisioningStateSucceeded),
	// 	},
	// }
}
