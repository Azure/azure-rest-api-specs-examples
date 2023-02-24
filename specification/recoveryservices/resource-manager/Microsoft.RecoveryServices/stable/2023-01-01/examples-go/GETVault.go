package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/GETVault.json
func ExampleVaultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservices.NewVaultsClient("77777777-b0c6-47a2-b37c-d8e65a629c18", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "Default-RecoveryServices-ResourceGroup", "swaggerExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Vault = armrecoveryservices.Vault{
	// 	Name: to.Ptr("swaggerExample"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults"),
	// 	Etag: to.Ptr("W/\"datetime'2017-12-15T12%3A36%3A51.68Z'\""),
	// 	ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/resourceGroups/Default-RecoveryServices-ResourceGroup/providers/Microsoft.RecoveryServices/vaults/swaggerExample"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"TestUpdatedKey": to.Ptr("TestUpdatedValue"),
	// 	},
	// 	Identity: &armrecoveryservices.IdentityData{
	// 		Type: to.Ptr(armrecoveryservices.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("3137d6c7-5d6c-411c-b934-7a2a729ee247"),
	// 		TenantID: to.Ptr("d676e86e-2206-4a7c-999c-ece52c144b5b"),
	// 	},
	// 	Properties: &armrecoveryservices.VaultProperties{
	// 		MonitoringSettings: &armrecoveryservices.MonitoringSettings{
	// 			AzureMonitorAlertSettings: &armrecoveryservices.AzureMonitorAlertSettings{
	// 				AlertsForAllJobFailures: to.Ptr(armrecoveryservices.AlertsStateEnabled),
	// 			},
	// 			ClassicAlertSettings: &armrecoveryservices.ClassicAlertSettings{
	// 				AlertsForCriticalOperations: to.Ptr(armrecoveryservices.AlertsStateDisabled),
	// 			},
	// 		},
	// 		PrivateEndpointConnections: []*armrecoveryservices.PrivateEndpointConnectionVaultProperties{
	// 			{
	// 				ID: to.Ptr("/subscriptions/6c48fa17-39c7-45f1-90ac-47a587128ace/resourceGroups/Default-RecoveryServices-ResourceGroup/providers/Microsoft.RecoveryServices/Vaults/pemsi-ecy-rsv2/privateEndpointConnections/pe114-pemsi-ecy-rsv.5944358949303501042.backup.75061caa-cba4-4849-8e09-608da4914aad"),
	// 				Properties: &armrecoveryservices.PrivateEndpointConnection{
	// 					GroupIDs: []*armrecoveryservices.VaultSubResourceType{
	// 						to.Ptr(armrecoveryservices.VaultSubResourceTypeAzureBackup)},
	// 						PrivateEndpoint: &armrecoveryservices.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/6c48fa17-39c7-45f1-90ac-47a587128ace/resourceGroups/Default-RecoveryServices-ResourceGroup/providers/Microsoft.Network/privateEndpoints/pe114-pemsi-ecy-rsv"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armrecoveryservices.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("None"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armrecoveryservices.PrivateEndpointConnectionStatusApproved),
	// 						},
	// 						ProvisioningState: to.Ptr(armrecoveryservices.ProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			PrivateEndpointStateForBackup: to.Ptr(armrecoveryservices.VaultPrivateEndpointStateEnabled),
	// 			PrivateEndpointStateForSiteRecovery: to.Ptr(armrecoveryservices.VaultPrivateEndpointStateNone),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PublicNetworkAccess: to.Ptr(armrecoveryservices.PublicNetworkAccessEnabled),
	// 			RedundancySettings: &armrecoveryservices.VaultPropertiesRedundancySettings{
	// 				CrossRegionRestore: to.Ptr(armrecoveryservices.CrossRegionRestoreEnabled),
	// 				StandardTierStorageRedundancy: to.Ptr(armrecoveryservices.StandardTierStorageRedundancyGeoRedundant),
	// 			},
	// 		},
	// 		SKU: &armrecoveryservices.SKU{
	// 			Name: to.Ptr(armrecoveryservices.SKUNameStandard),
	// 		},
	// 	}
}
