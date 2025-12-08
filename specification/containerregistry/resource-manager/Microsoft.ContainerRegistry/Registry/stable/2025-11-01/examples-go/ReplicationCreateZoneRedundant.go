package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/523ccabf440d8cf1c5b0ea18a8ad1ffedf4902ac/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/ReplicationCreateZoneRedundant.json
func ExampleReplicationsClient_BeginCreate_replicationCreateZoneRedundant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationsClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myReplication", armcontainerregistry.Replication{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
		Properties: &armcontainerregistry.ReplicationProperties{
			RegionEndpointEnabled: to.Ptr(true),
			ZoneRedundancy:        to.Ptr(armcontainerregistry.ZoneRedundancyEnabled),
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
	// res.Replication = armcontainerregistry.Replication{
	// 	Name: to.Ptr("myReplication"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/replications"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/replications/myReplication"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.ReplicationProperties{
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RegionEndpointEnabled: to.Ptr(true),
	// 		Status: &armcontainerregistry.Status{
	// 			DisplayStatus: to.Ptr("Ready"),
	// 			Message: to.Ptr("The replication is ready."),
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:15:37.070Z"); return t}()),
	// 		},
	// 		ZoneRedundancy: to.Ptr(armcontainerregistry.ZoneRedundancyEnabled),
	// 	},
	// }
}
