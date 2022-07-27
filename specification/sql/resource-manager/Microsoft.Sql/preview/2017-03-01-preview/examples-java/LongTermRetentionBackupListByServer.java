import com.azure.core.util.Context;

/** Samples for LongTermRetentionBackups ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/LongTermRetentionBackupListByServer.json
     */
    /**
     * Sample code: Get all long term retention backups under the server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllLongTermRetentionBackupsUnderTheServer(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getLongTermRetentionBackups()
            .listByServer("japaneast", "testserver", null, null, Context.NONE);
    }
}
