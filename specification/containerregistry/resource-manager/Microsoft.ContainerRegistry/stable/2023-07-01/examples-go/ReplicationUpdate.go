package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bbe1ea8bf5aa6cfbfa8855e03dbb9a93f8266bcd/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/examples/ReplicationUpdate.json
func ExampleReplicationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationsClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myReplication", armcontainerregistry.ReplicationUpdateParameters{
		Tags: map[string]*string{
			"key": to.Ptr("value"),
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
	// 		ZoneRedundancy: to.Ptr(armcontainerregistry.ZoneRedundancyDisabled),
	// 	},
	// }
}
