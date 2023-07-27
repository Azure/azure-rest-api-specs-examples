package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/eba34b9c764e877193788a87a81cebfa915eb858/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/PacketCoreControlPlaneVersionList.json
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
		// 			Name: to.Ptr("PMN-4-9-4"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlaneVersions"),
		// 			ID: to.Ptr("/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/PMN-4-9-4"),
		// 			Properties: &armmobilenetwork.PacketCoreControlPlaneVersionPropertiesFormat{
		// 				Platforms: []*armmobilenetwork.Platform{
		// 					{
		// 						MaximumPlatformSoftwareVersion: to.Ptr("2211"),
		// 						MinimumPlatformSoftwareVersion: to.Ptr("2209"),
		// 						PlatformType: to.Ptr(armmobilenetwork.PlatformTypeAKSHCI),
		// 						RecommendedVersion: to.Ptr(armmobilenetwork.RecommendedVersionNotRecommended),
		// 						VersionState: to.Ptr(armmobilenetwork.VersionStateActive),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PMN-4-10-2"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlaneVersions"),
		// 			ID: to.Ptr("/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/PMN-4-10-2"),
		// 			Properties: &armmobilenetwork.PacketCoreControlPlaneVersionPropertiesFormat{
		// 				Platforms: []*armmobilenetwork.Platform{
		// 					{
		// 						MaximumPlatformSoftwareVersion: to.Ptr("2212"),
		// 						MinimumPlatformSoftwareVersion: to.Ptr("2210"),
		// 						PlatformType: to.Ptr(armmobilenetwork.PlatformTypeAKSHCI),
		// 						RecommendedVersion: to.Ptr(armmobilenetwork.RecommendedVersionNotRecommended),
		// 						VersionState: to.Ptr(armmobilenetwork.VersionStateActive),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PMN-4-11-1"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlaneVersions"),
		// 			ID: to.Ptr("/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/PMN-4-11-1"),
		// 			Properties: &armmobilenetwork.PacketCoreControlPlaneVersionPropertiesFormat{
		// 				Platforms: []*armmobilenetwork.Platform{
		// 					{
		// 						MaximumPlatformSoftwareVersion: to.Ptr("2301"),
		// 						MinimumPlatformSoftwareVersion: to.Ptr("2211"),
		// 						PlatformType: to.Ptr(armmobilenetwork.PlatformTypeAKSHCI),
		// 						RecommendedVersion: to.Ptr(armmobilenetwork.RecommendedVersionRecommended),
		// 						VersionState: to.Ptr(armmobilenetwork.VersionStateActive),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
