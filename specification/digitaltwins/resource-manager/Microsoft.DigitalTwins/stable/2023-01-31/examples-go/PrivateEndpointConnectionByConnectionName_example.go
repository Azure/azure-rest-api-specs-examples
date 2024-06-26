package armdigitaltwins_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/PrivateEndpointConnectionByConnectionName_example.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdigitaltwins.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "resRg", "myDigitalTwinsService", "myPrivateConnection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armdigitaltwins.PrivateEndpointConnection{
	// 	ID: to.Ptr("/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourcegroups/resRg/providers/Microsoft.DigitalTwins/digitalTwinsInstances/myDigitalTwinsService/privateEndpointConnections/myPrivateConnection"),
	// 	Properties: &armdigitaltwins.ConnectionProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("digitalTwinsInstance")},
	// 			PrivateEndpoint: &armdigitaltwins.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourceGroups/resRg/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armdigitaltwins.ConnectionPropertiesPrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("Please approve my request, thanks."),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armdigitaltwins.PrivateLinkServiceConnectionStatusPending),
	// 			},
	// 		},
	// 	}
}
