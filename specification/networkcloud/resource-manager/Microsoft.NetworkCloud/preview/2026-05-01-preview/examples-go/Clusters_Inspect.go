package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: 2026-05-01-preview/Clusters_Inspect.json
func ExampleClustersClient_BeginInspect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("123e4567-e89b-12d3-a456-426655440000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginInspect(ctx, "resourceGroupName", "clusterName", &armnetworkcloud.ClustersClientBeginInspectOptions{
		ClusterInspectParameters: &armnetworkcloud.ClusterInspectParameters{
			AdditionalActions: []*armnetworkcloud.ClusterInspectAdditionalAction{
				to.Ptr(armnetworkcloud.ClusterInspectAdditionalActionResetHardware),
			},
			FilterDevices: &armnetworkcloud.FilterDevices{
				BareMetalMachineNames: []*string{
					to.Ptr("machine1"),
					to.Ptr("machine2"),
				},
				RackNames: []*string{
					to.Ptr("rack1"),
				},
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
