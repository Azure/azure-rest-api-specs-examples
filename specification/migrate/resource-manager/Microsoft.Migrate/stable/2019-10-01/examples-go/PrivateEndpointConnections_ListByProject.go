package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/PrivateEndpointConnections_ListByProject.json
func ExamplePrivateEndpointConnectionClient_ListByProject() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionClient().ListByProject(ctx, "abgoyal-westEurope", "abgoyalWEselfhostb72bproject", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionCollection = armmigrate.PrivateEndpointConnectionCollection{
	// 	Value: []*armmigrate.PrivateEndpointConnection{
	// 		{
	// 			Name: to.Ptr("custestpece80project3980pe.7e35576b-3df4-478e-9759-f64351cf4f43"),
	// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/privateEndpointConnections"),
	// 			ETag: to.Ptr("\"00009300-0000-0300-0000-602b967b0000\""),
	// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/madhavicus/providers/Microsoft.Migrate/assessmentprojects/custestpece80project/privateEndpointConnections/custestpece80project3980pe.7e35576b-3df4-478e-9759-f64351cf4f43"),
	// 			Properties: &armmigrate.PrivateEndpointConnectionProperties{
	// 				PrivateEndpoint: &armmigrate.ResourceID{
	// 					ID: to.Ptr("/subscriptions/31be0ff4-c932-4cb3-8efc-efa411d79280/resourceGroups/PrivLink-SelfHost/providers/Microsoft.Network/privateEndpoints/custestpece80project3980pe"),
	// 				},
	// 				PrivateLinkServiceConnectionState: &armmigrate.PrivateLinkServiceConnectionState{
	// 					ActionsRequired: to.Ptr(""),
	// 					Status: to.Ptr(armmigrate.PrivateLinkServiceConnectionStateStatusApproved),
	// 				},
	// 				ProvisioningState: to.Ptr(armmigrate.PrivateEndpointConnectionPropertiesProvisioningStateSucceeded),
	// 			},
	// 	}},
	// }
}
