
/**
 * Samples for LongTermRetentionBackups ListByResourceGroupServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ResourceGroupBasedLongTermRetentionBackupListByServer.json
     */
    /**
     * Sample code: Get all long term retention backups under the server based on resource group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAllLongTermRetentionBackupsUnderTheServerBasedOnResourceGroup(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups().listByResourceGroupServer("testResourceGroup",
            "japaneast", "testserver", null, null, com.azure.core.util.Context.NONE);
    }
}
