import com.azure.core.util.Context;

/** Samples for RestorableSqlDatabases List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBRestorableSqlDatabaseList.json
     */
    /**
     * Sample code: CosmosDBRestorableSqlDatabaseList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBRestorableSqlDatabaseList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getRestorableSqlDatabases()
            .list("WestUS", "d9b26648-2f53-4541-b3d8-3044f4f9810d", Context.NONE);
    }
}
