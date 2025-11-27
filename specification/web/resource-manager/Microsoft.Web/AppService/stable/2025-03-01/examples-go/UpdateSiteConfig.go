package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/UpdateSiteConfig.json
func ExampleWebAppsClient_CreateOrUpdateConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebAppsClient().CreateOrUpdateConfiguration(ctx, "testrg123", "sitef6141", armappservice.SiteConfigResource{
		Properties: &armappservice.SiteConfig{
			AcrUseManagedIdentityCreds: to.Ptr(false),
			AlwaysOn:                   to.Ptr(false),
			AppCommandLine:             to.Ptr(""),
			AutoHealEnabled:            to.Ptr(false),
			AzureStorageAccounts:       map[string]*armappservice.AzureStorageInfoValue{},
			DefaultDocuments: []*string{
				to.Ptr("Default.htm"),
				to.Ptr("Default.html"),
				to.Ptr("Default.asp"),
				to.Ptr("index.htm"),
				to.Ptr("index.html"),
				to.Ptr("iisstart.htm"),
				to.Ptr("default.aspx"),
				to.Ptr("index.php"),
				to.Ptr("hostingstart.html")},
			DetailedErrorLoggingEnabled:            to.Ptr(false),
			FtpsState:                              to.Ptr(armappservice.FtpsStateAllAllowed),
			FunctionAppScaleLimit:                  to.Ptr[int32](0),
			FunctionsRuntimeScaleMonitoringEnabled: to.Ptr(false),
			Http20Enabled:                          to.Ptr(false),
			HTTPLoggingEnabled:                     to.Ptr(false),
			LinuxFxVersion:                         to.Ptr(""),
			LoadBalancing:                          to.Ptr(armappservice.SiteLoadBalancingLeastRequests),
			LogsDirectorySizeLimit:                 to.Ptr[int32](35),
			ManagedPipelineMode:                    to.Ptr(armappservice.ManagedPipelineModeIntegrated),
			MinTLSVersion:                          to.Ptr(armappservice.SupportedTLSVersionsOne2),
			MinimumElasticInstanceCount:            to.Ptr[int32](0),
			NetFrameworkVersion:                    to.Ptr("v4.0"),
			NodeVersion:                            to.Ptr(""),
			NumberOfWorkers:                        to.Ptr[int32](1),
			PhpVersion:                             to.Ptr("5.6"),
			PowerShellVersion:                      to.Ptr(""),
			PythonVersion:                          to.Ptr(""),
			RemoteDebuggingEnabled:                 to.Ptr(false),
			RequestTracingEnabled:                  to.Ptr(false),
			ScmMinTLSVersion:                       to.Ptr(armappservice.SupportedTLSVersionsOne2),
			Use32BitWorkerProcess:                  to.Ptr(true),
			VirtualApplications: []*armappservice.VirtualApplication{
				{
					PhysicalPath:   to.Ptr("site\\wwwroot"),
					PreloadEnabled: to.Ptr(false),
					VirtualPath:    to.Ptr("/"),
				}},
			VnetName:              to.Ptr(""),
			VnetPrivatePortsCount: to.Ptr[int32](0),
			VnetRouteAllEnabled:   to.Ptr(false),
			WebSocketsEnabled:     to.Ptr(false),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SiteConfigResource = armappservice.SiteConfigResource{
	// 	Name: to.Ptr("web"),
	// 	Type: to.Ptr("Microsoft.Web/sites/config"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/sitef6141/config/web"),
	// 	Kind: to.Ptr("app"),
	// 	Properties: &armappservice.SiteConfig{
	// 		AcrUseManagedIdentityCreds: to.Ptr(false),
	// 		AlwaysOn: to.Ptr(false),
	// 		AppCommandLine: to.Ptr(""),
	// 		AutoHealEnabled: to.Ptr(false),
	// 		AzureStorageAccounts: map[string]*armappservice.AzureStorageInfoValue{
	// 		},
	// 		DefaultDocuments: []*string{
	// 			to.Ptr("Default.htm"),
	// 			to.Ptr("Default.html"),
	// 			to.Ptr("Default.asp"),
	// 			to.Ptr("index.htm"),
	// 			to.Ptr("index.html"),
	// 			to.Ptr("iisstart.htm"),
	// 			to.Ptr("default.aspx"),
	// 			to.Ptr("index.php"),
	// 			to.Ptr("hostingstart.html")},
	// 			DetailedErrorLoggingEnabled: to.Ptr(false),
	// 			FtpsState: to.Ptr(armappservice.FtpsStateAllAllowed),
	// 			FunctionAppScaleLimit: to.Ptr[int32](0),
	// 			FunctionsRuntimeScaleMonitoringEnabled: to.Ptr(false),
	// 			Http20Enabled: to.Ptr(false),
	// 			HTTPLoggingEnabled: to.Ptr(false),
	// 			LinuxFxVersion: to.Ptr(""),
	// 			LoadBalancing: to.Ptr(armappservice.SiteLoadBalancingLeastRequests),
	// 			LogsDirectorySizeLimit: to.Ptr[int32](35),
	// 			ManagedPipelineMode: to.Ptr(armappservice.ManagedPipelineModeIntegrated),
	// 			MinTLSVersion: to.Ptr(armappservice.SupportedTLSVersionsOne2),
	// 			MinimumElasticInstanceCount: to.Ptr[int32](0),
	// 			NetFrameworkVersion: to.Ptr("v4.0"),
	// 			NodeVersion: to.Ptr(""),
	// 			NumberOfWorkers: to.Ptr[int32](1),
	// 			PhpVersion: to.Ptr("5.6"),
	// 			PowerShellVersion: to.Ptr(""),
	// 			PythonVersion: to.Ptr(""),
	// 			RemoteDebuggingEnabled: to.Ptr(false),
	// 			RequestTracingEnabled: to.Ptr(false),
	// 			ScmMinTLSVersion: to.Ptr(armappservice.SupportedTLSVersionsOne2),
	// 			Use32BitWorkerProcess: to.Ptr(true),
	// 			VirtualApplications: []*armappservice.VirtualApplication{
	// 				{
	// 					PhysicalPath: to.Ptr("site\\wwwroot"),
	// 					PreloadEnabled: to.Ptr(false),
	// 					VirtualPath: to.Ptr("/"),
	// 			}},
	// 			VnetName: to.Ptr(""),
	// 			VnetPrivatePortsCount: to.Ptr[int32](0),
	// 			VnetRouteAllEnabled: to.Ptr(false),
	// 			WebSocketsEnabled: to.Ptr(false),
	// 		},
	// 	}
}
