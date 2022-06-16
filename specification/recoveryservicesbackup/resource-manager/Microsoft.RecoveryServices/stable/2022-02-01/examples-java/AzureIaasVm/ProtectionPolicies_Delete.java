import com.azure.core.util.Context;

/** Samples for ProtectionPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureIaasVm/ProtectionPolicies_Delete.json
     */
    /**
     * Sample code: Delete Azure Vm Protection Policy.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void deleteAzureVmProtectionPolicy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionPolicies().delete("NetSDKTestRsVault", "SwaggerTestRg", "testPolicy1", Context.NONE);
    }
}
