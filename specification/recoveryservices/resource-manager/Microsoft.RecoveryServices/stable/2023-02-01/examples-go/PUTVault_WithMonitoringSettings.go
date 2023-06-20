package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/PUTVault_WithMonitoringSettings.json
func ExampleVaultsClient_BeginCreateOrUpdate_createOrUpdateVaultWithMonitoringSetting() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVaultsClient().BeginCreateOrUpdate(ctx, "Default-RecoveryServices-ResourceGroup", "swaggerExample", armrecoveryservices.Vault{
		Location: to.Ptr("West US"),
		Identity: &armrecoveryservices.IdentityData{
			Type: to.Ptr(armrecoveryservices.ResourceIdentityTypeSystemAssigned),
		},
		Properties: &armrecoveryservices.VaultProperties{
			MonitoringSettings: &armrecoveryservices.MonitoringSettings{
				AzureMonitorAlertSettings: &armrecoveryservices.AzureMonitorAlertSettings{
					AlertsForAllJobFailures: to.Ptr(armrecoveryservices.AlertsStateEnabled),
				},
				ClassicAlertSettings: &armrecoveryservices.ClassicAlertSettings{
					AlertsForCriticalOperations: to.Ptr(armrecoveryservices.AlertsStateDisabled),
				},
			},
			PublicNetworkAccess: to.Ptr(armrecoveryservices.PublicNetworkAccessEnabled),
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
		log.Fatalf("failed to pull the result: %v", err)
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
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicNetworkAccess: to.Ptr(armrecoveryservices.PublicNetworkAccessEnabled),
	// 	},
	// 	SKU: &armrecoveryservices.SKU{
	// 		Name: to.Ptr(armrecoveryservices.SKUNameRS0),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}
