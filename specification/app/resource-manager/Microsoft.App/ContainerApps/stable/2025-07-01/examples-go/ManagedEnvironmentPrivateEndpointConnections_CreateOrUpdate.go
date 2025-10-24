package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b2965096067d6f8374b5485b0568fd36e7c9d099/specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/ManagedEnvironmentPrivateEndpointConnections_CreateOrUpdate.json
func ExampleManagedEnvironmentPrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedEnvironmentPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "examplerg", "managedEnv", "jlaw-demo1", armappcontainers.PrivateEndpointConnection{
		Properties: &armappcontainers.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armappcontainers.PrivateLinkServiceConnectionState{
				ActionsRequired: to.Ptr("None"),
				Status:          to.Ptr(armappcontainers.PrivateEndpointServiceConnectionStatusApproved),
			},
		},
	}, nil)
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
	// res.PrivateEndpointConnection = armappcontainers.PrivateEndpointConnection{
	// 	Name: to.Ptr("jlaw-demo1"),
	// 	Type: to.Ptr("Microsoft.App/managedEnvironments/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/managedEnv/privateEndpointConenctions/jlaw-demo1"),
	// 	Properties: &armappcontainers.PrivateEndpointConnectionProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("managedEnvironments")},
	// 			PrivateEndpoint: &armappcontainers.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.Network/privateEndpoints/pltest"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armappcontainers.PrivateLinkServiceConnectionState{
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armappcontainers.PrivateEndpointServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armappcontainers.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	}
}
