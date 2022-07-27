import com.azure.core.util.Context;

/** Samples for LongTermRetentionManagedInstanceBackups ListByResourceGroupDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByDatabase.json
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
            .getLongTermRetentionManagedInstanceBackups()
            .listByResourceGroupDatabase(
                "testResourceGroup", "japaneast", "testInstance", "testDatabase", null, null, Context.NONE);
    }
}
