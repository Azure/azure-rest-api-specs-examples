
/**
 * Samples for ValidateOperationStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * AzureIaasVm/ValidateOperationStatus.json
     */
    /**
     * Sample code: Get Operation Status of Validate Operation.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getOperationStatusOfValidateOperation(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.validateOperationStatuses().getWithResponse("NetSDKTestRsVault", "SwaggerTestRg",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
