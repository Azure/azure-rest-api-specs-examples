package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager("resourceGroupPS1", nil)
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
		// page.OperationsDiscoveryCollection = armrecoveryservicessiterecovery.OperationsDiscoveryCollection{
		// 	Value: []*armrecoveryservicessiterecovery.OperationsDiscovery{
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/vaultTokens/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The Vault Token operation can be used to get   Vault Token for vault level backend operations."),
		// 				Operation: to.Ptr("Vault Token"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/registeredIdentities/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The Register Service Container   operation can be used to register a container with Recovery Service."),
		// 				Operation: to.Ptr("Register Service Container"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/registeredIdentities/operationResults/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The Get Operation Results operation can   be used get the operation status and result for the asynchronously submitted operation"),
		// 				Operation: to.Ptr("Get Operation Results"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/registeredIdentities/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The Get Containers operation can be used get   the containers registered for a resource."),
		// 				Operation: to.Ptr("Get Containers"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/registeredIdentities/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The UnRegister Container   operation can be used to unregister a container."),
		// 				Operation: to.Ptr("Unregister Service Container"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/certificates/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The Update Resource Certificate   operation updates the resource/vault credential certificate."),
		// 				Operation: to.Ptr("Update Resource Certificate"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationAlertSettings/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Alerts Settings"),
		// 				Operation: to.Ptr("Read Alerts Settings"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Alerts   Settings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationAlertSettings/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create or Update Any Alerts Settings"),
		// 				Operation: to.Ptr("Create or Update Alerts Settings"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Alerts   Settings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationEvents/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Events"),
		// 				Operation: to.Ptr("Read Events"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Events"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationNetworks/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Networks"),
		// 				Operation: to.Ptr("Read Networks"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Networks"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationNetworks/replicationNetworkMappings/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Network   Mappings"),
		// 				Operation: to.Ptr("Read Network Mappings"),
		// 				Provider: to.Ptr("Microsoft   Recovery Services"),
		// 				Resource: to.Ptr("Network Mappings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationNetworks/replicationNetworkMappings/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create or Update Any Network   Mappings"),
		// 				Operation: to.Ptr("Create or Update Network Mappings"),
		// 				Provider: to.Ptr("Microsoft   Recovery Services"),
		// 				Resource: to.Ptr("Network Mappings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationNetworks/replicationNetworkMappings/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Delete Any Network   Mappings"),
		// 				Operation: to.Ptr("Delete Network Mappings"),
		// 				Provider: to.Ptr("Microsoft   Recovery Services"),
		// 				Resource: to.Ptr("Network Mappings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectableItems/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Protectable   Items"),
		// 				Operation: to.Ptr("Read Protectable Items"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protectable Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionContainerMappings/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Protection Container Mappings"),
		// 				Operation: to.Ptr("Read Protection Container   Mappings"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protection Container Mappings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionContainerMappings/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create or Update Any Protection Container Mappings"),
		// 				Operation: to.Ptr("Create or Update Protection Container   Mappings"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protection Container Mappings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionContainerMappings/remove/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Remove Protection Container Mapping"),
		// 				Operation: to.Ptr("Remove Protection Container   Mapping"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protection Container Mappings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionContainerMappings/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Delete Any Protection Container Mappings"),
		// 				Operation: to.Ptr("Delete Protection Container   Mappings"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protection Container Mappings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/recoveryPoints/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read   Any Replication Recovery Points"),
		// 				Operation: to.Ptr("Read Replication Recovery Points"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Replication Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Protected   Items"),
		// 				Operation: to.Ptr("Read Protected Items"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create or Update   Any Protected Items"),
		// 				Operation: to.Ptr("Create or Update Protected Items"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Delete Any Protected   Items"),
		// 				Operation: to.Ptr("Delete Protected Items"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/remove/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Remove Protected   Item"),
		// 				Operation: to.Ptr("Remove Protected Item"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/plannedFailover/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Planned   Failover"),
		// 				Operation: to.Ptr("Planned Failover"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/unplannedFailover/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Failover"),
		// 				Operation: to.Ptr("Failover"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/testFailover/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Test   Failover"),
		// 				Operation: to.Ptr("Test Failover"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/testFailoverCleanup/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Test Failover   Cleanup"),
		// 				Operation: to.Ptr("Test Failover Cleanup"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/failoverCommit/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Failover   Commit"),
		// 				Operation: to.Ptr("Failover Commit"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/reProtect/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("ReProtect Protected   Item"),
		// 				Operation: to.Ptr("ReProtect Protected Item"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/updateMobilityService/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Update Mobility   Service"),
		// 				Operation: to.Ptr("Update Mobility Service"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/repairReplication/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Repair   replication"),
		// 				Operation: to.Ptr("Repair replication"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/applyRecoveryPoint/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Apply Recovery   Point"),
		// 				Operation: to.Ptr("Apply Recovery Point"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationJobs/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Jobs"),
		// 				Operation: to.Ptr("Read Jobs"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationJobs/cancel/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Cancel Job"),
		// 				Operation: to.Ptr("Cancel Job"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationJobs/restart/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Restart job"),
		// 				Operation: to.Ptr("Restart job"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationJobs/resume/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Resume Job"),
		// 				Operation: to.Ptr("Resume Job"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Protection   Containers"),
		// 				Operation: to.Ptr("Read Protection Containers"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Protection Containers"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/discoverProtectableItem/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Discover   Protectable Item"),
		// 				Operation: to.Ptr("Discover Protectable Item"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Protection Containers"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create or Update Any Protection   Containers"),
		// 				Operation: to.Ptr("Create or Update Protection Containers"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Protection Containers"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/remove/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Remove Protection   Container"),
		// 				Operation: to.Ptr("Remove Protection Container"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Protection Containers"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/switchprotection/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Switch Protection   Container"),
		// 				Operation: to.Ptr("Switch Protection Container"),
		// 				Provider: to.Ptr("Microsoft   Recovery Services"),
		// 				Resource: to.Ptr("Protection Containers"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationPolicies/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Policies"),
		// 				Operation: to.Ptr("Read Policies"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Policies"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationPolicies/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create or Update Any Policies"),
		// 				Operation: to.Ptr("Create or Update Policies"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Policies"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationPolicies/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Delete Any Policies"),
		// 				Operation: to.Ptr("Delete Policies"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Policies"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Recovery Plans"),
		// 				Operation: to.Ptr("Read Recovery Plans"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Recovery   Plans"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create or Update Any Recovery Plans"),
		// 				Operation: to.Ptr("Create or Update Recovery Plans"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Recovery   Plans"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Delete Any Recovery Plans"),
		// 				Operation: to.Ptr("Delete Recovery Plans"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Recovery   Plans"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans/plannedFailover/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Planned Failover Recovery   Plan"),
		// 				Operation: to.Ptr("Planned Failover Recovery Plan"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Recovery Plans"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans/unplannedFailover/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Failover Recovery   Plan"),
		// 				Operation: to.Ptr("Failover Recovery Plan"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Recovery Plans"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans/testFailover/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Test Failover Recovery   Plan"),
		// 				Operation: to.Ptr("Test Failover Recovery Plan"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Recovery Plans"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans/testFailoverCleanup/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Test Failover   Cleanup Recovery Plan"),
		// 				Operation: to.Ptr("Test Failover Cleanup Recovery Plan"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Recovery Plans"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans/failoverCommit/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Failover Commit   Recovery Plan"),
		// 				Operation: to.Ptr("Failover Commit Recovery Plan"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Recovery Plans"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans/reProtect/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("ReProtect Recovery   Plan"),
		// 				Operation: to.Ptr("ReProtect Recovery Plan"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Recovery Plans"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationRecoveryServicesProviders/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read   Any Recovery Services Providers"),
		// 				Operation: to.Ptr("Read Recovery Services Providers"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Recovery Services Providers"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationRecoveryServicesProviders/remove/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Remove Recovery Services   Provider"),
		// 				Operation: to.Ptr("Remove Recovery Services Provider"),
		// 				Provider: to.Ptr("Microsoft   Recovery Services"),
		// 				Resource: to.Ptr("Recovery Services Providers"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationRecoveryServicesProviders/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Delete Any Recovery Services   Providers"),
		// 				Operation: to.Ptr("Delete Recovery Services Providers"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Recovery Services Providers"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationRecoveryServicesProviders/refreshProvider/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Refresh   Provider"),
		// 				Operation: to.Ptr("Refresh Provider"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Recovery Services Providers"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Fabrics"),
		// 				Operation: to.Ptr("Read Fabrics"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Fabrics"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create or Update Any Fabrics"),
		// 				Operation: to.Ptr("Create or Update Fabrics"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Fabrics"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/remove/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Remove Fabric"),
		// 				Operation: to.Ptr("Remove Fabric"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Fabrics"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/checkConsistency/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Checks Consistency of the   Fabric"),
		// 				Operation: to.Ptr("Checks Consistency of the Fabric"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Fabrics"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Delete Any Fabrics"),
		// 				Operation: to.Ptr("Delete Fabrics"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Fabrics"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/renewcertificate/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Renew Certificate for Fabric"),
		// 				Operation: to.Ptr("Renew Certificate for Fabric"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Fabrics"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/deployProcessServerImage/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Deploy Process Server Image"),
		// 				Operation: to.Ptr("Deploy Process Server Image"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Fabrics"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/reassociateGateway/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Reassociate Gateway"),
		// 				Operation: to.Ptr("Reassociate Gateway"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Fabrics"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationStorageClassifications/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Storage   Classifications"),
		// 				Operation: to.Ptr("Read Storage Classifications"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Storage Classifications"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationStorageClassifications/replicationStorageClassificationMappings/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Storage Classification Mappings"),
		// 				Operation: to.Ptr("Read Storage Classification   Mappings"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Storage Classification Mappings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationStorageClassifications/replicationStorageClassificationMappings/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create or Update Any Storage Classification Mappings"),
		// 				Operation: to.Ptr("Create or Update Storage Classification   Mappings"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Storage Classification Mappings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationStorageClassifications/replicationStorageClassificationMappings/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Delete Any Storage Classification Mappings"),
		// 				Operation: to.Ptr("Delete Storage Classification   Mappings"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Storage Classification Mappings"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/usages/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Vault Usages"),
		// 				Operation: to.Ptr("Read   Vault Usages"),
		// 				Provider: to.Ptr("Microsoft Recovery Services"),
		// 				Resource: to.Ptr("Vault Usages"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationvCenters/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Read Any Jobs"),
		// 				Operation: to.Ptr("Read Jobs"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationvCenters/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create or Update Any Jobs"),
		// 				Operation: to.Ptr("Create or Update Jobs"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationvCenters/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Delete Any Jobs"),
		// 				Operation: to.Ptr("Delete Jobs"),
		// 				Provider: to.Ptr("Microsoft Recovery   Services"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/usages/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns usage details for a Recovery Services Vault."),
		// 				Operation: to.Ptr("Recovery Services Vault usage details."),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vault   Usage"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupUsageSummaries/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns summaries for   Protected Items and Protected Servers for a Recovery Services ."),
		// 				Operation: to.Ptr("Recovery Services Protected Items and Protected Servers usage summaries details."),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup Usages   Summaries"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/storageConfig/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns Storage Configuration for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Resource Storage Config"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vault Storage   Config"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/storageConfig/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Updates Storage Configuration for Recovery Services Vault."),
		// 				Operation: to.Ptr("Write Resource Storage Config"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vault Storage   Config"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupconfig/vaultconfig/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns Configuration for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Resource Config"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vault   Config"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupconfig/vaultconfig/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Updates Configuration for Recovery Services Vault."),
		// 				Operation: to.Ptr("Update Resource Config"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vault   Config"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/tokenInfo/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns token information for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get   Vault Token Info"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Token Info"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupSecurityPIN/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns Security PIN   Information for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Security PIN Info"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("SecurityPINInfo"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupManagementMetaData/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns   Backup Management Metadata for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Backup Management Metadata"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup Management Metadata"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupOperationResults/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns Backup Operation Result for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Backup Operation Result"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup Operation   Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupOperations/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns Backup Operation Status for Recovery Services Vault."),
		// 				Operation: to.Ptr("Get Backup Operation Status"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup Operation   Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupJobs/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns all Job Objects"),
		// 				Operation: to.Ptr("Get Jobs"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup   Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupJobs/cancel/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Cancel the Job"),
		// 				Operation: to.Ptr("Cancel Jobs"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup   Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupJobsExport/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Export Jobs"),
		// 				Operation: to.Ptr("Export Jobs"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Export Backup   Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupJobs/operationResults/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns the Result of Job Operation."),
		// 				Operation: to.Ptr("Get Job Operation Result"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup Jobs   Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupJobsExport/operationResults/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns the Result of Export Job Operation."),
		// 				Operation: to.Ptr("Get Export Job Operation   Result"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Export Backup Jobs Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Get Recovery Points for   Protected Items."),
		// 				Operation: to.Ptr("Get Recovery Points"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints/restore/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Restore Recovery Points for   Protected Items."),
		// 				Operation: to.Ptr("Restore Recovery Points"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints/provisionInstantItemRecovery/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Provision Instant Item Recovery for Protected Item"),
		// 				Operation: to.Ptr("Provision Instant Item Recovery for Protected   Item"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints/revokeInstantItemRecovery/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Revoke Instant Item Recovery for Protected Item"),
		// 				Operation: to.Ptr("Revoke Instant Item Recovery for Protected   Item"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupPolicies/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns all Protection Policies"),
		// 				Operation: to.Ptr("Get Protection Policy"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup   Policies"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupPolicies/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Creates Protection Policy"),
		// 				Operation: to.Ptr("Create Protection Policy"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup   Policies"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupPolicies/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Delete a Protection Policy"),
		// 				Operation: to.Ptr("Delete Protection Policy"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup   Policies"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupPolicies/operationResults/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Get Results of Policy Operation."),
		// 				Operation: to.Ptr("Get Policy Operation Results"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup   Policy Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupPolicies/operationStatus/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Get Status of Policy Operation."),
		// 				Operation: to.Ptr("Get Policy Operation Status"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup   Policy Operation Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/protectedItems/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns object details of   the Protected Item"),
		// 				Operation: to.Ptr("Get Protected Item Details"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupProtectedItems/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns the list of all   Protected Items."),
		// 				Operation: to.Ptr("Get All Protected Items"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/protectedItems/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create a backup   Protected Item"),
		// 				Operation: to.Ptr("Create Backup Protected Item"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/protectedItems/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Deletes Protected   Item"),
		// 				Operation: to.Ptr("Delete Protected Items"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/protectedItems/operationResults/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Gets Result of Operation Performed on Protected Items."),
		// 				Operation: to.Ptr("Get Protected Items Operation   Results"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Protected Item Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/protectedItems/operationStatus/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns the status of Operation performed on Protected Items."),
		// 				Operation: to.Ptr("Get Protected Items operation   status"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Protected Item Operation Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/protectedItems/backup/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Performs Backup for Protected   Item."),
		// 				Operation: to.Ptr("Backup Protected Item"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Protected Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupProtectableItems/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns list of all   Protectable Items."),
		// 				Operation: to.Ptr("Get Protectable Items"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup Protectable Items"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/refreshContainers/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Refreshes the container   list"),
		// 				Operation: to.Ptr("Refresh container"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Refresh Containers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/operationResults/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns   status of the operation"),
		// 				Operation: to.Ptr("Get Operation Results"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Refresh Containers Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupProtectionContainers/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns   all containers belonging to the subscription"),
		// 				Operation: to.Ptr("Get Containers In Subscription"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup Protection Containers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns all   registered containers"),
		// 				Operation: to.Ptr("Get Registered Container"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Protection Containers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupFabrics/protectionContainers/operationResults/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Gets result of Operation performed on Protection Container."),
		// 				Operation: to.Ptr("Get Container Operation   Results"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Protection Containers Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupEngines"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Returns all the backup management servers registered with   vault."),
		// 				Operation: to.Ptr("List of backup management servers."),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup   Engines"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupStatus"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Check Backup Status for   Recovery Services Vaults"),
		// 				Operation: to.Ptr("Check Backup Status for Vault"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Backup Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("Create Vault operation creates an Azure resource   of type 'vault'"),
		// 				Operation: to.Ptr("Create Vault"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The Get Vault operation gets an object representing   the Azure resource of type 'vault'"),
		// 				Operation: to.Ptr("Get Vault"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The Delete Vault operation deletes the specified   Azure resource of type 'vault'"),
		// 				Operation: to.Ptr("Delete Vault"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/extendedInformation/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The Get Extended Info operation gets an   object's Extended Info representing the Azure resource of type ?vault?"),
		// 				Operation: to.Ptr("Get Extended Info"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/extendedInformation/write"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The Get Extended Info operation gets an   object's Extended Info representing the Azure resource of type ?vault?"),
		// 				Operation: to.Ptr("Get Extended Info"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/extendedInformation/delete"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("The Get Extended Info operation gets an   object's Extended Info representing the Azure resource of type ?vault?"),
		// 				Operation: to.Ptr("Get Extended Info"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/locations/allocatedStamp/read"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("GetAllocatedStamp is   internal operation used by service"),
		// 				Operation: to.Ptr("Get Allocated Stamp"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("locations/allocatedStamp"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/locations/allocateStamp/action"),
		// 			Display: &armrecoveryservicessiterecovery.Display{
		// 				Description: to.Ptr("AllocateStamp is   internal operation used by service"),
		// 				Operation: to.Ptr("Allocated Stamp Action"),
		// 				Provider: to.Ptr("Microsoft.RecoveryServices"),
		// 				Resource: to.Ptr("locations/allocateStamp"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 	}},
		// }
	}
}
