
/**
 * Samples for BackupInstances Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/BackupInstanceOperations/GetBackupInstance_BlobBackupAutoProtection.json
     */
    /**
     * Sample code: Get BackupInstance with BlobBackupAutoProtection.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getBackupInstanceWithBlobBackupAutoProtection(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().getWithResponse("blobrg", "blobvault", "blobbackupinstance",
            com.azure.core.util.Context.NONE);
    }
}
