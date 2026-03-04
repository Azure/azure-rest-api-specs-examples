package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v6"
)

// Generated from example definition: 2025-05-01/ListWebAppSlots.json
func ExampleWebAppsClient_NewListSlotsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebAppsClient().NewListSlotsPager("testrg123", "sitef6141", nil)
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
		// page = armappservice.WebAppsClientListSlotsResponse{
		// 	WebAppCollection: armappservice.WebAppCollection{
		// 		Value: []*armappservice.Site{
		// 			{
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/sitef6141/slots/staging"),
		// 				Name: to.Ptr("sitef6141/staging"),
		// 				Type: to.Ptr("Microsoft.Web/sites/slots"),
		// 				Kind: to.Ptr("app"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armappservice.SiteProperties{
		// 					State: to.Ptr("Running"),
		// 					HostNames: []*string{
		// 						to.Ptr("sitef6141-staging.azurewebsites.net"),
		// 					},
		// 					RepositorySiteName: to.Ptr("sitef6141"),
		// 					UsageState: to.Ptr(armappservice.UsageStateNormal),
		// 					Enabled: to.Ptr(true),
		// 					EnabledHostNames: []*string{
		// 						to.Ptr("sitef6141-staging.azurewebsites.net"),
		// 						to.Ptr("sitef6141-staging.scm.azurewebsites.net"),
		// 					},
		// 					AvailabilityState: to.Ptr(armappservice.SiteAvailabilityStateNormal),
		// 					HostNameSSLStates: []*armappservice.HostNameSSLState{
		// 						{
		// 							Name: to.Ptr("sitef6141-staging.azurewebsites.net"),
		// 							SSLState: to.Ptr(armappservice.SSLStateDisabled),
		// 							HostType: to.Ptr(armappservice.HostTypeStandard),
		// 						},
		// 						{
		// 							Name: to.Ptr("sitef6141-staging.scm.azurewebsites.net"),
		// 							SSLState: to.Ptr(armappservice.SSLStateDisabled),
		// 							HostType: to.Ptr(armappservice.HostTypeRepository),
		// 						},
		// 					},
		// 					ServerFarmID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/DefaultAsp"),
		// 					Reserved: to.Ptr(false),
		// 					IsXenon: to.Ptr(false),
		// 					HyperV: to.Ptr(false),
		// 					LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:33:11.641Z"); return t}()),
		// 					OutboundVnetRouting: &armappservice.OutboundVnetRouting{
		// 						AllTraffic: to.Ptr(false),
		// 						ApplicationTraffic: to.Ptr(false),
		// 						ContentShareTraffic: to.Ptr(false),
		// 						ImagePullTraffic: to.Ptr(false),
		// 						BackupRestoreTraffic: to.Ptr(false),
		// 					},
		// 					SiteConfig: &armappservice.SiteConfig{
		// 						NumberOfWorkers: to.Ptr[int32](1),
		// 						DefaultDocuments: []*string{
		// 							to.Ptr("Default.htm"),
		// 							to.Ptr("Default.html"),
		// 							to.Ptr("Default.asp"),
		// 							to.Ptr("index.htm"),
		// 							to.Ptr("index.html"),
		// 							to.Ptr("iisstart.htm"),
		// 							to.Ptr("default.aspx"),
		// 							to.Ptr("index.php"),
		// 							to.Ptr("hostingstart.html"),
		// 						},
		// 						NetFrameworkVersion: to.Ptr("v4.0"),
		// 						PhpVersion: to.Ptr("5.6"),
		// 						PythonVersion: to.Ptr(""),
		// 						NodeVersion: to.Ptr(""),
		// 						PowerShellVersion: to.Ptr(""),
		// 						LinuxFxVersion: to.Ptr(""),
		// 						RequestTracingEnabled: to.Ptr(false),
		// 						RemoteDebuggingEnabled: to.Ptr(false),
		// 						HTTPLoggingEnabled: to.Ptr(false),
		// 						AcrUseManagedIdentityCreds: to.Ptr(false),
		// 						LogsDirectorySizeLimit: to.Ptr[int32](35),
		// 						DetailedErrorLoggingEnabled: to.Ptr(false),
		// 						Use32BitWorkerProcess: to.Ptr(true),
		// 						WebSocketsEnabled: to.Ptr(false),
		// 						AlwaysOn: to.Ptr(false),
		// 						AppCommandLine: to.Ptr(""),
		// 						ManagedPipelineMode: to.Ptr(armappservice.ManagedPipelineModeIntegrated),
		// 						VirtualApplications: []*armappservice.VirtualApplication{
		// 							{
		// 								VirtualPath: to.Ptr("/"),
		// 								PhysicalPath: to.Ptr("site\\wwwroot"),
		// 								PreloadEnabled: to.Ptr(false),
		// 							},
		// 						},
		// 						LoadBalancing: to.Ptr(armappservice.SiteLoadBalancingLeastRequests),
		// 						AutoHealEnabled: to.Ptr(false),
		// 						VnetName: to.Ptr(""),
		// 						VnetRouteAllEnabled: to.Ptr(false),
		// 						VnetPrivatePortsCount: to.Ptr[int32](0),
		// 						Http20Enabled: to.Ptr(false),
		// 						MinTLSVersion: to.Ptr(armappservice.SupportedTLSVersionsOne2),
		// 						ScmMinTLSVersion: to.Ptr(armappservice.SupportedTLSVersionsOne2),
		// 						FtpsState: to.Ptr(armappservice.FtpsStateAllAllowed),
		// 						FunctionAppScaleLimit: to.Ptr[int32](0),
		// 						FunctionsRuntimeScaleMonitoringEnabled: to.Ptr(false),
		// 						MinimumElasticInstanceCount: to.Ptr[int32](0),
		// 						AzureStorageAccounts: map[string]*armappservice.AzureStorageInfoValue{
		// 						},
		// 					},
		// 					ScmSiteAlsoStopped: to.Ptr(false),
		// 					ClientAffinityEnabled: to.Ptr(true),
		// 					ClientCertEnabled: to.Ptr(false),
		// 					ClientCertMode: to.Ptr(armappservice.ClientCertModeRequired),
		// 					HostNamesDisabled: to.Ptr(false),
		// 					CustomDomainVerificationID: to.Ptr("7F3BB652450EF7AD0B6AA33064586E4A9CE823A46DF0B2EC6443A68086E84153"),
		// 					OutboundIPAddresses: to.Ptr("70.37.102.201,20.225.43.144"),
		// 					PossibleOutboundIPAddresses: to.Ptr("70.37.102.201,20.225.43.144,20.225.184.122,20.225.184.188"),
		// 					ContainerSize: to.Ptr[int32](0),
		// 					DailyMemoryTimeQuota: to.Ptr[int32](0),
		// 					ResourceGroup: to.Ptr("testrg123"),
		// 					DefaultHostName: to.Ptr("sitef6141-staging.azurewebsites.net"),
		// 					HTTPSOnly: to.Ptr(false),
		// 					RedundancyMode: to.Ptr(armappservice.RedundancyModeNone),
		// 					StorageAccountRequired: to.Ptr(false),
		// 					KeyVaultReferenceIdentity: to.Ptr("SystemAssigned"),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/sitef6141/slots/qa"),
		// 				Name: to.Ptr("sitef6141/qa"),
		// 				Type: to.Ptr("Microsoft.Web/sites/slots"),
		// 				Kind: to.Ptr("app"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armappservice.SiteProperties{
		// 					State: to.Ptr("Running"),
		// 					HostNames: []*string{
		// 						to.Ptr("sitef6141-qa.azurewebsites.net"),
		// 					},
		// 					RepositorySiteName: to.Ptr("sitef6141"),
		// 					UsageState: to.Ptr(armappservice.UsageStateNormal),
		// 					Enabled: to.Ptr(true),
		// 					EnabledHostNames: []*string{
		// 						to.Ptr("sitef6141-staging.azurewebsites.net"),
		// 						to.Ptr("sitef6141-staging.scm.azurewebsites.net"),
		// 					},
		// 					AvailabilityState: to.Ptr(armappservice.SiteAvailabilityStateNormal),
		// 					HostNameSSLStates: []*armappservice.HostNameSSLState{
		// 						{
		// 							Name: to.Ptr("sitef6141-qa.azurewebsites.net"),
		// 							SSLState: to.Ptr(armappservice.SSLStateDisabled),
		// 							HostType: to.Ptr(armappservice.HostTypeStandard),
		// 						},
		// 						{
		// 							Name: to.Ptr("sitef6141-qa.scm.azurewebsites.net"),
		// 							SSLState: to.Ptr(armappservice.SSLStateDisabled),
		// 							HostType: to.Ptr(armappservice.HostTypeRepository),
		// 						},
		// 					},
		// 					ServerFarmID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/DefaultAsp"),
		// 					Reserved: to.Ptr(false),
		// 					IsXenon: to.Ptr(false),
		// 					HyperV: to.Ptr(false),
		// 					LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:33:11.641Z"); return t}()),
		// 					OutboundVnetRouting: &armappservice.OutboundVnetRouting{
		// 						AllTraffic: to.Ptr(false),
		// 						ApplicationTraffic: to.Ptr(false),
		// 						ContentShareTraffic: to.Ptr(false),
		// 						ImagePullTraffic: to.Ptr(false),
		// 						BackupRestoreTraffic: to.Ptr(false),
		// 					},
		// 					SiteConfig: &armappservice.SiteConfig{
		// 						NumberOfWorkers: to.Ptr[int32](1),
		// 						DefaultDocuments: []*string{
		// 							to.Ptr("Default.htm"),
		// 							to.Ptr("Default.html"),
		// 							to.Ptr("Default.asp"),
		// 							to.Ptr("index.htm"),
		// 							to.Ptr("index.html"),
		// 							to.Ptr("iisstart.htm"),
		// 							to.Ptr("default.aspx"),
		// 							to.Ptr("index.php"),
		// 							to.Ptr("hostingstart.html"),
		// 						},
		// 						NetFrameworkVersion: to.Ptr("v4.0"),
		// 						PhpVersion: to.Ptr("5.6"),
		// 						PythonVersion: to.Ptr(""),
		// 						NodeVersion: to.Ptr(""),
		// 						PowerShellVersion: to.Ptr(""),
		// 						LinuxFxVersion: to.Ptr(""),
		// 						RequestTracingEnabled: to.Ptr(false),
		// 						RemoteDebuggingEnabled: to.Ptr(false),
		// 						HTTPLoggingEnabled: to.Ptr(false),
		// 						AcrUseManagedIdentityCreds: to.Ptr(false),
		// 						LogsDirectorySizeLimit: to.Ptr[int32](35),
		// 						DetailedErrorLoggingEnabled: to.Ptr(false),
		// 						Use32BitWorkerProcess: to.Ptr(true),
		// 						WebSocketsEnabled: to.Ptr(false),
		// 						AlwaysOn: to.Ptr(false),
		// 						AppCommandLine: to.Ptr(""),
		// 						ManagedPipelineMode: to.Ptr(armappservice.ManagedPipelineModeIntegrated),
		// 						VirtualApplications: []*armappservice.VirtualApplication{
		// 							{
		// 								VirtualPath: to.Ptr("/"),
		// 								PhysicalPath: to.Ptr("site\\wwwroot"),
		// 								PreloadEnabled: to.Ptr(false),
		// 							},
		// 						},
		// 						LoadBalancing: to.Ptr(armappservice.SiteLoadBalancingLeastRequests),
		// 						AutoHealEnabled: to.Ptr(false),
		// 						VnetName: to.Ptr(""),
		// 						VnetRouteAllEnabled: to.Ptr(false),
		// 						VnetPrivatePortsCount: to.Ptr[int32](0),
		// 						Http20Enabled: to.Ptr(false),
		// 						MinTLSVersion: to.Ptr(armappservice.SupportedTLSVersionsOne2),
		// 						ScmMinTLSVersion: to.Ptr(armappservice.SupportedTLSVersionsOne2),
		// 						FtpsState: to.Ptr(armappservice.FtpsStateAllAllowed),
		// 						FunctionAppScaleLimit: to.Ptr[int32](0),
		// 						FunctionsRuntimeScaleMonitoringEnabled: to.Ptr(false),
		// 						MinimumElasticInstanceCount: to.Ptr[int32](0),
		// 						AzureStorageAccounts: map[string]*armappservice.AzureStorageInfoValue{
		// 						},
		// 					},
		// 					ScmSiteAlsoStopped: to.Ptr(false),
		// 					ClientAffinityEnabled: to.Ptr(true),
		// 					ClientCertEnabled: to.Ptr(false),
		// 					ClientCertMode: to.Ptr(armappservice.ClientCertModeRequired),
		// 					HostNamesDisabled: to.Ptr(false),
		// 					CustomDomainVerificationID: to.Ptr("7F3BB652450EF7AD0B6AA33064586E4A9CE823A46DF0B2EC6443A68086E84153"),
		// 					OutboundIPAddresses: to.Ptr("70.37.102.201,20.225.43.144"),
		// 					PossibleOutboundIPAddresses: to.Ptr("70.37.102.201,20.225.43.144,20.225.184.122,20.225.184.188"),
		// 					ContainerSize: to.Ptr[int32](0),
		// 					DailyMemoryTimeQuota: to.Ptr[int32](0),
		// 					ResourceGroup: to.Ptr("testrg123"),
		// 					DefaultHostName: to.Ptr("sitef6141-qa.azurewebsites.net"),
		// 					HTTPSOnly: to.Ptr(false),
		// 					RedundancyMode: to.Ptr(armappservice.RedundancyModeNone),
		// 					StorageAccountRequired: to.Ptr(false),
		// 					KeyVaultReferenceIdentity: to.Ptr("SystemAssigned"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
