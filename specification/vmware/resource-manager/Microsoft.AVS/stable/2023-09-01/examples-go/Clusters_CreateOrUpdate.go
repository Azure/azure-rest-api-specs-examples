package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7678455846b1000fd31db27596e4ca3d299a872/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_CreateOrUpdate.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", "cluster1", armavs.Cluster{
		Properties: &armavs.ClusterProperties{
			ClusterSize: to.Ptr[int32](3),
		},
		SKU: &armavs.SKU{
			Name: to.Ptr("AV20"),
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
	// res.Cluster = armavs.Cluster{
	// 	Name: to.Ptr("cluster1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/clusters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1"),
	// 	Properties: &armavs.ClusterProperties{
	// 		ClusterSize: to.Ptr[int32](3),
	// 		Hosts: []*string{
	// 			to.Ptr("fakehost22.nyc1.kubernetes.center"),
	// 			to.Ptr("fakehost23.nyc1.kubernetes.center"),
	// 			to.Ptr("fakehost24.nyc1.kubernetes.center")},
	// 			ProvisioningState: to.Ptr(armavs.ClusterProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armavs.SKU{
	// 			Name: to.Ptr("AV20"),
	// 		},
	// 	}
}
