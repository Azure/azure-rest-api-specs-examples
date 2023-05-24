package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/ClusterManagers_ListBySubscription.json
func ExampleClusterManagersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterManagersClient().NewListBySubscriptionPager(nil)
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
		// page.ClusterManagerList = armnetworkcloud.ClusterManagerList{
		// 	Value: []*armnetworkcloud.ClusterManager{
		// 		{
		// 			Name: to.Ptr("clusterManagerName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/clusterManagers"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusterManagers/clusterManagerName"),
		// 			SystemData: &armnetworkcloud.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("identityA"),
		// 				CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identityB"),
		// 				LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("location"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("myvalue1"),
		// 				"key2": to.Ptr("myvalue2"),
		// 			},
		// 			Properties: &armnetworkcloud.ClusterManagerProperties{
		// 				AnalyticsWorkspaceID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/microsoft.operationalInsights/workspaces/logAnalyticsWorkspaceName"),
		// 				ClusterVersions: []*armnetworkcloud.ClusterAvailableVersion{
		// 					{
		// 						SupportExpiryDate: to.Ptr("2023-04-29"),
		// 						TargetClusterVersion: to.Ptr("1.0.0"),
		// 					},
		// 					{
		// 						SupportExpiryDate: to.Ptr("2025-01-01"),
		// 						TargetClusterVersion: to.Ptr("1.0.2"),
		// 				}},
		// 				DetailedStatus: to.Ptr(armnetworkcloud.ClusterManagerDetailedStatusAvailable),
		// 				DetailedStatusMessage: to.Ptr("cluster manager is up and running"),
		// 				FabricControllerID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/fabricControllerName"),
		// 				ManagedResourceGroupConfiguration: &armnetworkcloud.ManagedResourceGroupConfiguration{
		// 					Name: to.Ptr("my-managed-rg"),
		// 					Location: to.Ptr("East US"),
		// 				},
		// 				ManagerExtendedLocation: &armnetworkcloud.ExtendedLocation{
		// 					Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterManagerExtendedLocationName"),
		// 					Type: to.Ptr("CustomLocation"),
		// 				},
		// 				ProvisioningState: to.Ptr(armnetworkcloud.ClusterManagerProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
