import com.azure.core.util.Context;

/** Samples for ProtectionPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/AzureIaasVm/ProtectionPolicies_Get.json
     */
    /**
     * Sample code: Get Azure IaasVm Protection Policy Details.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getAzureIaasVmProtectionPolicyDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionPolicies().getWithResponse("NetSDKTestRsVault", "SwaggerTestRg", "testPolicy1", Context.NONE);
    }
}
