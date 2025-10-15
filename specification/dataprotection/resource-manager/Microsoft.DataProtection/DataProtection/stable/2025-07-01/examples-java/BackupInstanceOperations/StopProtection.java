
/**
 * Samples for BackupInstances StopProtection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BackupInstanceOperations/StopProtection.json
     */
    /**
     * Sample code: StopProtection.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void stopProtection(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().stopProtection("testrg", "testvault", "testbi", null,
            com.azure.core.util.Context.NONE);
    }
}
