package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdateSegments.json
func ExampleWorkloadNetworksClient_BeginUpdateSegments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewWorkloadNetworksClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateSegments(ctx,
		"group1",
		"cloud1",
		"segment1",
		armavs.WorkloadNetworkSegment{
			Properties: &armavs.WorkloadNetworkSegmentProperties{
				ConnectedGateway: to.Ptr("/infra/tier-1s/gateway"),
				Revision:         to.Ptr[int64](1),
				Subnet: &armavs.WorkloadNetworkSegmentSubnet{
					DhcpRanges: []*string{
						to.Ptr("40.20.0.0-40.20.0.1")},
					GatewayAddress: to.Ptr("40.20.20.20/16"),
				},
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
