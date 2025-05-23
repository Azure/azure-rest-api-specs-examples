
import com.azure.resourcemanager.sql.models.BackupStorageRedundancy;
import com.azure.resourcemanager.sql.models.UpdateLongTermRetentionBackupParameters;

/**
 * Samples for LongTermRetentionBackups UpdateByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ResourceGroupBasedLongTermRetentionBackupUpdate.json
     */
    /**
     * Sample code: Update the long term retention backup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTheLongTermRetentionBackup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getLongTermRetentionBackups()
            .updateByResourceGroup("testResourceGroup", "japaneast", "testserver", "testDatabase",
                "55555555-6666-7777-8888-999999999999;131637960820000000", new UpdateLongTermRetentionBackupParameters()
                    .withRequestedBackupStorageRedundancy(BackupStorageRedundancy.GEO),
                com.azure.core.util.Context.NONE);
    }
}
