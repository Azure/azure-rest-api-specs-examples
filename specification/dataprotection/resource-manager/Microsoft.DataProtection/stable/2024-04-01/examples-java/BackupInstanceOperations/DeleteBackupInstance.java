
/**
 * Samples for BackupInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * BackupInstanceOperations/DeleteBackupInstance.json
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
