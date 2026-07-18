
/**
 * Samples for LongTermRetentionManagedInstanceBackups ListByResourceGroupDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByDatabase.json
     */
    /**
     * Sample code: Get all long term retention backups under the database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAllLongTermRetentionBackupsUnderTheDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionManagedInstanceBackups().listByResourceGroupDatabase(
            "testResourceGroup", "japaneast", "testInstance", "testDatabase", null, null,
            com.azure.core.util.Context.NONE);
    }
}
