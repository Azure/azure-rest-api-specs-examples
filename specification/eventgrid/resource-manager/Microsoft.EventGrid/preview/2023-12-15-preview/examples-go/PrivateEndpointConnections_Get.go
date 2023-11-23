package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/PrivateEndpointConnections_Get.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "examplerg", armeventgrid.PrivateEndpointConnectionsParentTypeTopics, "exampletopic1", "BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armeventgrid.PrivateEndpointConnection{
	// 	Name: to.Ptr("BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B"),
	// 	Type: to.Ptr("Microsoft.EventGrid/topics/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1/privateEndpointConnections/BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B"),
	// 	Properties: &armeventgrid.PrivateEndpointConnectionProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("topic")},
	// 			PrivateEndpoint: &armeventgrid.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.Network/privateEndpoints/bmtpe5"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armeventgrid.ConnectionState{
	// 				Description: to.Ptr("Test"),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armeventgrid.PersistedConnectionStatusPending),
	// 			},
	// 			ProvisioningState: to.Ptr(armeventgrid.ResourceProvisioningStateSucceeded),
	// 		},
	// 	}
}
