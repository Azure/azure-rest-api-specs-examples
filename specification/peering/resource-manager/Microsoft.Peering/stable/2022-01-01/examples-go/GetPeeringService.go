package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/GetPeeringService.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Get(ctx, "rgName", "peeringServiceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Service = armpeering.Service{
	// 	Name: to.Ptr("peeringServiceName"),
	// 	Type: to.Ptr("Microsoft.Peering/peeringServices"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peeringServices/peeringServiceName"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpeering.ServiceProperties{
	// 		LogAnalyticsWorkspaceProperties: &armpeering.LogAnalyticsWorkspaceProperties{
	// 			ConnectedAgents: []*string{
	// 				to.Ptr("Agent1"),
	// 				to.Ptr("Agent2")},
	// 				Key: to.Ptr("key"),
	// 				WorkspaceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 			PeeringServiceLocation: to.Ptr("state1"),
	// 			PeeringServiceProvider: to.Ptr("serviceProvider1"),
	// 			ProviderBackupPeeringLocation: to.Ptr("peeringLocation2"),
	// 			ProviderPrimaryPeeringLocation: to.Ptr("peeringLocation1"),
	// 			ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
