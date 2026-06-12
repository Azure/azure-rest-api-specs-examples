package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: 2026-05-01-preview/ClusterManagers_Create_Uami.json
func ExampleClusterManagersClient_BeginCreateOrUpdate_createOrUpdateClusterManagerWithUserAssignedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("123e4567-e89b-12d3-a456-426655440000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterManagersClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "clusterManagerName", armnetworkcloud.ClusterManager{
		Identity: &armnetworkcloud.ManagedServiceIdentity{
			Type: to.Ptr(armnetworkcloud.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armnetworkcloud.UserAssignedIdentity{
				"/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1": {},
			},
		},
		Kind:     to.Ptr(armnetworkcloud.DeploymentTypeAzureLocal),
		Location: to.Ptr("location"),
		Properties: &armnetworkcloud.ClusterManagerProperties{
			AnalyticsWorkspaceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.OperationalInsights/workspaces/logAnalyticsWorkspaceName"),
			FabricControllerID:   to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/fabricControllerName"),
			ManagedResourceGroupConfiguration: &armnetworkcloud.ManagedResourceGroupConfiguration{
				Location: to.Ptr("East US"),
				Name:     to.Ptr("my-managed-rg"),
			},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetworkcloud.ClusterManagersClientCreateOrUpdateResponse{
	// 	ClusterManager: armnetworkcloud.ClusterManager{
	// 		ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusterManagers/clusterManagerName"),
	// 		Identity: &armnetworkcloud.ManagedServiceIdentity{
	// 			Type: to.Ptr(armnetworkcloud.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armnetworkcloud.UserAssignedIdentity{
	// 				"/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1": &armnetworkcloud.UserAssignedIdentity{
	// 					ClientID: to.Ptr("200e0000-e89b-12d3-a456-426655440000"),
	// 					PrincipalID: to.Ptr("321e7654-e89b-12d3-a456-426655449999"),
	// 				},
	// 			},
	// 		},
	// 		Kind: to.Ptr(armnetworkcloud.DeploymentTypeAzureLocal),
	// 		Location: to.Ptr("location"),
	// 		Name: to.Ptr("clusterManagerName"),
	// 		Properties: &armnetworkcloud.ClusterManagerProperties{
	// 			AnalyticsWorkspaceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.OperationalInsights/workspaces/logAnalyticsWorkspaceName"),
	// 			AvailabilityZones: []*string{
	// 				to.Ptr("1"),
	// 				to.Ptr("2"),
	// 			},
	// 			ClusterVersions: []*armnetworkcloud.ClusterAvailableVersion{
	// 				{
	// 					SupportExpiryDate: to.Ptr("2023-04-29"),
	// 					TargetClusterVersion: to.Ptr("1.0.0"),
	// 				},
	// 				{
	// 					SupportExpiryDate: to.Ptr("2025-01-01"),
	// 					TargetClusterVersion: to.Ptr("1.0.2"),
	// 				},
	// 			},
	// 			DetailedStatus: to.Ptr(armnetworkcloud.ClusterManagerDetailedStatusAvailable),
	// 			DetailedStatusMessage: to.Ptr("cluster manager is up and running"),
	// 			FabricControllerID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/fabricControllerName"),
	// 			ManagedResourceGroupConfiguration: &armnetworkcloud.ManagedResourceGroupConfiguration{
	// 				Location: to.Ptr("East US"),
	// 				Name: to.Ptr("my-managed-rg"),
	// 			},
	// 			ManagerExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 				Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterManagerExtendedLocationName"),
	// 				Type: to.Ptr("CustomLocation"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetworkcloud.ClusterManagerProvisioningStateAccepted),
	// 			RelayConfiguration: &armnetworkcloud.ClusterManagerRelayConfiguration{
	// 				RelayNamespaceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.Relay/namespaces/relayNamespaceName"),
	// 			},
	// 			VMSize: to.Ptr("Standard_D8s_v3"),
	// 		},
	// 		SystemData: &armnetworkcloud.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 			CreatedBy: to.Ptr("identityA"),
	// 			CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("identityB"),
	// 			LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("myvalue1"),
	// 			"key2": to.Ptr("myvalue2"),
	// 		},
	// 		Type: to.Ptr("Microsoft.NetworkCloud/clusterManagers"),
	// 	},
	// }
}
