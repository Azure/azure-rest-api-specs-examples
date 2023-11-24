package armservicefabric_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterPutOperation_example_max.json
func ExampleClustersClient_BeginCreateOrUpdate_putAClusterWithMaximumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", armservicefabric.Cluster{
		Location: to.Ptr("eastus"),
		Tags:     map[string]*string{},
		Properties: &armservicefabric.ClusterProperties{
			AddOnFeatures: []*armservicefabric.AddOnFeatures{
				to.Ptr(armservicefabric.AddOnFeaturesRepairManager),
				to.Ptr(armservicefabric.AddOnFeaturesDNSService),
				to.Ptr(armservicefabric.AddOnFeaturesBackupRestoreService),
				to.Ptr(armservicefabric.AddOnFeaturesResourceMonitorService)},
			ApplicationTypeVersionsCleanupPolicy: &armservicefabric.ApplicationTypeVersionsCleanupPolicy{
				MaxUnusedVersionsToKeep: to.Ptr[int64](2),
			},
			AzureActiveDirectory: &armservicefabric.AzureActiveDirectory{
				ClientApplication:  to.Ptr("d151ad89-4bce-4ae8-b3d1-1dc79679fa75"),
				ClusterApplication: to.Ptr("5886372e-7bf4-4878-a497-8098aba608ae"),
				TenantID:           to.Ptr("6abcc6a0-8666-43f1-87b8-172cf86a9f9c"),
			},
			CertificateCommonNames: &armservicefabric.ServerCertificateCommonNames{
				CommonNames: []*armservicefabric.ServerCertificateCommonName{
					{
						CertificateCommonName:       to.Ptr("abc.com"),
						CertificateIssuerThumbprint: to.Ptr("12599211F8F14C90AFA9532AD79A6F2CA1C00622"),
					}},
				X509StoreName: to.Ptr(armservicefabric.StoreNameMy),
			},
			ClientCertificateCommonNames: []*armservicefabric.ClientCertificateCommonName{
				{
					CertificateCommonName:       to.Ptr("abc.com"),
					CertificateIssuerThumbprint: to.Ptr("5F3660C715EBBDA31DB1FFDCF508302348DE8E7A"),
					IsAdmin:                     to.Ptr(true),
				}},
			ClientCertificateThumbprints: []*armservicefabric.ClientCertificateThumbprint{
				{
					CertificateThumbprint: to.Ptr("5F3660C715EBBDA31DB1FFDCF508302348DE8E7A"),
					IsAdmin:               to.Ptr(true),
				}},
			ClusterCodeVersion: to.Ptr("7.0.470.9590"),
			DiagnosticsStorageAccountConfig: &armservicefabric.DiagnosticsStorageAccountConfig{
				BlobEndpoint:            to.Ptr("https://diag.blob.core.windows.net/"),
				ProtectedAccountKeyName: to.Ptr("StorageAccountKey1"),
				QueueEndpoint:           to.Ptr("https://diag.queue.core.windows.net/"),
				StorageAccountName:      to.Ptr("diag"),
				TableEndpoint:           to.Ptr("https://diag.table.core.windows.net/"),
			},
			EventStoreServiceEnabled: to.Ptr(true),
			FabricSettings: []*armservicefabric.SettingsSectionDescription{
				{
					Name: to.Ptr("UpgradeService"),
					Parameters: []*armservicefabric.SettingsParameterDescription{
						{
							Name:  to.Ptr("AppPollIntervalInSeconds"),
							Value: to.Ptr("60"),
						}},
				}},
			InfrastructureServiceManager: to.Ptr(true),
			ManagementEndpoint:           to.Ptr("https://myCluster.eastus.cloudapp.azure.com:19080"),
			NodeTypes: []*armservicefabric.NodeTypeDescription{
				{
					Name: to.Ptr("nt1vm"),
					ApplicationPorts: &armservicefabric.EndpointRangeDescription{
						EndPort:   to.Ptr[int32](30000),
						StartPort: to.Ptr[int32](20000),
					},
					ClientConnectionEndpointPort: to.Ptr[int32](19000),
					DurabilityLevel:              to.Ptr(armservicefabric.DurabilityLevelSilver),
					EphemeralPorts: &armservicefabric.EndpointRangeDescription{
						EndPort:   to.Ptr[int32](64000),
						StartPort: to.Ptr[int32](49000),
					},
					HTTPGatewayEndpointPort:   to.Ptr[int32](19007),
					IsPrimary:                 to.Ptr(true),
					IsStateless:               to.Ptr(false),
					MultipleAvailabilityZones: to.Ptr(true),
					VMInstanceCount:           to.Ptr[int32](5),
				}},
			Notifications: []*armservicefabric.Notification{
				{
					IsEnabled:            to.Ptr(true),
					NotificationCategory: to.Ptr(armservicefabric.NotificationCategoryWaveProgress),
					NotificationLevel:    to.Ptr(armservicefabric.NotificationLevelCritical),
					NotificationTargets: []*armservicefabric.NotificationTarget{
						{
							NotificationChannel: to.Ptr(armservicefabric.NotificationChannelEmailUser),
							Receivers: []*string{
								to.Ptr("****@microsoft.com"),
								to.Ptr("****@microsoft.com")},
						},
						{
							NotificationChannel: to.Ptr(armservicefabric.NotificationChannelEmailSubscription),
							Receivers: []*string{
								to.Ptr("Owner"),
								to.Ptr("AccountAdmin")},
						}},
				},
				{
					IsEnabled:            to.Ptr(true),
					NotificationCategory: to.Ptr(armservicefabric.NotificationCategoryWaveProgress),
					NotificationLevel:    to.Ptr(armservicefabric.NotificationLevelAll),
					NotificationTargets: []*armservicefabric.NotificationTarget{
						{
							NotificationChannel: to.Ptr(armservicefabric.NotificationChannelEmailUser),
							Receivers: []*string{
								to.Ptr("****@microsoft.com"),
								to.Ptr("****@microsoft.com")},
						},
						{
							NotificationChannel: to.Ptr(armservicefabric.NotificationChannelEmailSubscription),
							Receivers: []*string{
								to.Ptr("Owner"),
								to.Ptr("AccountAdmin")},
						}},
				}},
			ReliabilityLevel: to.Ptr(armservicefabric.ReliabilityLevelPlatinum),
			ReverseProxyCertificateCommonNames: &armservicefabric.ServerCertificateCommonNames{
				CommonNames: []*armservicefabric.ServerCertificateCommonName{
					{
						CertificateCommonName:       to.Ptr("abc.com"),
						CertificateIssuerThumbprint: to.Ptr("12599211F8F14C90AFA9532AD79A6F2CA1C00622"),
					}},
				X509StoreName: to.Ptr(armservicefabric.StoreNameMy),
			},
			SfZonalUpgradeMode: to.Ptr(armservicefabric.SfZonalUpgradeModeHierarchical),
			UpgradeDescription: &armservicefabric.ClusterUpgradePolicy{
				DeltaHealthPolicy: &armservicefabric.ClusterUpgradeDeltaHealthPolicy{
					ApplicationDeltaHealthPolicies: map[string]*armservicefabric.ApplicationDeltaHealthPolicy{
						"fabric:/myApp1": {
							DefaultServiceTypeDeltaHealthPolicy: &armservicefabric.ServiceTypeDeltaHealthPolicy{
								MaxPercentDeltaUnhealthyServices: to.Ptr[int32](0),
							},
							ServiceTypeDeltaHealthPolicies: map[string]*armservicefabric.ServiceTypeDeltaHealthPolicy{
								"myServiceType1": {
									MaxPercentDeltaUnhealthyServices: to.Ptr[int32](0),
								},
							},
						},
					},
					MaxPercentDeltaUnhealthyApplications:       to.Ptr[int32](0),
					MaxPercentDeltaUnhealthyNodes:              to.Ptr[int32](0),
					MaxPercentUpgradeDomainDeltaUnhealthyNodes: to.Ptr[int32](0),
				},
				ForceRestart:              to.Ptr(false),
				HealthCheckRetryTimeout:   to.Ptr("00:05:00"),
				HealthCheckStableDuration: to.Ptr("00:00:30"),
				HealthCheckWaitDuration:   to.Ptr("00:00:30"),
				HealthPolicy: &armservicefabric.ClusterHealthPolicy{
					ApplicationHealthPolicies: map[string]*armservicefabric.ApplicationHealthPolicy{
						"fabric:/myApp1": {
							DefaultServiceTypeHealthPolicy: &armservicefabric.ServiceTypeHealthPolicy{
								MaxPercentUnhealthyServices: to.Ptr[int32](0),
							},
							ServiceTypeHealthPolicies: map[string]*armservicefabric.ServiceTypeHealthPolicy{
								"myServiceType1": {
									MaxPercentUnhealthyServices: to.Ptr[int32](100),
								},
							},
						},
					},
					MaxPercentUnhealthyApplications: to.Ptr[int32](0),
					MaxPercentUnhealthyNodes:        to.Ptr[int32](0),
				},
				UpgradeDomainTimeout:          to.Ptr("00:15:00"),
				UpgradeReplicaSetCheckTimeout: to.Ptr("00:10:00"),
				UpgradeTimeout:                to.Ptr("01:00:00"),
			},
			UpgradeMode:                   to.Ptr(armservicefabric.UpgradeModeManual),
			UpgradePauseEndTimestampUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-25T22:00:00.000Z"); return t }()),
			UpgradePauseStartTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-21T22:00:00.000Z"); return t }()),
			UpgradeWave:                   to.Ptr(armservicefabric.ClusterUpgradeCadenceWave1),
			VMImage:                       to.Ptr("Windows"),
			VmssZonalUpgradeMode:          to.Ptr(armservicefabric.VmssZonalUpgradeModeParallel),
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
	// res.Cluster = armservicefabric.Cluster{
	// 	Name: to.Ptr("myCluster"),
	// 	Type: to.Ptr("Microsoft.ServiceFabric/clusters"),
	// 	Etag: to.Ptr("W/\"636462502169240745\""),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armservicefabric.ClusterProperties{
	// 		AddOnFeatures: []*armservicefabric.AddOnFeatures{
	// 			to.Ptr(armservicefabric.AddOnFeaturesRepairManager),
	// 			to.Ptr(armservicefabric.AddOnFeaturesDNSService),
	// 			to.Ptr(armservicefabric.AddOnFeaturesBackupRestoreService),
	// 			to.Ptr(armservicefabric.AddOnFeaturesResourceMonitorService)},
	// 			ApplicationTypeVersionsCleanupPolicy: &armservicefabric.ApplicationTypeVersionsCleanupPolicy{
	// 				MaxUnusedVersionsToKeep: to.Ptr[int64](2),
	// 			},
	// 			AvailableClusterVersions: []*armservicefabric.ClusterVersionDetails{
	// 				{
	// 					CodeVersion: to.Ptr("7.0.470.9590"),
	// 					Environment: to.Ptr(armservicefabric.ClusterEnvironmentWindows),
	// 					SupportExpiryUTC: to.Ptr("2018-06-15T23:59:59.9999999"),
	// 			}},
	// 			AzureActiveDirectory: &armservicefabric.AzureActiveDirectory{
	// 				ClientApplication: to.Ptr("d151ad89-4bce-4ae8-b3d1-1dc79679fa75"),
	// 				ClusterApplication: to.Ptr("5886372e-7bf4-4878-a497-8098aba608ae"),
	// 				TenantID: to.Ptr("6abcc6a0-8666-43f1-87b8-172cf86a9f9c"),
	// 			},
	// 			CertificateCommonNames: &armservicefabric.ServerCertificateCommonNames{
	// 				CommonNames: []*armservicefabric.ServerCertificateCommonName{
	// 					{
	// 						CertificateCommonName: to.Ptr("abc.com"),
	// 						CertificateIssuerThumbprint: to.Ptr("12599211F8F14C90AFA9532AD79A6F2CA1C00622"),
	// 				}},
	// 				X509StoreName: to.Ptr(armservicefabric.StoreNameMy),
	// 			},
	// 			ClientCertificateCommonNames: []*armservicefabric.ClientCertificateCommonName{
	// 				{
	// 					CertificateCommonName: to.Ptr("abc.com"),
	// 					CertificateIssuerThumbprint: to.Ptr("5F3660C715EBBDA31DB1FFDCF508302348DE8E7A"),
	// 					IsAdmin: to.Ptr(true),
	// 			}},
	// 			ClientCertificateThumbprints: []*armservicefabric.ClientCertificateThumbprint{
	// 				{
	// 					CertificateThumbprint: to.Ptr("5F3660C715EBBDA31DB1FFDCF508302348DE8E7A"),
	// 					IsAdmin: to.Ptr(false),
	// 			}},
	// 			ClusterCodeVersion: to.Ptr("7.0.470.9590"),
	// 			ClusterEndpoint: to.Ptr("https://eastus.servicefabric.azure.com"),
	// 			ClusterID: to.Ptr("92584666-9889-4ae8-8d02-91902923d37f"),
	// 			ClusterState: to.Ptr(armservicefabric.ClusterStateWaitingForNodes),
	// 			DiagnosticsStorageAccountConfig: &armservicefabric.DiagnosticsStorageAccountConfig{
	// 				BlobEndpoint: to.Ptr("https://diag.blob.core.windows.net/"),
	// 				ProtectedAccountKeyName: to.Ptr("StorageAccountKey1"),
	// 				QueueEndpoint: to.Ptr("https://diag.queue.core.windows.net/"),
	// 				StorageAccountName: to.Ptr("diag"),
	// 				TableEndpoint: to.Ptr("https://diag.table.core.windows.net/"),
	// 			},
	// 			EventStoreServiceEnabled: to.Ptr(true),
	// 			FabricSettings: []*armservicefabric.SettingsSectionDescription{
	// 				{
	// 					Name: to.Ptr("UpgradeService"),
	// 					Parameters: []*armservicefabric.SettingsParameterDescription{
	// 						{
	// 							Name: to.Ptr("AppPollIntervalInSeconds"),
	// 							Value: to.Ptr("60"),
	// 					}},
	// 			}},
	// 			InfrastructureServiceManager: to.Ptr(true),
	// 			ManagementEndpoint: to.Ptr("https://myCluster.eastus.cloudapp.azure.com:19080"),
	// 			NodeTypes: []*armservicefabric.NodeTypeDescription{
	// 				{
	// 					Name: to.Ptr("nt1vm"),
	// 					ApplicationPorts: &armservicefabric.EndpointRangeDescription{
	// 						EndPort: to.Ptr[int32](30000),
	// 						StartPort: to.Ptr[int32](20000),
	// 					},
	// 					ClientConnectionEndpointPort: to.Ptr[int32](19000),
	// 					DurabilityLevel: to.Ptr(armservicefabric.DurabilityLevelSilver),
	// 					EphemeralPorts: &armservicefabric.EndpointRangeDescription{
	// 						EndPort: to.Ptr[int32](64000),
	// 						StartPort: to.Ptr[int32](49000),
	// 					},
	// 					HTTPGatewayEndpointPort: to.Ptr[int32](19007),
	// 					IsPrimary: to.Ptr(true),
	// 					IsStateless: to.Ptr(false),
	// 					MultipleAvailabilityZones: to.Ptr(true),
	// 					VMInstanceCount: to.Ptr[int32](5),
	// 			}},
	// 			Notifications: []*armservicefabric.Notification{
	// 				{
	// 					IsEnabled: to.Ptr(true),
	// 					NotificationCategory: to.Ptr(armservicefabric.NotificationCategoryWaveProgress),
	// 					NotificationLevel: to.Ptr(armservicefabric.NotificationLevelCritical),
	// 					NotificationTargets: []*armservicefabric.NotificationTarget{
	// 						{
	// 							NotificationChannel: to.Ptr(armservicefabric.NotificationChannelEmailUser),
	// 							Receivers: []*string{
	// 								to.Ptr("****@microsoft.com"),
	// 								to.Ptr("****@microsoft.com")},
	// 							},
	// 							{
	// 								NotificationChannel: to.Ptr(armservicefabric.NotificationChannelEmailSubscription),
	// 								Receivers: []*string{
	// 									to.Ptr("Owner"),
	// 									to.Ptr("AccountAdmin")},
	// 							}},
	// 						},
	// 						{
	// 							IsEnabled: to.Ptr(true),
	// 							NotificationCategory: to.Ptr(armservicefabric.NotificationCategoryWaveProgress),
	// 							NotificationLevel: to.Ptr(armservicefabric.NotificationLevelAll),
	// 							NotificationTargets: []*armservicefabric.NotificationTarget{
	// 								{
	// 									NotificationChannel: to.Ptr(armservicefabric.NotificationChannelEmailUser),
	// 									Receivers: []*string{
	// 										to.Ptr("****@microsoft.com"),
	// 										to.Ptr("****@microsoft.com")},
	// 									},
	// 									{
	// 										NotificationChannel: to.Ptr(armservicefabric.NotificationChannelEmailSubscription),
	// 										Receivers: []*string{
	// 											to.Ptr("Owner"),
	// 											to.Ptr("AccountAdmin")},
	// 									}},
	// 							}},
	// 							ProvisioningState: to.Ptr(armservicefabric.ProvisioningStateSucceeded),
	// 							ReliabilityLevel: to.Ptr(armservicefabric.ReliabilityLevelPlatinum),
	// 							ReverseProxyCertificateCommonNames: &armservicefabric.ServerCertificateCommonNames{
	// 								CommonNames: []*armservicefabric.ServerCertificateCommonName{
	// 									{
	// 										CertificateCommonName: to.Ptr("abc.com"),
	// 										CertificateIssuerThumbprint: to.Ptr("12599211F8F14C90AFA9532AD79A6F2CA1C00622"),
	// 								}},
	// 								X509StoreName: to.Ptr(armservicefabric.StoreNameMy),
	// 							},
	// 							SfZonalUpgradeMode: to.Ptr(armservicefabric.SfZonalUpgradeModeHierarchical),
	// 							UpgradeDescription: &armservicefabric.ClusterUpgradePolicy{
	// 								DeltaHealthPolicy: &armservicefabric.ClusterUpgradeDeltaHealthPolicy{
	// 									ApplicationDeltaHealthPolicies: map[string]*armservicefabric.ApplicationDeltaHealthPolicy{
	// 										"fabric:/myApp1": &armservicefabric.ApplicationDeltaHealthPolicy{
	// 											DefaultServiceTypeDeltaHealthPolicy: &armservicefabric.ServiceTypeDeltaHealthPolicy{
	// 												MaxPercentDeltaUnhealthyServices: to.Ptr[int32](0),
	// 											},
	// 											ServiceTypeDeltaHealthPolicies: map[string]*armservicefabric.ServiceTypeDeltaHealthPolicy{
	// 												"myServiceType1": &armservicefabric.ServiceTypeDeltaHealthPolicy{
	// 													MaxPercentDeltaUnhealthyServices: to.Ptr[int32](0),
	// 												},
	// 											},
	// 										},
	// 									},
	// 									MaxPercentDeltaUnhealthyApplications: to.Ptr[int32](0),
	// 									MaxPercentDeltaUnhealthyNodes: to.Ptr[int32](0),
	// 									MaxPercentUpgradeDomainDeltaUnhealthyNodes: to.Ptr[int32](0),
	// 								},
	// 								ForceRestart: to.Ptr(true),
	// 								HealthCheckRetryTimeout: to.Ptr("00:05:00"),
	// 								HealthCheckStableDuration: to.Ptr("00:00:30"),
	// 								HealthCheckWaitDuration: to.Ptr("00:00:30"),
	// 								HealthPolicy: &armservicefabric.ClusterHealthPolicy{
	// 									ApplicationHealthPolicies: map[string]*armservicefabric.ApplicationHealthPolicy{
	// 										"fabric:/myApp1": &armservicefabric.ApplicationHealthPolicy{
	// 											DefaultServiceTypeHealthPolicy: &armservicefabric.ServiceTypeHealthPolicy{
	// 												MaxPercentUnhealthyServices: to.Ptr[int32](0),
	// 											},
	// 											ServiceTypeHealthPolicies: map[string]*armservicefabric.ServiceTypeHealthPolicy{
	// 												"myServiceType1": &armservicefabric.ServiceTypeHealthPolicy{
	// 													MaxPercentUnhealthyServices: to.Ptr[int32](100),
	// 												},
	// 											},
	// 										},
	// 									},
	// 									MaxPercentUnhealthyApplications: to.Ptr[int32](0),
	// 									MaxPercentUnhealthyNodes: to.Ptr[int32](0),
	// 								},
	// 								UpgradeDomainTimeout: to.Ptr("00:15:00"),
	// 								UpgradeReplicaSetCheckTimeout: to.Ptr("00:10:00"),
	// 								UpgradeTimeout: to.Ptr("00:15:00"),
	// 							},
	// 							UpgradeMode: to.Ptr(armservicefabric.UpgradeModeManual),
	// 							UpgradePauseEndTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-25T22:00:00.000Z"); return t}()),
	// 							UpgradePauseStartTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-21T22:00:00.000Z"); return t}()),
	// 							UpgradeWave: to.Ptr(armservicefabric.ClusterUpgradeCadenceWave1),
	// 							VMImage: to.Ptr("Windows"),
	// 							VmssZonalUpgradeMode: to.Ptr(armservicefabric.VmssZonalUpgradeModeParallel),
	// 						},
	// 					}
}
