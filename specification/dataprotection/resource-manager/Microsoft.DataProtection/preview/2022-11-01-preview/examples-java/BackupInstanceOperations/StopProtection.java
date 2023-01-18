/** Samples for BackupInstances StopProtection. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/StopProtection.json
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
