import com.azure.core.util.Context;

/** Samples for LongTermRetentionBackups DeleteByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupDelete.json
     */
    /**
     * Sample code: Delete the long term retention backup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteTheLongTermRetentionBackup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getLongTermRetentionBackups()
            .deleteByResourceGroup(
                "testResourceGroup",
                "japaneast",
                "testserver",
                "testDatabase",
                "55555555-6666-7777-8888-999999999999;131637960820000000",
                Context.NONE);
    }
}
