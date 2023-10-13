/** Samples for BackupOperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/Common/ProtectedItem_Delete_OperationResult.json
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
                com.azure.core.util.Context.NONE);
    }
}
