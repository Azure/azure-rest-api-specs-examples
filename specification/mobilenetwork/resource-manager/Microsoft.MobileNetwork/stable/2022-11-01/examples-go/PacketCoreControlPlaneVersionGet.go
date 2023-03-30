package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/340d577969b7bff5ad0488d79543314bc17daa50/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/PacketCoreControlPlaneVersionGet.json
func ExamplePacketCoreControlPlaneVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPacketCoreControlPlaneVersionsClient().Get(ctx, "PMN-4-11-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PacketCoreControlPlaneVersion = armmobilenetwork.PacketCoreControlPlaneVersion{
	// 	Name: to.Ptr("PMN-4-11-1"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlaneVersions"),
	// 	ID: to.Ptr("/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/PMN-4-11-1"),
	// 	Properties: &armmobilenetwork.PacketCoreControlPlaneVersionPropertiesFormat{
	// 		Platforms: []*armmobilenetwork.Platform{
	// 			{
	// 				MaximumPlatformSoftwareVersion: to.Ptr("2211"),
	// 				MinimumPlatformSoftwareVersion: to.Ptr("2209"),
	// 				PlatformType: to.Ptr(armmobilenetwork.PlatformTypeAKSHCI),
	// 				RecommendedVersion: to.Ptr(armmobilenetwork.RecommendedVersionRecommended),
	// 				VersionState: to.Ptr(armmobilenetwork.VersionStateActive),
	// 		}},
	// 	},
	// }
}
