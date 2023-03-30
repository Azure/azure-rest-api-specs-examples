package armiotcentral_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/96e52e2b911d533f95a0ad8e324c828d556c5f2b/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateEndpointConnections_Get.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "resRg", "myIoTCentralApp", "myIoTCentralAppEndpoint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armiotcentral.PrivateEndpointConnection{
	// 	Name: to.Ptr("myIoTCentralAppEndpoint.a791c6b5-874d-4f03-9092-718490d33770"),
	// 	Type: to.Ptr("Microsoft.IoTCentral/iotApps/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.IoTCentral/iotApps/myIoTCentralApp/PrivateEndpointConnections/myIoTCentralAppEndpoint.a791c6b5-874d-4f03-9092-718490d33770"),
	// 	Properties: &armiotcentral.PrivateEndpointConnectionProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("iotApp")},
	// 			PrivateEndpoint: &armiotcentral.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/privateEndpoints/myIoTCentralAppEndpoint"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armiotcentral.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("Auto-approved"),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armiotcentral.PrivateEndpointServiceConnectionStatusApproved),
	// 			},
	// 		},
	// 	}
}
