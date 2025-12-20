package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/WorkloadNetworks_GetPortMirroring.json
func ExampleWorkloadNetworksClient_GetPortMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkloadNetworksClient().GetPortMirroring(ctx, "group1", "cloud1", "portMirroring1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.WorkloadNetworksClientGetPortMirroringResponse{
	// 	WorkloadNetworkPortMirroring: &armavs.WorkloadNetworkPortMirroring{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/portMirroringProfiles/portMirroring1"),
	// 		Name: to.Ptr("portMirroring1"),
	// 		Properties: &armavs.WorkloadNetworkPortMirroringProperties{
	// 			DisplayName: to.Ptr("portMirroring1"),
	// 			Direction: to.Ptr(armavs.PortMirroringDirectionEnumBIDIRECTIONAL),
	// 			Source: to.Ptr("vmGroup1"),
	// 			Destination: to.Ptr("vmGroup2"),
	// 			Status: to.Ptr(armavs.PortMirroringStatusEnumSUCCESS),
	// 			Revision: to.Ptr[int64](1),
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/portMirroringProfiles"),
	// 	},
	// }
}
