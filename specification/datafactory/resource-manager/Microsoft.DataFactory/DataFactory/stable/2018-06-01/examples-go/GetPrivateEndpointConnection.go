package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v11"
)

// Generated from example definition: 2018-06-01/GetPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "connection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatafactory.PrivateEndpointConnectionClientGetResponse{
	// 	PrivateEndpointConnectionResource: &armdatafactory.PrivateEndpointConnectionResource{
	// 		Name: to.Ptr("exampleFactoryName"),
	// 		Type: to.Ptr("Microsoft.DataFactory/factories/privateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
	// 		Properties: &armdatafactory.RemotePrivateEndpointConnection{
	// 			PrivateEndpoint: &armdatafactory.ArmIDWrapper{
	// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/privateEndpoints/myPrivateEndpoint"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armdatafactory.PrivateLinkConnectionState{
	// 				Description: to.Ptr("Approved by admin."),
	// 				ActionsRequired: to.Ptr(""),
	// 				Status: to.Ptr("Approved"),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}
