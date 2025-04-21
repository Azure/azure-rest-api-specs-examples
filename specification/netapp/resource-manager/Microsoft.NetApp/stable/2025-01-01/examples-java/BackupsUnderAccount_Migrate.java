
import com.azure.resourcemanager.netapp.models.BackupsMigrationRequest;

/**
 * Samples for BackupsUnderAccount MigrateBackups.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-01-01/examples/BackupsUnderAccount_Migrate.
     * json
     */
    /**
     * Sample code: BackupsUnderAccount_Migrate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsUnderAccountMigrate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupsUnderAccounts().migrateBackups("myRG", "account1",
            new BackupsMigrationRequest().withBackupVaultId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupVaults/backupVault1"),
            com.azure.core.util.Context.NONE);
    }
}
