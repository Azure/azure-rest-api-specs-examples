package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a5e7ff51c8af3781e7f6dd3b82a3fc922e2f6f93/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_CreatePortMirroringProfiles.json
func ExampleWorkloadNetworksClient_BeginCreatePortMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadNetworksClient().BeginCreatePortMirroring(ctx, "group1", "cloud1", "portMirroring1", armavs.WorkloadNetworkPortMirroring{
		Properties: &armavs.WorkloadNetworkPortMirroringProperties{
			Destination: to.Ptr("vmGroup2"),
			Direction:   to.Ptr(armavs.PortMirroringDirectionEnumBIDIRECTIONAL),
			DisplayName: to.Ptr("portMirroring1"),
			Revision:    to.Ptr[int64](1),
			Source:      to.Ptr("vmGroup1"),
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
	// res.WorkloadNetworkPortMirroring = armavs.WorkloadNetworkPortMirroring{
	// 	Name: to.Ptr("portMirroring1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/portMirroringProfiles"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/portMirroringProfiles/portMirroring1"),
	// 	Properties: &armavs.WorkloadNetworkPortMirroringProperties{
	// 		Destination: to.Ptr("vmGroup2"),
	// 		Direction: to.Ptr(armavs.PortMirroringDirectionEnumBIDIRECTIONAL),
	// 		DisplayName: to.Ptr("portMirroring1"),
	// 		Revision: to.Ptr[int64](1),
	// 		Source: to.Ptr("vmGroup1"),
	// 		Status: to.Ptr(armavs.PortMirroringStatusEnumSUCCESS),
	// 	},
	// }
}
