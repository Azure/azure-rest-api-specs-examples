
/**
 * Samples for BackupInstances Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/BackupInstanceOperations/GetBackupInstance_ADLSBlobBackupAutoProtection.json
     */
    /**
     * Sample code: Get BackupInstance with ADLSBlobBackupAutoProtection.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getBackupInstanceWithADLSBlobBackupAutoProtection(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().getWithResponse("adlsrg", "adlsvault", "adlsbackupinstance",
            com.azure.core.util.Context.NONE);
    }
}
