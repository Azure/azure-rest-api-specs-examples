package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v3"
)

// Generated from example definition: 2026-05-01/PUTVault_WithSourceScanConfiguration.json
func ExampleVaultsClient_BeginCreateOrUpdate_createOrUpdateVaultWithSourceScanConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("77777777-b0c6-47a2-b37c-d8e65a629c18", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVaultsClient().BeginCreateOrUpdate(ctx, "Default-RecoveryServices-ResourceGroup", "swaggerExample", armrecoveryservices.Vault{
		Identity: &armrecoveryservices.IdentityData{
			Type: to.Ptr(armrecoveryservices.ResourceIdentityTypeSystemAssigned),
		},
		Location: to.Ptr("West US"),
		Properties: &armrecoveryservices.VaultProperties{
			PublicNetworkAccess: to.Ptr(armrecoveryservices.PublicNetworkAccessEnabled),
			SecuritySettings: &armrecoveryservices.SecuritySettings{
				SourceScanConfiguration: &armrecoveryservices.SourceScanConfiguration{
					SourceScanIdentity: &armrecoveryservices.AssociatedIdentity{
						OperationIdentityType: to.Ptr(armrecoveryservices.IdentityTypeSystemAssigned),
					},
					State: to.Ptr(armrecoveryservices.StateEnabled),
				},
			},
		},
		SKU: &armrecoveryservices.SKU{
			Name: to.Ptr(armrecoveryservices.SKUNameStandard),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservices.VaultsClientCreateOrUpdateResponse{
	// 	Vault: armrecoveryservices.Vault{
	// 		Name: to.Ptr("swaggerExample"),
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults"),
	// 		Etag: to.Ptr("W/\"datetime'2025-02-15T12%3A36%3A51.68Z'\""),
	// 		ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/resourceGroups/Default-RecoveryServices-ResourceGroup/providers/Microsoft.RecoveryServices/vaults/swaggerExample"),
	// 		Identity: &armrecoveryservices.IdentityData{
	// 			Type: to.Ptr(armrecoveryservices.ResourceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("1be097b0-eb5e-4927-bac2-b24ee8716f64"),
	// 			TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armrecoveryservices.VaultProperties{
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PublicNetworkAccess: to.Ptr(armrecoveryservices.PublicNetworkAccessEnabled),
	// 			SecuritySettings: &armrecoveryservices.SecuritySettings{
	// 				SourceScanConfiguration: &armrecoveryservices.SourceScanConfiguration{
	// 					SourceScanIdentity: &armrecoveryservices.AssociatedIdentity{
	// 						OperationIdentityType: to.Ptr(armrecoveryservices.IdentityTypeSystemAssigned),
	// 					},
	// 					State: to.Ptr(armrecoveryservices.StateEnabled),
	// 				},
	// 			},
	// 		},
	// 		SKU: &armrecoveryservices.SKU{
	// 			Name: to.Ptr(armrecoveryservices.SKUNameStandard),
	// 		},
	// 		Tags: map[string]*string{
	// 			"TestUpdatedKey": to.Ptr("TestUpdatedValue"),
	// 		},
	// 	},
	// }
}
