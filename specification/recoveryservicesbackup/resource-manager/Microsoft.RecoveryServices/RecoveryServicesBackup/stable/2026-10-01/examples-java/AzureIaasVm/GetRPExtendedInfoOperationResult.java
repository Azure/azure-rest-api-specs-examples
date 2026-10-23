
/**
 * Samples for RecoveryPoints GetRPExtendedInfoOperationResult.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AzureIaasVm/GetRPExtendedInfoOperationResult.json
     */
    /**
     * Sample code: Get RP Extended Info Operation Result.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getRPExtendedInfoOperationResult(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.recoveryPoints().getRPExtendedInfoOperationResultWithResponse("testRG", "testVault", "Azure",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
