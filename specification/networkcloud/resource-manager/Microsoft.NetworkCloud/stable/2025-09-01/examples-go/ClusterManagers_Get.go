package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/46c51b03d99b113ecc3b38883e3cb2d395fe94a4/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/ClusterManagers_Get.json
func ExampleClusterManagersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClusterManagersClient().Get(ctx, "resourceGroupName", "clusterManagerName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterManager = armnetworkcloud.ClusterManager{
	// 	Name: to.Ptr("clusterManagerName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/clusterManagers"),
	// 	ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusterManagers/clusterManagerName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	Identity: &armnetworkcloud.ManagedServiceIdentity{
	// 		Type: to.Ptr(armnetworkcloud.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("321e7654-e89b-12d3-a456-426655449999"),
	// 		TenantID: to.Ptr("199e9999-e89b-12d3-a456-426655441111"),
	// 	},
	// 	Properties: &armnetworkcloud.ClusterManagerProperties{
	// 		AnalyticsWorkspaceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/microsoft.operationalInsights/workspaces/logAnalyticsWorkspaceName"),
	// 		AvailabilityZones: []*string{
	// 			to.Ptr("1"),
	// 			to.Ptr("2")},
	// 			ClusterVersions: []*armnetworkcloud.ClusterAvailableVersion{
	// 				{
	// 					SupportExpiryDate: to.Ptr("2023-04-29"),
	// 					TargetClusterVersion: to.Ptr("1.0.0"),
	// 				},
	// 				{
	// 					SupportExpiryDate: to.Ptr("2025-01-01"),
	// 					TargetClusterVersion: to.Ptr("1.0.2"),
	// 			}},
	// 			DetailedStatus: to.Ptr(armnetworkcloud.ClusterManagerDetailedStatusAvailable),
	// 			DetailedStatusMessage: to.Ptr("cluster manager is up and running"),
	// 			FabricControllerID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/fabricControllerName"),
	// 			ManagedResourceGroupConfiguration: &armnetworkcloud.ManagedResourceGroupConfiguration{
	// 				Name: to.Ptr("my-managed-rg"),
	// 				Location: to.Ptr("East US"),
	// 			},
	// 			ManagerExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 				Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterManagerExtendedLocationName"),
	// 				Type: to.Ptr("CustomLocation"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetworkcloud.ClusterManagerProvisioningStateSucceeded),
	// 			VMSize: to.Ptr("Standard_D8s_v3"),
	// 		},
	// 	}
}
