
/**
 * Samples for Backups ListByVault.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/BackupsUnderBackupVault_List.
     * json
     */
    /**
     * Sample code: Backups_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backups().listByVault("myRG", "account1", "backupVault1", null, com.azure.core.util.Context.NONE);
    }
}
