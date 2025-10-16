
/**
 * Samples for BackupInstances ResumeProtection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BackupInstanceOperations/ResumeProtection.json
     */
    /**
     * Sample code: ResumeProtection.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void resumeProtection(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().resumeProtection("testrg", "testvault", "testbi", com.azure.core.util.Context.NONE);
    }
}
