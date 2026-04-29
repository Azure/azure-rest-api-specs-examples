
/**
 * Samples for BackupOperationResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Common/ProtectedItem_Delete_OperationResult.json
     */
    /**
     * Sample code: Get Result for Protected Item Delete Operation.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getResultForProtectedItemDeleteOperation(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupOperationResults().getWithResponse("PySDKBackupTestRsVault", "PythonSDKBackupTestRg",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
