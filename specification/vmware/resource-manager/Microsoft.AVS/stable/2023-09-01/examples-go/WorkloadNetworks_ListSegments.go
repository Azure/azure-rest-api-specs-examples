package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7678455846b1000fd31db27596e4ca3d299a872/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_ListSegments.json
func ExampleWorkloadNetworksClient_NewListSegmentsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkloadNetworksClient().NewListSegmentsPager("group1", "cloud1", nil)
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
		// page.WorkloadNetworkSegmentsList = armavs.WorkloadNetworkSegmentsList{
		// 	Value: []*armavs.WorkloadNetworkSegment{
		// 		{
		// 			Name: to.Ptr("segment1"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/segments"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/segments/segment1"),
		// 			Properties: &armavs.WorkloadNetworkSegmentProperties{
		// 				ConnectedGateway: to.Ptr("/infra/tier-1s/gateway"),
		// 				DisplayName: to.Ptr("segment1"),
		// 				PortVif: []*armavs.WorkloadNetworkSegmentPortVif{
		// 					{
		// 						PortName: to.Ptr("vm1"),
		// 				}},
		// 				Revision: to.Ptr[int64](1),
		// 				Status: to.Ptr(armavs.SegmentStatusEnumSUCCESS),
		// 				Subnet: &armavs.WorkloadNetworkSegmentSubnet{
		// 					DhcpRanges: []*string{
		// 						to.Ptr("40.20.0.0-40.20.0.1")},
		// 						GatewayAddress: to.Ptr("40.20.20.20/16"),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
