
/**
 * Samples for BackupPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PolicyCRUD/GetBackupPolicy.json
     */
    /**
     * Sample code: Get BackupPolicy.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getBackupPolicy(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupPolicies().getWithResponse("000pikumar", "PrivatePreviewVault", "OSSDBPolicy",
            com.azure.core.util.Context.NONE);
    }
}
