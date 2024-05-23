
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureResourceProtectionIntent;

/**
 * Samples for ProtectionIntent CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * AzureIaasVm/ProtectionIntent_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Azure Vm Protection Intent.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void createOrUpdateAzureVmProtectionIntent(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionIntents().define("vm;iaasvmcontainerv2;chamsrgtest;chamscandel").withRegion((String) null)
            .withExistingBackupFabric("myVault", "myRG", "Azure")
            .withProperties(new AzureResourceProtectionIntent().withSourceResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/chamsrgtest/providers/Microsoft.Compute/virtualMachines/chamscandel")
                .withPolicyId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupPolicies/myPolicy"))
            .create();
    }
}
