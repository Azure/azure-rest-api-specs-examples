package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdatePortMirroringProfiles.json
func ExampleWorkloadNetworksClient_BeginUpdatePortMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewWorkloadNetworksClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdatePortMirroring(ctx,
		"group1",
		"cloud1",
		"portMirroring1",
		armavs.WorkloadNetworkPortMirroring{
			Properties: &armavs.WorkloadNetworkPortMirroringProperties{
				Destination: to.Ptr("vmGroup2"),
				Direction:   to.Ptr(armavs.PortMirroringDirectionEnumBIDIRECTIONAL),
				Revision:    to.Ptr[int64](1),
				Source:      to.Ptr("vmGroup1"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
