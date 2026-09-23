
/**
 * Samples for LongTermRetentionBackups ListByResourceGroupDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ResourceGroupBasedLongTermRetentionBackupListByDatabaseWithPagination.json
     */
    /**
     * Sample code: Get long term retention backups under the database based on resource group with pagination.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getLongTermRetentionBackupsUnderTheDatabaseBasedOnResourceGroupWithPagination(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups().listByResourceGroupDatabase("testResourceGroup",
            "japaneast", "testserver", "testDatabase", null, null, com.azure.core.util.Context.NONE);
    }
}
