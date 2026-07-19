
import com.azure.resourcemanager.sql.models.BackupStorageRedundancy;
import com.azure.resourcemanager.sql.models.UpdateLongTermRetentionBackupParameters;

/**
 * Samples for LongTermRetentionBackups UpdateByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ResourceGroupBasedLongTermRetentionBackupUpdate.json
     */
    /**
     * Sample code: Update the long term retention backup.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateTheLongTermRetentionBackup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups()
            .updateByResourceGroup("testResourceGroup", "japaneast", "testserver", "testDatabase",
                "55555555-6666-7777-8888-999999999999;131637960820000000", new UpdateLongTermRetentionBackupParameters()
                    .withRequestedBackupStorageRedundancy(BackupStorageRedundancy.GEO),
                com.azure.core.util.Context.NONE);
    }
}
