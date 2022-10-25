import com.azure.core.util.Context;

/** Samples for ProtectionPolicyOperationStatuses Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/AzureIaasVm/ProtectionPolicyOperationStatuses_Get.json
     */
    /**
     * Sample code: Get Protection Policy Operation Status.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectionPolicyOperationStatus(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionPolicyOperationStatuses()
            .getWithResponse(
                "NetSDKTestRsVault",
                "SwaggerTestRg",
                "testPolicy1",
                "00000000-0000-0000-0000-000000000000",
                Context.NONE);
    }
}
