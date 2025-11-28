package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ServiceActionRestartReplica_example.json
func ExampleServicesClient_BeginRestartReplica() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginRestartReplica(ctx, "resRg", "myCluster", "myApp", "myService", armservicefabricmanagedclusters.RestartReplicaRequest{
		PartitionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		ReplicaIDs: []*int64{
			to.Ptr[int64](123456789012345680),
		},
		RestartKind:  to.Ptr(armservicefabricmanagedclusters.RestartKindSimultaneous),
		ForceRestart: to.Ptr(false),
		Timeout:      to.Ptr[int64](60),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
