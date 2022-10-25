import com.azure.core.util.Context;

/** Samples for ProtectionPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/AzureIaasVm/V2Policy/v2-Get-Policy.json
     */
    /**
     * Sample code: Get Azure IaasVm Enhanced Protection Policy Details.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getAzureIaasVmEnhancedProtectionPolicyDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionPolicies()
            .getWithResponse("NetSDKTestRsVault", "SwaggerTestRg", "v2-daily-sample", Context.NONE);
    }
}
