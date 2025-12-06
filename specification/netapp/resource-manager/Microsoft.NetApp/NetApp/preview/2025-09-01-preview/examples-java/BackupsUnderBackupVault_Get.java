
/**
 * Samples for Backups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/BackupsUnderBackupVault_Get.json
     */
    /**
     * Sample code: BackupsUnderBackupVault_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsUnderBackupVaultGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backups().getWithResponse("myRG", "account1", "backupVault1", "backup1",
            com.azure.core.util.Context.NONE);
    }
}
