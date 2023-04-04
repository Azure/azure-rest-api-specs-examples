package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/DefenderSettings/PackageDownloads.json
func ExampleDefenderSettingsClient_PackageDownloads() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefenderSettingsClient().PackageDownloads(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PackageDownloads = armiotsecurity.PackageDownloads{
	// 	AuthorizedDevicesImportTemplate: []*armiotsecurity.PackageDownloadInfo{
	// 		{
	// 			Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 			Version: to.Ptr("2020.02.10"),
	// 	}},
	// 	CentralManager: &armiotsecurity.PackageDownloadsCentralManager{
	// 		Full: &armiotsecurity.PackageDownloadsCentralManagerFull{
	// 			Iso: []*armiotsecurity.PackageDownloadInfo{
	// 				{
	// 					Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 					Version: to.Ptr("3.1"),
	// 					VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 				},
	// 				{
	// 					Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 					Version: to.Ptr("2.8.11"),
	// 					VersionKind: to.Ptr(armiotsecurity.VersionKindPrevious),
	// 			}},
	// 			Ovf: &armiotsecurity.PackageDownloadsCentralManagerFullOvf{
	// 				Enterprise: []*armiotsecurity.PackageDownloadInfo{
	// 					{
	// 						Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 						Version: to.Ptr("3.1"),
	// 						VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 				}},
	// 				EnterpriseHighAvailability: []*armiotsecurity.PackageDownloadInfo{
	// 					{
	// 						Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 						Version: to.Ptr("3.1"),
	// 						VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 				}},
	// 				Medium: []*armiotsecurity.PackageDownloadInfo{
	// 					{
	// 						Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 						Version: to.Ptr("3.1"),
	// 						VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 				}},
	// 				MediumHighAvailability: []*armiotsecurity.PackageDownloadInfo{
	// 					{
	// 						Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 						Version: to.Ptr("3.1"),
	// 						VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 				}},
	// 			},
	// 		},
	// 		Upgrade: []*armiotsecurity.UpgradePackageDownloadInfo{
	// 			{
	// 				Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 				Version: to.Ptr("2.8.2"),
	// 				VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 				FromVersion: to.Ptr("2.8.0"),
	// 			},
	// 			{
	// 				Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 				Version: to.Ptr("2.8.10"),
	// 				VersionKind: to.Ptr(armiotsecurity.VersionKindPrevious),
	// 				FromVersion: to.Ptr("2.8.0"),
	// 		}},
	// 	},
	// 	DeviceInformationUpdateImportTemplate: []*armiotsecurity.PackageDownloadInfo{
	// 		{
	// 			Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 			Version: to.Ptr("2020.02.10"),
	// 	}},
	// 	Sensor: &armiotsecurity.PackageDownloadsSensor{
	// 		Full: &armiotsecurity.PackageDownloadsSensorFull{
	// 			Iso: []*armiotsecurity.PackageDownloadInfo{
	// 				{
	// 					Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 					Version: to.Ptr("3.1"),
	// 					VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 				},
	// 				{
	// 					Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 					Version: to.Ptr("2.8.11"),
	// 					VersionKind: to.Ptr(armiotsecurity.VersionKindPrevious),
	// 			}},
	// 			Ovf: &armiotsecurity.PackageDownloadsSensorFullOvf{
	// 				Enterprise: []*armiotsecurity.PackageDownloadInfo{
	// 					{
	// 						Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 						Version: to.Ptr("3.1"),
	// 						VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 					},
	// 					{
	// 						Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 						Version: to.Ptr("2.8.11"),
	// 						VersionKind: to.Ptr(armiotsecurity.VersionKindPrevious),
	// 				}},
	// 				Line: []*armiotsecurity.PackageDownloadInfo{
	// 					{
	// 						Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 						Version: to.Ptr("3.1"),
	// 						VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 					},
	// 					{
	// 						Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 						Version: to.Ptr("2.8.11"),
	// 						VersionKind: to.Ptr(armiotsecurity.VersionKindPrevious),
	// 				}},
	// 				Medium: []*armiotsecurity.PackageDownloadInfo{
	// 					{
	// 						Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 						Version: to.Ptr("3.1"),
	// 						VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 					},
	// 					{
	// 						Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 						Version: to.Ptr("2.8.11"),
	// 						VersionKind: to.Ptr(armiotsecurity.VersionKindPrevious),
	// 				}},
	// 			},
	// 		},
	// 		Upgrade: []*armiotsecurity.UpgradePackageDownloadInfo{
	// 			{
	// 				Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 				Version: to.Ptr("2.8.2"),
	// 				VersionKind: to.Ptr(armiotsecurity.VersionKindLatest),
	// 				FromVersion: to.Ptr("2.8.0"),
	// 			},
	// 			{
	// 				Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 				Version: to.Ptr("2.8.10"),
	// 				VersionKind: to.Ptr(armiotsecurity.VersionKindPrevious),
	// 				FromVersion: to.Ptr("2.8.0"),
	// 		}},
	// 	},
	// 	Snmp: []*armiotsecurity.PackageDownloadInfo{
	// 		{
	// 			Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 			Version: to.Ptr("2020.02.10"),
	// 	}},
	// 	ThreatIntelligence: []*armiotsecurity.PackageDownloadInfo{
	// 		{
	// 			Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 			Version: to.Ptr("2020.02.10"),
	// 	}},
	// 	WmiTool: []*armiotsecurity.PackageDownloadInfo{
	// 		{
	// 			Link: to.Ptr("http://microsoft.com/downloadLocation"),
	// 			Version: to.Ptr("2020.02.10"),
	// 	}},
	// }
}
