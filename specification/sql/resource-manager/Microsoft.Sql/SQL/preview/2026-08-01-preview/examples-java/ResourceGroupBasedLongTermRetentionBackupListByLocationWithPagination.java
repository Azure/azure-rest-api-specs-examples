
/**
 * Samples for LongTermRetentionBackups ListByResourceGroupLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ResourceGroupBasedLongTermRetentionBackupListByLocationWithPagination.json
     */
    /**
     * Sample code: Get long term retention backups under the location based on resource group with pagination.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getLongTermRetentionBackupsUnderTheLocationBasedOnResourceGroupWithPagination(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups().listByResourceGroupLocation("testResourceGroup",
            "japaneast", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
