package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ManagedClusterListBySubscriptionOperation_example.json
func ExampleManagedClustersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedClustersClient().NewListBySubscriptionPager(nil)
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
		// page = armservicefabricmanagedclusters.ManagedClustersClientListBySubscriptionResponse{
		// 	ManagedClusterListResult: armservicefabricmanagedclusters.ManagedClusterListResult{
		// 		NextLink: to.Ptr("http://examplelink.com"),
		// 		Value: []*armservicefabricmanagedclusters.ManagedCluster{
		// 			{
		// 				Name: to.Ptr("myCluster"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/managedClusters"),
		// 				Etag: to.Ptr("W/\"636462502169240745\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedClusters/myCluster"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armservicefabricmanagedclusters.ManagedClusterProperties{
		// 					AdminUserName: to.Ptr("vmadmin"),
		// 					ClientConnectionPort: to.Ptr[int32](19000),
		// 					ClusterCertificateThumbprints: []*string{
		// 						to.Ptr("12599211F8F14C90AFA9532AD79A6F2CA1C00622"),
		// 					},
		// 					ClusterCodeVersion: to.Ptr("7.1.168.9494"),
		// 					ClusterID: to.Ptr("92584666-9889-4ae8-8d02-91902923d37f"),
		// 					ClusterState: to.Ptr(armservicefabricmanagedclusters.ClusterStateWaitingForNodes),
		// 					ClusterUpgradeCadence: to.Ptr(armservicefabricmanagedclusters.ClusterUpgradeCadenceWave0),
		// 					ClusterUpgradeMode: to.Ptr(armservicefabricmanagedclusters.ClusterUpgradeModeAutomatic),
		// 					DNSName: to.Ptr("myCluster"),
		// 					FabricSettings: []*armservicefabricmanagedclusters.SettingsSectionDescription{
		// 						{
		// 							Name: to.Ptr("ManagedIdentityTokenService"),
		// 							Parameters: []*armservicefabricmanagedclusters.SettingsParameterDescription{
		// 								{
		// 									Name: to.Ptr("IsEnabled"),
		// 									Value: to.Ptr("true"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Fqdn: to.Ptr("MyCluster.eastus.cloudapp.azure.com"),
		// 					HTTPGatewayConnectionPort: to.Ptr[int32](19080),
		// 					LoadBalancingRules: []*armservicefabricmanagedclusters.LoadBalancingRule{
		// 					},
		// 					ProvisioningState: to.Ptr(armservicefabricmanagedclusters.ManagedResourceProvisioningStateSucceeded),
		// 				},
		// 				SKU: &armservicefabricmanagedclusters.SKU{
		// 					Name: to.Ptr(armservicefabricmanagedclusters.SKUNameStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myCluster2"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/managedClusters"),
		// 				Etag: to.Ptr("W/\"636462502164040075\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedClusters/myCluster2"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armservicefabricmanagedclusters.ManagedClusterProperties{
		// 					AdminUserName: to.Ptr("vmadmin"),
		// 					ClientConnectionPort: to.Ptr[int32](19000),
		// 					ClusterCertificateThumbprints: []*string{
		// 						to.Ptr("12599211F8F14C90AFA9532AD79A6F2CA1C00622"),
		// 					},
		// 					ClusterCodeVersion: to.Ptr("7.1.168.9494"),
		// 					ClusterID: to.Ptr("2747e469-b24e-4039-8a0a-46151419523f"),
		// 					ClusterState: to.Ptr(armservicefabricmanagedclusters.ClusterStateWaitingForNodes),
		// 					ClusterUpgradeCadence: to.Ptr(armservicefabricmanagedclusters.ClusterUpgradeCadenceWave0),
		// 					ClusterUpgradeMode: to.Ptr(armservicefabricmanagedclusters.ClusterUpgradeModeAutomatic),
		// 					DNSName: to.Ptr("myCluster2"),
		// 					FabricSettings: []*armservicefabricmanagedclusters.SettingsSectionDescription{
		// 						{
		// 							Name: to.Ptr("ManagedIdentityTokenService"),
		// 							Parameters: []*armservicefabricmanagedclusters.SettingsParameterDescription{
		// 								{
		// 									Name: to.Ptr("IsEnabled"),
		// 									Value: to.Ptr("true"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Fqdn: to.Ptr("MyCluster.eastus.cloudapp.azure.com"),
		// 					HTTPGatewayConnectionPort: to.Ptr[int32](19080),
		// 					LoadBalancingRules: []*armservicefabricmanagedclusters.LoadBalancingRule{
		// 					},
		// 					ProvisioningState: to.Ptr(armservicefabricmanagedclusters.ManagedResourceProvisioningStateSucceeded),
		// 				},
		// 				SKU: &armservicefabricmanagedclusters.SKU{
		// 					Name: to.Ptr(armservicefabricmanagedclusters.SKUNameStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
