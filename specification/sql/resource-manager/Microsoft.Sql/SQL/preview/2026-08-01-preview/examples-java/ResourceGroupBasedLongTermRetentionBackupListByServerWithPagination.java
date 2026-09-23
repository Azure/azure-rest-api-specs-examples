
/**
 * Samples for LongTermRetentionBackups ListByResourceGroupServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ResourceGroupBasedLongTermRetentionBackupListByServerWithPagination.json
     */
    /**
     * Sample code: Get long term retention backups under the server based on resource group with pagination.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getLongTermRetentionBackupsUnderTheServerBasedOnResourceGroupWithPagination(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups().listByResourceGroupServer("testResourceGroup",
            "japaneast", "testserver", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
