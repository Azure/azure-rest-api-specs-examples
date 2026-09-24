
/**
 * Samples for LongTermRetentionManagedInstanceBackups ListByResourceGroupLocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-08-01-preview/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByLocation.json
     */
    /**
     * Sample code: Get all long term retention backups under the location.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAllLongTermRetentionBackupsUnderTheLocation(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionManagedInstanceBackups().listByResourceGroupLocation(
            "testResourceGroup", "japaneast", null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
