package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/ListOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.ClientDiscoveryResponse = armrecoveryservices.ClientDiscoveryResponse{
		// 	Value: []*armrecoveryservices.ClientDiscoveryValueForSingleAPI{
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/usages/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns usage details for a Recovery Services Vault."),
		// 				Operation: to.Ptr("Recovery Services Vault usage details."),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Vault Usage"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupUsageSummaries/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns summaries for Protected Items and Protected Servers for a Recovery Services ."),
		// 				Operation: to.Ptr("Recovery Services Protected Items and Protected Servers usage summaries details."),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Usages Summaries"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/storageConfig/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Storage Configuration for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Resource Storage Config"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Vault Storage Config"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/storageConfig/write"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Updates Storage Configuration for Recovery Services Vault."),
		// 				Operation: to.Ptr("Write Resource Storage Config"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Vault Storage Config"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupconfig/vaultconfig/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Configuration for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Resource Config"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Vault Config"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupconfig/vaultconfig/write"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Updates Configuration for Recovery Services Vault."),
		// 				Operation: to.Ptr("Update Resource Config"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Vault Config"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/tokenInfo/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns token information for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Vault Token Info"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Token Info"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupSecurityPIN/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Security PIN Information for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Security PIN Info"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("SecurityPINInfo"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupManagementMetaData/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Backup Management Metadata for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Backup Management Metadata"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Management Metadata"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupOperationResults/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Backup Operation Result for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Backup Operation Result"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupOperations/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Backup Operation Status for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Backup Operation Status"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Operation Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupJobs/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns all Job Objects"),
		// 				Operation: to.Ptr("Get Jobs"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupJobs/cancel/action"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Cancel the Job"),
		// 				Operation: to.Ptr("Cancel Jobs"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupJobsExport/action"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Export Jobs"),
		// 				Operation: to.Ptr("Export Jobs"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Export Backup Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupJobs/operationResults/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns the Result of Job Operation."),
		// 				Operation: to.Ptr("Get Job Operation Result"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Jobs Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupJobsExport/operationResults/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns the Result of Export Job Operation."),
		// 				Operation: to.Ptr("Get Export Job Operation Result"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Export Backup Jobs Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Get Recovery Points for Protected Items."),
		// 				Operation: to.Ptr("Get Recovery Points"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints/restore/action"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Restore Recovery Points for Protected Items."),
		// 				Operation: to.Ptr("Restore Recovery Points"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints/provisionInstantItemRecovery/action"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Provision Instant Item Recovery for Protected Item"),
		// 				Operation: to.Ptr("Provision Instant Item Recovery for Protected Item"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints/revokeInstantItemRecovery/action"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Revoke Instant Item Recovery for Protected Item"),
		// 				Operation: to.Ptr("Revoke Instant Item Recovery for Protected Item"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupPolicies/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns all Protection Policies"),
		// 				Operation: to.Ptr("Get Protection Policy"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Policies"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupPolicies/write"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Creates Protection Policy"),
		// 				Operation: to.Ptr("Create Protection Policy"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Policies"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupPolicies/delete"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Delete a Protection Policy"),
		// 				Operation: to.Ptr("Delete Protection Policy"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Policies"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupPolicies/operationResults/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Get Results of Policy Operation."),
		// 				Operation: to.Ptr("Get Policy Operation Results"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Policy Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupPolicies/operationsStatus/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Get Status of Policy Operation."),
		// 				Operation: to.Ptr("Get Policy Operation Status"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Policy Operation Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/protectedItems/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns object details of the Protected Item"),
		// 				Operation: to.Ptr("Get Protected Item Details"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupProtectedItems/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns the list of all Protected Items."),
		// 				Operation: to.Ptr("Get All Protected Items"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/protectedItems/write"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Create a backup Protected Item"),
		// 				Operation: to.Ptr("Create Backup Protected Item"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/protectedItems/delete"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Deletes Protected Item"),
		// 				Operation: to.Ptr("Delete Protected Items"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/protectedItems/operationResults/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Gets Result of Operation Performed on Protected Items."),
		// 				Operation: to.Ptr("Get Protected Items Operation Results"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protected Item Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/protectedItems/operationsStatus/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns the status of Operation performed on Protected Items."),
		// 				Operation: to.Ptr("Get Protected Items operation status"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protected Item Operation Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/protectedItems/backup/action"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Performs Backup for Protected Item."),
		// 				Operation: to.Ptr("Backup Protected Item"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupProtectableItems/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns list of all Protectable Items."),
		// 				Operation: to.Ptr("Get Protectable Items"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Protectable Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/refreshContainers/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Refreshes the container list"),
		// 				Operation: to.Ptr("Refresh container"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Refresh Containers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/operationResults/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns status of the operation"),
		// 				Operation: to.Ptr("Get Operation Results"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Refresh Containers Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupProtectionContainers/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns all containers belonging to the subscription"),
		// 				Operation: to.Ptr("Get Containers In Subscription"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Protection Containers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns all registered containers"),
		// 				Operation: to.Ptr("Get Registered Container"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protection Containers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/operationResults/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Gets result of Operation performed on Protection Container."),
		// 				Operation: to.Ptr("Get Container Operation Results"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protection Containers Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupEngines"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns all the backup management servers registered with vault."),
		// 				Operation: to.Ptr("List of backup management servers."),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Engines"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupStatus"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Check Backup Status for Recovery Services Vaults"),
		// 				Operation: to.Ptr("Check Backup Status for Vault"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Backup Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupPreValidateProtection"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr(""),
		// 				Operation: to.Ptr("Pre Validate Enable Protection"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("PreValidate Protection"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupValidateFeatures"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Validate Features"),
		// 				Operation: to.Ptr("Validate Features"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Validate Features"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/backupProtectionIntent/write"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Create a backup Protection Intent"),
		// 				Operation: to.Ptr("Create backup Protection Intent"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protection Intent"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/{fabricName}/protectionContainers/{containerName}/items/read"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Get all items in a container"),
		// 				Operation: to.Ptr("Get all items in a container"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Workload Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.recoveryservices/vaults/backupFabrics/protectionContainers/inquire/action"),
		// 			Display: &armrecoveryservices.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Get all items in a container"),
		// 				Operation: to.Ptr("Get all items in a container"),
		// 				Provider: to.Ptr("microsoft.recoveryservices"),
		// 				Resource: to.Ptr("Protection Containers Inquire"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 	}},
		// }
	}
}
