package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/UpdateWebAppSlot.json
func ExampleWebAppsClient_UpdateSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebAppsClient().UpdateSlot(ctx, "testrg123", "sitef6141", "staging", armappservice.SitePatchResource{
		Properties: &armappservice.SitePatchResourceProperties{
			ServerFarmID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/DefaultAsp"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Site = armappservice.Site{
	// 	Name: to.Ptr("sitef6141/staging"),
	// 	Type: to.Ptr("Microsoft.Web/sites/slots"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/sitef6141/slots/staging"),
	// 	Kind: to.Ptr("app"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.SiteProperties{
	// 		AvailabilityState: to.Ptr(armappservice.SiteAvailabilityStateNormal),
	// 		ClientAffinityEnabled: to.Ptr(true),
	// 		ClientCertEnabled: to.Ptr(false),
	// 		ClientCertMode: to.Ptr(armappservice.ClientCertModeRequired),
	// 		ContainerSize: to.Ptr[int32](0),
	// 		CustomDomainVerificationID: to.Ptr("7F3BB652450EF7AD0B6AA33064586E4A9CE823A46DF0B2EC6443A68086E84153"),
	// 		DailyMemoryTimeQuota: to.Ptr[int32](0),
	// 		DefaultHostName: to.Ptr("sitef6141.azurewebsites.net"),
	// 		Enabled: to.Ptr(true),
	// 		EnabledHostNames: []*string{
	// 			to.Ptr("sitef6141-staging.azurewebsites.net"),
	// 			to.Ptr("sitef6141-staging.scm.azurewebsites.net")},
	// 			HostNameSSLStates: []*armappservice.HostNameSSLState{
	// 				{
	// 					Name: to.Ptr("sitef6141-staging.azurewebsites.net"),
	// 					HostType: to.Ptr(armappservice.HostTypeStandard),
	// 					SSLState: to.Ptr(armappservice.SSLStateDisabled),
	// 				},
	// 				{
	// 					Name: to.Ptr("sitef6141-staging.scm.azurewebsites.net"),
	// 					HostType: to.Ptr(armappservice.HostTypeRepository),
	// 					SSLState: to.Ptr(armappservice.SSLStateDisabled),
	// 			}},
	// 			HostNames: []*string{
	// 				to.Ptr("sitef6141-staging.azurewebsites.net")},
	// 				HostNamesDisabled: to.Ptr(false),
	// 				HTTPSOnly: to.Ptr(false),
	// 				HyperV: to.Ptr(false),
	// 				IsXenon: to.Ptr(false),
	// 				KeyVaultReferenceIdentity: to.Ptr("SystemAssigned"),
	// 				LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:33:11.641Z"); return t}()),
	// 				OutboundIPAddresses: to.Ptr("70.37.102.201,20.225.43.144"),
	// 				OutboundVnetRouting: &armappservice.OutboundVnetRouting{
	// 					AllTraffic: to.Ptr(false),
	// 					ApplicationTraffic: to.Ptr(false),
	// 					BackupRestoreTraffic: to.Ptr(false),
	// 					ContentShareTraffic: to.Ptr(false),
	// 					ImagePullTraffic: to.Ptr(false),
	// 				},
	// 				PossibleOutboundIPAddresses: to.Ptr("70.37.102.201,20.225.43.144,20.225.184.122,20.225.184.188"),
	// 				RedundancyMode: to.Ptr(armappservice.RedundancyModeNone),
	// 				RepositorySiteName: to.Ptr("sitef6141"),
	// 				Reserved: to.Ptr(false),
	// 				ResourceGroup: to.Ptr("testrg123"),
	// 				ScmSiteAlsoStopped: to.Ptr(false),
	// 				ServerFarmID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/DefaultAsp"),
	// 				SiteConfig: &armappservice.SiteConfig{
	// 					AcrUseManagedIdentityCreds: to.Ptr(false),
	// 					AlwaysOn: to.Ptr(false),
	// 					AppCommandLine: to.Ptr(""),
	// 					AutoHealEnabled: to.Ptr(false),
	// 					AzureStorageAccounts: map[string]*armappservice.AzureStorageInfoValue{
	// 					},
	// 					DefaultDocuments: []*string{
	// 						to.Ptr("Default.htm"),
	// 						to.Ptr("Default.html"),
	// 						to.Ptr("Default.asp"),
	// 						to.Ptr("index.htm"),
	// 						to.Ptr("index.html"),
	// 						to.Ptr("iisstart.htm"),
	// 						to.Ptr("default.aspx"),
	// 						to.Ptr("index.php"),
	// 						to.Ptr("hostingstart.html")},
	// 						DetailedErrorLoggingEnabled: to.Ptr(false),
	// 						FtpsState: to.Ptr(armappservice.FtpsStateAllAllowed),
	// 						FunctionAppScaleLimit: to.Ptr[int32](0),
	// 						FunctionsRuntimeScaleMonitoringEnabled: to.Ptr(false),
	// 						Http20Enabled: to.Ptr(false),
	// 						HTTPLoggingEnabled: to.Ptr(false),
	// 						LinuxFxVersion: to.Ptr(""),
	// 						LoadBalancing: to.Ptr(armappservice.SiteLoadBalancingLeastRequests),
	// 						LogsDirectorySizeLimit: to.Ptr[int32](35),
	// 						ManagedPipelineMode: to.Ptr(armappservice.ManagedPipelineModeIntegrated),
	// 						MinTLSVersion: to.Ptr(armappservice.SupportedTLSVersionsOne2),
	// 						MinimumElasticInstanceCount: to.Ptr[int32](0),
	// 						NetFrameworkVersion: to.Ptr("v4.0"),
	// 						NodeVersion: to.Ptr(""),
	// 						NumberOfWorkers: to.Ptr[int32](1),
	// 						PhpVersion: to.Ptr("5.6"),
	// 						PowerShellVersion: to.Ptr(""),
	// 						PythonVersion: to.Ptr(""),
	// 						RemoteDebuggingEnabled: to.Ptr(false),
	// 						RequestTracingEnabled: to.Ptr(false),
	// 						ScmMinTLSVersion: to.Ptr(armappservice.SupportedTLSVersionsOne2),
	// 						Use32BitWorkerProcess: to.Ptr(true),
	// 						VirtualApplications: []*armappservice.VirtualApplication{
	// 							{
	// 								PhysicalPath: to.Ptr("site\\wwwroot"),
	// 								PreloadEnabled: to.Ptr(false),
	// 								VirtualPath: to.Ptr("/"),
	// 						}},
	// 						VnetName: to.Ptr(""),
	// 						VnetPrivatePortsCount: to.Ptr[int32](0),
	// 						VnetRouteAllEnabled: to.Ptr(false),
	// 						WebSocketsEnabled: to.Ptr(false),
	// 					},
	// 					State: to.Ptr("Running"),
	// 					StorageAccountRequired: to.Ptr(false),
	// 					UsageState: to.Ptr(armappservice.UsageStateNormal),
	// 				},
	// 			}
}
