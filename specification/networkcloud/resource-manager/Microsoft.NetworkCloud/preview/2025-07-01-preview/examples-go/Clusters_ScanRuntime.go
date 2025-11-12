package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d3ac611f503e05650fb85520582b06140d2599e/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2025-07-01-preview/examples/Clusters_ScanRuntime.json
func ExampleClustersClient_BeginScanRuntime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginScanRuntime(ctx, "resourceGroupName", "clusterName", &armnetworkcloud.ClustersClientBeginScanRuntimeOptions{ClusterScanRuntimeParameters: &armnetworkcloud.ClusterScanRuntimeParameters{
		ScanActivity: to.Ptr(armnetworkcloud.ClusterScanRuntimeParametersScanActivityScan),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
