
import com.azure.resourcemanager.netapp.models.BackupsMigrationRequest;

/**
 * Samples for BackupsUnderVolume MigrateBackups.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/
     * BackupsUnderVolume_Migrate.json
     */
    /**
     * Sample code: BackupsUnderVolume_Migrate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsUnderVolumeMigrate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupsUnderVolumes().migrateBackups("myRG", "account1", "pool1", "volume1",
            new BackupsMigrationRequest().withBackupVaultId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupVaults/backupVault1"),
            com.azure.core.util.Context.NONE);
    }
}
