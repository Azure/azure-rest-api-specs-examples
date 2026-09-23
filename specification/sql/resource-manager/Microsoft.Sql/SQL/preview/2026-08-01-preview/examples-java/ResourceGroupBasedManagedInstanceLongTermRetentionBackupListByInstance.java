
/**
 * Samples for LongTermRetentionManagedInstanceBackups ListByResourceGroupInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-08-01-preview/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByInstance.json
     */
    /**
     * Sample code: Get all long term retention backups under the managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAllLongTermRetentionBackupsUnderTheManagedInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionManagedInstanceBackups().listByResourceGroupInstance(
            "testResourceGroup", "japaneast", "testInstance", null, null, com.azure.core.util.Context.NONE);
    }
}
