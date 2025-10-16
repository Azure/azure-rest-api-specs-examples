
/**
 * Samples for BackupInstances SuspendBackups.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BackupInstanceOperations/SuspendBackups.json
     */
    /**
     * Sample code: SuspendBackups.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void suspendBackups(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().suspendBackups("testrg", "testvault", "testbi", null,
            com.azure.core.util.Context.NONE);
    }
}
