
import com.azure.resourcemanager.netapp.models.Backup;

/**
 * Samples for Backups Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/BackupsUnderBackupVault_Update.
     * json
     */
    /**
     * Sample code: BackupsUnderBackupVault_Update.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsUnderBackupVaultUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        Backup resource = manager.backups()
            .getWithResponse("myRG", "account1", "backupVault1", "backup1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
