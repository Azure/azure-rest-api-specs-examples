package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/PacketCoreControlPlaneVersionList.json
func ExamplePacketCoreControlPlaneVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPacketCoreControlPlaneVersionsClient().NewListPager(nil)
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
		// page.PacketCoreControlPlaneVersionListResult = armmobilenetwork.PacketCoreControlPlaneVersionListResult{
		// 	Value: []*armmobilenetwork.PacketCoreControlPlaneVersion{
		// 		{
		// 			Name: to.Ptr("2404.0-0"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlaneVersions"),
		// 			ID: to.Ptr("/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/2404.0-0"),
		// 			Properties: &armmobilenetwork.PacketCoreControlPlaneVersionPropertiesFormat{
		// 				Platforms: []*armmobilenetwork.Platform{
		// 					{
		// 						HaUpgradesAvailable: []*string{
		// 							to.Ptr("2404.1"),
		// 							to.Ptr("2407.0")},
		// 							MaximumPlatformSoftwareVersion: to.Ptr("2211"),
		// 							MinimumPlatformSoftwareVersion: to.Ptr("2209"),
		// 							PlatformType: to.Ptr(armmobilenetwork.PlatformTypeAKSHCI),
		// 							RecommendedVersion: to.Ptr(armmobilenetwork.RecommendedVersionNotRecommended),
		// 							VersionState: to.Ptr(armmobilenetwork.VersionStateActive),
		// 					}},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("2404.1-1"),
		// 				Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlaneVersions"),
		// 				ID: to.Ptr("/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/2404.1-1"),
		// 				Properties: &armmobilenetwork.PacketCoreControlPlaneVersionPropertiesFormat{
		// 					Platforms: []*armmobilenetwork.Platform{
		// 						{
		// 							HaUpgradesAvailable: []*string{
		// 								to.Ptr("2407.0")},
		// 								MaximumPlatformSoftwareVersion: to.Ptr("2212"),
		// 								MinimumPlatformSoftwareVersion: to.Ptr("2210"),
		// 								PlatformType: to.Ptr(armmobilenetwork.PlatformTypeAKSHCI),
		// 								RecommendedVersion: to.Ptr(armmobilenetwork.RecommendedVersionNotRecommended),
		// 								VersionState: to.Ptr(armmobilenetwork.VersionStateActive),
		// 						}},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("2407.0-1"),
		// 					Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlaneVersions"),
		// 					ID: to.Ptr("/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/2407.0-1"),
		// 					Properties: &armmobilenetwork.PacketCoreControlPlaneVersionPropertiesFormat{
		// 						Platforms: []*armmobilenetwork.Platform{
		// 							{
		// 								HaUpgradesAvailable: []*string{
		// 									to.Ptr("2410.1")},
		// 									MaximumPlatformSoftwareVersion: to.Ptr("2301"),
		// 									MinimumPlatformSoftwareVersion: to.Ptr("2211"),
		// 									PlatformType: to.Ptr(armmobilenetwork.PlatformTypeAKSHCI),
		// 									RecommendedVersion: to.Ptr(armmobilenetwork.RecommendedVersionRecommended),
		// 									VersionState: to.Ptr(armmobilenetwork.VersionStateActive),
		// 							}},
		// 						},
		// 				}},
		// 			}
	}
}
