/** Samples for BackupInstances ResumeBackups. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/BackupInstanceOperations/ResumeBackups.json
     */
    /**
     * Sample code: ResumeBackups.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void resumeBackups(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().resumeBackups("testrg", "testvault", "testbi", com.azure.core.util.Context.NONE);
    }
}
