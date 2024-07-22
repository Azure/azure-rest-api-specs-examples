
/**
 * Samples for BackupInstances SuspendBackups.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * BackupInstanceOperations/SuspendBackups.json
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
