/** Samples for BackupPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/PolicyCRUD/DeleteBackupPolicy.json
     */
    /**
     * Sample code: Delete BackupPolicy.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void deleteBackupPolicy(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .backupPolicies()
            .deleteWithResponse("000pikumar", "PrivatePreviewVault", "OSSDBPolicy", com.azure.core.util.Context.NONE);
    }
}
