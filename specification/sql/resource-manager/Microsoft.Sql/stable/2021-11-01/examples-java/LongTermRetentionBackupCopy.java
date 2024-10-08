
import com.azure.resourcemanager.sql.models.BackupStorageRedundancy;
import com.azure.resourcemanager.sql.models.CopyLongTermRetentionBackupParameters;

/**
 * Samples for LongTermRetentionBackups Copy.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/LongTermRetentionBackupCopy.json
     */
    /**
     * Sample code: Copy the long term retention backup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void copyTheLongTermRetentionBackup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getLongTermRetentionBackups().copy("japaneast", "testserver",
            "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000",
            new CopyLongTermRetentionBackupParameters().withTargetServerResourceId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/resourceGroups/resourceGroup/servers/testserver2")
                .withTargetDatabaseName("testDatabase2").withTargetBackupStorageRedundancy(BackupStorageRedundancy.GEO),
            com.azure.core.util.Context.NONE);
    }
}
