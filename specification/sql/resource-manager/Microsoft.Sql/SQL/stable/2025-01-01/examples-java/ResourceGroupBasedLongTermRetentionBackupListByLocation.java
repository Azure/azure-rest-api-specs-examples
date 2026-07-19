
/**
 * Samples for LongTermRetentionBackups ListByResourceGroupLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ResourceGroupBasedLongTermRetentionBackupListByLocation.json
     */
    /**
     * Sample code: Get all long term retention backups under the location based on resource group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAllLongTermRetentionBackupsUnderTheLocationBasedOnResourceGroup(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups().listByResourceGroupLocation("testResourceGroup",
            "japaneast", null, null, com.azure.core.util.Context.NONE);
    }
}
