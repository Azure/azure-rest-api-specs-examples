
/**
 * Samples for LongTermRetentionManagedInstanceBackups ListByResourceGroupDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-08-01-preview/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByDatabase.json
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
