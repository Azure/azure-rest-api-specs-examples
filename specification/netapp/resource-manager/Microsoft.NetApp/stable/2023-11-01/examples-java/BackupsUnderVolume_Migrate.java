
import com.azure.resourcemanager.netapp.models.BackupsMigrationRequest;

/**
 * Samples for BackupsUnderVolume MigrateBackups.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/BackupsUnderVolume_Migrate.json
     */
    /**
     * Sample code: BackupsUnderVolume_Migrate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsUnderVolumeMigrate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupsUnderVolumes().migrateBackups("myRG", "account1", "pool1", "volume1",
            new BackupsMigrationRequest().withBackupVaultId(
                "/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupVaults/backupVault1"),
            com.azure.core.util.Context.NONE);
    }
}
