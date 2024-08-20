
/**
 * Samples for Backups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/BackupsUnderBackupVault_Get.
     * json
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
