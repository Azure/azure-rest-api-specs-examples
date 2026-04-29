
/**
 * Samples for ResourceProvider GetOperationStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/BackupDataMove/BackupDataMoveOperationStatus_Get.json
     */
    /**
     * Sample code: Get OperationStatus.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        getOperationStatus(com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.resourceProviders().getOperationStatusWithResponse("source-rsv", "sourceRG",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
