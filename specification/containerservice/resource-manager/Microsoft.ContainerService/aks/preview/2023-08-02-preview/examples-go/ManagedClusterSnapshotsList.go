package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/89260be1a92c914b7b48af8e8f75938d5e76851d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-08-02-preview/examples/ManagedClusterSnapshotsList.json
func ExampleManagedClusterSnapshotsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedClusterSnapshotsClient().NewListPager(nil)
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
		// page.ManagedClusterSnapshotListResult = armcontainerservice.ManagedClusterSnapshotListResult{
		// 	Value: []*armcontainerservice.ManagedClusterSnapshot{
		// 		{
		// 			Name: to.Ptr("snapshot1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/ManagedClusterSnapshots"),
		// 			ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedclustersnapshots/snapshot1"),
		// 			SystemData: &armcontainerservice.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-09T20:13:23.298420761Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("val1"),
		// 				"key2": to.Ptr("val2"),
		// 			},
		// 			Properties: &armcontainerservice.ManagedClusterSnapshotProperties{
		// 				CreationData: &armcontainerservice.CreationData{
		// 					SourceResourceID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1"),
		// 				},
		// 				ManagedClusterPropertiesReadOnly: &armcontainerservice.ManagedClusterPropertiesForSnapshot{
		// 					EnableRbac: to.Ptr(true),
		// 					KubernetesVersion: to.Ptr("1.20.5"),
		// 					NetworkProfile: &armcontainerservice.NetworkProfileForSnapshot{
		// 						LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
		// 						NetworkMode: to.Ptr(armcontainerservice.NetworkModeBridge),
		// 						NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
		// 						NetworkPolicy: to.Ptr(armcontainerservice.NetworkPolicyCalico),
		// 					},
		// 					SKU: &armcontainerservice.ManagedClusterSKU{
		// 						Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
		// 						Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
		// 					},
		// 				},
		// 				SnapshotType: to.Ptr(armcontainerservice.SnapshotTypeManagedCluster),
		// 			},
		// 	}},
		// }
	}
}
