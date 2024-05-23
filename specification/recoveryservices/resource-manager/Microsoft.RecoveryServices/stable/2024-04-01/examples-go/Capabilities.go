package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/Capabilities.json
func ExampleClient_Capabilities() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Capabilities(ctx, "westus", armrecoveryservices.ResourceCapabilities{
		Type: to.Ptr("Microsoft.RecoveryServices/Vaults"),
		Properties: &armrecoveryservices.CapabilitiesProperties{
			DNSZones: []*armrecoveryservices.DNSZone{
				{
					SubResource: to.Ptr(armrecoveryservices.VaultSubResourceTypeAzureBackup),
				},
				{
					SubResource: to.Ptr(armrecoveryservices.VaultSubResourceTypeAzureSiteRecovery),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CapabilitiesResponse = armrecoveryservices.CapabilitiesResponse{
	// 	Type: to.Ptr("Microsoft.RecoveryServices/Vaults"),
	// 	Properties: &armrecoveryservices.CapabilitiesResponseProperties{
	// 		DNSZones: []*armrecoveryservices.DNSZoneResponse{
	// 			{
	// 				SubResource: to.Ptr(armrecoveryservices.VaultSubResourceTypeAzureBackup),
	// 				RequiredZoneNames: []*string{
	// 					to.Ptr("privatelink.wus.backup.windowsazure.com"),
	// 					to.Ptr("privatelink.queue.core.windows.net"),
	// 					to.Ptr("privatelink.blob.core.windows.net")},
	// 				},
	// 				{
	// 					SubResource: to.Ptr(armrecoveryservices.VaultSubResourceTypeAzureSiteRecovery),
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.siterecovery.windowsazure.com")},
	// 				}},
	// 			},
	// 		}
}
