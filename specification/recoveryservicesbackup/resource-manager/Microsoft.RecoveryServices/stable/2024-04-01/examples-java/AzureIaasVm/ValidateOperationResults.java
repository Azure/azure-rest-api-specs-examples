
/**
 * Samples for ValidateOperationResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * AzureIaasVm/ValidateOperationResults.json
     */
    /**
     * Sample code: Get Operation Results of Validate Operation.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getOperationResultsOfValidateOperation(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.validateOperationResults().getWithResponse("NetSDKTestRsVault", "SwaggerTestRg",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
