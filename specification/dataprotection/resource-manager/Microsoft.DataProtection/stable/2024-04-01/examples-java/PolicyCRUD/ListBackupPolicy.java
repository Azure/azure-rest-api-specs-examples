
/**
 * Samples for BackupPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/PolicyCRUD/
     * ListBackupPolicy.json
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
