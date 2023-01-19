/** Samples for BackupPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/PolicyCRUD/GetBackupPolicy.json
     */
    /**
     * Sample code: Get BackupPolicy.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getBackupPolicy(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .backupPolicies()
            .getWithResponse("000pikumar", "PrivatePreviewVault", "OSSDBPolicy", com.azure.core.util.Context.NONE);
    }
}
