
/**
 * Samples for ProtectionPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * AzureIaasVm/ProtectionPolicies_Get.json
     */
    /**
     * Sample code: Get Azure IaasVm Protection Policy Details.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getAzureIaasVmProtectionPolicyDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionPolicies().getWithResponse("NetSDKTestRsVault", "SwaggerTestRg", "testPolicy1",
            com.azure.core.util.Context.NONE);
    }
}
