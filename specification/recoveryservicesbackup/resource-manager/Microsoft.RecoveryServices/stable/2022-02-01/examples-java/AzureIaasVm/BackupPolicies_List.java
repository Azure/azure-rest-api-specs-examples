import com.azure.core.util.Context;

/** Samples for BackupPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureIaasVm/BackupPolicies_List.json
     */
    /**
     * Sample code: List protection policies with backupManagementType filter as AzureIaasVm.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectionPoliciesWithBackupManagementTypeFilterAsAzureIaasVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupPolicies()
            .list("NetSDKTestRsVault", "SwaggerTestRg", "backupManagementType eq 'AzureIaasVM'", Context.NONE);
    }
}
