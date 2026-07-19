
/**
 * Samples for LongTermRetentionManagedInstanceBackups ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceLongTermRetentionBackupListByLocation.json
     */
    /**
     * Sample code: Get all long term retention backups under the location.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAllLongTermRetentionBackupsUnderTheLocation(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionManagedInstanceBackups().listByLocation("japaneast", null, null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
