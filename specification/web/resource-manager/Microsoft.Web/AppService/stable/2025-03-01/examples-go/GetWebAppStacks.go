package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetWebAppStacks.json
func ExampleProviderClient_NewGetWebAppStacksPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProviderClient().NewGetWebAppStacksPager(&armappservice.ProviderClientGetWebAppStacksOptions{StackOsType: nil})
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
		// page.WebAppStackCollection = armappservice.WebAppStackCollection{
		// 	Value: []*armappservice.WebAppStack{
		// 		{
		// 			Name: to.Ptr("dotnet"),
		// 			Type: to.Ptr("Microsoft.Web/webAppStacks"),
		// 			ID: to.Ptr("/providers/Microsoft.Web/webAppStacks/dotnet"),
		// 			Properties: &armappservice.WebAppStackProperties{
		// 				DisplayText: to.Ptr(".NET"),
		// 				MajorVersions: []*armappservice.WebAppMajorVersion{
		// 					{
		// 						DisplayText: to.Ptr(".NET 5"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr(".NET 5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("5.0.x"),
		// 										},
		// 										IsEarlyAccess: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("DOTNETCORE|5.0"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("5.0.x"),
		// 										},
		// 										IsEarlyAccess: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("v5.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("5"),
		// 						}},
		// 						Value: to.Ptr("5"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr(".NET Core 3"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr(".NET Core 3.1 (LTS)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("3.1.301"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("DOTNETCORE|3.1"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("3.1.301"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("3.1"),
		// 									},
		// 								},
		// 								Value: to.Ptr("3.1"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr(".NET Core 3.0"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-03T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("3.0.103"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("DOTNETCORE|3.0"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-03T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("3.0.103"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("3.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("3.0"),
		// 						}},
		// 						Value: to.Ptr("3"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr(".NET Core 2"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr(".NET Core 2.2"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-23T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("2.2.207"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("DOTNETCORE|2.2"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-23T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("2.2.207"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("2.2"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.2"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr(".NET Core 2.1 (LTS)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-21T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("2.1.807"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("DOTNETCORE|2.1"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-21T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("2.1.807"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("2.1"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.1"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr(".NET Core 2.0"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("2.1.202"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("DOTNETCORE|2.0"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("2.1.202"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("2.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.0"),
		// 						}},
		// 						Value: to.Ptr("dotnetcore2"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr(".NET Core 1"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr(".NET Core 1.1"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-27T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("1.1.14"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("DOTNETCORE|1.1"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-27T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("1.1.14"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.1"),
		// 									},
		// 								},
		// 								Value: to.Ptr("1.1"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr(".NET Core 1.0"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-27T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("1.1.14"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("DOTNETCORE|1.0"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-27T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("1.1.14"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("1.0"),
		// 						}},
		// 						Value: to.Ptr("1"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("ASP.NET V4"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("ASP.NET V4.8"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("3.1"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("v4.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("v4.8"),
		// 						}},
		// 						Value: to.Ptr("v4"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("ASP.NET V3"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("ASP.NET V3.5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("2.1"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("v2.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("v3.5"),
		// 						}},
		// 						Value: to.Ptr("v3"),
		// 				}},
		// 				PreferredOs: to.Ptr(armappservice.StackPreferredOsWindows),
		// 				Value: to.Ptr("dotnet"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("node"),
		// 			Type: to.Ptr("Microsoft.Web/webAppStacks"),
		// 			ID: to.Ptr("/providers/Microsoft.Web/webAppStacks/node"),
		// 			Properties: &armappservice.WebAppStackProperties{
		// 				DisplayText: to.Ptr("Node"),
		// 				MajorVersions: []*armappservice.WebAppMajorVersion{
		// 					{
		// 						DisplayText: to.Ptr("Node LTS"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Node LTS"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|lts"),
		// 									},
		// 								},
		// 								Value: to.Ptr("lts"),
		// 						}},
		// 						Value: to.Ptr("lts"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Node 14"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Node 14 LTS"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("14.x"),
		// 										},
		// 										IsPreview: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|14-lts"),
		// 									},
		// 								},
		// 								Value: to.Ptr("14-lts"),
		// 						}},
		// 						Value: to.Ptr("14"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Node 12"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Node 12 LTS"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("12.x"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|12-lts"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("12.13.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("12-lts"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 12.9"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("12.x"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(true),
		// 										RuntimeVersion: to.Ptr("NODE|12.9"),
		// 									},
		// 								},
		// 								Value: to.Ptr("12.9"),
		// 						}},
		// 						Value: to.Ptr("12"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Node 10"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Node 10 LTS"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("10.x"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|10-lts"),
		// 									},
		// 								},
		// 								Value: to.Ptr("10-LTS"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 10.16"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("10.x"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|10.16"),
		// 									},
		// 								},
		// 								Value: to.Ptr("10.16"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 10.15"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("10.x"),
		// 										},
		// 										IsHidden: to.Ptr(true),
		// 										IsPreview: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("10.15.2"),
		// 									},
		// 								},
		// 								Value: to.Ptr("10.15"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 10.14"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("10.x"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|10.14"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("10.x"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("10.14.1"),
		// 									},
		// 								},
		// 								Value: to.Ptr("10.14"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 10.12"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("10.x"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|10.12"),
		// 									},
		// 								},
		// 								Value: to.Ptr("10.12"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 10.10"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("10.x"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|10.10"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("10.x"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("10.0.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("10.10"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 10.6"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("10.x"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|10.6"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("10.6.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("10.6"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 10.1"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("10.x"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|10.1"),
		// 									},
		// 								},
		// 								Value: to.Ptr("10.1"),
		// 						}},
		// 						Value: to.Ptr("10"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Node 9"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Node 9.4"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|9.4"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.4"),
		// 						}},
		// 						Value: to.Ptr("9"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Node 8"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Node 8 LTS"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|8-lts"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8-lts"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 8.12"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|8.12"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.12"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 8.11"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|8.11"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("8.11"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.11"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 8.10"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("8.10"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.10"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 8.9"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|8.9"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("8.9"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.9"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 8.8"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|8.8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.8"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 8.5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("8.5"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 8.4"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("8.4"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.4"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 8.2"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|8.2"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.2"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 8.1"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|8.1"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("8.1.4"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.1"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 8.0"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-31T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|8.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0"),
		// 						}},
		// 						Value: to.Ptr("8"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Node 7"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Node 7.10"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("7.10.1"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.10"),
		// 						}},
		// 						Value: to.Ptr("7"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Node 6"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Node 6 LTS"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|6-lts"),
		// 									},
		// 								},
		// 								Value: to.Ptr("6-LTS"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 6.12"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("6.12"),
		// 									},
		// 								},
		// 								Value: to.Ptr("6.12"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 6.11"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|6.11"),
		// 									},
		// 								},
		// 								Value: to.Ptr("6.11"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 6.10"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|6.10"),
		// 									},
		// 								},
		// 								Value: to.Ptr("6.10"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 6.9"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|6.9"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("6.9.5"),
		// 									},
		// 								},
		// 								Value: to.Ptr("6.9"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 6.6"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|6.6"),
		// 									},
		// 								},
		// 								Value: to.Ptr("6.6"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 6.5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("6.5.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("6.5"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 6.2"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|6.2"),
		// 									},
		// 								},
		// 								Value: to.Ptr("6.2"),
		// 						}},
		// 						Value: to.Ptr("6"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Node 4"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Node 4.8"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|4.8"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("4.8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("4.8"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 4.5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|4.5"),
		// 									},
		// 								},
		// 								Value: to.Ptr("4.5"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Node 4.4"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("NODE|4.4"),
		// 									},
		// 								},
		// 								Value: to.Ptr("4.4"),
		// 						}},
		// 						Value: to.Ptr("4"),
		// 				}},
		// 				PreferredOs: to.Ptr(armappservice.StackPreferredOsLinux),
		// 				Value: to.Ptr("node"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("python"),
		// 			Type: to.Ptr("Microsoft.Web/webAppStacks"),
		// 			ID: to.Ptr("/providers/Microsoft.Web/webAppStacks/python"),
		// 			Properties: &armappservice.WebAppStackProperties{
		// 				DisplayText: to.Ptr("Python"),
		// 				MajorVersions: []*armappservice.WebAppMajorVersion{
		// 					{
		// 						DisplayText: to.Ptr("Python 3"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Python 3.8"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("3.8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("PYTHON|3.8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("3.8"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Python 3.7"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("3.7"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("PYTHON|3.7"),
		// 									},
		// 								},
		// 								Value: to.Ptr("3.7"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Python 3.6"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("3.6"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("PYTHON|3.6"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("3.6"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("3.4.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("3.6"),
		// 						}},
		// 						Value: to.Ptr("3"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Python 2"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Python 2.7"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("2.7"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("PYTHON|2.7"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("2.7"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("2.7.3"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.7"),
		// 						}},
		// 						Value: to.Ptr("2"),
		// 				}},
		// 				PreferredOs: to.Ptr(armappservice.StackPreferredOsLinux),
		// 				Value: to.Ptr("python"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("php"),
		// 			Type: to.Ptr("Microsoft.Web/webAppStacks"),
		// 			ID: to.Ptr("/providers/Microsoft.Web/webAppStacks/php"),
		// 			Properties: &armappservice.WebAppStackProperties{
		// 				DisplayText: to.Ptr("PHP"),
		// 				MajorVersions: []*armappservice.WebAppMajorVersion{
		// 					{
		// 						DisplayText: to.Ptr("PHP 7"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("PHP 7.4"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-28T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("PHP|7.4"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-28T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("7.4"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.4"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("PHP 7.3"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-06T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("PHP|7.3"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-06T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("7.3"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.3"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("PHP 7.2"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("PHP|7.2"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-30T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("7.2"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.2"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("PHP 7.1"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("7.1"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.1"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("7.0"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("PHP|7.0"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("7.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0"),
		// 						}},
		// 						Value: to.Ptr("7"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("PHP 5"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("PHP 5.6"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("PHP|5.6"),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("5.6"),
		// 									},
		// 								},
		// 								Value: to.Ptr("5.6"),
		// 						}},
		// 						Value: to.Ptr("5"),
		// 				}},
		// 				PreferredOs: to.Ptr(armappservice.StackPreferredOsLinux),
		// 				Value: to.Ptr("php"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ruby"),
		// 			Type: to.Ptr("Microsoft.Web/webAppStacks"),
		// 			ID: to.Ptr("/providers/Microsoft.Web/webAppStacks/ruby"),
		// 			Properties: &armappservice.WebAppStackProperties{
		// 				DisplayText: to.Ptr("Ruby"),
		// 				MajorVersions: []*armappservice.WebAppMajorVersion{
		// 					{
		// 						DisplayText: to.Ptr("Ruby 2"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Ruby 2.6"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("RUBY|2.6"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.6"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Ruby 2.6.2"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("RUBY|2.6.2"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.6.2"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Ruby 2.5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("RUBY|2.5"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.5"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Ruby 2.5.5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("RUBY|2.5.5"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.5.5"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Ruby 2.4"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("RUBY|2.4"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.4"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Ruby 2.4.5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("RUBY|2.4.5"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.4.5"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Ruby 2.3"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("RUBY|2.3"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.3"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Ruby 2.3.8"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("RUBY|2.3.8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.3.8"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Ruby 2.3.3"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("RUBY|2.3.3"),
		// 									},
		// 								},
		// 								Value: to.Ptr("2.3.3"),
		// 						}},
		// 						Value: to.Ptr("2"),
		// 				}},
		// 				PreferredOs: to.Ptr(armappservice.StackPreferredOsLinux),
		// 				Value: to.Ptr("ruby"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("java"),
		// 			Type: to.Ptr("Microsoft.Web/webAppStacks"),
		// 			ID: to.Ptr("/providers/Microsoft.Web/webAppStacks/java"),
		// 			Properties: &armappservice.WebAppStackProperties{
		// 				DisplayText: to.Ptr("Java"),
		// 				MajorVersions: []*armappservice.WebAppMajorVersion{
		// 					{
		// 						DisplayText: to.Ptr("Java 11"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Java 11"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("11"),
		// 										},
		// 										IsAutoUpdate: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr(""),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("11"),
		// 										},
		// 										IsAutoUpdate: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("11"),
		// 									},
		// 								},
		// 								Value: to.Ptr("11.0"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 11.0.7"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("11"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr(""),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("11"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("11.0.7"),
		// 									},
		// 								},
		// 								Value: to.Ptr("11.0.7"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 11.0.6"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("11"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr(""),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("11"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("11.0.6"),
		// 									},
		// 								},
		// 								Value: to.Ptr("11.0.6"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 11.0.5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("11"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr(""),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("11"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("11.0.5_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("11.0.5"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 11.0.3"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("11"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("11.0.3_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("11.0.3"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 11.0.2"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("11"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("11.0.2_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("11.0.2"),
		// 						}},
		// 						Value: to.Ptr("11"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Java 8"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Java 8"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										IsAutoUpdate: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr(""),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										IsAutoUpdate: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_252"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr(""),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_252"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.252"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_242"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr(""),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_242"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.242"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_232"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr(""),
		// 									},
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_232_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.232"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_212"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_212_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.212"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_202"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_202_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.202"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_202 (Oracle)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_202"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.202 (Oracle)"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_181"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_181_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.181"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_181 (Oracle)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_181"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.181 (Oracle)"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_172"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_172_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.172"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_172 (Oracle)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_172"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.172 (Oracle)"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_144"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_144"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.144"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_111 (Oracle)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_111"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.111 (Oracle)"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_102"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_102"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.102"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_92"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_92"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.92"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_73 (Oracle)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_73"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.73 (Oracle)"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_60 (Oracle)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_60"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.60 (Oracle)"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.8.0_25 (Oracle)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("8"),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.8.0_25"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.25 (Oracle)"),
		// 						}},
		// 						Value: to.Ptr("8"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Java 7"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Java 7"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsAutoUpdate: to.Ptr(true),
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.7"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.7.0_262"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.7.0_262_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.262"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.7.0_242"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.7.0_242_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.242"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.7.0_222"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.7.0_222_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.222"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.7.0_191"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.7.0_191_ZULU"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.191"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.7.0_80 (Oracle)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.7.0_80"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.80 (Oracle)"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.7.0_71 (Oracle)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.7.0_71"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.71 (Oracle)"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java 1.7.0_51 (Oracle)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsRuntimeSettings: &armappservice.WebAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(false),
		// 										},
		// 										IsDeprecated: to.Ptr(true),
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("1.7.0_51"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.51 (Oracle)"),
		// 						}},
		// 						Value: to.Ptr("7"),
		// 				}},
		// 				PreferredOs: to.Ptr(armappservice.StackPreferredOsLinux),
		// 				Value: to.Ptr("java"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("javacontainers"),
		// 			Type: to.Ptr("Microsoft.Web/webAppStacks"),
		// 			ID: to.Ptr("/providers/Microsoft.Web/webAppStacks/javacontainers"),
		// 			Properties: &armappservice.WebAppStackProperties{
		// 				DisplayText: to.Ptr("Java Containers"),
		// 				MajorVersions: []*armappservice.WebAppMajorVersion{
		// 					{
		// 						DisplayText: to.Ptr("Java SE (Embedded Web Server)"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Java SE (Embedded Web Server)"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										Java11Runtime: to.Ptr("JAVA|11-java11"),
		// 										Java8Runtime: to.Ptr("JAVA|8-jre8"),
		// 									},
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										JavaContainer: to.Ptr("JAVA"),
		// 										JavaContainerVersion: to.Ptr("SE"),
		// 									},
		// 								},
		// 								Value: to.Ptr("SE"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java SE 11.0.7"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java11Runtime: to.Ptr("JAVA|11.0.7"),
		// 									},
		// 								},
		// 								Value: to.Ptr("11.0.7"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java SE 11.0.6"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java11Runtime: to.Ptr("JAVA|11.0.6"),
		// 									},
		// 								},
		// 								Value: to.Ptr("11.0.6"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java SE 11.0.5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java11Runtime: to.Ptr("JAVA|11.0.5"),
		// 									},
		// 								},
		// 								Value: to.Ptr("11.0.5"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java SE 8u252"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java8Runtime: to.Ptr("JAVA|8u252"),
		// 									},
		// 								},
		// 								Value: to.Ptr("1.8.252"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java SE 8u242"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java8Runtime: to.Ptr("JAVA|8u242"),
		// 									},
		// 								},
		// 								Value: to.Ptr("1.8.242"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Java SE 8u232"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java8Runtime: to.Ptr("JAVA|8u232"),
		// 									},
		// 								},
		// 								Value: to.Ptr("1.8.232"),
		// 						}},
		// 						Value: to.Ptr("javase"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("JBoss EAP"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("JBoss EAP 7.2"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										IsPreview: to.Ptr(true),
		// 										Java8Runtime: to.Ptr("JBOSSEAP|7.2-java8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.2"),
		// 						}},
		// 						Value: to.Ptr("jbosseap"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Tomcat 9.0"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										Java11Runtime: to.Ptr("TOMCAT|9.0-java11"),
		// 										Java8Runtime: to.Ptr("TOMCAT|9.0-jre8"),
		// 									},
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("9.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0.37"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java11Runtime: to.Ptr("TOMCAT|9.0.37-java11"),
		// 										Java8Runtime: to.Ptr("TOMCAT|9.0.37-java8"),
		// 									},
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("9.0.37"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0.37"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0.33"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java11Runtime: to.Ptr("TOMCAT|9.0.33-java11"),
		// 										Java8Runtime: to.Ptr("TOMCAT|9.0.33-java8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0.33"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0.31"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("9.0.31"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0.31"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0.27"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("9.0.27"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0.27"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0.21"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("9.0.21"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0.21"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0.20"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java11Runtime: to.Ptr("TOMCAT|9.0.20-java11"),
		// 										Java8Runtime: to.Ptr("TOMCAT|9.0.20-java8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0.20"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0.14"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("9.0.14"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0.14"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0.12"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("9.0.12"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0.12"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0.8"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("9.0.8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0.8"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 9.0.0"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("9.0.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.0.0"),
		// 						}},
		// 						Value: to.Ptr("tomcat9.0"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Tomcat 8.5"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										Java11Runtime: to.Ptr("TOMCAT|8.5-java11"),
		// 										Java8Runtime: to.Ptr("TOMCAT|8.5-jre8"),
		// 									},
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.5"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.6"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.5.6"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.6"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.57"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java11Runtime: to.Ptr("TOMCAT|8.5.57-java11"),
		// 										Java8Runtime: to.Ptr("TOMCAT|8.5.57-java8"),
		// 									},
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.5.57"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.57"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.53"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java11Runtime: to.Ptr("TOMCAT|8.5.53-java11"),
		// 										Java8Runtime: to.Ptr("TOMCAT|8.5.53-java8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.53"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.51"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.5.51"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.51"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.47"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.5.47"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.47"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.42"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.5.42"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.42"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.41"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										Java11Runtime: to.Ptr("TOMCAT|8.5.41-java11"),
		// 										Java8Runtime: to.Ptr("TOMCAT|8.5.41-java8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.41"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.37"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.5.37"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.37"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.34"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.5.34"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.34"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.31"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.5.31"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.31"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.5.20"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.5.20"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.5.20"),
		// 						}},
		// 						Value: to.Ptr("tomcat8.5"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Tomcat 8.0"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.0"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.0.53"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.0.53"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.53"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.0.46"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.0.46"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.46"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 8.0.23"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("8.0.23"),
		// 									},
		// 								},
		// 								Value: to.Ptr("8.0.23"),
		// 						}},
		// 						Value: to.Ptr("tomcat8.0"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Tomcat 7.0"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 7.0"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("7.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 7.0.94"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("7.0.94"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.94"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 7.0.81"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("7.0.81"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.81"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 7.0.62"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("7.0.62"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.62"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Tomcat 7.0.50"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										JavaContainer: to.Ptr("TOMCAT"),
		// 										JavaContainerVersion: to.Ptr("7.0.50"),
		// 									},
		// 								},
		// 								Value: to.Ptr("7.0.50"),
		// 						}},
		// 						Value: to.Ptr("tomcat7.0"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Jetty 9.3"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Jetty 9.3"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										IsDeprecated: to.Ptr(true),
		// 										JavaContainer: to.Ptr("JETTY"),
		// 										JavaContainerVersion: to.Ptr("9.3"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.3"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Jetty 9.3.25"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										IsDeprecated: to.Ptr(true),
		// 										JavaContainer: to.Ptr("JETTY"),
		// 										JavaContainerVersion: to.Ptr("9.3.25"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.3.25"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Jetty 9.3.13"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										IsDeprecated: to.Ptr(true),
		// 										JavaContainer: to.Ptr("JETTY"),
		// 										JavaContainerVersion: to.Ptr("9.3.13"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.3.13"),
		// 						}},
		// 						Value: to.Ptr("jetty9.3"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("Jetty 9.1"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("Jetty 9.1"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										IsDeprecated: to.Ptr(true),
		// 										JavaContainer: to.Ptr("JETTY"),
		// 										JavaContainerVersion: to.Ptr("9.1"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.1"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("Jetty 9.1.0"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									WindowsContainerSettings: &armappservice.WindowsJavaContainerSettings{
		// 										IsDeprecated: to.Ptr(true),
		// 										JavaContainer: to.Ptr("JETTY"),
		// 										JavaContainerVersion: to.Ptr("9.1.0"),
		// 									},
		// 								},
		// 								Value: to.Ptr("9.1.0"),
		// 						}},
		// 						Value: to.Ptr("jetty9.1"),
		// 					},
		// 					{
		// 						DisplayText: to.Ptr("WildFly 14"),
		// 						MinorVersions: []*armappservice.WebAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr("WildFly 14"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										IsAutoUpdate: to.Ptr(true),
		// 										IsDeprecated: to.Ptr(true),
		// 										Java8Runtime: to.Ptr("WILDFLY|14-jre8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("14"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr("WildFly 14.0.1"),
		// 								StackSettings: &armappservice.WebAppRuntimes{
		// 									LinuxContainerSettings: &armappservice.LinuxJavaContainerSettings{
		// 										IsDeprecated: to.Ptr(true),
		// 										Java8Runtime: to.Ptr("WILDFLY|14.0.1-java8"),
		// 									},
		// 								},
		// 								Value: to.Ptr("14.0.1"),
		// 						}},
		// 						Value: to.Ptr("wildfly14"),
		// 				}},
		// 				Value: to.Ptr("javacontainers"),
		// 			},
		// 	}},
		// }
	}
}
