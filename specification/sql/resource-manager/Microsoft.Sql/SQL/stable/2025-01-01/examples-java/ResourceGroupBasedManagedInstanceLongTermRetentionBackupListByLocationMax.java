
/**
 * Samples for LongTermRetentionManagedInstanceBackups ListByResourceGroupLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByLocationMax.json
     */
    /**
     * Sample code: Get all long term retention backups under the location with maximal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAllLongTermRetentionBackupsUnderTheLocationWithMaximalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionManagedInstanceBackups().listByResourceGroupLocation(
            "testResourceGroup", "japaneast", null, null, 0L, 2L, "Properties/ManagedInstanceName eq 'testInstance1'",
            com.azure.core.util.Context.NONE);
    }
}
