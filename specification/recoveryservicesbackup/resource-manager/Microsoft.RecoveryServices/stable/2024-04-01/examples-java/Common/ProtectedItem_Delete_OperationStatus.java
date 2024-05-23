
/**
 * Samples for BackupOperationStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * Common/ProtectedItem_Delete_OperationStatus.json
     */
    /**
     * Sample code: Get Protected Item Delete Operation Status.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectedItemDeleteOperationStatus(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupOperationStatuses().getWithResponse("PySDKBackupTestRsVault", "PythonSDKBackupTestRg",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
