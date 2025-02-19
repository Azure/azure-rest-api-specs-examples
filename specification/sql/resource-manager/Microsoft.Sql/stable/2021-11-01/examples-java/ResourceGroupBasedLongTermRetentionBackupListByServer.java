
/**
 * Samples for LongTermRetentionBackups ListByResourceGroupServer.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ResourceGroupBasedLongTermRetentionBackupListByServer.json
     */
    /**
     * Sample code: Get all long term retention backups under the server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAllLongTermRetentionBackupsUnderTheServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getLongTermRetentionBackups().listByResourceGroupServer(
            "testResourceGroup", "japaneast", "testserver", null, null, com.azure.core.util.Context.NONE);
    }
}
