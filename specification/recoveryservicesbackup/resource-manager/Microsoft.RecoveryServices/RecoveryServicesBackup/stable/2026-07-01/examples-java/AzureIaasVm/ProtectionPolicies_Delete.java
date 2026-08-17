
/**
 * Samples for ProtectionPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/AzureIaasVm/ProtectionPolicies_Delete.json
     */
    /**
     * Sample code: Delete Azure Vm Protection Policy.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void deleteAzureVmProtectionPolicy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionPolicies().delete("NetSDKTestRsVault", "SwaggerTestRg", "testPolicy1",
            com.azure.core.util.Context.NONE);
    }
}
