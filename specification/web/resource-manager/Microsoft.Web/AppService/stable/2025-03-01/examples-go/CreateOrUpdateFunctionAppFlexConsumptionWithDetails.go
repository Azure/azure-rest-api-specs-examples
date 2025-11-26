package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/CreateOrUpdateFunctionAppFlexConsumptionWithDetails.json
func ExampleWebAppsClient_BeginCreateOrUpdate_createOrUpdateFlexConsumptionFunctionAppWithDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWebAppsClient().BeginCreateOrUpdate(ctx, "testrg123", "sitef6141", armappservice.Site{
		Kind:     to.Ptr("functionapp,linux"),
		Location: to.Ptr("East US"),
		Properties: &armappservice.SiteProperties{
			FunctionAppConfig: &armappservice.FunctionAppConfig{
				Deployment: &armappservice.FunctionsDeployment{
					Storage: &armappservice.FunctionsDeploymentStorage{
						Type: to.Ptr(armappservice.FunctionsDeploymentStorageTypeBlobContainer),
						Authentication: &armappservice.FunctionsDeploymentStorageAuthentication{
							Type:                               to.Ptr(armappservice.AuthenticationTypeStorageAccountConnectionString),
							StorageAccountConnectionStringName: to.Ptr("TheAppSettingName"),
						},
						Value: to.Ptr("https://storageAccountName.blob.core.windows.net/containername"),
					},
				},
				Runtime: &armappservice.FunctionsRuntime{
					Name:    to.Ptr(armappservice.RuntimeNamePython),
					Version: to.Ptr("3.11"),
				},
				ScaleAndConcurrency: &armappservice.FunctionsScaleAndConcurrency{
					AlwaysReady: []*armappservice.FunctionsAlwaysReadyConfig{
						{
							Name:          to.Ptr("http"),
							InstanceCount: to.Ptr[int32](2),
						}},
					InstanceMemoryMB:     to.Ptr[int32](2048),
					MaximumInstanceCount: to.Ptr[int32](100),
					Triggers: &armappservice.FunctionsScaleAndConcurrencyTriggers{
						HTTP: &armappservice.FunctionsScaleAndConcurrencyTriggersHTTP{
							PerInstanceConcurrency: to.Ptr[int32](16),
						},
					},
				},
			},
			SiteConfig: &armappservice.SiteConfig{
				AppSettings: []*armappservice.NameValuePair{
					{
						Name:  to.Ptr("AzureWebJobsStorage"),
						Value: to.Ptr("DefaultEndpointsProtocol=https;AccountName=StorageAccountName;AccountKey=Sanitized;EndpointSuffix=core.windows.net"),
					},
					{
						Name:  to.Ptr("APPLICATIONINSIGHTS_CONNECTION_STRING"),
						Value: to.Ptr("InstrumentationKey=Sanitized;IngestionEndpoint=Sanitized;LiveEndpoint=Sanitized"),
					}},
			},
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
	// res.Site = armappservice.Site{
	// 	Name: to.Ptr("sitef6141"),
	// 	Type: to.Ptr("Microsoft.Web/sites"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/sitef6141"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.SiteProperties{
	// 		AvailabilityState: to.Ptr(armappservice.SiteAvailabilityStateNormal),
	// 		ClientAffinityEnabled: to.Ptr(false),
	// 		ClientCertEnabled: to.Ptr(false),
	// 		ContainerSize: to.Ptr[int32](2048),
	// 		DailyMemoryTimeQuota: to.Ptr[int32](0),
	// 		DefaultHostName: to.Ptr("sitef6141.azurewebsites.net"),
	// 		Enabled: to.Ptr(true),
	// 		EnabledHostNames: []*string{
	// 			to.Ptr("sitef6141.azurewebsites.net"),
	// 			to.Ptr("sitef6141.scm.azurewebsites.net")},
	// 			FunctionAppConfig: &armappservice.FunctionAppConfig{
	// 				Deployment: &armappservice.FunctionsDeployment{
	// 					Storage: &armappservice.FunctionsDeploymentStorage{
	// 						Type: to.Ptr(armappservice.FunctionsDeploymentStorageTypeBlobContainer),
	// 						Authentication: &armappservice.FunctionsDeploymentStorageAuthentication{
	// 							Type: to.Ptr(armappservice.AuthenticationTypeStorageAccountConnectionString),
	// 							StorageAccountConnectionStringName: to.Ptr("TheAppSettingName"),
	// 						},
	// 						Value: to.Ptr("https://storageAccountName.blob.core.windows.net/containername"),
	// 					},
	// 				},
	// 				Runtime: &armappservice.FunctionsRuntime{
	// 					Name: to.Ptr(armappservice.RuntimeNamePython),
	// 					Version: to.Ptr("3.11"),
	// 				},
	// 				ScaleAndConcurrency: &armappservice.FunctionsScaleAndConcurrency{
	// 					AlwaysReady: []*armappservice.FunctionsAlwaysReadyConfig{
	// 						{
	// 							Name: to.Ptr("http"),
	// 							InstanceCount: to.Ptr[int32](2),
	// 					}},
	// 					InstanceMemoryMB: to.Ptr[int32](2048),
	// 					MaximumInstanceCount: to.Ptr[int32](100),
	// 					Triggers: &armappservice.FunctionsScaleAndConcurrencyTriggers{
	// 						HTTP: &armappservice.FunctionsScaleAndConcurrencyTriggersHTTP{
	// 							PerInstanceConcurrency: to.Ptr[int32](16),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			HostNameSSLStates: []*armappservice.HostNameSSLState{
	// 				{
	// 					Name: to.Ptr("sitef6141.azurewebsites.net"),
	// 					HostType: to.Ptr(armappservice.HostTypeStandard),
	// 					SSLState: to.Ptr(armappservice.SSLStateDisabled),
	// 				},
	// 				{
	// 					Name: to.Ptr("sitef6141.scm.azurewebsites.net"),
	// 					HostType: to.Ptr(armappservice.HostTypeRepository),
	// 					SSLState: to.Ptr(armappservice.SSLStateDisabled),
	// 			}},
	// 			HostNames: []*string{
	// 				to.Ptr("sitef6141.azurewebsites.net")},
	// 				HTTPSOnly: to.Ptr(true),
	// 				HyperV: to.Ptr(false),
	// 				IsXenon: to.Ptr(false),
	// 				LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-08T16:24:54.800Z"); return t}()),
	// 				OutboundIPAddresses: to.Ptr("70.37.102.201,20.225.43.144"),
	// 				OutboundVnetRouting: &armappservice.OutboundVnetRouting{
	// 					AllTraffic: to.Ptr(false),
	// 					ApplicationTraffic: to.Ptr(false),
	// 					BackupRestoreTraffic: to.Ptr(false),
	// 					ContentShareTraffic: to.Ptr(false),
	// 					ImagePullTraffic: to.Ptr(false),
	// 				},
	// 				PossibleOutboundIPAddresses: to.Ptr("70.37.102.201,20.225.43.144,20.225.184.122,20.225.184.188"),
	// 				PublicNetworkAccess: to.Ptr("Enabled"),
	// 				RedundancyMode: to.Ptr(armappservice.RedundancyModeNone),
	// 				RepositorySiteName: to.Ptr("sitef6141"),
	// 				ResourceConfig: &armappservice.ResourceConfig{
	// 					CPU: to.Ptr[float64](1),
	// 					Memory: to.Ptr("2.0Gi"),
	// 				},
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
	// 					DetailedErrorLoggingEnabled: to.Ptr(false),
	// 					FtpsState: to.Ptr(armappservice.FtpsStateAllAllowed),
	// 					FunctionAppScaleLimit: to.Ptr[int32](0),
	// 					FunctionsRuntimeScaleMonitoringEnabled: to.Ptr(false),
	// 					Http20Enabled: to.Ptr(false),
	// 					HTTPLoggingEnabled: to.Ptr(false),
	// 					IPSecurityRestrictions: []*armappservice.IPSecurityRestriction{
	// 						{
	// 							Name: to.Ptr("Allow all"),
	// 							Description: to.Ptr("Allow all access"),
	// 							Action: to.Ptr("Allow"),
	// 							IPAddress: to.Ptr("Any"),
	// 							Priority: to.Ptr[int32](2147483647),
	// 					}},
	// 					KeyVaultReferenceIdentity: to.Ptr(""),
	// 					LinuxFxVersion: to.Ptr(""),
	// 					LoadBalancing: to.Ptr(armappservice.SiteLoadBalancingLeastRequests),
	// 					LogsDirectorySizeLimit: to.Ptr[int32](35),
	// 					ManagedPipelineMode: to.Ptr(armappservice.ManagedPipelineModeIntegrated),
	// 					MinTLSVersion: to.Ptr(armappservice.SupportedTLSVersionsOne2),
	// 					MinimumElasticInstanceCount: to.Ptr[int32](0),
	// 					NetFrameworkVersion: to.Ptr(""),
	// 					NodeVersion: to.Ptr(""),
	// 					NumberOfWorkers: to.Ptr[int32](1),
	// 					PhpVersion: to.Ptr(""),
	// 					PowerShellVersion: to.Ptr(""),
	// 					PythonVersion: to.Ptr(""),
	// 					RemoteDebuggingEnabled: to.Ptr(false),
	// 					RequestTracingEnabled: to.Ptr(false),
	// 					ScmIPSecurityRestrictions: []*armappservice.IPSecurityRestriction{
	// 						{
	// 							Name: to.Ptr("Allow all"),
	// 							Description: to.Ptr("Allow all access"),
	// 							Action: to.Ptr("Allow"),
	// 							IPAddress: to.Ptr("Any"),
	// 							Priority: to.Ptr[int32](2147483647),
	// 					}},
	// 					ScmMinTLSVersion: to.Ptr(armappservice.SupportedTLSVersionsOne2),
	// 					Use32BitWorkerProcess: to.Ptr(false),
	// 					VirtualApplications: []*armappservice.VirtualApplication{
	// 						{
	// 							PhysicalPath: to.Ptr("site\\wwwroot"),
	// 							PreloadEnabled: to.Ptr(false),
	// 							VirtualPath: to.Ptr("/"),
	// 					}},
	// 					VnetName: to.Ptr(""),
	// 					VnetPrivatePortsCount: to.Ptr[int32](0),
	// 					VnetRouteAllEnabled: to.Ptr(false),
	// 					WebSocketsEnabled: to.Ptr(false),
	// 				},
	// 				State: to.Ptr("Running"),
	// 				StorageAccountRequired: to.Ptr(false),
	// 				UsageState: to.Ptr(armappservice.UsageStateNormal),
	// 			},
	// 		}
}
