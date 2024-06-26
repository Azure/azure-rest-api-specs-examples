import com.azure.core.util.Context;

/** Samples for BackupOperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/Common/ProtectedItem_Delete_OperationResult.json
     */
    /**
     * Sample code: Get Result for Protected Item Delete Operation.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getResultForProtectedItemDeleteOperation(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupOperationResults()
            .getWithResponse(
                "PySDKBackupTestRsVault",
                "PythonSDKBackupTestRg",
                "00000000-0000-0000-0000-000000000000",
                Context.NONE);
    }
}
