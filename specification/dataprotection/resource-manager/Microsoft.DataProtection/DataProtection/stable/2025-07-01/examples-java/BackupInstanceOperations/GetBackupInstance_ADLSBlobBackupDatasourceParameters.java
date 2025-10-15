
/**
 * Samples for BackupInstances Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BackupInstanceOperations/GetBackupInstance_ADLSBlobBackupDatasourceParameters.json
     */
    /**
     * Sample code: Get BackupInstance for ADLS Blob.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        getBackupInstanceForADLSBlob(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().getWithResponse("adlsrg", "adlsvault", "adlsbackupinstance",
            com.azure.core.util.Context.NONE);
    }
}
