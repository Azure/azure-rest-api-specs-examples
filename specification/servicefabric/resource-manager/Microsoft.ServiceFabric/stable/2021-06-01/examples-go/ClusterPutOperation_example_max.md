Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabric%2Farmservicefabric%2Fv0.3.0/sdk/resourcemanager/servicefabric/armservicefabric/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicefabric_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterPutOperation_example_max.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armservicefabric.Cluster{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armservicefabric.ClusterProperties{
				AddOnFeatures: []*armservicefabric.AddOnFeatures{
					armservicefabric.AddOnFeatures("RepairManager").ToPtr(),
					armservicefabric.AddOnFeatures("DnsService").ToPtr(),
					armservicefabric.AddOnFeatures("BackupRestoreService").ToPtr(),
					armservicefabric.AddOnFeatures("ResourceMonitorService").ToPtr()},
				ApplicationTypeVersionsCleanupPolicy: &armservicefabric.ApplicationTypeVersionsCleanupPolicy{
					MaxUnusedVersionsToKeep: to.Int64Ptr(2),
				},
				AzureActiveDirectory: &armservicefabric.AzureActiveDirectory{
					ClientApplication:  to.StringPtr("<client-application>"),
					ClusterApplication: to.StringPtr("<cluster-application>"),
					TenantID:           to.StringPtr("<tenant-id>"),
				},
				CertificateCommonNames: &armservicefabric.ServerCertificateCommonNames{
					CommonNames: []*armservicefabric.ServerCertificateCommonName{
						{
							CertificateCommonName:       to.StringPtr("<certificate-common-name>"),
							CertificateIssuerThumbprint: to.StringPtr("<certificate-issuer-thumbprint>"),
						}},
					X509StoreName: armservicefabric.StoreName("My").ToPtr(),
				},
				ClientCertificateCommonNames: []*armservicefabric.ClientCertificateCommonName{
					{
						CertificateCommonName:       to.StringPtr("<certificate-common-name>"),
						CertificateIssuerThumbprint: to.StringPtr("<certificate-issuer-thumbprint>"),
						IsAdmin:                     to.BoolPtr(true),
					}},
				ClientCertificateThumbprints: []*armservicefabric.ClientCertificateThumbprint{
					{
						CertificateThumbprint: to.StringPtr("<certificate-thumbprint>"),
						IsAdmin:               to.BoolPtr(true),
					}},
				ClusterCodeVersion: to.StringPtr("<cluster-code-version>"),
				DiagnosticsStorageAccountConfig: &armservicefabric.DiagnosticsStorageAccountConfig{
					BlobEndpoint:            to.StringPtr("<blob-endpoint>"),
					ProtectedAccountKeyName: to.StringPtr("<protected-account-key-name>"),
					QueueEndpoint:           to.StringPtr("<queue-endpoint>"),
					StorageAccountName:      to.StringPtr("<storage-account-name>"),
					TableEndpoint:           to.StringPtr("<table-endpoint>"),
				},
				EventStoreServiceEnabled: to.BoolPtr(true),
				FabricSettings: []*armservicefabric.SettingsSectionDescription{
					{
						Name: to.StringPtr("<name>"),
						Parameters: []*armservicefabric.SettingsParameterDescription{
							{
								Name:  to.StringPtr("<name>"),
								Value: to.StringPtr("<value>"),
							}},
					}},
				InfrastructureServiceManager: to.BoolPtr(true),
				ManagementEndpoint:           to.StringPtr("<management-endpoint>"),
				NodeTypes: []*armservicefabric.NodeTypeDescription{
					{
						Name: to.StringPtr("<name>"),
						ApplicationPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(30000),
							StartPort: to.Int32Ptr(20000),
						},
						ClientConnectionEndpointPort: to.Int32Ptr(19000),
						DurabilityLevel:              armservicefabric.DurabilityLevel("Silver").ToPtr(),
						EphemeralPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(64000),
							StartPort: to.Int32Ptr(49000),
						},
						HTTPGatewayEndpointPort:   to.Int32Ptr(19007),
						IsPrimary:                 to.BoolPtr(true),
						IsStateless:               to.BoolPtr(false),
						MultipleAvailabilityZones: to.BoolPtr(true),
						VMInstanceCount:           to.Int32Ptr(5),
					}},
				Notifications: []*armservicefabric.Notification{
					{
						IsEnabled:            to.BoolPtr(true),
						NotificationCategory: armservicefabric.NotificationCategory("WaveProgress").ToPtr(),
						NotificationLevel:    armservicefabric.NotificationLevel("Critical").ToPtr(),
						NotificationTargets: []*armservicefabric.NotificationTarget{
							{
								NotificationChannel: armservicefabric.NotificationChannel("EmailUser").ToPtr(),
								Receivers: []*string{
									to.StringPtr("****@microsoft.com"),
									to.StringPtr("****@microsoft.com")},
							},
							{
								NotificationChannel: armservicefabric.NotificationChannel("EmailSubscription").ToPtr(),
								Receivers: []*string{
									to.StringPtr("Owner"),
									to.StringPtr("AccountAdmin")},
							}},
					},
					{
						IsEnabled:            to.BoolPtr(true),
						NotificationCategory: armservicefabric.NotificationCategory("WaveProgress").ToPtr(),
						NotificationLevel:    armservicefabric.NotificationLevel("All").ToPtr(),
						NotificationTargets: []*armservicefabric.NotificationTarget{
							{
								NotificationChannel: armservicefabric.NotificationChannel("EmailUser").ToPtr(),
								Receivers: []*string{
									to.StringPtr("****@microsoft.com"),
									to.StringPtr("****@microsoft.com")},
							},
							{
								NotificationChannel: armservicefabric.NotificationChannel("EmailSubscription").ToPtr(),
								Receivers: []*string{
									to.StringPtr("Owner"),
									to.StringPtr("AccountAdmin")},
							}},
					}},
				ReliabilityLevel: armservicefabric.ReliabilityLevel("Platinum").ToPtr(),
				ReverseProxyCertificateCommonNames: &armservicefabric.ServerCertificateCommonNames{
					CommonNames: []*armservicefabric.ServerCertificateCommonName{
						{
							CertificateCommonName:       to.StringPtr("<certificate-common-name>"),
							CertificateIssuerThumbprint: to.StringPtr("<certificate-issuer-thumbprint>"),
						}},
					X509StoreName: armservicefabric.StoreName("My").ToPtr(),
				},
				SfZonalUpgradeMode: armservicefabric.SfZonalUpgradeMode("Hierarchical").ToPtr(),
				UpgradeDescription: &armservicefabric.ClusterUpgradePolicy{
					DeltaHealthPolicy: &armservicefabric.ClusterUpgradeDeltaHealthPolicy{
						ApplicationDeltaHealthPolicies: map[string]*armservicefabric.ApplicationDeltaHealthPolicy{
							"fabric:/myApp1": {
								DefaultServiceTypeDeltaHealthPolicy: &armservicefabric.ServiceTypeDeltaHealthPolicy{
									MaxPercentDeltaUnhealthyServices: to.Int32Ptr(0),
								},
								ServiceTypeDeltaHealthPolicies: map[string]*armservicefabric.ServiceTypeDeltaHealthPolicy{
									"myServiceType1": {
										MaxPercentDeltaUnhealthyServices: to.Int32Ptr(0),
									},
								},
							},
						},
						MaxPercentDeltaUnhealthyApplications:       to.Int32Ptr(0),
						MaxPercentDeltaUnhealthyNodes:              to.Int32Ptr(0),
						MaxPercentUpgradeDomainDeltaUnhealthyNodes: to.Int32Ptr(0),
					},
					ForceRestart:              to.BoolPtr(false),
					HealthCheckRetryTimeout:   to.StringPtr("<health-check-retry-timeout>"),
					HealthCheckStableDuration: to.StringPtr("<health-check-stable-duration>"),
					HealthCheckWaitDuration:   to.StringPtr("<health-check-wait-duration>"),
					HealthPolicy: &armservicefabric.ClusterHealthPolicy{
						ApplicationHealthPolicies: map[string]*armservicefabric.ApplicationHealthPolicy{
							"fabric:/myApp1": {
								DefaultServiceTypeHealthPolicy: &armservicefabric.ServiceTypeHealthPolicy{
									MaxPercentUnhealthyServices: to.Int32Ptr(0),
								},
								ServiceTypeHealthPolicies: map[string]*armservicefabric.ServiceTypeHealthPolicy{
									"myServiceType1": {
										MaxPercentUnhealthyServices: to.Int32Ptr(100),
									},
								},
							},
						},
						MaxPercentUnhealthyApplications: to.Int32Ptr(0),
						MaxPercentUnhealthyNodes:        to.Int32Ptr(0),
					},
					UpgradeDomainTimeout:          to.StringPtr("<upgrade-domain-timeout>"),
					UpgradeReplicaSetCheckTimeout: to.StringPtr("<upgrade-replica-set-check-timeout>"),
					UpgradeTimeout:                to.StringPtr("<upgrade-timeout>"),
				},
				UpgradeMode:                   armservicefabric.UpgradeMode("Manual").ToPtr(),
				UpgradePauseEndTimestampUTC:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-25T22:00:00Z"); return t }()),
				UpgradePauseStartTimestampUTC: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-21T22:00:00Z"); return t }()),
				UpgradeWave:                   armservicefabric.ClusterUpgradeCadence("Wave1").ToPtr(),
				VMImage:                       to.StringPtr("<vmimage>"),
				VmssZonalUpgradeMode:          armservicefabric.VmssZonalUpgradeMode("Parallel").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClustersClientCreateOrUpdateResult)
}
```
