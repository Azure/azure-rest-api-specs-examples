/** Samples for BackupPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/PolicyCRUD/ListBackupPolicy.json
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
