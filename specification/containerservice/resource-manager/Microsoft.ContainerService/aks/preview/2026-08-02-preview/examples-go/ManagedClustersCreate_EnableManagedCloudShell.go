package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-08-02-preview/ManagedClustersCreate_EnableManagedCloudShell.json
func ExampleManagedClustersClient_BeginCreateOrUpdate_createManagedPrivateClusterWithAksManagedCloudShellResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedClustersClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
		Location: to.Ptr("location1"),
		Properties: &armcontainerservice.ManagedClusterProperties{
			APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
				EnablePrivateCluster: to.Ptr(true),
			},
			NetworkProfile: &armcontainerservice.NetworkProfile{
				ClusterAPIConnectivityProfile: &armcontainerservice.ClusterAPIConnectivityProfile{
					CloudShell: &armcontainerservice.CloudShellProfile{
						Mode: to.Ptr(armcontainerservice.CloudShellProfileModeManaged),
						Storage: &armcontainerservice.CloudShellStorageProfile{
							Mode: to.Ptr(armcontainerservice.CloudShellStorageProfileModeManaged),
						},
					},
				},
			},
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
	// res = armcontainerservice.ManagedClustersClientCreateOrUpdateResponse{
	// 	ManagedCluster: armcontainerservice.ManagedCluster{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	// 		Location: to.Ptr("location1"),
	// 		Name: to.Ptr("clustername1"),
	// 		Properties: &armcontainerservice.ManagedClusterProperties{
	// 			APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
	// 				EnablePrivateCluster: to.Ptr(true),
	// 				PrivateDNSZone: to.Ptr("system"),
	// 			},
	// 			NetworkProfile: &armcontainerservice.NetworkProfile{
	// 				ClusterAPIConnectivityProfile: &armcontainerservice.ClusterAPIConnectivityProfile{
	// 					CloudShell: &armcontainerservice.CloudShellProfile{
	// 						Mode: to.Ptr(armcontainerservice.CloudShellProfileModeManaged),
	// 						Network: &armcontainerservice.CloudShellNetworkProfile{
	// 							RelayNamespaceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MC_rg1_clustername1_location1/providers/Microsoft.Relay/namespaces/cloudshell-relay"),
	// 							SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MC_rg1_clustername1_location1/providers/Microsoft.Network/virtualNetworks/cloudshell-vnet/subnets/cloudshell-subnet"),
	// 						},
	// 						Storage: &armcontainerservice.CloudShellStorageProfile{
	// 							FileShareName: to.Ptr("cloudshell"),
	// 							Mode: to.Ptr(armcontainerservice.CloudShellStorageProfileModeManaged),
	// 							StorageAccountID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MC_rg1_clustername1_location1/providers/Microsoft.Storage/storageAccounts/cloudshellstorage"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		Type: to.Ptr("Microsoft.ContainerService/managedClusters"),
	// 	},
	// }
}
