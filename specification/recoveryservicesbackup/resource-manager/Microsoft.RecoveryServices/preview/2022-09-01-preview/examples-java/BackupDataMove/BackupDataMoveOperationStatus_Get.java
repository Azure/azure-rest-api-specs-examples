import com.azure.core.util.Context;

/** Samples for ResourceProvider GetOperationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/BackupDataMove/BackupDataMoveOperationStatus_Get.json
     */
    /**
     * Sample code: Get OperationStatus.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getOperationStatus(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .resourceProviders()
            .getOperationStatusWithResponse(
                "source-rsv", "sourceRG", "00000000-0000-0000-0000-000000000000", Context.NONE);
    }
}
