package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_ListPortMirroring.json
func ExampleWorkloadNetworksClient_NewListPortMirroringPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkloadNetworksClient().NewListPortMirroringPager("group1", "cloud1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.WorkloadNetworkPortMirroringList = armavs.WorkloadNetworkPortMirroringList{
		// 	Value: []*armavs.WorkloadNetworkPortMirroring{
		// 		{
		// 			Name: to.Ptr("cloud1"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/portMirroringProfiles"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/portMirroringProfiles/portMirroring1"),
		// 			Properties: &armavs.WorkloadNetworkPortMirroringProperties{
		// 				Destination: to.Ptr("vmGroup2"),
		// 				Direction: to.Ptr(armavs.PortMirroringDirectionEnumBIDIRECTIONAL),
		// 				DisplayName: to.Ptr("portMirroring1"),
		// 				Revision: to.Ptr[int64](1),
		// 				Source: to.Ptr("vmGroup1"),
		// 				Status: to.Ptr(armavs.PortMirroringStatusEnumSUCCESS),
		// 			},
		// 	}},
		// }
	}
}
