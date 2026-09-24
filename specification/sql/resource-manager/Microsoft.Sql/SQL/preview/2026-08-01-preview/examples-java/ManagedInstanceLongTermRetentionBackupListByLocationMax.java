
/**
 * Samples for LongTermRetentionManagedInstanceBackups ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedInstanceLongTermRetentionBackupListByLocationMax.json
     */
    /**
     * Sample code: Get all long term retention backups under the location with maximal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAllLongTermRetentionBackupsUnderTheLocationWithMaximalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionManagedInstanceBackups().listByLocation("japaneast", null, null, 0L,
            2L, "Properties/ManagedInstanceName eq 'testInstance1'", com.azure.core.util.Context.NONE);
    }
}
