package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-01-02-preview/ManagedClusterSnapshotsUpdateTags.json
func ExampleManagedClusterSnapshotsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClusterSnapshotsClient().UpdateTags(ctx, "rg1", "snapshot1", armcontainerservice.TagsObject{
		Tags: map[string]*string{
			"key2": to.Ptr("new-val2"),
			"key3": to.Ptr("val3"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservice.ManagedClusterSnapshotsClientUpdateTagsResponse{
	// 	ManagedClusterSnapshot: &armcontainerservice.ManagedClusterSnapshot{
	// 		Name: to.Ptr("snapshot1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/managedClusterSnapshots"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedclustersnapshots/snapshot1"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armcontainerservice.ManagedClusterSnapshotProperties{
	// 			CreationData: &armcontainerservice.CreationData{
	// 				SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1"),
	// 			},
	// 			ManagedClusterPropertiesReadOnly: &armcontainerservice.ManagedClusterPropertiesForSnapshot{
	// 				EnableRbac: to.Ptr(true),
	// 				KubernetesVersion: to.Ptr("1.20.5"),
	// 				NetworkProfile: &armcontainerservice.NetworkProfileForSnapshot{
	// 					LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	// 					NetworkMode: to.Ptr(armcontainerservice.NetworkModeBridge),
	// 					NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
	// 					NetworkPolicy: to.Ptr(armcontainerservice.NetworkPolicyCalico),
	// 				},
	// 				SKU: &armcontainerservice.ManagedClusterSKU{
	// 					Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	// 					Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	// 				},
	// 			},
	// 			SnapshotType: to.Ptr(armcontainerservice.SnapshotTypeManagedCluster),
	// 		},
	// 		SystemData: &armcontainerservice.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-09T20:13:23.298420761Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("val1"),
	// 			"key2": to.Ptr("val2"),
	// 		},
	// 	},
	// }
}
