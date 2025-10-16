
/**
 * Samples for BackupPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PolicyCRUD/ListBackupPolicy.json
     */
    /**
     * Sample code: List BackupPolicy.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void listBackupPolicy(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupPolicies().list("000pikumar", "PrivatePreviewVault", com.azure.core.util.Context.NONE);
    }
}
