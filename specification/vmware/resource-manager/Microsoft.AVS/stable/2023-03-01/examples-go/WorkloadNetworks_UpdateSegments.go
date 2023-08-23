package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a5e7ff51c8af3781e7f6dd3b82a3fc922e2f6f93/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_UpdateSegments.json
func ExampleWorkloadNetworksClient_BeginUpdateSegments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadNetworksClient().BeginUpdateSegments(ctx, "group1", "cloud1", "segment1", armavs.WorkloadNetworkSegment{
		Properties: &armavs.WorkloadNetworkSegmentProperties{
			ConnectedGateway: to.Ptr("/infra/tier-1s/gateway"),
			Revision:         to.Ptr[int64](1),
			Subnet: &armavs.WorkloadNetworkSegmentSubnet{
				DhcpRanges: []*string{
					to.Ptr("40.20.0.0-40.20.0.1")},
				GatewayAddress: to.Ptr("40.20.20.20/16"),
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
	// res.WorkloadNetworkSegment = armavs.WorkloadNetworkSegment{
	// 	Name: to.Ptr("segment1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/segments"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/segments/segment1"),
	// 	Properties: &armavs.WorkloadNetworkSegmentProperties{
	// 		ConnectedGateway: to.Ptr("/infra/tier-1s/gateway"),
	// 		DisplayName: to.Ptr("segment1"),
	// 		PortVif: []*armavs.WorkloadNetworkSegmentPortVif{
	// 			{
	// 				PortName: to.Ptr("vm1"),
	// 		}},
	// 		Revision: to.Ptr[int64](2),
	// 		Status: to.Ptr(armavs.SegmentStatusEnumSUCCESS),
	// 		Subnet: &armavs.WorkloadNetworkSegmentSubnet{
	// 			DhcpRanges: []*string{
	// 				to.Ptr("40.20.0.0-40.20.0.1")},
	// 				GatewayAddress: to.Ptr("40.20.20.20/16"),
	// 			},
	// 		},
	// 	}
}
