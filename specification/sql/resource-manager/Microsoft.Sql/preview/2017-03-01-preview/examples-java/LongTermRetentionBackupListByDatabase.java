import com.azure.core.util.Context;

/** Samples for LongTermRetentionBackups ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/LongTermRetentionBackupListByDatabase.json
     */
    /**
     * Sample code: Get all long term retention backups under the database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllLongTermRetentionBackupsUnderTheDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getLongTermRetentionBackups()
            .listByDatabase("japaneast", "testserver", "testDatabase", null, null, Context.NONE);
    }
}
