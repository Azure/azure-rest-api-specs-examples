
/**
 * Samples for BackupInstances ResumeBackups.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BackupInstanceOperations/ResumeBackups.json
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
