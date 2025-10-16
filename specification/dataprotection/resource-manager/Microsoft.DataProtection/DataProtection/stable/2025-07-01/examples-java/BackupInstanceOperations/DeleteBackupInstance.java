
/**
 * Samples for BackupInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BackupInstanceOperations/DeleteBackupInstance.json
     */
    /**
     * Sample code: Delete BackupInstance.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void deleteBackupInstance(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().delete("000pikumar", "PratikPrivatePreviewVault1", "testInstance1",
            com.azure.core.util.Context.NONE);
    }
}
