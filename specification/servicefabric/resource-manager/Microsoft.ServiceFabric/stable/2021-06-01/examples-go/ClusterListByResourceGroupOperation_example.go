package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterListByResourceGroupOperation_example.json
func ExampleClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListByResourceGroupPager("resRg", nil)
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
		// page.ClusterListResult = armservicefabric.ClusterListResult{
		// 	Value: []*armservicefabric.Cluster{
		// 		{
		// 			Name: to.Ptr("myCluster"),
		// 			Type: to.Ptr("Microsoft.ServiceFabric/clusters"),
		// 			Etag: to.Ptr("W/\"636462502169240745\""),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicefabric.ClusterProperties{
		// 				AddOnFeatures: []*armservicefabric.AddOnFeatures{
		// 					to.Ptr(armservicefabric.AddOnFeaturesRepairManager),
		// 					to.Ptr(armservicefabric.AddOnFeaturesDNSService),
		// 					to.Ptr(armservicefabric.AddOnFeaturesBackupRestoreService),
		// 					to.Ptr(armservicefabric.AddOnFeaturesResourceMonitorService)},
		// 					AvailableClusterVersions: []*armservicefabric.ClusterVersionDetails{
		// 						{
		// 							CodeVersion: to.Ptr("6.1.480.9494"),
		// 							Environment: to.Ptr(armservicefabric.ClusterEnvironmentWindows),
		// 							SupportExpiryUTC: to.Ptr("2018-06-15T23:59:59.9999999"),
		// 					}},
		// 					AzureActiveDirectory: &armservicefabric.AzureActiveDirectory{
		// 						ClientApplication: to.Ptr("d151ad89-4bce-4ae8-b3d1-1dc79679fa75"),
		// 						ClusterApplication: to.Ptr("5886372e-7bf4-4878-a497-8098aba608ae"),
		// 						TenantID: to.Ptr("6abcc6a0-8666-43f1-87b8-172cf86a9f9c"),
		// 					},
		// 					CertificateCommonNames: &armservicefabric.ServerCertificateCommonNames{
		// 						CommonNames: []*armservicefabric.ServerCertificateCommonName{
		// 							{
		// 								CertificateCommonName: to.Ptr("abc.com"),
		// 								CertificateIssuerThumbprint: to.Ptr("12599211F8F14C90AFA9532AD79A6F2CA1C00622"),
		// 						}},
		// 						X509StoreName: to.Ptr(armservicefabric.StoreNameMy),
		// 					},
		// 					ClientCertificateCommonNames: []*armservicefabric.ClientCertificateCommonName{
		// 						{
		// 							CertificateCommonName: to.Ptr("abc.com"),
		// 							CertificateIssuerThumbprint: to.Ptr("5F3660C715EBBDA31DB1FFDCF508302348DE8E7A"),
		// 							IsAdmin: to.Ptr(true),
		// 					}},
		// 					ClientCertificateThumbprints: []*armservicefabric.ClientCertificateThumbprint{
		// 						{
		// 							CertificateThumbprint: to.Ptr("5F3660C715EBBDA31DB1FFDCF508302348DE8E7A"),
		// 							IsAdmin: to.Ptr(false),
		// 					}},
		// 					ClusterCodeVersion: to.Ptr("6.1.480.9494"),
		// 					ClusterEndpoint: to.Ptr("https://eastus.servicefabric.azure.com"),
		// 					ClusterID: to.Ptr("92584666-9889-4ae8-8d02-91902923d37f"),
		// 					ClusterState: to.Ptr(armservicefabric.ClusterStateWaitingForNodes),
		// 					DiagnosticsStorageAccountConfig: &armservicefabric.DiagnosticsStorageAccountConfig{
		// 						BlobEndpoint: to.Ptr("https://diag.blob.core.windows.net/"),
		// 						ProtectedAccountKeyName: to.Ptr("StorageAccountKey1"),
		// 						QueueEndpoint: to.Ptr("https://diag.queue.core.windows.net/"),
		// 						StorageAccountName: to.Ptr("diag"),
		// 						TableEndpoint: to.Ptr("https://diag.table.core.windows.net/"),
		// 					},
		// 					FabricSettings: []*armservicefabric.SettingsSectionDescription{
		// 						{
		// 							Name: to.Ptr("UpgradeService"),
		// 							Parameters: []*armservicefabric.SettingsParameterDescription{
		// 								{
		// 									Name: to.Ptr("AppPollIntervalInSeconds"),
		// 									Value: to.Ptr("60"),
		// 							}},
		// 					}},
		// 					ManagementEndpoint: to.Ptr("https://myCluster.eastus.cloudapp.azure.com:19080"),
		// 					NodeTypes: []*armservicefabric.NodeTypeDescription{
		// 						{
		// 							Name: to.Ptr("nt1vm"),
		// 							ApplicationPorts: &armservicefabric.EndpointRangeDescription{
		// 								EndPort: to.Ptr[int32](30000),
		// 								StartPort: to.Ptr[int32](20000),
		// 							},
		// 							ClientConnectionEndpointPort: to.Ptr[int32](19000),
		// 							DurabilityLevel: to.Ptr(armservicefabric.DurabilityLevelBronze),
		// 							EphemeralPorts: &armservicefabric.EndpointRangeDescription{
		// 								EndPort: to.Ptr[int32](64000),
		// 								StartPort: to.Ptr[int32](49000),
		// 							},
		// 							HTTPGatewayEndpointPort: to.Ptr[int32](19007),
		// 							IsPrimary: to.Ptr(true),
		// 							VMInstanceCount: to.Ptr[int32](5),
		// 					}},
		// 					ProvisioningState: to.Ptr(armservicefabric.ProvisioningStateSucceeded),
		// 					ReliabilityLevel: to.Ptr(armservicefabric.ReliabilityLevelSilver),
		// 					ReverseProxyCertificateCommonNames: &armservicefabric.ServerCertificateCommonNames{
		// 						CommonNames: []*armservicefabric.ServerCertificateCommonName{
		// 							{
		// 								CertificateCommonName: to.Ptr("abc.com"),
		// 								CertificateIssuerThumbprint: to.Ptr("12599211F8F14C90AFA9532AD79A6F2CA1C00622"),
		// 						}},
		// 						X509StoreName: to.Ptr(armservicefabric.StoreNameMy),
		// 					},
		// 					UpgradeDescription: &armservicefabric.ClusterUpgradePolicy{
		// 						DeltaHealthPolicy: &armservicefabric.ClusterUpgradeDeltaHealthPolicy{
		// 							ApplicationDeltaHealthPolicies: map[string]*armservicefabric.ApplicationDeltaHealthPolicy{
		// 								"fabric:/myApp1": &armservicefabric.ApplicationDeltaHealthPolicy{
		// 									DefaultServiceTypeDeltaHealthPolicy: &armservicefabric.ServiceTypeDeltaHealthPolicy{
		// 										MaxPercentDeltaUnhealthyServices: to.Ptr[int32](0),
		// 									},
		// 									ServiceTypeDeltaHealthPolicies: map[string]*armservicefabric.ServiceTypeDeltaHealthPolicy{
		// 										"myServiceType1": &armservicefabric.ServiceTypeDeltaHealthPolicy{
		// 											MaxPercentDeltaUnhealthyServices: to.Ptr[int32](0),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							MaxPercentDeltaUnhealthyApplications: to.Ptr[int32](0),
		// 							MaxPercentDeltaUnhealthyNodes: to.Ptr[int32](0),
		// 							MaxPercentUpgradeDomainDeltaUnhealthyNodes: to.Ptr[int32](0),
		// 						},
		// 						ForceRestart: to.Ptr(false),
		// 						HealthCheckRetryTimeout: to.Ptr("00:05:00"),
		// 						HealthCheckStableDuration: to.Ptr("00:00:30"),
		// 						HealthCheckWaitDuration: to.Ptr("00:00:30"),
		// 						HealthPolicy: &armservicefabric.ClusterHealthPolicy{
		// 							ApplicationHealthPolicies: map[string]*armservicefabric.ApplicationHealthPolicy{
		// 								"fabric:/myApp1": &armservicefabric.ApplicationHealthPolicy{
		// 									DefaultServiceTypeHealthPolicy: &armservicefabric.ServiceTypeHealthPolicy{
		// 										MaxPercentUnhealthyServices: to.Ptr[int32](0),
		// 									},
		// 									ServiceTypeHealthPolicies: map[string]*armservicefabric.ServiceTypeHealthPolicy{
		// 										"myServiceType1": &armservicefabric.ServiceTypeHealthPolicy{
		// 											MaxPercentUnhealthyServices: to.Ptr[int32](100),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							MaxPercentUnhealthyApplications: to.Ptr[int32](0),
		// 							MaxPercentUnhealthyNodes: to.Ptr[int32](0),
		// 						},
		// 						UpgradeDomainTimeout: to.Ptr("00:15:00"),
		// 						UpgradeReplicaSetCheckTimeout: to.Ptr("00:10:00"),
		// 						UpgradeTimeout: to.Ptr("01:00:00"),
		// 					},
		// 					UpgradeMode: to.Ptr(armservicefabric.UpgradeModeManual),
		// 					VMImage: to.Ptr("Windows"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myCluster2"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/clusters"),
		// 				Etag: to.Ptr("W/\"636462502164040075\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster2"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armservicefabric.ClusterProperties{
		// 					AddOnFeatures: []*armservicefabric.AddOnFeatures{
		// 						to.Ptr(armservicefabric.AddOnFeaturesRepairManager)},
		// 						AvailableClusterVersions: []*armservicefabric.ClusterVersionDetails{
		// 							{
		// 								CodeVersion: to.Ptr("6.1.187.1"),
		// 								Environment: to.Ptr(armservicefabric.ClusterEnvironmentLinux),
		// 								SupportExpiryUTC: to.Ptr("2018-06-15T23:59:59.9999999"),
		// 						}},
		// 						ClientCertificateCommonNames: []*armservicefabric.ClientCertificateCommonName{
		// 						},
		// 						ClientCertificateThumbprints: []*armservicefabric.ClientCertificateThumbprint{
		// 						},
		// 						ClusterCodeVersion: to.Ptr("6.1.187.1"),
		// 						ClusterEndpoint: to.Ptr("https://eastus.servicefabric.azure.com"),
		// 						ClusterID: to.Ptr("2747e469-b24e-4039-8a0a-46151419523f"),
		// 						ClusterState: to.Ptr(armservicefabric.ClusterStateWaitingForNodes),
		// 						DiagnosticsStorageAccountConfig: &armservicefabric.DiagnosticsStorageAccountConfig{
		// 							BlobEndpoint: to.Ptr("https://diag.blob.core.windows.net/"),
		// 							ProtectedAccountKeyName: to.Ptr("StorageAccountKey1"),
		// 							QueueEndpoint: to.Ptr("https://diag.queue.core.windows.net/"),
		// 							StorageAccountName: to.Ptr("diag"),
		// 							TableEndpoint: to.Ptr("https://diag.table.core.windows.net/"),
		// 						},
		// 						FabricSettings: []*armservicefabric.SettingsSectionDescription{
		// 							{
		// 								Name: to.Ptr("UpgradeService"),
		// 								Parameters: []*armservicefabric.SettingsParameterDescription{
		// 									{
		// 										Name: to.Ptr("AppPollIntervalInSeconds"),
		// 										Value: to.Ptr("60"),
		// 								}},
		// 						}},
		// 						ManagementEndpoint: to.Ptr("http://myCluster2.eastus.cloudapp.azure.com:19080"),
		// 						NodeTypes: []*armservicefabric.NodeTypeDescription{
		// 							{
		// 								Name: to.Ptr("nt1vm"),
		// 								ApplicationPorts: &armservicefabric.EndpointRangeDescription{
		// 									EndPort: to.Ptr[int32](30000),
		// 									StartPort: to.Ptr[int32](20000),
		// 								},
		// 								ClientConnectionEndpointPort: to.Ptr[int32](19000),
		// 								DurabilityLevel: to.Ptr(armservicefabric.DurabilityLevelBronze),
		// 								EphemeralPorts: &armservicefabric.EndpointRangeDescription{
		// 									EndPort: to.Ptr[int32](64000),
		// 									StartPort: to.Ptr[int32](49000),
		// 								},
		// 								HTTPGatewayEndpointPort: to.Ptr[int32](19007),
		// 								IsPrimary: to.Ptr(true),
		// 								VMInstanceCount: to.Ptr[int32](5),
		// 						}},
		// 						ProvisioningState: to.Ptr(armservicefabric.ProvisioningStateSucceeded),
		// 						ReliabilityLevel: to.Ptr(armservicefabric.ReliabilityLevelSilver),
		// 						UpgradeDescription: &armservicefabric.ClusterUpgradePolicy{
		// 							DeltaHealthPolicy: &armservicefabric.ClusterUpgradeDeltaHealthPolicy{
		// 								MaxPercentDeltaUnhealthyApplications: to.Ptr[int32](0),
		// 								MaxPercentDeltaUnhealthyNodes: to.Ptr[int32](0),
		// 								MaxPercentUpgradeDomainDeltaUnhealthyNodes: to.Ptr[int32](0),
		// 							},
		// 							ForceRestart: to.Ptr(false),
		// 							HealthCheckRetryTimeout: to.Ptr("00:05:00"),
		// 							HealthCheckStableDuration: to.Ptr("00:00:30"),
		// 							HealthCheckWaitDuration: to.Ptr("00:00:30"),
		// 							HealthPolicy: &armservicefabric.ClusterHealthPolicy{
		// 								MaxPercentUnhealthyApplications: to.Ptr[int32](0),
		// 								MaxPercentUnhealthyNodes: to.Ptr[int32](0),
		// 							},
		// 							UpgradeDomainTimeout: to.Ptr("00:15:00"),
		// 							UpgradeReplicaSetCheckTimeout: to.Ptr("00:10:00"),
		// 							UpgradeTimeout: to.Ptr("01:00:00"),
		// 						},
		// 						UpgradeMode: to.Ptr(armservicefabric.UpgradeModeManual),
		// 						VMImage: to.Ptr("Ubuntu"),
		// 					},
		// 			}},
		// 		}
	}
}
