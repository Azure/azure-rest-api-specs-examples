package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/df863270270ad5b54fa8cce71d2c33becee0c097/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2023-01-02-preview/examples/ManagedClusterSnapshotsUpdateTags.json
func ExampleManagedClusterSnapshotsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.ManagedClusterSnapshot = armcontainerservice.ManagedClusterSnapshot{
	// 	Name: to.Ptr("snapshot1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/ManagedClusterSnapshots"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedclustersnapshots/snapshot1"),
	// 	SystemData: &armcontainerservice.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-09T20:13:23.298420761Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("val1"),
	// 		"key2": to.Ptr("val2"),
	// 	},
	// 	Properties: &armcontainerservice.ManagedClusterSnapshotProperties{
	// 		CreationData: &armcontainerservice.CreationData{
	// 			SourceResourceID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1"),
	// 		},
	// 		ManagedClusterPropertiesReadOnly: &armcontainerservice.ManagedClusterPropertiesForSnapshot{
	// 			EnableRbac: to.Ptr(true),
	// 			KubernetesVersion: to.Ptr("1.20.5"),
	// 			NetworkProfile: &armcontainerservice.NetworkProfileForSnapshot{
	// 				LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	// 				NetworkMode: to.Ptr(armcontainerservice.NetworkModeBridge),
	// 				NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
	// 				NetworkPolicy: to.Ptr(armcontainerservice.NetworkPolicyCalico),
	// 			},
	// 			SKU: &armcontainerservice.ManagedClusterSKU{
	// 				Name: to.Ptr(armcontainerservice.ManagedClusterSKUNameBasic),
	// 				Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	// 			},
	// 		},
	// 		SnapshotType: to.Ptr(armcontainerservice.SnapshotTypeManagedCluster),
	// 	},
	// }
}
