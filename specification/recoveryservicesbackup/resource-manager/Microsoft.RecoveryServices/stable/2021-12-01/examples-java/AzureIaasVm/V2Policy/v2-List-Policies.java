import com.azure.core.util.Context;

/** Samples for BackupPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/V2Policy/v2-List-Policies.json
     */
    /**
     * Sample code: List protection policies with backupManagementType filter as AzureIaasVm with both V1 and V2
     * policies.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectionPoliciesWithBackupManagementTypeFilterAsAzureIaasVmWithBothV1AndV2Policies(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupPolicies()
            .list("NetSDKTestRsVault", "SwaggerTestRg", "backupManagementType eq 'AzureIaasVM'", Context.NONE);
    }
}
