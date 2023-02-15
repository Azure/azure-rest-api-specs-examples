/** Samples for BackupInstances StopProtection. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-01-01/examples/BackupInstanceOperations/StopProtection.json
     */
    /**
     * Sample code: StopProtection.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void stopProtection(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().stopProtection("testrg", "testvault", "testbi", com.azure.core.util.Context.NONE);
    }
}
