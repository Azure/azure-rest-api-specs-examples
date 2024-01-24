
/**
 * Samples for BmsPrepareDataMoveOperationResult Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * BackupDataMove/PrepareDataMoveOperationResult_Get.json
     */
    /**
     * Sample code: Get operation result for PrepareDataMove.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getOperationResultForPrepareDataMove(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.bmsPrepareDataMoveOperationResults().getWithResponse("source-rsv", "sourceRG",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
