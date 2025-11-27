package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ManagedClusterPutOperation_example_min.json
func ExampleManagedClustersClient_BeginCreateOrUpdate_putAClusterWithMinimumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedClustersClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", armservicefabricmanagedclusters.ManagedCluster{
		Location: to.Ptr("eastus"),
		Properties: &armservicefabricmanagedclusters.ManagedClusterProperties{
			AdminPassword:         to.Ptr("{vm-password}"),
			AdminUserName:         to.Ptr("vmadmin"),
			ClusterUpgradeCadence: to.Ptr(armservicefabricmanagedclusters.ClusterUpgradeCadenceWave1),
			ClusterUpgradeMode:    to.Ptr(armservicefabricmanagedclusters.ClusterUpgradeModeAutomatic),
			DNSName:               to.Ptr("myCluster"),
			FabricSettings: []*armservicefabricmanagedclusters.SettingsSectionDescription{
				{
					Name: to.Ptr("ManagedIdentityTokenService"),
					Parameters: []*armservicefabricmanagedclusters.SettingsParameterDescription{
						{
							Name:  to.Ptr("IsEnabled"),
							Value: to.Ptr("true"),
						},
					},
				},
			},
		},
		SKU: &armservicefabricmanagedclusters.SKU{
			Name: to.Ptr(armservicefabricmanagedclusters.SKUNameBasic),
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
	// res = armservicefabricmanagedclusters.ManagedClustersClientCreateOrUpdateResponse{
	// 	ManagedCluster: &armservicefabricmanagedclusters.ManagedCluster{
	// 		Name: to.Ptr("myCluster"),
	// 		Type: to.Ptr("Microsoft.ServiceFabric/managedClusters"),
	// 		Etag: to.Ptr("W/\"636462502169240745\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedClusters/myCluster"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armservicefabricmanagedclusters.ManagedClusterProperties{
	// 			AdminUserName: to.Ptr("vmadmin"),
	// 			ClientConnectionPort: to.Ptr[int32](19000),
	// 			ClusterCertificateThumbprints: []*string{
	// 				to.Ptr("12599211F8F14C90AFA9532AD79A6F2CA1C00622"),
	// 			},
	// 			ClusterID: to.Ptr("92584666-9889-4ae8-8d02-91902923d37f"),
	// 			ClusterState: to.Ptr(armservicefabricmanagedclusters.ClusterStateWaitingForNodes),
	// 			ClusterUpgradeCadence: to.Ptr(armservicefabricmanagedclusters.ClusterUpgradeCadenceWave1),
	// 			ClusterUpgradeMode: to.Ptr(armservicefabricmanagedclusters.ClusterUpgradeModeAutomatic),
	// 			DNSName: to.Ptr("myCluster"),
	// 			FabricSettings: []*armservicefabricmanagedclusters.SettingsSectionDescription{
	// 				{
	// 					Name: to.Ptr("ManagedIdentityTokenService"),
	// 					Parameters: []*armservicefabricmanagedclusters.SettingsParameterDescription{
	// 						{
	// 							Name: to.Ptr("IsEnabled"),
	// 							Value: to.Ptr("true"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Fqdn: to.Ptr("MyCluster.eastus.cloudapp.azure.com"),
	// 			HTTPGatewayConnectionPort: to.Ptr[int32](19080),
	// 			LoadBalancingRules: []*armservicefabricmanagedclusters.LoadBalancingRule{
	// 			},
	// 			ProvisioningState: to.Ptr(armservicefabricmanagedclusters.ManagedResourceProvisioningStateCreating),
	// 		},
	// 		SKU: &armservicefabricmanagedclusters.SKU{
	// 			Name: to.Ptr(armservicefabricmanagedclusters.SKUNameBasic),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
