
/**
 * Samples for LongTermRetentionBackups ListByResourceGroupDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ResourceGroupBasedLongTermRetentionBackupListByDatabase.json
     */
    /**
     * Sample code: Get all long term retention backups under the database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAllLongTermRetentionBackupsUnderTheDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups().listByResourceGroupDatabase("testResourceGroup",
            "japaneast", "testserver", "testDatabase", null, null, com.azure.core.util.Context.NONE);
    }
}
