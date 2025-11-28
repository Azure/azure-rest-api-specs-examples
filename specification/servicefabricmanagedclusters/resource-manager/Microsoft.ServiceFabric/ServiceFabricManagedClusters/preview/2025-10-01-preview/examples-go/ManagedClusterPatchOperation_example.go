package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ManagedClusterPatchOperation_example.json
func ExampleManagedClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedClustersClient().BeginUpdate(ctx, "resRg", "myCluster", armservicefabricmanagedclusters.ManagedClusterUpdateParameters{
		Tags: map[string]*string{
			"a": to.Ptr("b"),
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
	// res = armservicefabricmanagedclusters.ManagedClustersClientUpdateResponse{
	// 	ManagedCluster: &armservicefabricmanagedclusters.ManagedCluster{
	// 		Name: to.Ptr("myCluster"),
	// 		Type: to.Ptr("Microsoft.ServiceFabric/managedClusters"),
	// 		Etag: to.Ptr("W/\"636462502169240745\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedClusters/myCluster"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armservicefabricmanagedclusters.ManagedClusterProperties{
	// 			AddonFeatures: []*armservicefabricmanagedclusters.ManagedClusterAddOnFeature{
	// 				to.Ptr(armservicefabricmanagedclusters.ManagedClusterAddOnFeatureDNSService),
	// 				to.Ptr(armservicefabricmanagedclusters.ManagedClusterAddOnFeatureBackupRestoreService),
	// 				to.Ptr(armservicefabricmanagedclusters.ManagedClusterAddOnFeatureResourceMonitorService),
	// 			},
	// 			AdminUserName: to.Ptr("vmadmin"),
	// 			AllowRdpAccess: to.Ptr(true),
	// 			ClientConnectionPort: to.Ptr[int32](19001),
	// 			ClusterCertificateThumbprints: []*string{
	// 				to.Ptr("12599211F8F14C90AFA9532AD79A6F2CA1C00622"),
	// 			},
	// 			ClusterCodeVersion: to.Ptr("7.1.168.9494"),
	// 			ClusterID: to.Ptr("92584666-9889-4ae8-8d02-91902923d37f"),
	// 			ClusterState: to.Ptr(armservicefabricmanagedclusters.ClusterStateWaitingForNodes),
	// 			ClusterUpgradeCadence: to.Ptr(armservicefabricmanagedclusters.ClusterUpgradeCadenceWave0),
	// 			ClusterUpgradeMode: to.Ptr(armservicefabricmanagedclusters.ClusterUpgradeModeAutomatic),
	// 			DNSName: to.Ptr("myCluster"),
	// 			EnableAutoOSUpgrade: to.Ptr(true),
	// 			FabricSettings: []*armservicefabricmanagedclusters.SettingsSectionDescription{
	// 				{
	// 					Name: to.Ptr("ManagedIdentityTokenService"),
	// 					Parameters: []*armservicefabricmanagedclusters.SettingsParameterDescription{
	// 						{
	// 							Name: to.Ptr("IsEnabled"),
	// 							Value: to.Ptr("false"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Fqdn: to.Ptr("MyCluster.eastus.cloudapp.azure.com"),
	// 			HTTPGatewayConnectionPort: to.Ptr[int32](19081),
	// 			LoadBalancingRules: []*armservicefabricmanagedclusters.LoadBalancingRule{
	// 				{
	// 					BackendPort: to.Ptr[int32](80),
	// 					FrontendPort: to.Ptr[int32](80),
	// 					ProbePort: to.Ptr[int32](80),
	// 					ProbeProtocol: to.Ptr(armservicefabricmanagedclusters.ProbeProtocolHTTP),
	// 					Protocol: to.Ptr(armservicefabricmanagedclusters.Protocol("http")),
	// 				},
	// 				{
	// 					BackendPort: to.Ptr[int32](443),
	// 					FrontendPort: to.Ptr[int32](443),
	// 					ProbePort: to.Ptr[int32](443),
	// 					ProbeProtocol: to.Ptr(armservicefabricmanagedclusters.ProbeProtocolHTTP),
	// 					Protocol: to.Ptr(armservicefabricmanagedclusters.Protocol("http")),
	// 				},
	// 				{
	// 					BackendPort: to.Ptr[int32](10000),
	// 					FrontendPort: to.Ptr[int32](10000),
	// 					ProbePort: to.Ptr[int32](10000),
	// 					ProbeProtocol: to.Ptr(armservicefabricmanagedclusters.ProbeProtocolHTTP),
	// 					Protocol: to.Ptr(armservicefabricmanagedclusters.ProtocolTCP),
	// 				},
	// 			},
	// 			NetworkSecurityRules: []*armservicefabricmanagedclusters.NetworkSecurityRule{
	// 				{
	// 					Name: to.Ptr("TestName"),
	// 					Description: to.Ptr("Test description"),
	// 					Access: to.Ptr(armservicefabricmanagedclusters.AccessAllow),
	// 					DestinationAddressPrefixes: []*string{
	// 						to.Ptr("*"),
	// 					},
	// 					DestinationPortRanges: []*string{
	// 						to.Ptr("*"),
	// 					},
	// 					Direction: to.Ptr(armservicefabricmanagedclusters.DirectionInbound),
	// 					Priority: to.Ptr[int32](1010),
	// 					SourceAddressPrefixes: []*string{
	// 						to.Ptr("*"),
	// 					},
	// 					SourcePortRanges: []*string{
	// 						to.Ptr("*"),
	// 					},
	// 					Protocol: to.Ptr(armservicefabricmanagedclusters.NsgProtocolTCP),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armservicefabricmanagedclusters.ManagedResourceProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armservicefabricmanagedclusters.SKU{
	// 			Name: to.Ptr(armservicefabricmanagedclusters.SKUNameStandard),
	// 		},
	// 		Tags: map[string]*string{
	// 			"a": to.Ptr("b"),
	// 		},
	// 	},
	// }
}
